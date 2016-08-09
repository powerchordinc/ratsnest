package ratsnest

import (
	"errors"
	"strings"
)

type (
	// Node is particular nested branch of a map[string]interface{}
	Node struct {
		// Value indicates either the value of an existing node or the desired value of a node
		// to be validated, depending on the context in which it is used.
		Value             interface{}
		// Key is either the key of the map for the node or the key that the desired value must
		// be within within the node.
		Key               string
		// Case indicates whether case sensitivity should be required for a match to occur. Defaults
		// to case-sensitive.
		Case              CaseSensitivity
		// MaxDepth is used to validate that a value appears no farther down the hierarchy of a node
		// than MaxDepth levels. 0 is infinite, and is the default. The depth of values whose keys
		// are directly inside the node is 1. MaxDepth should be thought of as applying to the depth
		// of the Value rather than the depth of the Key.
		MaxDepth          int

		// childRequirements houses added-on nodes for validation of nodes within other nodes.
		childRequirements []*Node
		// childNodes houses Nodes that are are children of a node initialized with the New function.
		childNodes        []*Node
		// sourceData points to the original data which was handed to the New initializer.
		sourceData        map[string]interface{}
	}

	// CaseSensitivity is a placeholder for an enum case sensitivity value.
	CaseSensitivity int
)

const (
	CaseSensitive CaseSensitivity = iota
	CaseInsensitive
)

// addChildren adds the childNodes to an initialized ratsnest Node.
func (n *Node) addChildren() {
	switch n.Value.(type) {
	case map[string]interface{}:
		for k, v := range n.Value.(map[string]interface{}) {
			cn := &Node{
				Key: k,
				Value: v,
				sourceData: n.sourceData,
			}
			cn.addChildren()
			n.childNodes = append(n.childNodes, cn)
		}
	case []interface{}:
		for _, v := range n.Value.([]interface{}) {
			cn := &Node{
				Value: v,
				sourceData: n.sourceData,
			}
			cn.addChildren()
			n.childNodes = append(n.childNodes, cn)
		}
		// append the entire array as a child if that's what you're looking for
		n.childNodes = append(n.childNodes, &Node{
			Value: n.Value,
		})
	default:
		cn := &Node{
			Value: n.Value,
			sourceData: n.sourceData,
		}
		n.childNodes = append(n.childNodes, cn)
	}
}

// Require registers a new child node with additional values criteria within the receiving node.
func (n *Node) Require(node Node) (*Node, error) {
	rn := &node
	rn.sourceData = n.sourceData
	n.childRequirements = append(n.childRequirements, rn)
	err := rn.isValid()
	if err != nil {
		return rn, err
	}
	found := n.checkChildren(node, 1)
	if found == nil {
		return nil, NodeNotFoundError{}
	}
	return found, nil
}

// checkChildren checks children of a node for the existence of another node.
func (n *Node) checkChildren(reqNode Node, depth int) *Node {
	if reqNode.MaxDepth > 0 && depth > reqNode.MaxDepth {
		return nil
	}
	for _, cn := range n.childNodes {
		switch cn.Value.(type) {
		case []interface{}:
			if _, reqNodeIsArr := reqNode.Value.([]interface{}); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range cn.Value.([]interface{}) {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		default:
			if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, reqNode.Case) {
				return cn
			}
		}
		if fn := cn.checkChildren(reqNode, depth+1); fn != nil {
			return fn
		}
	}
	return nil
}

// checkKeyVal simply checks whether a key and a value match.
func checkKeyVal(key, expKey string, val, expVal interface{}, cs CaseSensitivity) bool {
	switch key {
	case "":
		return isMatch(expVal, val, cs)
	default:
		if cs == CaseInsensitive {
			key, expKey = strings.ToLower(key), strings.ToLower(expKey)
		}
		switch expVal {
		case nil:
			if key == expKey {
				return true
			}
		default:
			if key == expKey && isMatch(expVal, val, cs) {
				return true
			}
		}
	}
	return false
}

func isMatch(t1, t2 interface{}, cs CaseSensitivity) bool {
	switch t1.(type) {
	case map[string]interface{}:
		t1Map := t1.(map[string]interface{})
		t2Map, t2IsMap := t2.(map[string]interface{})
		if !t2IsMap || len(t1Map) != len(t2Map) {
			return false
		}
		for k, v := range t1Map {
			if t2Val, t2OK := t2Map[k]; !t2OK || !isMatch(v, t2Val, cs) {
				return false
			}
		}
		return true
	case []interface{}:
		t1Arr := t1.([]interface{})
		t2Arr, t2IsArr := t2.([]interface{})
		if !t2IsArr || len(t1Arr) != len(t2Arr) {
			return false
		}
		T1Check:
			for _, v := range t1.([]interface{}) {
				for _, t2V := range t2Arr {
					if isMatch(v, t2V, cs) {
						continue T1Check
					}
				}
				return false
			}
		return true
	case string:
		if cs == CaseSensitive {
			return t1 == t2
		}
		t1Str := t1.(string)
		t2Str, t2IsStr := t2.(string)
		if !t2IsStr {
			return false
		}
		t1Str, t2Str = strings.ToLower(t1Str), strings.ToLower(t2Str)
		return t1Str == t2Str
	default:
		return t1 == t2
	}
}

// isValid validates that the instance of the Node type itself is valid, not that the value being sought
// is present within the node.
func (n *Node) isValid() error {
	switch {
	case n.Value == nil && strings.TrimSpace(n.Key) == "":
		return errors.New("Nodes must have a key or a value, or both")
	case n.MaxDepth < 0:
		return errors.New("MaxDepth must be a positive integer")
	case n.sourceData == nil:
		return errors.New("No source data was found. Be sure to initialize ratnest with your data by using ratsnest.New(...)")
	}
	return nil
}