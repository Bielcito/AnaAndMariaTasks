package main

import "fmt"
import "sync"
import "time"

type Person struct{
	name string // Nome da pessoa
	accessedObjects []*Object; // Lista de objetos que já foram acessados
}

func (p *Person) runTaskList (t *TaskList, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	fmt.Print(p.name, " está se aprontando\n")

	for {
		// Tranca a lista:
		if(DEBUGMODE) {	fmt.Print(p.name, " quer a lista\n")	}
		t.Lock()
		if(DEBUGMODE) {	fmt.Print(p.name, " acessou a lista\n")	}

		// Se não houver mais tarefas a serem concluídas...
		if(len(t.tasks) == 0) {

			// Destranca a lista:
			if(DEBUGMODE) {	fmt.Print(p.name, " quer liberar a lista\n") }
			t.Unlock()
			if(DEBUGMODE) {	fmt.Print(p.name, " liberou a lista.\n") }

			duration := time.Now().Sub(start)
			duration = duration / 100000000

			fmt.Print(p.name, " demorou ", int64(duration), " segundos para se preparar\n")

			// Termina a execução desta thread:
			break
		}

		// Tenta escolher uma tarefa para ser acessada...
		taskNumber := p.chooseTask(t)
		// Caso não consiga, DESTRANCA_LISTA
		if(taskNumber == -1) {
			t.Unlock();
			continue;
		}
		// Caso consiga, TRANCA OBJETO E DESTRANCA_LISTA
		choicedTask := &t.tasks[taskNumber]
		if(DEBUGMODE) { fmt.Print("Tarefa escolhida para " + p.name + " foi " + choicedTask.name + "\n") }

		if(DEBUGMODE) {	fmt.Print(p.name, " quer liberar a lista\n") }
		t.Unlock()
		if(DEBUGMODE) {	fmt.Print(p.name, " liberou a lista\n") }

		// Escolhe algum objeto para ser acessado, TRANCA_OBJETO
		objectNumber := random(0, len(t.tasks[taskNumber].objects)-1)
		choicedObject := &choicedTask.objects[objectNumber]
		if(DEBUGMODE) {	fmt.Print(p.name, " quer o objeto ", choicedObject.name, "\n") }
		choicedObject.Lock()
		if(DEBUGMODE) {	fmt.Print(p.name, " acessou o objeto ", choicedObject.name, "\n") }

		if(DEBUGMODE) {	fmt.Print(p.name, " quer deletar o objeto ", choicedObject.name, "\n") }
		choicedTask.delete(objectNumber)
		choicedTask.Unlock()
		p.randomWaitTime(500,1000)
		if(DEBUGMODE) {	fmt.Print(p.name, " deletou o objeto ", choicedObject.name, "\n") }
		if(DEBUGMODE) {	fmt.Print(p.name, " quer liberar a tarefa ", choicedTask.name, "\n") }
		if(DEBUGMODE) {	fmt.Print(p.name, " liberou a tarefa ", choicedTask.name, "\n") }

		/*// Caso o objeto já estiver sido completo, e DESTRANCA_TAREFA
		if(choicedObject.completed) {
			// Caso deleteOnAccess = true, deleta o objeto
			if(choicedTask.deleteOnAccess) {
				choicedTask.delete(objectNumber)
			}
			// Caso concludeOnAccess = true, verifica se quem concluiu foi esta pessoa, se sim, procura outra tarefa.
			if(choicedTask.concludeOnAccess) {

			}

			choicedTask.Unlock()
		// Caso não, destranca tarefa, diz que o objeto foi completo e destranca o objeto, depois continua para a próxima iteração! (ALTERAR DEPOIS PARA QUE DEMORE MAIS TEMPO PARA COMPLETAR O OBJETO)
		} else {
			choicedTask.Unlock()
			choicedObject.completed = true
			// Caso seja necessário deletar o objeto, deleta-o:
			if(choicedTask.deleteOnAccess) {
				t.delete(taskNumber)
			}
			choicedObject.Unlock()
			fmt.Print("Completando o objeto " + choicedObject.name + "\n")
			time.Sleep(time.Duration(random(300, 600)) * time.Millisecond)
			continue
		}*/
	}
}

func (p *Person) runAlarm(alarm *Alarm, wg *sync.WaitGroup, wg2 *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()

	alarm.mutex.Lock()

	if(!alarm.turnedOn)	{
		alarm.turnedOn = true;
		wg2.Add(1)
		go alarm.run(wg2)
		if(DEBUGMODE) { fmt.Print(p.name, " ligou o alarme\n") }
		fmt.Print("O alarme foi acionado\n")
	}

	fmt.Print(p.name, " começou a calçar o tênis\n")

	alarm.mutex.Unlock()

	p.randomWaitTime(3000, 4500)

	duration := time.Now().Sub(start)
	duration = duration / 100000000

	fmt.Print(p.name, " demorou ", int64(duration), " segundos para calçar seus tênis\n")
}

func (p *Person) completeObject(o Object) {
	//p.accessedObjects = append(p.accessedObjects, o)
}

// Usa o método do embaralhamento de cartas para pegar tarefas randômicas sem repetí-las:
func (p *Person) chooseTask(t *TaskList) int {
	if(DEBUGMODE){ fmt.Print("Escolhendo tarefa para " + p.name + ".\n") }
	var listSize int
	var numbers []int
	var choicedTask *Task
	var taskNumber int

	// Finalmente escolhe a tarefa percorrendo o array e TRANCA TAREFA
	for {

		listSize = len(t.tasks)

		if(listSize == 0) {
			return -1;
		}

		numbers = p.shuffle(listSize)
		taskNumber = numbers[0]
		choicedTask = &t.tasks[taskNumber]
		

		if(DEBUGMODE) {	fmt.Print(p.name, " quer a tarefa ", choicedTask.name, ".\n")	}
		choicedTask.Lock()
		if(DEBUGMODE) {	fmt.Print(p.name, " acessou a tarefa ", choicedTask.name, ".\n")	}

		// Caso a tarefa atual não tenha mais objetos, deleta ela e pega a próxima
		if(len(choicedTask.objects) == 0) {
			t.delete(taskNumber)
			continue;
		}

		return taskNumber;

		// Caso concludeOnAccess = true...
		/*if(choicedTask.concludeOnAccess) {
			// Verifica se esta pessoa já completou a tarefa da iteração:
			for j := 0; j < len(choicedTask.completed); j++ {
				// Caso sim...
				if(choicedTask.completed[j] == p) {
					// Caso esta tarefa já tenha sido completa por todo mundo, deleta ela
					if(len(choicedTask.completed) == t.PersonsDoingThisTask) {
						t.delete(taskNumber)
					}

					break;
				}
			}

			// Caso não, retorna essa tarefa:
			return taskNumber;

		} else {

		}*/
	}
}

func (p *Person) shuffle (size int) []int {

	numbers := make([]int, size) // lista de inteiros do tamanho do número de tarefas

	// Inicializa os valores do slice
	for i := 0; i < size; i++ {
		numbers[i] = i;
	}

	// Embaralha os números
	for i := 0; i < size; i++ {
		aux := random(0, size-1)
		numbers[i], numbers[aux] = numbers[aux], numbers[i]
	}

	return numbers
}

func (p *Person) waitTime (delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func (p *Person) randomWaitTime (de int, ate int) {
	time.Sleep(time.Duration(random(de, ate)) * time.Millisecond)	
}