package handler

import (
	"fmt"
	"net/http"
	"github.com/Abiji-2020/go-microservice.git/repository/order"
)
type Order struct{
	Repo *order.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter,r *http.Request){
	
}
func (o *Order) List(w http.ResponseWriter,r *http.Request){
	fmt.Println("List all Orders")
}
func (o *Order) GetById(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get Order by ID")
}

func (o *Order) UpdateBYID(w http.ResponseWriter,r *http.Request){
	fmt.Println("Update Order By Id")
}
func (o *Order) DeleteByID(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete Order By Id")
}