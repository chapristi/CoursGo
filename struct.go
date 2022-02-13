package main

import (
	"fmt"
)
type Adress struct{
	street,city string
}
type User struct {
	Id int
	Name string
	Mail string
	Date Time
	addr Adress
}
func main() {
	fmt.Println("A little copying is better than a little dependency.")
	var u User
	u.Name = "Louis"
	u.addr.street = "Limoges"
}
