package main

import "fmt"

func main() {
	app := Application{}
	app.initializeTaskList()
	app.initializePersons()
	fmt.Println(app)
	app.run()
}

func printObjectName(object Object) {
	fmt.Println(object.name)
}