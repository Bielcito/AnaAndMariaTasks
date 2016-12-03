package main

import "sync"

type Task struct
{
	name string
	objects[] Object
	mutex sync.Mutex
	concludeOnAccess bool
	deleteOnAccess bool
	accessEverything bool
	completed []*Person
}

// Deleta um objeto da lista de objetos
func (t *Task) delete (i int) {
	t.objects = append(t.objects[:i], t.objects[i+1:]...)
}

// Adiciona um objeto na lista de objetos
func (t *Task) append (o Object) {
	t.objects = append(t.objects, o)
}

func (t *Task) Lock () {
	t.mutex.Lock()
}

func (t *Task) Unlock () {
	t.mutex.Unlock()
}