package main

import "sync"

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
			accessEverything: true,
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
			accessEverything: true,
		},
	)
	a.tl.append(
		Task{
			name: "Pegar as chaves e os celulares", 
			objects: []Object{
				{name: "Chaves"},
			},
			concludeOnAccess: true,
		},
	)
}

func (a *Application) initializePersons () {
	a.persons = append(a.persons, Person{name: "Ana"})
	a.persons = append(a.persons, Person{name: "Maria"})
}

func (a *Application) run () {

	var wg sync.WaitGroup

	for i := 0; i < len(a.persons); i++ {
		wg.Add(1)
		go a.persons[i].run(&a.tl, &wg)
	}

	wg.Wait()
}