package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type packageData struct {
	ImportPath string
	Deps       []string
}

func main() {
	given, err := list(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}

	all, err := list("...")
	if err != nil {
		log.Fatal(err)
	}

	printDependents(given, all)
}

func printDependents(given, all []*packageData) {
loopAll:
	for _, a := range all {
		for _, g := range given {
			if !contains(a.Deps, g.ImportPath) {
				continue loopAll
			}
		}
		fmt.Println(a.ImportPath)
	}
}

func contains(ss []string, s string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}

func list(args ...string) ([]*packageData, error) {
	cmdArgs := []string{"list", "-json"}
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command("go", cmdArgs...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("fail to get pipe: %v", err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("fail to start cmd: %v", err)
	}

	decoder := json.NewDecoder(stdout)
	var pdata []*packageData
	for decoder.More() {
		var pd packageData
		err := decoder.Decode(&pd)
		if err != nil {
			return nil, fmt.Errorf("fail to decode json: %v", err)
		}
		pdata = append(pdata, &pd)
	}

	return pdata, nil
}
