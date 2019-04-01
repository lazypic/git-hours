package main

import (
	"fmt"
	"time"
)

func timeZoneOffset() string {
	_, offset := time.Now().Zone()
	if offset > 0 {
		return fmt.Sprintf("+%04d", offset/60/60*100)
	}
	return fmt.Sprintf("-%04d", (-1*offset)/60/60*100)
}
