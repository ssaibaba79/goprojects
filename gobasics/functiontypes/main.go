package main

import (
	"fmt"
	"sort"
)

type sorter func([]string) []string

var AscendSort = func(input []string) []string {
	sort.Strings(input)
	return input
}

var DescendSort = func(input []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(input)))
	return input
}

func main() {
	v := []string{"Sarvan", "Anicha", "Shashti", "Poornima"}
	printOrdered(v, AscendSort)
	printOrdered(v, DescendSort)

}

func printOrdered(v []string, sorter sorter) {
	fmt.Println(sorter(v))
}
