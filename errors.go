package ratsnest

type (
	// NoCriteriaError is returned when a root node has been initialized but no other nodes with values criterion
	// have been added to the root.
	NoCriteriaError struct{}

	// NodeNotFoundError is returned when a Node has be Required but not found in the parent node.
	NodeNotFoundError struct{}
)

// Error implements the error interface.
func (err NoCriteriaError) Error() string {
	return "No requirements have been added to the root node."
}

// Error implements the error interface.
func (err NodeNotFoundError) Error() string {
	return "The required node was not found in the parent."
}