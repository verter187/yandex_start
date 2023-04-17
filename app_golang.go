package main

import (
	"fmt"
	"time"
)

// Аналог конструктора для создания персоны
func NewPerson(name, email string, dobYear, dobMonth, dobDay int) Person {
	return Person{
		Name:        name,
		Email:       email,
		dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
	}
}

type Person struct {
	Name        string
	Email       string
	dateOfBirth time.Time
}

func main() {
	man := Person{
		Name:        "John",
		Email:       "John@example.com",
		dateOfBirth: time.Now(),
	}
	fmt.Printf("Man %#v", man)
	fmt.Println(man.Name)
	fmt.Println(man.Email)
	fmt.Println(man.dateOfBirth)

	// У каждого поля структуры может быть набор аннотаций, которые называются тегами (tags):
	type GetUserRequest struct {
		UserId    string `json:"user_id" yaml: "user_id" format:"uuid" example:"2e263a90-b74b-11eb-8529-0242ac130003"`
		IsDeleted *bool  `json:"is_deleted,omitempty" yaml:"is_deleted"`
	}
}
