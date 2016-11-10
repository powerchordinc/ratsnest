package ratsnest

import (
	"strings"

	"testing"
)

var (
	data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"elit": []interface{}{
				"donec",
				"hendrerit",
				"turpis",
				map[string]interface{}{
					"vel": "sem",
				},
			},
			"gravida": true,
			"vestibulum": []interface{}{
				36954.02,
				true,
				72,
			},
			"dictum": []string{
				"mi",
				"eu",
				"ultrices",
				"imperdiet",
			},
			"fusce": []int{
				1,
				2,
				4,
				3,
			},
			"nonQuam": []int64{
				5,
				6,
				7,
				8,
			},
			"sedQuam": []int32{
				9,
				10,
				12,
				11,
			},
			"rutrum": []float64{
				1.123,
				2.123,
				4.123,
				3.123,
			},
			"quisque": []float32{
				5.123,
				6.123,
				8.123,
				7.123,
			},
			"tristique": []bool{
				false,
				true,
				true,
				false,
				true,
			},
		},
	}
	root *Node
	vn *Node
	n *Node
)

func reset() {
	n = &Node{
		Value: "ipsum",
		Key:   "lorem",
	}
	vn = &Node{
		Value: "foobar",
		Key:   "bazbat",
		sourceData: map[string]interface{}{
			"foo": "bar",
		},
	}
	root, _ = New(data)
}

