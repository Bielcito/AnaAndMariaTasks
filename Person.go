package main

import "fmt"
import "sync"

type Person struct{
	name string
}

func (p *Person) run (t *TaskList, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print(p.name, " está se aprontando\n")

	for {
		// Tranca a lista:
		t.mutex.Lock()

		// Se não houver mais tarefas a serem concluídas...
		if(len(t.tasks) == 0) {

			// Destranca a lista:
			t.mutex.Unlock()

			fmt.Print(p.name, " demorou x segundos para se preparar.\n")

			// Termina a execução desta thread:
			break
		}

		// Escolha uma tarefa para ser acessada e TRANCA_TAREFA
		taskNumber := random(0, len(t.tasks)-1)
		choicedTask := &t.tasks[taskNumber]
		choicedTask.Lock()

		// Caso haja algum objeto dentro dela, DESTRANCA_LISTA
		if(len(choicedTask.objects) > 0) {
			t.Unlock()
		// Caso não, deleta a tarefa, DESTRANCA_LISTA e repete
		} else {
			t.delete(taskNumber)
			t.Unlock()
			continue
		}

		// Escolhe algum objeto para ser acessado, TRANCA_OBJETO
		objectNumber := random(0, len(t.tasks[taskNumber].objects)-1)
		choicedObject := &choicedTask.objects[objectNumber]
		choicedObject.Lock()

		// Caso o objeto já estiver sido completo, deleta ele e DESTRANCA_TAREFA
		if(choicedObject.completed) {
			choicedTask.delete(objectNumber)
			choicedTask.Unlock()
		// Caso não, destranca tarefa, diz que o objeto foi completo, e destranca o objeto, depois continua para a próxima iteração! (ALTERAR DEPOIS PARA QUE DEMORE MAIS TEMPO PARA COMPLETAR O OBJETO)
		} else {
			choicedTask.Unlock()
			choicedObject.completed = true
			choicedObject.Unlock()
			continue
		}

		//fmt.Print(p.name, " acessou o objeto \"", choicedObject.name , "\" na tarefa \"", choicedTask.name ,"\".\n")

		// Caso seja necessário deletar o objeto, deleta-o:
		/*if(choicedTask.deleteOnAccess) {
			t.delete(taskNumber)
		}*/
	}
}