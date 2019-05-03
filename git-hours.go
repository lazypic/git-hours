package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	since, before := beforeMonth()
	sincePtr := flag.String("since", since+" 00:00:00 "+timeZoneOffset(), "since(after) date")
	beforePtr := flag.String("before", before+" 23:59:59 "+timeZoneOffset(), "before date")
	authorPtr := flag.String("author", "", "author name") // git option : --author="\(Adam\)\|\(Jon\)"
	debugPtr := flag.Bool("debug", false, "debug mode")
	helpPtr := flag.Bool("help", false, "print help")
	flag.Parse()
	if *helpPtr {
		flag.PrintDefaults()
		os.Exit(0)
	}
	//checkMultiname
	var author string
	if strings.Contains(*authorPtr, ",") {
		author += `\(`
		author += strings.Join(strings.Split(*authorPtr, ","), `\)\|\(`)
		author += `\)`
	} else {
		author = *authorPtr
	}
	cmd := exec.Command(
		"git",
		"--no-pager",
		"log",
		"--reverse",
		"--date=iso-local",
		`--pretty=format:%ad %an %s`,
		fmt.Sprintf(`--author=%s`, author),
		fmt.Sprintf(`--since="%s"`, *sincePtr),
		fmt.Sprintf(`--before="%s"`, *beforePtr),
	)
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
	total, err := time.ParseDuration("0h")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if stdout.String() == "" {
		fmt.Printf("From %q to %q : %s\n", *sincePtr, *beforePtr, total)
		os.Exit(0)
	}

	var beforeCommitTime time.Time
	for n, l := range strings.Split(stdout.String(), "\n") {
		getTime := findISO8601.FindString(l)
		if getTime == "" {
			continue
		}
		rfctime, err := ISO8601_to_RFC3339(getTime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		t, err := time.Parse(time.RFC3339, rfctime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		elapsed := t.Sub(beforeCommitTime)
		if *debugPtr {
			if n != 0 {
				fmt.Println(elapsed, ">")
			}
			fmt.Println("\t", l)
		}
		h, err := time.ParseDuration("1h")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		if elapsed < h*2 {
			total += elapsed
		} else {
			total += h
		}
		beforeCommitTime = t
	}
	fmt.Printf("From %q to %q : %s\n", *sincePtr, *beforePtr, total)
}
