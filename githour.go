package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"strings"
	)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// zone
	// start
	// end
	cmd := exec.Command("git", "--no-pager", "log", "--reverse", "--date=iso", `--pretty=format:%ad %an %s`, `--after="2018-01-01 00:00:00 +0900"`, `--before="2019-12-31 23:59:59 +0900"`)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if stderr.String() != "" {
		fmt.Fprintf(os.Stderr, stderr.String())
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, stderr.String())
		os.Exit(1)
	}
	total, err := time.ParseDuration("1h")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if stdout.String() == "" {
		fmt.Println("0h")
		os.Exit(0)
	}

	var before time.Time
	for n, l := range strings.Split(stdout.String(), "\n") {
		lists := strings.Split(l, " ")
		iso2rfc3339 := fmt.Sprintf("%sT%s%s:%s", lists[0], lists[1], lists[2][0:3], lists[2][3:5])
		t, err := time.Parse(time.RFC3339, iso2rfc3339)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		if n == 0 {
			before = t
			continue
		}
		elapsed := t.Sub(before)
		fmt.Println("\t",elapsed)
		h, err := time.ParseDuration("1h")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		if elapsed < h*2 {
			total += elapsed
		} else {
			total += h
		}
		before = t
	}
	fmt.Println(total)
	// how to get local timezone offset value
	zone, offset := time.Now().Zone()
	fmt.Println(zone, offset/60/60)
}
