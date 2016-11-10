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
		Value interface{}
		// Key is either the key of the map for the node or the key that the desired value must
		// be within within the node.
		Key string
		// Case indicates whether case sensitivity should be required for a match to occur. Defaults
		// to case-sensitive.
		Case CaseSensitivity
		// MaxDepth is used to validate that a value appears no farther down the hierarchy of a node
		// than MaxDepth levels. 0 is infinite, and is the default. The depth of values whose keys
		// are directly inside the node is 1. MaxDepth should be thought of as applying to the depth
		// of the Value rather than the depth of the Key.
		MaxDepth int

		// childRequirements houses added-on nodes for validation of nodes within other nodes.
		childRequirements []*Node
		// childNodes houses Nodes that are are children of a node initialized with the New function.
		childNodes []*Node
		// sourceData points to the original data which was handed to the New initializer.
		sourceData map[string]interface{}
	}

	// CaseSensitivity is a placeholder for an enum case sensitivity value.
	CaseSensitivity int
)

const (
	// CaseSensitive is the default.
	CaseSensitive CaseSensitivity = iota
	// CaseInsensitive will not consider case when matching.
	CaseInsensitive
)

// addChildren adds the childNodes to an initialized ratsnest Node.
func (n *Node) addChildren() {
	switch n.Value.(type) {
	case map[string]interface{}:
		for k, v := range n.Value.(map[string]interface{}) {
			cn := &Node{
				Key:        k,
				Value:      v,
				sourceData: n.sourceData,
			}
			cn.addChildren()
			n.childNodes = append(n.childNodes, cn)
		}
	case []interface{}:
		for _, v := range n.Value.([]interface{}) {
			cn := &Node{
				Value:      v,
				sourceData: n.sourceData,
			}
			cn.addChildren()
			n.childNodes = append(n.childNodes, cn)
		}
	default:
		cn := &Node{
			Value:      n.Value,
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
	for _, cn := range n.childNodes {
		switch castedVal := cn.Value.(type) {
		case []interface{}:
			if _, reqNodeIsArr := reqNode.Value.([]interface{}); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		case []string:
			if _, reqNodeIsArr := reqNode.Value.([]string); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		case []int:
			if _, reqNodeIsArr := reqNode.Value.([]int); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		case []int64:
			if _, reqNodeIsArr := reqNode.Value.([]int64); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		case []int32:
			if _, reqNodeIsArr := reqNode.Value.([]int32); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		case []float64:
			if _, reqNodeIsArr := reqNode.Value.([]float64); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		case []float32:
			if _, reqNodeIsArr := reqNode.Value.([]float32); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
					if checkKeyVal(cn.Key, reqNode.Key, v, reqNode.Value, reqNode.Case) {
						return cn
					}
				}
			}
		case []bool:
			if _, reqNodeIsArr := reqNode.Value.([]bool); reqNodeIsArr {
				if checkKeyVal(cn.Key, reqNode.Key, cn.Value, reqNode.Value, cn.Case) {
					return cn
				}
			} else {
				for _, v := range castedVal {
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
		if reqNode.MaxDepth < 1 || depth+1 <= reqNode.MaxDepth {
			if fn := cn.checkChildren(reqNode, depth + 1); fn != nil {
				return fn
			}
		}
	}
	return nil
}

// checkKeyVal simply checks whether a key and a value match.
func checkKeyVal(key, expKey string, val, expVal interface{}, cs CaseSensitivity) bool {
	switch expKey {
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
	switch casted := t1.(type) {
	case map[string]interface{}:
		t2Map, t2IsMap := t2.(map[string]interface{})
		if !t2IsMap || len(casted) != len(t2Map) {
			return false
		}
		for k, v := range casted {
			if t2Val, t2OK := t2Map[k]; !t2OK || !isMatch(v, t2Val, cs) {
				return false
			}
		}
		return true
	case []interface{}:
		t2Arr, t2IsArr := t2.([]interface{})
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
	T1IfaceArr:
		for _, v := range casted {
			for _, t2V := range t2Arr {
				if isMatch(v, t2V, cs) {
					continue T1IfaceArr
				}
			}
			return false
		}
		return true
	case []string:
		t2Arr, t2IsArr := t2.([]string)
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
	T1StringArr:
		for _, v := range casted {
			for _, t2V := range t2Arr {
				if isMatch(v, t2V, cs) {
					continue T1StringArr
				}
			}
			return false
		}
		return true
	case []int:
		t2Arr, t2IsArr := t2.([]int)
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
	T1IntArr:
		for _, v := range casted {
			for _, t2V := range t2Arr {
				if v == t2V {
					continue T1IntArr
				}
			}
			return false
		}
		return true
	case []int64:
		t2Arr, t2IsArr := t2.([]int64)
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
	T1Int64Arr:
		for _, v := range casted {
			for _, t2V := range t2Arr {
				if v == t2V {
					continue T1Int64Arr
				}
			}
			return false
		}
		return true
	case []int32:
		t2Arr, t2IsArr := t2.([]int32)
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
	T1Int32Arr:
		for _, v := range casted {
			for _, t2V := range t2Arr {
				if v == t2V {
					continue T1Int32Arr
				}
			}
			return false
		}
		return true
	case []float64:
		t2Arr, t2IsArr := t2.([]float64)
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
	T1Float64Arr:
		for _, v := range casted {
			for _, t2V := range t2Arr {
				if v == t2V {
					continue T1Float64Arr
				}
			}
			return false
		}
		return true
	case []float32:
		t2Arr, t2IsArr := t2.([]float32)
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
	T1Float32Arr:
		for _, v := range casted {
			for _, t2V := range t2Arr {
				if v == t2V {
					continue T1Float32Arr
				}
			}
			return false
		}
		return true
	case []bool:
		t2Arr, t2IsArr := t2.([]bool)
		if !t2IsArr || len(casted) != len(t2Arr) {
			return false
		}
		trues, expTrues := 0, 0
		falses, expFalses := 0, 0
		for _, t1V := range casted {
			if t1V {
				expTrues++
			} else {
				expFalses++
			}
		}
		for _, t2V := range t2Arr {
			if t2V {
				trues++
			} else {
				falses++
			}
		}
		return expTrues == trues && expFalses == falses
	case string:
		if cs == CaseSensitive {
			return t1 == t2
		}
		t2Str, t2IsStr := t2.(string)
		if !t2IsStr {
			return false
		}
		return strings.EqualFold(casted, t2Str)
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
