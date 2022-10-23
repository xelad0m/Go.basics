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
	paris := Address{"France", "Paris"}
	emma := Person{Name: "Emma"}
	grace := Person{Name: "Grace"}
	alice := Person{
		Name:      "Alice",
		IsAwesome: true,
		Residence: &paris,
		Friends:   []*Person{&emma, &grace},
	}

	b, err := xml.MarshalIndent(alice, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	/*
		<person awesome="true">
		    <name>Alice</name>
		    <residence country="France" city="Paris"></residence>
		    <friends>
		        <person>
		            <name>Emma</name>
		            <friends></friends>
		        </person>
		        <person>
		            <name>Grace</name>
		            <friends></friends>
		        </person>
		    </friends>
		</person>
	*/
}
