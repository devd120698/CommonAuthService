package main

import "github.com/labstack/echo"

func main() {
	//http.HandleFunc("/signIn", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, World!")
	//})
	//
	//http.ListenAndServe(":8080", nil)
	e := echo.New()
	e.Logger.Fatal(e.Start(":8000"))

}
