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

```go
// ...

import "github.com/powerchordinc/ratsnest"

var theData = []map[string]interface{}{
	{
		"firstName": "James",
		"lastName": "Beam",
		"age": 81
	},
	{
		"name": "Maduro",
		"manufacturer": "Cigar City",
		"depth": "dark",
		"quantity": 12,
		"attributes": {
			"manufacturedIn": {
				"US": []string{
					"Gulf coast",
					"FL",
					"Tampa"
				}
			}
		}
	}
}

func main() {
	for _, thisD := range theData {
		root, err := ratsnest.New(thisD)
		if err != nil {
			panic(fmt.Errorf("Error creating a new root node: %v", err))
		}
		_, err = root.Require(ratsnest.Node{ // returns a ratsnest.Node, but we don't care in this case
			Value: "Cigar City",
			Key: "manufacturer", // defaults to any key
			MaxDepth: 1, // defaults to 0, which is infinite
		}
		if err != nil {
			 continue // node not found
		}
		_, err = root.Require(ratsnest.Node{
			Value: 12,
			Key: "quantity",
		})
		if err != nil {
			continue // node not found
		}
		usa, err := root.Require(ratsnest.Node{
			Value: "US", // we don't care what `Key` has the value
		})
		if err != nil {
			continue // node not found
		}
		_, err = usa.Require(ratsnest.Node{
			Value: "Tampa",
			Case: ratsnest.CaseInsensitive, // defaults to case-sensitive
			MaxDepth: 2,
		})
		if err != nil {
			continue // node not found
		}
		
		fmt.Printf("'%v %v' appears to be what you are seeking.", thisD["manufacturer"], thisD["name"])
	}
}
```

### Comparison of Maps and Slices

Where maps are concerned, you can ask for just a `Key`, a `Key` and `Value`, or the entire map as `Value: map[string]interface{}{...}`. Order of appearance does not matter. Case [in]sensitivity applies to both keys and values of the maps themselves (should the values be strings).

For slices, you can ask for just one value within the slice, you can declare `Value` to be a `[]interface{}`, which will ensure the lengths match and all elements are present (regardless of order), or you can declare a `Key` and a `Value` (for searching `map[string][]interface{}`).