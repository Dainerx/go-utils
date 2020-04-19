package pipelines

import (
	"log"
	"math/rand"
	"time"
)

type Chappati struct {
	Verbose       bool
	Count         int           // number of chappati to make
	NumBakers     int           // number of cooks doing baking
	BakeTime      time.Duration // time to prepare dough for one chappati
	BakeStdDev    time.Duration // standard deviation of preparing dough for one chappati
	BakeBuf       int           // buffer slots between baking and making
	MakeTime      time.Duration // time to make one chappati
	MakeStdDev    time.Duration
	MakeBuf       int           // buffer slots between making and packaging
	PackageTime   time.Duration // time to inscribe one cake
	PackageStdDev time.Duration // standard deviation of inscribing time
}

type chappati int

func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}

func (ch *Chappati) baker(baked chan<- chappati) {
	for i := 0; i < ch.Count; i++ {
		c := chappati(i)
		if ch.Verbose {
			log.Println("baking", c)
		}
		work(ch.BakeTime, ch.BakeStdDev)
		baked <- c
	}
}

func (ch *Chappati) maker(made chan<- chappati, baked <-chan chappati) {
	for c := range baked {
		if ch.Verbose {
			log.Println("making", c)
		}
		work(ch.MakeTime, ch.MakeStdDev)
		made <- c
	}
}

func (ch *Chappati) packager(made <-chan chappati) {
	for i := 0; i < ch.Count; i++ {
		c := <-made
		if ch.Verbose {
			log.Println("packaging", c)
		}
		work(ch.PackageTime, ch.PackageStdDev)
		if ch.Verbose {
			log.Println("sold", c)
		}
	}
}

// Run.
func (ch *Chappati) Run(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan chappati, ch.BakeBuf)
		made := make(chan chappati, ch.MakeBuf)

		for i := 0; i < ch.NumBakers; i++ {
			go ch.baker(baked)
		}

		go ch.maker(made, baked)
		ch.packager(made)

		log.Println("We closed")
	}
}
