package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x int
	y int
}

var stdio = bufio.NewReadWriter(
	bufio.NewReader(os.Stdin),
	bufio.NewWriter(os.Stdout),
)

func main() {
	defer stdio.Flush()

	var n int;

	fmt.Fscanln(stdio, &n)

	var points [100000]Point;

	for index := 0; index < n; index++ {
		fmt.Fscanln(stdio, &points[index].x, &points[index].y)
	}
	
	sort.Slice(points[:n], func(i, j int) bool {
		if points[i].y == points[j].y {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})

	for i := 0; i < n; i++ {
		fmt.Fprintln(stdio, points[i].x, points[i].y)
	}
}