package main

import "sync"

type Object struct {
	name string
	mutex sync.Mutex
	completed bool
}

func (o *Object) Lock () {
	o.mutex.Lock()
}

func (o *Object) Unlock () {
	o.mutex.Unlock()
}