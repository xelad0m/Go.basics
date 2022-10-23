package main

import (
	"encoding/xml"
	"fmt"
)

type Address struct {
	Country string `xml:"country,attr"`
	City    string `xml:"city,attr"`
}

type Person struct {
	XMLName   xml.Name  `xml:"person"`
	Name      string    `xml:"name"`
	IsAwesome bool      `xml:"awesome,attr,omitempty"`
	Residence *Address  `xml:"residence"`
	Friends   []*Person `xml:"friends>person"`
}

func main() {
	src := `
		<person awesome="true">
		  <name>Alice</name>
		  <residence country="France" city="Paris"></residence>
		  <friends>
		    <person>
		      <name>Emma</name>
		    </person>
		    <person>
		      <name>Grace</name>
		    </person>
		  </friends>
		</person>`

	var alice Person
	err := xml.Unmarshal([]byte(src), &alice)
	if err != nil {
		panic(err)
	}

	fmt.Printf(alice.Name)
	if alice.IsAwesome {
		fmt.Printf(" ✓ awesome")
	}
	fmt.Println()
	fmt.Printf("from %s, %s\n", alice.Residence.City, alice.Residence.Country)
	fmt.Println("friends:")
	for _, person := range alice.Friends {
		fmt.Printf("- %s\n", person.Name)
	}
	// Alice ✓ awesome
	// from Paris, France
	// friends:
	// - Emma
	// - Grace
}
