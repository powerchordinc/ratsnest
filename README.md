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

An [example application](https://github.com/powerchordinc/ratsnest/example) can be found in the [`example`](https://github.com/powerchordinc/ratsnest/example) directory

### Comparison of Maps and Slices

Where maps are concerned, you can ask for just a `Key`, a `Key` and `Value`, or the entire map as `Value: map[string]interface{}{...}`. Order of appearance does not matter. Case [in]sensitivity applies to both keys and values of the maps themselves (should the values be strings).

For slices, you can ask for just one value within the slice, you can declare `Value` to be a `[]interface{}`, which will ensure the lengths match and all elements are present (regardless of order), or you can declare a `Key` and a `Value` (for searching `map[string][]interface{}`).