package ratsnest

import (
	"testing"
)

var (
	node *Node
	e    error
)

var (
	complexData = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"maximus": map[string]interface{}{
				"integer": "vehicula",
				"nisl":    "in tempus",
			},
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

	complexMin1Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"maximus": map[string]interface{}{
				"integer": "vehicula",
				"nisl":    "in tempus",
			},
			"elit": []interface{}{
				"donec",
				"hendrerit",
				"turpis",
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

	complexMin2Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"maximus": map[string]interface{}{
				"integer": "vehicula",
				"nisl":    "in tempus",
			},
			"elit": []interface{}{
				"donec",
				"hendrerit",
				"turpis",
			},
			"gravida": true,
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

	complexMin3Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"maximus": map[string]interface{}{
				"integer": "vehicula",
				"nisl":    "in tempus",
			},
			"elit": []interface{}{
				"donec",
				"hendrerit",
				"turpis",
			},
			"gravida": true,
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

	complexMin4Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"maximus": map[string]interface{}{
				"integer": "vehicula",
				"nisl":    "in tempus",
			},
			"gravida": true,
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

	complexMin5Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"gravida":     true,
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

	complexMin6Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"gravida":     true,
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

	complexMin7Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"gravida":     true,
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

	complexMin8Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"gravida":     true,
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

	complexMin9Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"gravida":     true,
			"tristique": []bool{
				false,
				true,
				true,
				false,
				true,
			},
		},
	}

	complexMin10Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
		"amet": map[string]interface{}{
			"consectetur": "adipiscing",
			"gravida":     true,
		},
	}

	complexMin11Data = map[string]interface{}{
		"lorem": "ipsum",
		"dolor": 592,
	}

	complexMin12Data = map[string]interface{}{
		"lorem": "ipsum",
	}

	nodeComplex      *Node
	nodeComplexMin1  *Node
	nodeComplexMin2  *Node
	nodeComplexMin3  *Node
	nodeComplexMin4  *Node
	nodeComplexMin5  *Node
	nodeComplexMin6  *Node
	nodeComplexMin7  *Node
	nodeComplexMin8  *Node
	nodeComplexMin9  *Node
	nodeComplexMin10 *Node
	nodeComplexMin11 *Node
	nodeComplexMin12 *Node
)

func init() {
	nodeComplex, _ = New(complexData)
	nodeComplexMin1, _ = New(complexMin1Data)
	nodeComplexMin2, _ = New(complexMin2Data)
	nodeComplexMin3, _ = New(complexMin3Data)
	nodeComplexMin4, _ = New(complexMin4Data)
	nodeComplexMin5, _ = New(complexMin5Data)
	nodeComplexMin6, _ = New(complexMin6Data)
	nodeComplexMin7, _ = New(complexMin7Data)
	nodeComplexMin8, _ = New(complexMin8Data)
	nodeComplexMin9, _ = New(complexMin9Data)
	nodeComplexMin10, _ = New(complexMin10Data)
	nodeComplexMin11, _ = New(complexMin11Data)
	nodeComplexMin12, _ = New(complexMin12Data)
}

func benchNew(d map[string]interface{}, b *testing.B) {
	var (
		n   *Node
		err error
	)
	for i := 0; i < b.N; i++ {
		n, err = New(d)
	}
	node, e = n, err
}

func benchRequire(node *Node, reqNode Node, b *testing.B) {
	var (
		n   *Node
		err error
	)
	for i := 0; i < b.N; i++ {
		n, err = node.Require(reqNode)
	}
	node, e = n, err
}

func BenchmarkNewComplex(b *testing.B) {
	benchNew(complexData, b)
}

func BenchmarkNewComplexMin1(b *testing.B) {
	benchNew(complexMin1Data, b)
}

func BenchmarkNewComplexMin2(b *testing.B) {
	benchNew(complexMin2Data, b)
}

func BenchmarkNewComplexMin3(b *testing.B) {
	benchNew(complexMin3Data, b)
}

