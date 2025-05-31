
# 🎓 Student Management REST API with Golang and MySQL

This project is a simple **RESTful API** built with **Go (Golang)** and **MySQL** to manage student data. It is designed using a modular architecture to separate concerns across configuration, controllers, models, routes, and utilities.

---
### go_students_crud_mysql (PASTI)

## 📁 Project Structure
student-management-api/
├── cmd/
│ └── main/
│ └── main.go # Application entry point
│
├── pkg/
│ ├── config/
│ │ └── config.go # Database configuration and connection
│
│ ├── controller/
│ │ └── student_controller.go # Handles request and response logic
│
│ ├── models/
│ │ └── student.go # Student model struct and database mapping
│
│ ├── routes/
│ │ └── student_routes.go # Defines all route endpoints
│
│ └── utils/
│ └── response.go # JSON unmarshal and helper functions
│
├── go.mod # Go module file
├── go.sum # Go dependencies lock file
└── README.md # Project documentation

---

## 📌 Features

- 📄 `GET /students` – Retrieve all student records  
- 🔍 `GET /students/:nim` – Get student by NIM  
- ➕ `POST /students` – Add a new student  
- ✏️ `PUT /students/:nim` – Update student information  
- ❌ `DELETE /students/:nim` – Delete student by NIM

---

## ⚙️ Technologies Used

- **Golang** – Programming language  
- **MySQL** – Relational database  
- **net/http** – Built-in package to create the web server  
- **GORM (optional)** – If ORM is used for DB interaction  
- **JSON** – Data exchange format

