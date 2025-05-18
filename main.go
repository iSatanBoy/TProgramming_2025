package main

import "fmt"

type Character struct {
	Class string
	Level int
	HP    int
}

func New(class string, level int, hp int) Character {
	return Character{
		Class: class,
		Level: level,
		HP:    hp,
	}
}

func (c Character) Info() {
	fmt.Println("Класс:", c.Class, "Уровень:", c.Level, "HP:", c.HP)
}

func (c *Character) LevelUp() {
	c.Level++
	c.HP += 10
	fmt.Println("Новый уровень!")
}

func (c *Character) Damage(damage int) {
	c.HP -= damage
	if c.HP <= 0 {
		fmt.Println("Ваш персонраж погиб")
	} else {
		fmt.Println("Персонаж получил", damage, "урона.")
	}
}

func main() {
	hero := New("Маг", 1, 100)
	hero.Info()
	hero.LevelUp()
	hero.Damage(10)
	hero.Info()
}
