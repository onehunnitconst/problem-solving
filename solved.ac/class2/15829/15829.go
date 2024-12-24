package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdio = bufio.NewReadWriter(
	bufio.NewReader(os.Stdin),
	bufio.NewWriter(os.Stdout),
)

func main() {
	defer stdio.Flush()

	var len uint64
	var str string

	var r uint64 = 31
	var m uint64 = 1234567891

	var sum uint64 = 0

	fmt.Fscanln(stdio, &len)
	fmt.Fscanln(stdio, &str)

	var index uint64

	for index = 0; index < len; index++ {
		id := uint64(str[index] - 'a' + 1)

		var coef uint64 = 1;

		var times uint64
		for times = 0; times < index; times++ {
			coef = (coef * r) % m
		}

		sum += id * coef % m
	}

	var result uint64 = sum % m

	fmt.Fprintln(stdio, result)
}