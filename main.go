package main

import (
	"fmt"
	"context"
	"github.com/Abiji-2020/go-microservice/application"
)	

func main() {
	app:= application.New()
	err:= app.Start(context.TODO())
}
