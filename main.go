package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"github.com/Abiji-2020/go-microservice.git/application"
	
)	

func main() {
	app:= application.New()
	ctx, cancel:= signal.NotifyContext(context.Background(),os.Interrupt)
	defer cancel()
	err:= app.Start(ctx)
	if err !=nil{
		fmt.Println("Failed to Start app:",err)
	}
}
