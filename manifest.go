package shovel

// A Manifest tells the shovel how to find the collection on the page
type Manifest struct {
	URL                string
	SingleItemSelector string
	AttributeSelectors map[string]string
}
