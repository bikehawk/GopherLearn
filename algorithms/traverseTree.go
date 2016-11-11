package main

import "fmt"

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

func traverseTree(tree []*Record) {
	for i, v := range tree {
		fmt.Printf("(%d) %d: %s\n", i, v.ID, v.Name)
		traverseTree(v.Children)
	}
}

func main() {
	data := Data{
		Record{1, "1", 0, nil},
		Record{2, "1.1", 1, nil},
		Record{3, "1.1.1", 2, nil},
		Record{4, "1.1.1.1.1", 7, nil},
		Record{5, "2", 0, nil},
		Record{6, "1.1.2", 2, nil},
		Record{7, "1.1.1.1", 3, nil},
		Record{8, "2.1", 5, nil},
		Record{9, "2.2", 5, nil},
		Record{10, "2.3", 5, nil},
		Record{11, "1.2", 1, nil},
		Record{12, "1.3", 1, nil},
		Record{13, "1.1.3", 2, nil},
	}

	tree := makeTree(data)

	traverseTree(tree)
}
