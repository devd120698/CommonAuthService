package main

import (
	"commonauthsvc/delivery"
	repoImpl "commonauthsvc/repository/impl"
	"commonauthsvc/service/impl"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/labstack/gommon/log"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()

	dbConn, err := GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repoImpl.Repository{
		DbConn: dbConn,
	}

	userSvc := impl.UserServiceImpl{
		UserRepo: &userRepo,
	}

	delivery.ConfigureHTTPHandler(e, &userSvc)

	//Start server
	go e.Logger.Fatal(e.Start(":9008"))

	// Graceful Shutdown
	gracefullShutdown(e)
}

func gracefullShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func GetDbConnection() (*sqlx.DB, error) {
	//connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "Divyansh", "Divyansh", "localhost", "3306", "CommonAuth")
	connection := fmt.Sprintf("%s:%s@/%s", "Divyansh", "Divyansh", "CommonAuth")
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "True")
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sqlx.Open(`mysql`, dsn)
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
