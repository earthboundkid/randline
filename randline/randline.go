package randline

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/carlmjohnson/exitcode"
	"github.com/carlmjohnson/flagext"
)

func CLI(args []string) error {
	fl := flag.NewFlagSet("randline", flag.ContinueOnError)
	cnt := fl.Int("lines", 1, "number of lines to show")
	src := flagext.FileOrURL(flagext.StdIO, nil)
	fl.Var(src, "src", "source file or URL")
	fl.Usage = func() {
		fmt.Fprintf(fl.Output(), `randline - display random line(s) from a file

Usage:

	randline [options]

Options:
`)
		fl.PrintDefaults()
	}
	if err := fl.Parse(args); err != nil {
		return err
	}

	p, err := NewPicker(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Load error: %v\n", err)
		return err
	}
	if err = p.Output(os.Stdout, *cnt); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
	}
	return nil
}

type Picker struct {
	ss []string
	r  rand.Rand
}

func NewPicker(rc io.ReadCloser) (*Picker, error) {
	var p Picker
	p.r = *rand.New(rand.NewSource(time.Now().UnixNano()))

	sc := bufio.NewScanner(rc)
	defer rc.Close()

	for sc.Scan() {
		p.ss = append(p.ss, sc.Text())
	}
	return &p, exitcode.Set(sc.Err(), 3)
}

func (p *Picker) Output(w io.Writer, cnt int) error {
	for i := 0; i < cnt; i++ {
		n := p.r.Intn(len(p.ss))
		s := p.ss[n]
		if _, err := fmt.Fprintln(w, s); err != nil {
			return err
		}
	}
	return nil
}
