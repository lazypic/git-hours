package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
	)

func main() {
	cmd := exec.Command("git", "--no-pager", "log", "--reverse", "--date=iso", `--pretty=format:%ad %an %s`, `--after="2018-01-01 00:00:00 +0900"`, `--before="2018-12-31 23:59:59 +0900"`)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(stdout))
	fmt.Println(time.Now().Format(time.RFC3339))
}
