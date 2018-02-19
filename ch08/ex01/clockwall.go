package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	loc  string
	uri  string
	time string
}

func main() {
	clocks, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clocks {
		go c.getTime()
	}

	times := make([]string, len(clocks))
	for {
		time.Sleep(time.Second)
		for i, c := range clocks {
			times[i] = fmt.Sprintf("%s: %s", c.loc, c.time)
		}

		fmt.Printf("\r%s", strings.Join(times, ", "))
	}
}

func (c *clock) getTime() {
	conn, err := net.Dial("tcp", c.uri)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		c.time = string(b)
	}
}

func parseArgs() ([]*clock, error) {
	var res []*clock

	if len(os.Args) == 1 {
		return res, nil
	}

	for _, arg := range os.Args[1:] {
		argfrags := strings.Split(arg, "=")
		if len(argfrags) != 2 {
			return nil, fmt.Errorf("invalid args: %q", arg)
		}

		c := &clock{loc: argfrags[0], uri: argfrags[1], time: ""}
		res = append(res, c)
	}

	return res, nil
}
