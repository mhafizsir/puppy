package puppy

import "github.com/mhafizsir/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}
