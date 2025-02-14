/*
 * Desafio conductor
 *
 * Api para controle de trasacoes de contas
 *
 * API version: 1.0.0
 * Contact: thiagorodriguescamara@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ThiagoRodriguesdeSantana/desafio_conductor/go/file"
	"github.com/gorilla/mux"
)

//TransactionsPDF return pdf from transactions
func (c *Controller) TransactionsPDF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	acc, _ := c.Db.FindTransactionByAccountID(id)

	file := file.NewPDF()

	path := file.GeneratePDF(acc)

	filePdf, _ := ioutil.ReadFile(path)

	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(filePdf)

}
