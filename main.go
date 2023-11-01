package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()

	e.POST("/createUser", createUser)
	e.GET("/getUser", getUser)
	/*
		Login API
		Logout API
	*/
	go e.Logger.Fatal(e.Start(":9008"))

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit)
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

func getUser(c echo.Context) error {
	userId := c.QueryParam("id")
	db, err := GetDbConnection()
	if err != nil {
		return c.JSON(http.StatusBadGateway, "Couldn't get DB connection")
	}

}

type UserInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phoneNo"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func createUser(c echo.Context) error {

	userInfo := UserInfo{}
	err := json.NewDecoder(c.Request().Body).Decode(&userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't decode request body")
	}

	//encrypt the password
	salt := randSeq(64)

	saltedPassword := fmt.Sprintf("%s%s", salt, userInfo.Password)

	hasher := sha256.New()
	hasher.Write([]byte(saltedPassword))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	db, err := GetDbConnection()
	if err != nil {
		log.Fatal("Coulnt get db connection", err)
	}
	query := "insert into Users (name,email, phoneNo, addedOn, updatedOn, salt, encPassword, address) values (?, ?, ?, ? , ?, ? , ? , ?)"

	response, err := db.Exec(query, userInfo.Name, userInfo.Email, userInfo.PhoneNo, time.Now(), time.Now(),
		salt, hashedPassword, userInfo.Address)
	if err != nil {
		fmt.Println("could not insert into db -> ", err)
		return err
	}

	lastId, err := response.LastInsertId()

	type Response struct {
		Message string `json:"message"`
		UserId  int    `json:"userId"`
	}

	msg := Response{
		Message: "User created",
		UserId:  int(lastId),
	}
	return c.JSON(200, msg)
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
