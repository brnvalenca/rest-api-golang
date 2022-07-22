package entities

type UniqueEntityID = string


/* 
	Ele faz uma funcao de validacao do UUID aqui, mas nao entendi de onde ele puxa
	a função. Ele faz um import de 
	"errors"
	ERROR "go_clean_api/api/shared/constants/errors"
	v "go_clean_api/api/shared/validators"
	mas eu n consegui ter acesso a essas duas ultimas libs pra tentar entender oq ta 
	acontecendo. a lib errors é bultin do golang. 
	Ele diz que essa entidade define um id unico q vai ser usado pra casa usuario
	num contexto de banco de dados relacional ou nao. 
*/