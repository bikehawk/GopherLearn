package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Record struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Parent   int       `json:"-"`
	Children []*Record `json:"children,omitempty"`
}

type Data []Record

func makeTree(data Data) []*Record {

	var nodes []*Record
	var topNodes []*Record
	var lookupList = make(map[int]*Record)

	for i, v := range data {
		lookupList[v.ID] = &data[i]
		nodes = append(nodes, &data[i])
		if v.Parent == 0 {
			topNodes = append(topNodes, &data[i])
		}
	}

	for _, v := range nodes {
		if v.Parent != 0 {
			lookupList[v.Parent].Children = append(lookupList[v.Parent].Children, v)
		}
	}
	return topNodes
}

func main() {
	data := Data{
		Record{1, "item1", 0, nil},
		Record{2, "item2", 1, nil},
		Record{3, "item3", 2, nil},
		Record{4, "item4", 7, nil},
		Record{5, "item5", 0, nil},
		Record{6, "item6", 2, nil},
		Record{7, "item7", 3, nil},
	}

	tree := makeTree(data)
	b, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
