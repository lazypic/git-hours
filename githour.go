package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var initZone string
var timeFormat = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// how to get local timezone offset value
	_, offset := time.Now().Zone()
	if offset > 0 {
		initZone = fmt.Sprintf("+%04d", offset/60/60*100)
	} else {
		initZone = fmt.Sprintf("-%04d", offset/60/60*100)
	}
}

func main() {
	startPtr := flag.String("start", "2018-01-01", "start date")
	endPtr := flag.String("end", "2019-12-31", "end date")
	zonePtr := flag.String("zone", initZone, "zone offset time")
	debugPtr := flag.Bool("debug", false, "debug mode")
	helpPtr := flag.Bool("help", false, "print help")
	flag.Parse()
	if *helpPtr {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !timeFormat.MatchString(*startPtr) {
		fmt.Println("not matching start date format. must be 0000-00-00")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !timeFormat.MatchString(*endPtr) {
		fmt.Println("not matching end date format. must be 0000-00-00")
		flag.PrintDefaults()
		os.Exit(1)
	}
	cmd := exec.Command(
		"git",
		"--no-pager",
		"log",
		"--reverse",
		"--date=iso",
		`--pretty=format:%ad %an %s`,
		fmt.Sprintf(`--after="%s 00:00:00 %s"`, *startPtr, *zonePtr),
		fmt.Sprintf(`--before="%s 23:59:59 %s"`, *endPtr, *zonePtr),
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
		elapsed := t.Sub(before)
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
		before = t
	}
	fmt.Println(*startPtr, "~", *endPtr, total)
}
