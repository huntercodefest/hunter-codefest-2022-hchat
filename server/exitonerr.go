package main

import (
	"fmt"
	"os"
)

// For convienience sake only
func ExitOnErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}
