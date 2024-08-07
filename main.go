package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type Level

type Level int

const (
	Info Level = iota
	Warn
	Error
)

func main() {
	fmt.Printf("Level %s", Info)
}