func BenchmarkNewComplexMin4(b *testing.B) {
	benchNew(complexMin4Data, b)
}

func BenchmarkNewComplexMin5(b *testing.B) {
	benchNew(complexMin5Data, b)
}

func BenchmarkNewComplexMin6(b *testing.B) {
	benchNew(complexMin6Data, b)
}

func BenchmarkNewComplexMin7(b *testing.B) {
	benchNew(complexMin7Data, b)
}

func BenchmarkNewComplexMin8(b *testing.B) {
	benchNew(complexMin8Data, b)
}

func BenchmarkNewComplexMin9(b *testing.B) {
	benchNew(complexMin9Data, b)
}

func BenchmarkNewComplexMin10(b *testing.B) {
	benchNew(complexMin10Data, b)
}

func BenchmarkNewComplexMin11(b *testing.B) {
	benchNew(complexMin11Data, b)
}

func BenchmarkNewComplexMin12(b *testing.B) {
	benchNew(complexMin12Data, b)
}

func BenchmarkRequireComplex_1Depth_Key(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key: "lorem",
	}, b)
}
func BenchmarkRequireComplex_1Depth_Key_Value(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:   "lorem",
		Value: "ipsum",
	}, b)
}
func BenchmarkRequireComplex_1Depth_Key_Value_Max(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:      "lorem",
		Value:    "ipsum",
		MaxDepth: 1,
	}, b)
}
func BenchmarkRequireComplex_2Depth_Key(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key: "consectetur",
	}, b)
}
func BenchmarkRequireComplex_2Depth_Key_Value(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:   "consectetur",
		Value: "adipiscing",
	}, b)
}
func BenchmarkRequireComplex_2Depth_Key_Value_Max(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:      "consectetur",
		Value:    "adipiscing",
		MaxDepth: 2,
	}, b)
}
func BenchmarkRequireComplex_3Depth_Key(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key: "nisl",
	}, b)
}
func BenchmarkRequireComplex_3Depth_Key_Value(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:   "nisl",
		Value: "in tempus",
	}, b)
}
func BenchmarkRequireComplex_3Depth_Key_Value_Max(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:      "nisl",
		Value:    "in tempus",
		MaxDepth: 3,
	}, b)
}
func BenchmarkRequireComplex_Array_Value(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Value: []int{1, 2, 3, 4},
	}, b)
}
func BenchmarkRequireComplex_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplex_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplex_NonExist_Key_Value_Max1(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:      "foobar",
		Value:    "bazbat",
		MaxDepth: 1,
	}, b)
}
func BenchmarkRequireComplex_NonExist_Key_Value_Max2(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:      "foobar",
		Value:    "bazbat",
		MaxDepth: 2,
	}, b)
}
func BenchmarkRequireComplex_NonExist_Key_Value_Max3(b *testing.B) {
	benchRequire(nodeComplex, Node{
		Key:      "foobar",
		Value:    "bazbat",
		MaxDepth: 3,
	}, b)
}

func BenchmarkRequireComplexMin1_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin1, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin1_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin1, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin2_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin2, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin2_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin2, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin3_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin3, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin3_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin3, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin4_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin4, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin4_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin4, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin5_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin5, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin5_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin5, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin6_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin6, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin6_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin6, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin7_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin7, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin7_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin7, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin8_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin8, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin8_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin8, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin9_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin9, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin9_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin9, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin10_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin10, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin10_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin10, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin11_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin11, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin11_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin11, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
func BenchmarkRequireComplexMin12_NonExist_Key(b *testing.B) {
	benchRequire(nodeComplexMin12, Node{
		Key: "foobar",
	}, b)
}
func BenchmarkRequireComplexMin12_NonExist_Key_Value(b *testing.B) {
	benchRequire(nodeComplexMin12, Node{
		Key:   "foobar",
		Value: "bazbat",
	}, b)
}
