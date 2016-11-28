Para o projeto em questão, levei em conta que:

- Há apenas 2 óculos solares, e podem ser pegos separadamente ao mesmo tempo.
- Há apenas 1 protetor solar, e só podem ser usados um por vez.
- Há 8 janelas na casa. Elas não podem ser fechadas ao mesmo tempo.
- O mesmo vale para as portas, mas são apenas 4.
- Há apenas uma chave para a casa, e dois celulares. Como cada uma tem o seu próprio celular, nunca vai acontecer das duas pegarem o mesmo.

Portanto...
Cada tarefa possui uma lista de objetos, e características próprias.
As características são definidas por variáveis boleanas internas, que são:
	- concludeOnAccess: a tarefa é concluída assim que alguém conseguir acessar qualquer objeto.
	- deleteOnAccess: o objeto acessado é deletado assim que alguém conseguir acessá-lo.
	- accessEverything: a tarefa deve ser acessada até que não sobre mais objetos na lista, neste caso o objeto também será deletado após o acesso.

A tarefa 1 {concludeOnAccess : deleteOnAccess}:
	- 2 óculos;
	- Quando acessado, o óculos é removido por quem o acessou; e
	- Tarefa concluída quando alguém conseguir acessar algum óculos.
A tarefa 2 {concludeOnAccess}:
	- 1 protetor solar; e
	- Concluído quando alguém conseguir acessar o protetor.
A tarefa 3 {accessEverything}:
	- 8 janelas;
	- Quando acessada, a janela é removida por quem a acessou.
	- A tarefa é concluída quando todas as janelas forem removidas.
A tarefa 4 {accessEverything}:
	- 4 portas;
	- Quando acessada, a porta é removida por quem a acessou; e
	- A tarefa é concluída quando todas as portas forem removidas.
A tarefa 5 {concludeOnAccess}:
	- 1 chave;
	- O ato de pegar o celular é feito internamente na classe pessoa, pois como cada uma tem seu próprio celular, eles devem ser visíveis apenas para quem o possui; e
	- A tarefa é concluída quando a chave for pega por alguém.

Algoritmo para Pessoa:

Caso ainda haja alguma tarefa para ser concluída...

	1) Acessar alguma tarefa randômica, identificando suas características.
	2) Realiza sua parte na tarefa, caso for necessário deletar algum objeto, primeiro tranca-se o mutex daquele objeto, depois verifica se ele ainda não foi deletado. Caso não, deleta-o. Caso sim, volta para o passo 1.
		Como realizar:
			- Busca algum objeto randômico, locka o mutex, trabalha nele, e depois destrava o mutex.
	3) Caso a tarefa seja concluída neste processo, tranca o mutex da tarefa, verifica se ele já não foi deletado. Caso não, deleta-o. Caso sim, volta para o passo 1.