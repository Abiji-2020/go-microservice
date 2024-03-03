package main

import (
	"context"
	"fmt"

	"github.com/Abiji-2020/go-microservice.git/application"
	
)	

func main() {
	app:= application.New()
	err:= app.Start(context.TODO())
	if err !=nil{
		fmt.Println("Failed to Start app:",err)
	}
}
