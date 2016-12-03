package main

func main() {
	app := Application{}
	app.initializeTaskList()
	app.initializePersons()
	app.runTaskList()
	app.runAlarm()
}