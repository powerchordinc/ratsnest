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
				"US": {
					"FL": {
						"Gulf Coast": {
							"Tampa": {
								"Ybor City"
							}
						}
					}
				}
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
`map[string]interface{}`) meets your presence criteria. In other words, it will help you loop
through all of those endless arbitrary objects to get to your dark, Gulf Coast dreams.

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
		"attributes": {
			"manufacturedIn": {
				"US": {
					"Gulf coast",
					"FL",
					"Tampa"
				]
			}
		}
	}
}

func main() {
	for _, thisD := range theData {
		root := ratsnest.New(thisD)
		root.Require(ratsnest.Node{ // returns a ratsnest.Node, but we don't care in this case
			Value: "Cigar City",
			Key: "manufacturer", // defaults to any key
			MaxDepth: ratsnest.ROOT, // defaults to infinite depth
		}
		root.Require(ratsnest.Node{
			Value: 12,
			Key: "quantity" // I know, I know, but 6?!
		})
		usa := root.Require(ratsnest.Node{
			Value: "US",
			// we don't care what `Key` has the value. We own that particular two-letter combination.
			// default of `Depth: ratsnest.INFINITE` works here
		})
		usa.Require(ratsnest.Node{
			Value: "Tampa",
			Case: ratsnest.CASE_INSENSITIVE, // defaults to case-sensitive
			// unless you know of another "Tampa" in the "Gulf Coast" then we're safe with any key--unless--Louisiana...
			MaxDepth: 2,
		})
		
		if root.IsSatisfied {
			fmt.Printf("'%v %v' appears to be what you are seeking.", thisD["manufacturer"], thisD["name"])
		}
	}
}


// btw, I googled for "Tampa, Louisiana" and I think we're safe.
```