package models

import "loja/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := db.ConectaBd()

	selectProdutos, err := db.Query("select * from produtos order by id desc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		err = selectProdutos.Scan(&p.Id, &p.Nome, &p.Preco, &p.Quantidade, &p.Descricao)
		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func BuscarProduto(produtoId string) Produto {
	db := db.ConectaBd()

	queryProduto, err := db.Query("select * from produtos where id=$1", produtoId)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for queryProduto.Next(){
		err = queryProduto.Scan(&produto.Id, &produto.Nome, &produto.Preco, &produto.Quantidade, &produto.Descricao)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()
	return produto
}

func CriarProduto(produto Produto) {
	db := db.ConectaBd()

	inserirDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	inserirDados.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	defer db.Close()
}

func AtualizarProduto(produto Produto) {
	db := db.ConectaBd()

	atualizarProduto, err := db.Prepare("update produtos set nome=$1, preco=$2, quantidade=$3, descricao=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizarProduto.Exec(produto.Nome, produto.Preco, produto.Quantidade, produto.Descricao, produto.Id)
	defer db.Close()
}

func DeletarProduto(produtoId string) {
	db := db.ConectaBd()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(produtoId)
	defer db.Close()
}
