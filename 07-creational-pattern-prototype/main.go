package main

import (
	"fmt"
	"golang/pattern/07-creational-pattern-prototype/prototype"
)

func main() {
	file1 := &prototype.File{Name: "File 1"}
	file2 := &prototype.File{Name: "File 2"}
	file3 := &prototype.File{Name: "File 3"}

	folder1 := &prototype.Folder{
		Childrens: []prototype.INode{file1},
		Name:      "Folder 1",
	}

	folder2 := &prototype.Folder{
		Childrens: []prototype.INode{folder1, file2, file3},
		Name:      "Folder 2",
	}

	fmt.Println("\n Printing for folder 2")
	folder2.Print("  ")
	cloneFolder := folder2.Clone()
	fmt.Println("\n Printing for clone folder 2")
	cloneFolder.Print("  ")
}
