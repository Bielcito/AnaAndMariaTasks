package main

import "time"
import "math/rand"

// Função que gera números aleatórios
func random(de int, ate int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ate - de + 1) + de
}