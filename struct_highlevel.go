package main

import "fmt"

type User struct {
ID int
Name string
Email string
}

type Admin struct {
User
IsAdmin bool
}

func main() {
user := User{ID: 1, Name: "Kathleen Booth", Email: "kathleenBooth@assembly.com"}
admin := Admin{User: user, IsAdmin: true}
fmt.Printf("User: %+v\n", user)
fmt.Printf("Admin: %+v\n", admin)
}