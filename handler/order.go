package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/Abiji-2020/go-microservice.git/model"
	"github.com/Abiji-2020/go-microservice.git/repository/order"
	"github.com/google/uuid"
	"math/rand"
)

type Order struct {
	Repo *order.RedisRepo
}

func (h *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		customerId uuid.UUID        `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}
	if err:= json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()
	order:= model.Order{
		OrderId : rand.Uint64(),
		CustomerId: body.customerId,
		LineItems: body.LineItems,
		CreatedAt: &now,
	}

	err := h.Repo.Insert(r.Context(), order)
	if err != nil{
		fmt.Println("failed to insert: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

}
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Orders")
}
func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Order by ID")
}

func (o *Order) UpdateBYID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Order By Id")
}
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Order By Id")
}
