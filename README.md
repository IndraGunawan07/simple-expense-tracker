# 📘 Simple Expense Tracker API

This is a RESTful API for a simple expense tracker application built using **Go (Golang)**. The API allows users to register, log in, and manage their expense data including categories and reports.

---

## 🚀 Features

- User registration and authentication
- CRUD operations for expense categories
- CRUD operations for expenses
- Expense reports generation

---

## 🛠 Tech Stack

- **Backend**: Go (Golang)
- **Routing**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **Authentication**: JWT (JSON Web Tokens)
- **Database**: PostgreSQL

---

## 📦 API Endpoints

### 🧑 User Authentication

| Method | Endpoint    | Description                                   |
| ------ | ----------- | --------------------------------------------- |
| POST   | `/register` | Register new user                             |
| POST   | `/login`    | Log in and receive JWT token                  |
| GET    | `/user`     | Get current authenticated user (requires JWT) |

### 📂 Categories (requires JWT)

| Method | Endpoint          | Description             |
| ------ | ----------------- | ----------------------- |
| GET    | `/categories`     | Get all categories      |
| POST   | `/categories`     | Create a new category   |
| PUT    | `/categories/:id` | Update a category by ID |
| DELETE | `/categories/:id` | Delete a category by ID |

### 💸 Expenses (requires JWT)

| Method | Endpoint        | Description             |
| ------ | --------------- | ----------------------- |
| GET    | `/expenses`     | Get all expenses        |
| POST   | `/expenses`     | Create a new expense    |
| PUT    | `/expenses/:id` | Update an expense by ID |
| DELETE | `/expenses/:id` | Delete an expense by ID |

### 📊 Reports (requires JWT)

| Method | Endpoint   | Description                 |
| ------ | ---------- | --------------------------- |
| GET    | `/reports` | Get expense summary reports |

---

## 🧪 How to Run Locally

1. **Clone the Repository**

   ```bash
   git clone https://github.com/IndraGunawan07/simple-expense-tracker.git
   cd simple-expense-tracker
   ```

2. **Install Go Dependencies**

   ```bash
   go mod tidy
   ```

3. **Configure Environment Variables**

   Create a `.env` file or set the environment variables manually:

   ```env
   PORT=8080
   DB_URL=your_database_connection_string
   JWT_SECRET=your_jwt_secret_key
   ```

4. **Run the Application**
   ```bash
   go run main.go
   ```
