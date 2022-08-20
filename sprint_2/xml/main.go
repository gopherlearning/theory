package main

import (
	"encoding/xml"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Person struct {
	// XMLName xml.Name `xml:"Person"`
	ID     int `xml:"id,attr"`
	Name   string
	Email  string
	Phones []string `xml:"Phones>Phone"`
}

type List struct {
	// XMLName xml.Name `xml:"List"`
	Persons []Person `xml:"Person"`
}

func main() {
	var v List
	data := `
    <List>
        <Person id="1">
            <Name>Carla Mitchel</Name>
            <Phones>
                <Phone>123-45-67</Phone>
                <Phone>890-12-34</Phone>
            </Phones>
        </Person>
        <Person id="2">
            <Name>Michael Smith</Name>
            <Email>msmith@example.com</Email>
        </Person>
    </List>
    `
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		logrus.Fatal(err)
	}
	for _, item := range v.Persons {
		fmt.Println(item.ID, item.Name, item.Email, item.Phones)
	}
}
