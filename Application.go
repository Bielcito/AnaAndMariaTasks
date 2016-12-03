package main

import "sync"
import "fmt"

type Application struct {
	tl TaskList // Lista de Tasks
	persons []Person // Lista de pessoas
}

func (a *Application) initializeTaskList () {
	a.tl.append(
		Task{
			name: "Pegar óculos de sol", 
			objects: []Object{
				{name: "Óculos de sol"}, 
				{name: "Óculos de sol"},
			},
			concludeOnAccess: true,
			deleteOnAccess: true,
		},
	)
	a.tl.append(
		Task{
			name: "Pegar protetor solar", 
			objects: []Object{
				{name: "Protetor solar"},
			},
			concludeOnAccess: true,
		},
	)
	a.tl.append(
		Task{
			name: "Verificar se as janelas estão fechadas", 
			objects: []Object{
				{name: "Janela"},
				{name: "Janela"},
				{name: "Janela"},
				{name: "Janela"},
				{name: "Janela"},
				{name: "Janela"},
				{name: "Janela"},
				{name: "Janela"},
			},
			deleteOnAccess: true,
		},
	)
	a.tl.append(
		Task{
			name: "Verificar se as portas estão fechadas", 
			objects: []Object{
				{name: "Porta"},
				{name: "Porta"},
				{name: "Porta"},
				{name: "Porta"},
			},
			deleteOnAccess: true,
		},
	)
	a.tl.append(
		Task{
			name: "Pegar as chaves e os celulares", 
			objects: []Object{
				{name: "Chaves"},
			},
			deleteOnAccess: true,
		},
	)
}

func (a *Application) initializePersons () {
	a.persons = append(a.persons, Person{name: "Ana"})
	a.persons = append(a.persons, Person{name: "Maria"})
}

func (a *Application) runTaskList () {

	var wg sync.WaitGroup

	for i := 0; i < len(a.persons); i++ {
		wg.Add(1)
		go a.persons[i].runTaskList(&a.tl, &wg)
	}

	wg.Wait()
}

func (a *Application) runAlarm () {
	alarm := Alarm{turnedOn: false}

	var wg, wg2 sync.WaitGroup

	for i := 0; i < len(a.persons); i++ {
		wg.Add(1)
		go a.persons[i].runAlarm(&alarm, &wg, &wg2)
	}

	wg.Wait()

	fmt.Print("Ana e Maria saíram e trancaram a porta\n")

	wg2.Wait()
}