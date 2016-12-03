package main

import "sync"
import "fmt"
import "time"

type Alarm struct {
	mutex sync.Mutex
	turnedOn bool
}

func (a *Alarm) run (wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Print("O alarme está em contagem regressiva\n")

	a.waitTime(6000)

	fmt.Print("O alarme foi ativado\n")
}

func (a *Alarm) waitTime (delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
}