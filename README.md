# Rat's Nest
Simple validation for deeply-nested, arbitrary maps in Golang

## The Need

Let's try working with some arbitrary JSON:

```json
{
	"ID": 420,
	"name": "Maduro",
	"manufacturer": "Cigar City",
	"quantity": 12,
	"values": {
		"color": {
			"family": "brown",
			"depth": "dark"
		},
		"attributes": {
			"easeOfDescription": 0.1,
			"classification": {
				"ale"
			},
			"location": {
				"US": [
					"Gulf Coast",
					"Tampa",
					"FL"
				]
			}
		}
	}
}

```

We'll play out a scenario that you're working with data for which the format is not completely known to you, but you do
know that your application is looking specifically for a `depth` of `dark` in the `Gulf Coast` (specifically `Tampa`)
of the `US` where the `manufacturer` is `Cigar City`.

At this point, you're probably just unmarshaling the JSON into a `map[string]interface{}`. Now it's time to look for
all the values that you desire.

## The Solution

This package aims to help verify that a JSON resource (or any other resource that can be unmarshaled to a 
`map[string]interface{}`) meets your presence criteria. In other words, it will help you loop and walk
through all of those endless arbitrary objects.

### Usage

An [example application](https://github.com/powerchordinc/ratsnest/blob/master/example/main.go) can be found in the [`example` directory](https://github.com/powerchordinc/ratsnest/tree/master/example).

Initialize your Rat's Nest with your `map[string]interface{}` data:

```go
data := map[string]interface{}{
	"foo": "bar",
	"baz": map[string]interface{}{
		"bat": []interface{}{
			42,
			12.345,
			false,
		},
	},
}

root, err := ratsnest.New(data)
if err != nil {
	// an issue with data prevented initialization
}
```

After obtaining the root "Node" of your data, you can begin to `Require` other Nodes:

```
bazNode, err := root.Require(ratsnest.Node{
	Key: "baz",
	Value: 42,
})
if err != nil {
	// Node not found or invalid Node passed to Require()
}
```

You can then add Nodes requirements onto the newly-obtained nodes or continue adding requirement Nodes onto the root node.

## Comparison of Maps and Slices

Where maps are concerned, you can ask for just a `Key`, a `Key` and `Value`, or the entire map as `Value: map[string]interface{}{...}`. Order of appearance does not matter. Case [in]sensitivity applies to both keys and values of the maps themselves (should the values be strings).

For slices, you can ask for just one value within the slice, you can declare `Value` to be a `[]interface{}`, which will ensure the lengths match and all elements are present (regardless of order), or you can declare a `Key` and a `Value` (for searching `map[string][]interface{}`).

## Testing and Benchmarks

The package has 100% unit test coverage with Ginkgo tests. There are also benchmarks in `benchmark_test.go`. The benchmarks test maps with different complexities and depths. 

### Benchmark Results

These results are generated on a MacBook Pro 2.5GHz Core i7.

The `New` function initializing a root `Node` with maps from most to least complexity:

```
BenchmarkNewComplex-8                            	  200000	      7467 ns/op
BenchmarkNewComplexMin1-8                        	  200000	      6831 ns/op
BenchmarkNewComplexMin2-8                        	  200000	      5639 ns/op
BenchmarkNewComplexMin3-8                        	  200000	      5932 ns/op
BenchmarkNewComplexMin4-8                        	  300000	      4495 ns/op
BenchmarkNewComplexMin5-8                        	  500000	      3102 ns/op
BenchmarkNewComplexMin6-8                        	  500000	      2891 ns/op
BenchmarkNewComplexMin7-8                        	  500000	      2618 ns/op
BenchmarkNewComplexMin8-8                        	 1000000	      2261 ns/op
BenchmarkNewComplexMin9-8                        	  500000	      2003 ns/op
BenchmarkNewComplexMin10-8                       	 1000000	      1631 ns/op
BenchmarkNewComplexMin11-8                       	 2000000	       727 ns/op
BenchmarkNewComplexMin12-8                       	 3000000	       428 ns/op
```

The `Require` function, requiring different things from the most complex data defined:

```
BenchmarkRequireComplex_1Depth_Key-8             	  300000	      4832 ns/op
BenchmarkRequireComplex_1Depth_Key_Value-8       	  300000	      4953 ns/op
BenchmarkRequireComplex_1Depth_Key_Value_Max-8   	 5000000	       353 ns/op
BenchmarkRequireComplex_2Depth_Key-8             	  300000	      4601 ns/op
BenchmarkRequireComplex_2Depth_Key_Value-8       	  300000	      4559 ns/op
BenchmarkRequireComplex_2Depth_Key_Value_Max-8   	 1000000	      2095 ns/op
BenchmarkRequireComplex_3Depth_Key-8             	 1000000	      1500 ns/op
BenchmarkRequireComplex_3Depth_Key_Value-8       	 1000000	      1443 ns/op
BenchmarkRequireComplex_3Depth_Key_Value_Max-8   	 1000000	      1484 ns/op
BenchmarkRequireComplex_Array_Value-8            	  500000	      2993 ns/op
```

Finally, since non-existent `Keys`/`Values` result in the longest `Require` times, benchmarks for requiring non-existent nodes from each of 12 maps with varying complexity:

```
BenchmarkRequireComplex_NonExist_Key-8           	  300000	      4916 ns/op
BenchmarkRequireComplex_NonExist_Key_Value-8     	  300000	      4980 ns/op
BenchmarkRequireComplex_NonExist_Key_Value_Max1-8	 5000000	       225 ns/op
BenchmarkRequireComplex_NonExist_Key_Value_Max2-8	  500000	      2588 ns/op
BenchmarkRequireComplex_NonExist_Key_Value_Max3-8	  300000	      5170 ns/op
BenchmarkRequireComplexMin1_NonExist_Key-8       	  300000	      5479 ns/op
BenchmarkRequireComplexMin1_NonExist_Key_Value-8 	  300000	      5703 ns/op
BenchmarkRequireComplexMin2_NonExist_Key-8       	  300000	      5094 ns/op
BenchmarkRequireComplexMin2_NonExist_Key_Value-8 	  300000	      5333 ns/op
BenchmarkRequireComplexMin3_NonExist_Key-8       	  200000	      5002 ns/op
BenchmarkRequireComplexMin3_NonExist_Key_Value-8 	  300000	      4551 ns/op
BenchmarkRequireComplexMin4_NonExist_Key-8       	  300000	      4795 ns/op
BenchmarkRequireComplexMin4_NonExist_Key_Value-8 	  300000	      4400 ns/op
BenchmarkRequireComplexMin5_NonExist_Key-8       	  500000	      3594 ns/op
BenchmarkRequireComplexMin5_NonExist_Key_Value-8 	  300000	      3973 ns/op
BenchmarkRequireComplexMin6_NonExist_Key-8       	  500000	      2986 ns/op
BenchmarkRequireComplexMin6_NonExist_Key_Value-8 	  500000	      3113 ns/op
BenchmarkRequireComplexMin7_NonExist_Key-8       	 1000000	      2401 ns/op
BenchmarkRequireComplexMin7_NonExist_Key_Value-8 	  500000	      2819 ns/op
BenchmarkRequireComplexMin8_NonExist_Key-8       	 1000000	      1849 ns/op
BenchmarkRequireComplexMin8_NonExist_Key_Value-8 	 1000000	      1835 ns/op
BenchmarkRequireComplexMin9_NonExist_Key-8       	 1000000	      1246 ns/op
BenchmarkRequireComplexMin9_NonExist_Key_Value-8 	 1000000	      1384 ns/op
BenchmarkRequireComplexMin10_NonExist_Key-8      	 2000000	       520 ns/op
BenchmarkRequireComplexMin10_NonExist_Key_Value-8	 3000000	       556 ns/op
BenchmarkRequireComplexMin11_NonExist_Key-8      	 5000000	       410 ns/op
BenchmarkRequireComplexMin11_NonExist_Key_Value-8	 5000000	       392 ns/op
BenchmarkRequireComplexMin12_NonExist_Key-8      	 5000000	       380 ns/op
BenchmarkRequireComplexMin12_NonExist_Key_Value-8	 5000000	       334 ns/op
```