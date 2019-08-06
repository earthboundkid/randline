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
	cnt := fl.Int("lines", 1, "number of lines to show (<1 for same as input)")
	replace := fl.Bool("replace", false, "allow the same line to appear more than once")
	byWord := fl.Bool("split-words", false, "split on words rather than lines")
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

	p, err := NewPicker(src, *byWord, *replace)
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
	ss      []string
	r       rand.Rand
	replace bool
}

func NewPicker(rc io.ReadCloser, byWord, replace bool) (*Picker, error) {
	p := Picker{
		r:       *rand.New(rand.NewSource(time.Now().UnixNano())),
		replace: replace,
	}

	sc := bufio.NewScanner(rc)
	defer rc.Close()

	if byWord {
		sc.Split(bufio.ScanWords)
	}

	for sc.Scan() {
		p.ss = append(p.ss, sc.Text())
	}
	return &p, exitcode.Set(sc.Err(), 3)
}

func (p *Picker) Pick() string {
	if len(p.ss) == 0 {
		return ""
	}

	n := p.r.Intn(len(p.ss))
	if !p.replace {
		p.ss[0], p.ss[n] = p.ss[n], p.ss[0]
		r := p.ss[0]
		p.ss = p.ss[1:]
		return r
	}
	return p.ss[n]
}

func (p *Picker) Output(w io.Writer, cnt int) error {
	if cnt < 1 {
		cnt = len(p.ss)
	}
	for i := 0; i < cnt; i++ {
		s := p.Pick()
		if _, err := fmt.Fprintln(w, s); err != nil {
			return err
		}
	}
	return nil
}
