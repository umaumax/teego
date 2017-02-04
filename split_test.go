package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func TestSplitLongLineIncludingR(t *testing.T) {
	w := &bytes.Buffer{}
	scanner := bufio.NewScanner(w)
	//	NOTE bufio.Scanner token too long error
	//	scanner.Split(bufio.ScanLines)
	scanner.Split(ScanLines)
	for i := 0; i < 10000; i++ {
		fmt.Fprintf(w, "i=%d\r", i)
	}
	for scanner.Scan() {
		text := scanner.Text()
		_ = text
	}
	if err := scanner.Err(); err != nil {
		t.Errorf("%s", err)
	}
}
