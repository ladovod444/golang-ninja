package user

import "fmt"

type Age int

func (a Age) IsAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	Age    Age
	sex    string
	weight int
	height int
	//getAge func() int {return 222}
}

func NewUser(name string, Age Age, sex string, weight int, height int) User {
	return User{
		name, Age, sex, weight, height,
	}
}

func (u User) PrintUserInfo(name string) {
	u.name = name
	fmt.Println(u.name, u.Age, u.sex)
}

func (u User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
