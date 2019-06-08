package goexpand

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var matcher = regexp.MustCompile("^(.*)\\[([\\d]+):([\\d]+)](.*)$")

// Expand instantiates strings from template pattern.
func Expand(pattern string) []string {
	results := []string{}
	for _, s := range strings.Split(pattern, ",") {
		if m := matcher.FindStringSubmatch(s); m != nil {
			start, _ := strconv.Atoi(m[2])
			end, _ := strconv.Atoi(m[3])
			width := len(m[3])
			f := "%0" + strconv.Itoa(width) + "d"
			for i := start; i <= end; i++ {
				r := m[1] + fmt.Sprintf(f, i) + m[4]
				if strings.Contains(r, "[") {
					results = append(results, Expand(r)...)
				} else {
					results = append(results, r)
				}
			}
		} else {
			results = append(results, s)
		}
	}
	return results
}
