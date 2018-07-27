package main

import (
	"fmt"
	"os"
)

// TerminalWriter writes to a terminal, breaking the lines
// at the given width.
type TerminalWriter struct {
	width int
}

// Write writes slice p to stdout, in chunks of tw.width bytes,
// separated by newline.
// It returns the number of successfully written bytes, and
// any error that occurred.
// If the complete slice is written, Write returns error io.EOF
func (tw *TerminalWriter) Write(p []byte) (n int, err error) {
	length := len(p)
	newLineByte := []byte("\n")
	idx := 1
	for ; idx*tw.width < length; idx++ {
		os.Stdout.Write(p[(idx-1)*tw.width : idx*tw.width])
		os.Stdout.Write(newLineByte)
	}

	os.Stdout.Write(p[(idx-1)*tw.width : length])
	// for idx := 0; idx < tw.width; idx++ {
	// 	fmt.Printf("%s", string(p[n]))
	// 	n++
	// }
	n = tw.width
	fmt.Printf("\n")
	return n, err
}

// NewTerminalWriter creates a new TerminalWriter. width is
// the terminal's width.
func NewTerminalWriter(width int) *TerminalWriter {
	return &TerminalWriter{width: width}
}

func main() {
	s := []byte("This is a long string converted into a byte slice for testing the TerminalWriter.")

	tw := NewTerminalWriter(31)
	n, err := tw.Write(s)
	fmt.Println(n, "bytes written. Error:", err)

}
