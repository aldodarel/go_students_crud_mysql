package routes

import (
	"github.com/aldodarel/go_students_crud_mysql/pkg/controllers"
	"github.com/gorilla/mux"
)

// Setting ROUTES untuk Request di Postman
var RegisterStudentsRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", 
	controllers.CreateStudent).Methods("POST")	// route endpoint untuk method POST
	router.HandleFunc("/student/", controllers.GetStudent).Methods("GET")		// route endpoint untuk method GET
	router.HandleFunc("/student/{studentId}", 	// route endpoint untuk method GET by NIM
	controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/student/{studentId}", 
	controllers.UpdateStudent).Methods("PUT")	// route endpoint untuk method PUT (Update)
	router.HandleFunc("/student/{studentId}", 
	controllers.DeleteStudent).Methods("DELETE")	// route endpoint untuk method DELETE
}