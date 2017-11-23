package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/relsa/exercise-gopl.io/ch02/ex02/unitconv"
)

var (
	t = flag.Bool("t", false, "templature")
	l = flag.Bool("l", false, "length")
	w = flag.Bool("w", false, "weight")
)

func main() {
	flag.Parse()

	if flag.NFlag() > 0 {
		arg := flag.Arg(0)
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail to ParseFloat: %s", arg)
			os.Exit(1)
		}

		printConv(val, *t, *l, *w)
	} else {
		var opt string
		var val float64

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("input type [t, l, w]: ")
		if scanner.Scan() {
			opt = scanner.Text()
		}

		fmt.Printf("input value: ")
		var tmp string
		if scanner.Scan() {
			tmp = scanner.Text()
		}
		val, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail to ParseFloat: %s", tmp)
			os.Exit(1)
		}

		printConv(val, opt == "t", opt == "l", opt == "w")
	}
}

func printConv(val float64, t, l, w bool) {
	if t {
		f := unitconv.Fahrenheit(val)
		c := unitconv.Celsius(val)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FToC(f), c, unitconv.CToF(c))
	}
	if l {
		f := unitconv.Feet(val)
		m := unitconv.Metre(val)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FToM(f), m, unitconv.MToF(m))
	}
	if w {
		p := unitconv.Pound(val)
		k := unitconv.Kilogram(val)
		fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PToK(p), k, unitconv.KToP(k))
	}
}
