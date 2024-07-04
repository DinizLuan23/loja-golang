package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFormatado, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na formatação do preço", err)
		}

		quantidadeFormatada, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na formatação do preço", err)
		}

		produto := models.Produto{
			Nome: nome,
			Descricao: descricao,
			Preco: precoFormatado,
			Quantidade: quantidadeFormatada,
		}

		models.CriarProduto(produto)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	produtoId := r.URL.Query().Get("id")
	models.DeletarProduto(produtoId)
	http.Redirect(w, r, "/", 301)
}
