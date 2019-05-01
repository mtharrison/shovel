package scraper

import (
	"log"
	"regexp"
)

type Rule interface {
	Transform(map[string]string) map[string]string
}

type RegexReplace struct {
	Field   string
	Find    string
	Replace string
}

func (rule RegexReplace) Transform(m map[string]string) map[string]string {

	r, err := regexp.Compile(rule.Find)
	if err != nil {
		log.Fatal(err)
	}

	m[rule.Field] = r.ReplaceAllString(m[rule.Field], rule.Replace)

	return m
}

type Split struct {
	Field   string
	Find    string
	Targets []string
}

func (rule Split) Transform(m map[string]string) map[string]string {

	r, err := regexp.Compile(rule.Find)
	if err != nil {
		log.Fatal(err)
	}

	parts := r.Split(m[rule.Field], -1)

	for i, target := range rule.Targets {
		if target == "_" {
			continue
		}

		if len(parts) > i {
			m[target] = parts[i]
		}
	}

	delete(m, rule.Field)

	return m
}
