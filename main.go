package main

import (
	"flag"
	"fmt"
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
	in, err := dcap.Open(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open source: %v\n", err)
		os.Exit(2)
	}
	defer in.Close()

	out, err := dcap.Open(dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open destination: %v\n", err)
		os.Exit(3)
	}
	defer out.Close()

	copydata(in, out)
}

func copydata(in interface{}, out interface{}) int {
	return 5
}
