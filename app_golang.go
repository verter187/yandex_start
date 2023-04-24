package main

import "fmt"

// Person — структура, описывающая человека.
type Person struct {
	Name string
	Year int
}

// NewPerson возвращает новую структуру Person.
func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

// String возвращает информацию о человеке.
func (p Person) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

// Print выводит информацию о человеке.
func (p Person) Print() {
	// вызовется метод String() для Person
	fmt.Println(p)
}

// Student описывает студента с использованием вложенной структуры Person. То есть структура Student описывает.
type Student struct {
	Person // вложенный объект Person
	Group  string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year), // Явно создаём структуру Person
		Group:  group,
	}
}

func (s Student) String() string {
	return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

func (s *Student) Debug() {
	// доступ к методам объекта Person
	s.Print()
	// или
	s.Person.Print()

	// доступ к полю 'Name' объекта Person
	s.Name = "Mark Smith"
	// или
	s.Person.Name = "Mark Smith"

	// вызовется метод String объекта Student
	fmt.Println(s)
	// вызовется метод String объекта Person
	fmt.Println(s.Person)
}

func main() {
	s := NewStudent("John Doe", 1980, "701")
	s.Debug()
	// s.Print()
	// // вызовется метод String() для Student
	// fmt.Println(s)
	// fmt.Println(s.Name, s.Year, s.Group)
}
