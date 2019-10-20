package sample

import (
	"fmt"
	"testing"
)

type IAnimal interface {
	Say(msg string) string
}

type Duck struct {
	Name string
}

type Dog struct {
	Name string
}

func (duck *Duck) Say(msg string) string {
	return fmt.Sprintf("duck duck! my name is %s,%s", duck.Name, msg)
}

func (dog *Dog) Say(msg string) string {
	return fmt.Sprintf("wang wang! my name is %s,%s", dog.Name, msg)
}

func TestAnimal(t *testing.T) {
	dog1 := Dog{Name: "Kitty"}
	dog2 := new(Dog)
	dog2.Name = "sdfsdf"

	t.Log(dog1.Say("hello"))

	var zoo = make([]IAnimal, 10, 20)

	animal := append(zoo, &dog1)
	t.Log(animal)
}
