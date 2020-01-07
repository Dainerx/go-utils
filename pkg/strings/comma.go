package strings

import (
	"bytes"
	"strings"
)

func CommaRec(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return CommaRec(s[:n-3]) + "," + s[n-3:]
}

func CommaBytesBuffer(s string) string {
	var buf bytes.Buffer
	n := len(s)
	left := n % 3
	if left == 0 {
		left = 3
	}
	buf.WriteString(s[:left])
	for i := left; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func CommaFloatEnhanced(s string) string {
	var buf bytes.Buffer
	// locate the start of significant digits part
	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart = 1
	}
	// locate the end of significant digits part
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	buf.WriteString(CommaBytesBuffer(s[mantissaStart:mantissaEnd]))
	buf.WriteString(s[mantissaEnd:])
	return buf.String()
}

// From https://github.com/torbiak/gopl/blob/master/ex3.11/main.go
func CommaFloatEnhancedSafe(s string) string {
	b := bytes.Buffer{}
	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3
	if pre > 0 {
		b.Write([]byte(mantissa[:pre]))
		if len(mantissa) > pre {
			b.WriteString(",")
		}
	}
	for i, c := range mantissa[pre:] {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	b.WriteString(s[mantissaEnd:])
	return b.String()
}
