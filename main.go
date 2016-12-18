package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	dcap "github.com/kofemann/dcap4go/dcap"
)

func main() {

	flag.Usage = func() {
		app := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "%s: [options] <src> <dest>\n", app)
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	src := flag.Arg(0)
	dest := flag.Arg(1)

	fmt.Println(src)
	fmt.Println(dest)
	in, err := dcap.Open(src, os.O_RDONLY, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open source: %v\n", err)
		os.Exit(2)
	}
	defer in.Close()

	out, err := dcap.Open(dest, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open destination: %v\n", err)
		os.Exit(3)
	}
	defer out.Close()

	copydata(in.Reader, out.Writer)
}

func copydata(src io.Reader, dst io.Writer) (int64, error) {
	return io.Copy(dst, src)
}
