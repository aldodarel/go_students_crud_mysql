package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aldodarel/go_students_crud_mysql/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterStudentsRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9800")
	log.Fatal(http.ListenAndServe("localhost:9800", r))
}