package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()

	fmt.Println("hwuhuwhehuew")
	e.GET("/signIn", userSignIn)

	go e.Logger.Fatal(e.Start(":9008"))

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func userSignIn(c echo.Context) error {

	type Response struct {
		Message string `json:"message"`
	}

	msg := Response{
		Message: "Sign in route is working",
	}

	return c.JSON(http.StatusOK, msg)
}
