package main

import "fmt"

// Defining behavior interfaces
type Flyer interface {
	Fly() string
}

type Swimmer interface {
	Swim() string
}

// Implementing behavior structs
type CanFly struct{}

func (c CanFly) Fly() string {
	return "I'm flying!"
}

type CanSwim struct{}

func (c CanSwim) Swim() string {
	return "I'm swimming!"
}

// Defining Animal struct with embedded behaviors
type Animal struct {
	name string
	Flyer
	Swimmer
}

// Function to create an animal
func NewAnimal(name string, flyer Flyer, swimmer Swimmer) Animal {
	return Animal{
		name:    name,
		Flyer:   flyer,
		Swimmer: swimmer,
	}
}

// Main function to demonstrate the usage
func main() {
	duck := NewAnimal("Duck", CanFly{}, CanSwim{})
	penguin := NewAnimal("Penguin", nil, CanSwim{})

	fmt.Printf("%s: %s\n", duck.name, duck.Fly())
	fmt.Printf("%s: %s\n", duck.name, duck.Swim())

	fmt.Printf("%s: %s\n", penguin.name, penguin.Swim())
	if penguin.Flyer != nil {
		fmt.Printf("%s: %s\n", penguin.name, penguin.Fly())
	} else {
		fmt.Printf("%s can't fly.\n", penguin.name)
	}
}
