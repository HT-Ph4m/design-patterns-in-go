package main

import (
	"fmt"
	"golang/pattern/05-creational-pattern-builder/builder"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor : %d\n", normalHouse.GetFloor())

	director2 := builder.NewDirector(iglooBuilder)
	iglooHouse := director2.BuildHouse()

	fmt.Println()
	fmt.Printf("Normal House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor : %d\n", iglooHouse.GetFloor())
}
