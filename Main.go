package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                   string
	age                    uint16
	money                  int16
	avg_grandes, happiness float64
}

func (u User) gerAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money "+
		"equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.gerAllInfo())
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func blog_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Blog page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about", about_page)
	http.HandleFunc("/blog", blog_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//bob := User{name: "Bob", age: 25, money: -50, avg_grandes: 4.2, happiness: 0.8}

	handleRequest()
}
