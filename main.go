package main

import (
	"Agnerft/github.com/loja/internal/adapter/conexao"
	"Agnerft/github.com/loja/internal/ports/conexaoInterface"
	"fmt"
)

func main() {

	var con conexaoInterface.Conexao

	con = &conexao.ConexaoHttp{}

	fmt.Println(con.BuscaDeputado())

}
