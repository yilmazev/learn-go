package main

import "fmt"

type User struct {
	id      int64
	name    string
	age     int
	details UserDetails
}

type UserDetails struct {
	address string
}

func main() {
	user1 := User{
		id:      1,
		name:    "Yılmaz Ev",
		age:     19,
		details: UserDetails{address: "İstanbul"},
	}

	changeName(&user1, "Mehmet Cengiz")
	fmt.Println(user1)
}

func changeName(user *User, newName string) {
	user.name = newName
}
