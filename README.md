

# Go Contact API

A simple contact management REST API built with Go.  
Supports full CRUD (Create, Read, Update, Delete) operations using HTTP requests.

---

## 🚀 Features

- ✅ Create a new contact (`POST /contact`)
- 📥 Get all contacts (`GET /contact`)
- 📝 Update a contact (`PUT /contact/{id}`)
- ❌ Delete a contact (`DELETE /contact/{id}`)

---

## 🛠 Tech Stack

- Go (Golang)
- Built-in `net/http` package
- JSON for data exchange

---

## 📦 Setup Instructions

1. **Clone the repository**  
   ```bash
   git clone https://github.com/odog1234/go-contact-api.git
   cd go-contact-api
   ```

2. **Run the API**  
   ```bash
   go run main.go
   ```

3. **Test with Postman or Curl**  
   - `GET` → http://localhost:8080/contact  
   - `POST` → Send a JSON body to `http://localhost:8080/contact`  
   - `PUT` → Update via `http://localhost:8080/contact/{id}`  
   - `DELETE` → Remove via `http://localhost:8080/contact/{id}`

---

## 🧠 Example JSON Payload

```json
{
  "name": "Owen Yesuf",
  "email": "owen@example.com",
  "number": "123-456-7890"
}
```

---

## 👨‍💻 Author

**Owen Yesuf**  
[LinkedIn Profile](https://www.linkedin.com/in/owen-yesuf-24b8842a6)

---

## 📌 Notes

This project is part of my backend learning journey using Go.  
More features like persistent storage (PostgreSQL), validation, and deployment will be added later.