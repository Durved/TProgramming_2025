package main

func main() {

}

type Bunny struct {
	name   string
	age    int
	weight int
}

func (b *Bunny) SetAge(age int) {
	b.age = age
}

func (b *Bunny) GetAge() int {
	return b.age
}

func (b *Bunny) GetWeight() int {
	return b.weight
}
