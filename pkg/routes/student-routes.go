package routes

import (
	"github.com/aldodarel/go_students_crud_mysql/pkg/controllers"
	"github.com/gorilla/mux"
)

// Setting ROUTES untuk Request di Postman
var RegisterStudentsRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", 
	controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/student/", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/student/{studentId}", 
	controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/student/{studentId}", 
	controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentId}", 
	controllers.DeleteStudent).Methods("DELETE")
}