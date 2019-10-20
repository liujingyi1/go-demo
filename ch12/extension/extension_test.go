package extension

import (
	"fmt"
	"testing"
)

//Pet
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

func (p *Pet) SayHello(name string) {
	fmt.Println("hello ", name)
}

//Dog
type Dog struct {
	age int
	Pet
}

func (d *Dog) Speak() {
	d.Pet.Speak()
	fmt.Print("Wang!")
}

func (d *Dog) SpeakTo(name string) {
	d.Speak()
	fmt.Println(name)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("jingyi")
	dog.SayHello("jingyi")
	dog.Speak()

	dog.Speak()
}
