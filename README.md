# Healthcare Database Management (Go, MongoDB, Gin)

[![Go](https://img.shields.io/badge/Go-1.18-blue.svg)](https://golang.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-7.0-green.svg)](https://www.mongodb.com/)
[![Gin](https://img.shields.io/badge/Gin-1.9.1-yellow.svg)](https://gin-gonic.com/)

A CRUD API for managing patient records using Go, MongoDB, and Gin framework.

---

## 🚀 How It's Useful

- Efficiently manage patient data with automatic ID generation.
- Easy integration with MongoDB for scalable and reliable data storage.
- Utilizes the Gin framework for building robust RESTful APIs.
- Highly customizable and extensible for various healthcare applications.

## 📝 How to Use It

1. Clone the repository:

   ```sh
   git clone https://github.com/samriddha-basu-cloud/Healtcare-db-management-go.git
   cd Healtcare-db-management-go
   ```

2. Initialize your MongoDB server and update the `mongoURI` variable in `main.go` with your MongoDB connection string.

3. Install Go dependencies:

   ```sh
   go mod tidy
   ```

4. Start the application:

   ```sh
   go run main.go
   ```

5. Access the API endpoints via [http://localhost:8080](http://localhost:8080).
    
    - Get List of All Patients (GET): http://localhost:8080/patients
    - Get a Patient by ID (GET): http://localhost:8080/patients/{id}
    - Add a New Patient (POST): http://localhost:8080/patients
    - Update a Patient (PUT): http://localhost:8080/patients/{id}
    - Delete a Patient by ID (DELETE): http://localhost:8080/patients/{id}
    - Delete All Patients (DELETE): http://localhost:8080/patients

## 🔗 How to Further Use in Other Projects

- Integrate this project as a backend for your healthcare applications.
- Extend the API with additional features like authentication and data analytics.
- Adapt the data models and routes to fit specific healthcare use cases.

## ❗ License

This project is not licensed. Feel free to use and modify it as needed.

---

