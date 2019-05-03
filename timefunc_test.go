package main

import (
	"testing"
)

func TestISO8601ToRFC3339(t *testing.T) {
	cases := []struct {
		in   string
		want string
		err  string
	}{{
		in:   "2019-03-30 10:00:00 +0900",
		want: "2019-03-30T10:00:00+09:00",
		err:  "",
	}, {
		in:   "",
		want: "",
		err:  "time string is not ISO8601 format",
	}}
	for _, c := range cases {
		got, err := ISO8601ToRFC3339(c.in)
		// check err
		if err != nil {
			if err.Error() != c.err {
				t.Fatalf("ISO8601ToRFC3339(%v): got err %s, want err %v", c.in, err.Error(), c.err)
			}
		}
		// check value
		if got != c.want {
			t.Fatalf("ISO8601ToRFC3339(%v): got %s, want %s", c.in, got, c.want)
		}
	}
}
