package ratsnest

// New creates a new Node with sourceData to be referenced in subsequent calls.
func New(source map[string]interface{}) (*Node, error) {
	n := &Node{
		Value:      source,
		sourceData: source,
	}
	valErr := n.isValid()
	if valErr != nil {
		return n, valErr
	}
	n.addChildren()
	return n, nil
}
