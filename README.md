# randline [![GoDoc](https://godoc.org/github.com/carlmjohnson/randline?status.svg)](https://godoc.org/github.com/carlmjohnson/randline) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/randline)](https://goreportcard.com/report/github.com/carlmjohnson/randline)

Chooses random line(s) from a file

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN="$(pwd)" GOPATH="$(mktemp -d)" go get github.com/carlmjohnson/randline
```

## Screenshots
```bash
$ randline -h
randline - display random line(s) from a file

Usage:

        randline [options]

Options:
  -lines int
        number of lines to show (default 1)
  -src value
        source file or URL (default stdin)

$ cat lunch.txt
Chinese
Thai
Indian
Italian

$ randline -src lunch.txt
Chinese

$ randline -lines 10 -src /usr/share/dict/words
preponderant
overpublic
amphitheatric
centifolious
fishworm
unhealthsomeness
repercussion
Samanid
subaerial
aculeolus
```
