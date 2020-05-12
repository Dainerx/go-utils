package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

const (
	MaxConcurrentRead = 20
	KiloBytes         = "kb"
	MegaBytes         = "mb"
	GigaBytes         = "gb"
)

type dirdata struct {
	name string
	size int64
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var unit = flag.String("u", "kb", "unit to display dir size in")

var done = make(chan struct{})

func cancelled() bool {
	select {
	// Recall that after a channel was closed or drained all values
	// subsequent receive operations will proceed immediately and give zero values.
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	flag.Parse()

	dirs := flag.Args() // take input dirs
	if len(dirs) == 0 {
		dirs = []string{"."} // run du on launching dir
	}

	var waitgroup sync.WaitGroup
	directories := make(map[string]*dirdata)
	fileSizes := make(chan *dirdata)
	for _, dir := range dirs {
		d := &dirdata{dir, 0}
		directories[dir] = d
	}

	// Background walking routines
	for dir := range directories {
		waitgroup.Add(1)
		go walkDir(dir, dir, &waitgroup, fileSizes)
	}
	go func() {
		waitgroup.Wait()
		close(fileSizes)
	}()

	// Cancellation routine
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// If verbose flag is not set the channel will remain with
	// default value nil, thus will never receive an event on its channel.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	tablewriter := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	printDirectoryNames(dirs, tablewriter)

	// Main routine to compute totals and display
Loop:
	for {
		select {
		case <-done:
			// drain the fileSizes channel to unblock go routines
			for range fileSizes {
				// do nothing just drain values
			}

			fmt.Println("Exit.")
			return
		case dirInfo, ok := <-fileSizes:
			if !ok {
				break Loop
			}
			directories[dirInfo.name].size += dirInfo.size
		case <-tick: // print
			printDiskUsage(dirs, directories, tablewriter)
		}
	}
	printDiskUsage(dirs, directories, tablewriter)
	fmt.Printf("\n")
}

func printDirectoryNames(roots []string, tablewriter *tabwriter.Writer) {
	for _, root := range roots {
		fmt.Fprintf(tablewriter, "%v\t", root)
	}
	fmt.Fprintf(tablewriter, "\n")

	for i := 0; i < len(roots); i++ {
		fmt.Fprintf(tablewriter, "%v\t", "-----")
	}
	fmt.Fprintf(tablewriter, "\n")
	tablewriter.Flush()
}

func printDiskUsage(roots []string, directories map[string]*dirdata, tablewriter *tabwriter.Writer) {
	fmt.Printf("\r") // consumes last displayed output
	for _, root := range roots {
		switch strings.ToLower(*unit) {
		case KiloBytes:
			fmt.Fprintf(tablewriter, "%v\t", fmt.Sprintf("%.1f KB", float64(directories[root].size)/1e3))
		case MegaBytes:
			fmt.Fprintf(tablewriter, "%v\t", fmt.Sprintf("%.1f MB", float64(directories[root].size)/1E6))
		case GigaBytes:
			fmt.Fprintf(tablewriter, "%v\t", fmt.Sprintf("%.1f GB", float64(directories[root].size)/1e9))
		default:
			fmt.Fprintf(tablewriter, "%v\t", fmt.Sprintf("%.1f KB", float64(directories[root].size)/1e3))
		}
	}
	tablewriter.Flush()
}

func walkDir(dir string, root string, waitgroup *sync.WaitGroup, fileSizes chan<- *dirdata) {
	defer waitgroup.Done()

	if cancelled() { // stop event detected
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subpath := filepath.Join(dir, entry.Name())
			waitgroup.Add(1)
			go walkDir(subpath, root, waitgroup, fileSizes)
		}

		fileSizes <- &dirdata{root, entry.Size()}
	}
}

var sema = make(chan struct{}, MaxConcurrentRead) // Limiting to the concurrency for opening files

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // get token
	case <-done:
		return nil
	}

	defer func() { <-sema }() // release token after return

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}

	return entries
}
