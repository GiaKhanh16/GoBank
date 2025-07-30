# 🏦 GoLang Bank Management API

A simple bank management API built with **Go**, demonstrating various Go features and best practices.

---

## 🚀 Features

### 🌐 HTTP Server
- Uses **`net/http`** to create and run an HTTP server (`http.ListenAndServe`).
- Implements REST API routes with different HTTP methods (**GET**, **POST**, **DELETE**).

### 🔀 Routing with Gorilla Mux
- Uses **`github.com/gorilla/mux`** to handle dynamic URL parameters (e.g., `/account/{id}`).
- Extracts path variables with `mux.Vars`.

### 📦 JSON Handling
- Encodes responses using `json.NewEncoder().Encode`.
- Decodes request bodies with `json.NewDecoder().Decode`.
- Sets appropriate `Content-Type` headers.

### 🛡️ Custom Middleware
- Implements a custom **JWT authentication** middleware (`WithJWTAuth`) to secure endpoints.

### 🔑 JWT (JSON Web Tokens) Authentication
- Uses **`github.com/golang-jwt/jwt/v5`** to create and validate JWT tokens.
- Demonstrates signing, parsing, and validating tokens with **HMAC**.

### 🔒 Environment Variables
- Reads the secret key from an environment variable (`os.Getenv("JWT_SECRET")`).

### ⚠️ Error Handling
- Centralized error handling using a custom `apiFunc` type and `makeHTTPHandleFunc` wrapper.
- Returns structured JSON error responses (`ApiError`).

### 🏗️ Structs and Interfaces
- Defines `APIServer` struct with methods as route handlers.
- Uses a `Storage` interface (implementation implied) to abstract data storage.

### 📌 Pointer Usage
- Uses pointers for request struct decoding (e.g., `new(CreateAccountRequest)`, `new(TransferRequest)`).

### 🛠️ Custom Types
- Implements custom request types (`CreateAccountRequest`, `TransferRequest`) with proper JSON handling.

### 🔁 Routing by HTTP Method
- Checks `r.Method` to handle different operations on the same endpoint.

### 🧩 Closures in Handlers
- Uses closures in `makeHTTPHandleFunc` and `WithJWTAuth` to wrap other functions elegantly.

### 🧹 Deferred Cleanup
- Uses `defer r.Body.Close()` to ensure request body is closed after processing.

---

## ✅ How to Run

```bash
# Clone the repository
git clone https://github.com/yourusername/GoLang-Bank-Management.git
cd GoLang-Bank-Management

# Set environment variable for JWT secret
export JWT_SECRET=your_secret_key

# Run the server
go run main.go