func ensureValidationError(t *testing.T) {
	err := vn.isValid()
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestNew(t *testing.T) {
	defer reset()
	_, err := New(nil)
	if err == nil {
		t.Error("Expected an error to occur when called with a nil value, but none occurred")
	}
}

func TestIsValid_WithoutKeyOrValue(t *testing.T) {
	defer reset()
	vn.Key = ""
	vn.Value = nil
	ensureValidationError(t)
}

func TestIsValid_WithNegativeMaxDepth(t *testing.T) {
	defer reset()
	vn.MaxDepth = -1
	ensureValidationError(t)
}

func TestIsValid_WithoutSourceData(t *testing.T) {
	defer reset()
	vn.sourceData = nil
	ensureValidationError(t)
}

func TestRequire_ValidationWhenInvalid(t *testing.T) {
	defer reset()
	n.Key = ""
	n.Value = nil
	_, err := root.Require(*n)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestRequire_ValidationWhenValid(t *testing.T) {
	defer reset()
	_, err := root.Require(*n)
	if err != nil {
		t.Errorf("Was expecting no error, but got %s", err.Error())
	}
	if len(root.childRequirements) != 1 {
		t.Errorf("Was expecting childRequirements of the root to be 1, but was %d", len(root.childRequirements))
	}
}

func TestRequire_Satisfaction(t *testing.T) {
	defer reset()
	for _, cases := range allValidationCases {
		testSatisfactionCases(cases, t)
	}
}

func TestErrorStrings_NoCriteria(t *testing.T) {
	if (NoCriteriaError{}).Error() != "No requirements have been added to the root node." {
		t.Errorf("Expected 'No requirements have been added to the root node.', but got '%s'", (NoCriteriaError{}).Error() )
	}
}

func TestErrorStrings_UnfoundNode(t *testing.T) {
	if (NodeNotFoundError{}).Error() != "The required node was not found in the parent." {
		t.Errorf("Expected 'The required node was not found in the parent.', but got '%s'", (NodeNotFoundError{}).Error() )
	}
}

type (
	validationCase struct {
		toRegister          Node
		expErr              error
		noChildRequirements bool
	}

	validationCases []validationCase
)

func testSatisfactionCases(cases validationCases, t *testing.T) {
	pn := root
	for _, c := range cases {
		rn, err := pn.Require(c.toRegister)
		if c.expErr == nil {
			if err != nil {
				t.Errorf("Expected error to be nil, but got %s", err.Error())
			}
		} else {
			if err != c.expErr {
				t.Errorf("Expected error to match %T, but got %T", c.expErr, err)
			}
		}
		c.toRegister.Case = CaseInsensitive
		c.toRegister.Key = strings.ToUpper(c.toRegister.Key)
		valStr, valIsStr := c.toRegister.Value.(string)
		if valIsStr {
			c.toRegister.Value = strings.ToUpper(valStr)
		}
		rn, err = pn.Require(c.toRegister)
		if c.expErr == nil {
			if err != nil {
				t.Errorf("Expected error to be nil, but got %s", err.Error())
			}
		} else {
			if err != c.expErr {
				t.Errorf("Expected error to match %T, but got %T", c.expErr, err)
			}
		}
		if !c.noChildRequirements {
			pn = rn
		}
	}
}

var allValidationCases = []validationCases{
	{
		{
			toRegister: Node{Value: 592},
		},
	},
	{
		{
			toRegister: Node{Value: "adipiscing"},
		},
	},
	{
		{
			toRegister: Node{Value: "donec"},
		},
	},
	{
		{
			toRegister: Node{Value: "sem"},
		},
	},
	{
		{
			toRegister: Node{
				Value: 592,
				Key:   "dolor",
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: "adipiscing",
				Key:   "consectetur",
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: "sem",
				Key:   "vel",
			},
		},
	},
	{
		{
			toRegister: Node{
				Key: "dolor",
			},
		},
	},
	{
		{
			toRegister: Node{
				Key: "consectetur",
			},
		},
	},
	{
		{
			toRegister: Node{
				Key: "vel",
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: 592,
			},
			noChildRequirements: true,
		},
		{
			toRegister: Node{
				Value: "adipiscing",
			},
		},
	},
	{
		{
			toRegister: Node{
				Key: "amet",
			},
		},
		{
			toRegister: Node{
				Key: "elit",
			},
		},
		{
			toRegister: Node{
				Value: "donec",
			},
		},
	},
	{
		{
			toRegister: Node{Key: "amet"},
		},
		{
			toRegister: Node{Key: "elit"},
		},
		{
			toRegister: Node{
				Value: map[string]interface{}{
					"vel": "sem",
				},
			},
		},
		{
			toRegister: Node{Value: "sem"},
		},
	},
	{
		{
			toRegister: Node{
				Value: []interface{}{
					true,
					72,
					36954.02,
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []interface{}{
					true,
					72,
					"foobar",
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: []string{
					"mi",
					"eu",
					"ultrices",
					"imperdiet",
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: "imperdiet",
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []string{
					"mi",
					"eu",
					"ultrices",
					"foobar",
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: []int{
					4,
					3,
					2,
					1,
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: 3,
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []int{
					4,
					3,
					2,
					5,
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: []int64{
					8,
					7,
					6,
					5,
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: int64(6),
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []int64{
					7,
					6,
					5,
					3,
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: []int32{
					12,
					11,
					9,
					10,
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: int32(11),
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []int32{
					9,
					10,
					11,
					5,
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: []float64{
					1.123,
					2.123,
					3.123,
					4.123,
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: float64(2.123),
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []float64{
					1.123,
					2.123,
					3.123,
					5.123,
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: []float32{
					5.123,
					6.123,
					7.123,
					8.123,
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: float32(6.123),
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []float32{
					5.123,
					6.123,
					7.123,
					9.123,
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: []bool{
					true,
					true,
					true,
					false,
					false,
				},
			},
		},
	},
	{
		{
			toRegister: Node{
				Value: []bool{
					true,
					true,
					false,
					false,
					false,
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{
				Value: map[string]interface{}{
					"vel": "bar",
				},
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{Key: "amet"},
		},
		{
			toRegister: Node{Value: "hendrerit"},
		},
	},
	{
		{
			toRegister: Node{Key: "amet"},
		},
		{
			toRegister: Node{
				Value:    "sem",
				MaxDepth: 1,
			},
			expErr: NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{Key: "amet"},
		},
		{
			toRegister: Node{Key: "vel"},
		},
	},
	{
		{
			toRegister: Node{Key: "amet"},
		},
		{
			toRegister: Node{Key: "foobar"},
			expErr:     NodeNotFoundError{},
		},
	},
	{
		{
			toRegister: Node{Key: "foobar", Value: "adipiscing"},
			expErr: NodeNotFoundError{},
		},
	},
}