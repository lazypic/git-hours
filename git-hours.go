package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const DTF string = "2006-01-02" // Default Time Format
var timeFormat = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var ISO8601 = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [+-]\d{4}$`)
var findISO8601 = regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [+-]\d{4}`)
var RFC3339 = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$`)

func ISO8601_to_RFC3339(t string) (string, error) {
	if !ISO8601.MatchString(t) {
		return t, errors.New("time string is not ISO8601 format.")
	}
	return fmt.Sprintf("%sT%s%s:%s", t[0:10], t[11:19], t[20:23], t[23:25]), nil
}

func beforeMonth() (string, string) {
	y, m, _ := time.Now().Date()
	if m == 1 {
		y -= 1
		m = 12
	} else {
		m -= 1
	}
	since := time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location())
	return fmt.Sprintf(since.Format(DTF)), fmt.Sprintf(since.AddDate(0, 1, 0).Add(-time.Nanosecond).Format(DTF))
}

func thisMonth() (string, string) {
	y, m, _ := time.Now().Date()
	since := time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location())
	return fmt.Sprintf(since.Format(DTF)), fmt.Sprintf(since.AddDate(0, 1, 0).Add(-time.Nanosecond).Format(DTF))
}

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
