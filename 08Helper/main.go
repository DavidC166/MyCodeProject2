package main

import "fmt"

func Checkin(name string) string{
	return "How are you"+ name + "!"
}
func main() {
	message := Checkin("David")
fmt.Println(message)
}
	