package main

import "sync"

type TaskList struct {
	tasks []Task
	mutex sync.Mutex
}

// Deleta uma tarefa da lista de tarefas
func (tl *TaskList) delete (i int) {
	tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
}

// Adiciona uma tarefa na lista de tarefas
func (tl *TaskList) append (t Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) Lock () {
	if(DEBUGMODE) {
		tl.mutex.Lock()
	}
}

func (tl *TaskList) Unlock () {
	tl.mutex.Unlock()
}