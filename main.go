package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/AstroKiR/kstruct/xml2json"
)

func ParseXML(filePath string) {

	var resultStructure xml2json.Node
	xml2json.Stack = nil

	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			currentNode := xml2json.Node{Name: se.Name.Local}
			xml2json.Push(&currentNode)
			fmt.Printf("push\t: %v \n", xml2json.Stack)
		case xml.EndElement:
			fmt.Printf("pop\t: %v \n", xml2json.Stack)
			xml2json.Pop(xml2json.Stack)
		}
	}
	fmt.Printf("%+v\n", resultStructure)
}

func main() {
	if len(os.Args) == 1 {
		panic("Provide the file path!")
	}
	ParseXML(os.Args[1])
}
