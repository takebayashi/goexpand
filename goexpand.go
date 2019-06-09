package goexpand

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Expand returns expanded values from template pattern with default expander.
func Expand(pattern string) []string {
	return DefaultExpander.Expand(pattern)
}

// Expander provides functions for expanding template patterns.
type Expander interface {
	// Expand returns expanded values from template pattern.
	Expand(pattern string) []string
}

// DefaultExpander provides default function for expanding template patterns.
var DefaultExpander Expander

// NewExpander creates new Expander with specified configurations.
func NewExpander(startBracket, endBracket, rangeDelimiter string) (Expander, error) {
	p := "^(.*)" + regexp.QuoteMeta(startBracket) +
		"([\\d]+)" + regexp.QuoteMeta(rangeDelimiter) + "([\\d]+)" +
		regexp.QuoteMeta(endBracket) + "(.*)$"
	m, err := regexp.Compile(p)
	return expander{matcher: m}, err
}

type expander struct {
	matcher *regexp.Regexp
}

func (ex expander) Expand(pattern string) []string {
	results := []string{}
	for _, s := range strings.Split(pattern, ",") {
		if m := ex.matcher.FindStringSubmatch(s); m != nil {
			start, _ := strconv.Atoi(m[2])
			end, _ := strconv.Atoi(m[3])
			width := len(m[3])
			f := "%0" + strconv.Itoa(width) + "d"
			for i := start; i <= end; i++ {
				r := m[1] + fmt.Sprintf(f, i) + m[4]
				if ex.matcher.MatchString(r) {
					results = append(results, ex.Expand(r)...)
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

func init() {
	var err error
	DefaultExpander, err = NewExpander("[", "]", ":")
	if err != nil {
		panic(err)
	}
}
