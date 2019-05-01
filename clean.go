package scraper

func Clean(input map[string]string, rules []Rule) map[string]string {

	for _, rule := range rules {
		input = rule.Transform(input)
	}

	return input
}

func CleanAll(input []map[string]string, rules []Rule) []map[string]string {

	var output []map[string]string

	for _, item := range input {
		output = append(output, Clean(item, rules))
	}

	return output
}
