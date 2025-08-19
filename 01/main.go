package main

import "fmt"

type Human struct {
	Name    string
	Age     int
	Country string
	Height  float64
	Weight  float64
}

func (h Human) Hello() {
	fmt.Printf("Привет! Меня зовут %s, мне %d лет\n", h.Name, h.Age)
}

func (h Human) GetInfo() {
	fmt.Printf("Имя: %s, Возраст: %d, Страна: %s, Рост: %.1f см, Вес: %.1f кг\n",
		h.Name, h.Age, h.Country, h.Height, h.Weight)
}

func (h Human) IsAdult() bool {
	return h.Age >= 18
}

type Action struct {
	Human
	Profession string
	Experience int
}

func (a Action) Work() {
	fmt.Printf("%s работает %s с опытом %d лет\n", a.Name, a.Profession, a.Experience)
}

func (a Action) GetFullInfo() {
	a.GetInfo()
	fmt.Printf("возраст: %d, Опыт: %d лет\n", a.Age, a.Experience)
}

func main() {
	human := Human{
		Name:    "Арсен",
		Age:     18,
		Country: "Россия",
		Height:  169.0,
		Weight:  169.0,
	}

	fmt.Println("1 человек:")
	human.Hello()
	human.GetInfo()
	fmt.Printf("Совершеннолетний: %t\n\n", human.IsAdult())
	action := Action{
		Human: Human{
			Name:    "Даниил",
			Age:     17,
			Country: "Россия",
			Height:  190.0,
			Weight:  90.0,
		},
		Profession: "Программист",
		Experience: 5,
	}

	fmt.Println("2 человек:")
	action.Hello()
	action.GetInfo()
	fmt.Printf("Совершеннолетний: %t\n", action.IsAdult())
	fmt.Printf("\n")
	action.Work()
	action.GetFullInfo()

}
