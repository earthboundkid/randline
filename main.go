package main

import (
	"os"

	"github.com/carlmjohnson/exitcode"
	"github.com/carlmjohnson/randline/randline"
)

func main() {
	exitcode.Exit(randline.CLI(os.Args[1:]))
}
