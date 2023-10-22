package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/signIn", createUser)

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

func GetDbConnection() (*sql.DB, error) {
	//connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "Divyansh", "Divyansh", "localhost", "3306", "CommonAuth")
	connection := fmt.Sprintf("%s:%s@/%s", "Divyansh", "Divyansh", "CommonAuth")
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "True")
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		return nil, err
	}
	dbConn.SetMaxOpenConns(20)
	dbConn.SetConnMaxLifetime(5 * time.Minute)
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func createUser(c echo.Context) error {

	//ctx := c.Request().Context()
	name := c.QueryParam("name")
	password := c.QueryParam("password")

	db, err := GetDbConnection()
	if err != nil {
		log.Fatal("Coulnt get db connection", err)
	}
	fmt.Println("db connection got ->", db)

	query := "insert into Users (name, password) values (?, ?)"

	response, err := db.Exec(query, name, password)
	if err != nil {
		fmt.Println("could not insert into db -> ", err)
		return err
	}

	type Response struct {
		Message string      `json:"message"`
		Result  interface{} `json:"result"`
	}

	msg := Response{
		Message: "Create User route is working",
		Result:  response,
	}

	return c.JSON(http.StatusOK, msg)
}
