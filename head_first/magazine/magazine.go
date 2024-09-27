package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

func (s Subscriber) GetSubscriber() {
	fmt.Printf("Subscruber %s has rate %f active %t\n", s.Name, s.Rate, s.Active)
}

func PrintInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
	fmt.Println(s.City, s.Street, s.PostalCode, s.State)
}
func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func ApplyDiscount(s *Subscriber) {
	s.Rate = 4.99
}

type Employee struct {
	Name   string
	Salary float64
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
