// Package main is an example application for Rat's Nest (https://github.com/powerchordinc/ratsnest).
package main

import (
	"github.com/powerchordinc/ratsnest"
	"fmt"
)

var theData = []map[string]interface{}{
	{
		"firstName": "James",
		"lastName": "Beam",
		"age": 81,
	},
	{
		"name": "Maduro",
		"manufacturer": "Cigar City",
		"depth": "dark",
		"quantity": 12,
		"attributes": map[string]interface{}{
			"manufacturedIn": map[string]interface{}{
				"US": []string{
					"Gulf coast",
					"FL",
					"Tampa",
				},
			},
		},
	},
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
		})
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
			Key: "US",
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

		fmt.Printf("'%v %v' appears to be what you are seeking!", thisD["manufacturer"], thisD["name"])
	}
}
