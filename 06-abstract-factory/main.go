package main

import (
	"fmt"
	"golang/pattern/06-abstract-factory/abstractfactory"
)

func main() {
	adidasFactory := abstractfactory.GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)

	fmt.Println()

	nikeFactory := abstractfactory.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
}

func printShoeDetails(s abstractfactory.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShortDetails(s abstractfactory.IShort) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
