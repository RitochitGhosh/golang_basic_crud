# 📝 Notes API

A simple RESTful CRUD API for managing notes, built with **Go**, **Gin**, and **MongoDB**.

---

## 🛠️ Tech Stack

- **Language:** Go
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database:** MongoDB (via [mongo-driver v2](https://github.com/mongodb/mongo-go-driver))

---

## 📁 Project Structure

```
notes/
├── model.go       # Data models and request types
├── repo.go        # Data access layer (MongoDB operations)
├── handler.go     # HTTP handlers (Gin controllers)
└── routes.go      # Route registration
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.21+
- [MongoDB](https://www.mongodb.com/try/download/community) running locally or a MongoDB Atlas URI

### Installation

```bash
git clone https://github.com/your-username/notes-api.git
cd notes-api
go mod tidy
```

### Running the Server

```bash
go run main.go
```

The server starts on `http://localhost:8080` by default.

---

## 🔌 API Endpoints

| Method   | Endpoint      | Description          |
|----------|---------------|----------------------|
| `POST`   | `/notes`      | Create a new note    |
| `GET`    | `/notes`      | List all notes       |
| `GET`    | `/notes/:id`  | Get a note by ID     |
| `PUT`    | `/notes/:id`  | Update a note by ID  |
| `DELETE` | `/notes/:id`  | Delete a note by ID  |

---

## 📦 Request & Response Examples

### Create a Note

**POST** `/notes`

```json
// Request Body
{
  "title": "My First Note",
  "content": "This is the content of my note.",
  "pinned": false
}
```

```json
// Response — 201 Created
{
  "id": "665f1a2b3c4d5e6f7a8b9c0d",
  "title": "My First Note",
  "content": "This is the content of my note.",
  "pinned": false,
  "createdAt": "2024-06-04T10:00:00Z",
  "updatedAt": "2024-06-04T10:00:00Z"
}
```

---

### List All Notes

**GET** `/notes`

```json
// Response — 200 OK
{
  "notes": [
    {
      "id": "665f1a2b3c4d5e6f7a8b9c0d",
      "title": "My First Note",
      "content": "This is the content of my note.",
      "pinned": false,
      "createdAt": "2024-06-04T10:00:00Z",
      "updatedAt": "2024-06-04T10:00:00Z"
    }
  ]
}
```

---

### Get a Note by ID

**GET** `/notes/:id`

```json
// Response — 200 OK
{
  "note": {
    "id": "665f1a2b3c4d5e6f7a8b9c0d",
    "title": "My First Note",
    "content": "This is the content of my note.",
    "pinned": false,
    "createdAt": "2024-06-04T10:00:00Z",
    "updatedAt": "2024-06-04T10:00:00Z"
  }
}
```

---

### Update a Note

**PUT** `/notes/:id`

```json
// Request Body
{
  "title": "Updated Title",
  "content": "Updated content.",
  "pinned": true
}
```

```json
// Response — 200 OK
{
  "id": "665f1a2b3c4d5e6f7a8b9c0d",
  "title": "Updated Title",
  "content": "Updated content.",
  "pinned": true,
  "createdAt": "2024-06-04T10:00:00Z",
  "updatedAt": "2024-06-04T11:00:00Z"
}
```

---

### Delete a Note

**DELETE** `/notes/:id`

```json
// Response — 200 OK
{
  "ok": true
}
```

---

## ⚠️ Error Responses

All error responses follow this structure:

```json
{
  "error": "Error message here"
}
```

| Status Code | Meaning                          |
|-------------|----------------------------------|
| `400`       | Bad request / invalid input      |
| `404`       | Note not found                   |
| `500`       | Internal server error            |

---

## 🗃️ Data Model

```go
type Note struct {
    ID        primitive.ObjectID `bson:"_id"       json:"id"`
    Title     string             `bson:"title"     json:"title"`
    Content   string             `bson:"content"   json:"content"`
    Pinned    bool               `bson:"pinned"    json:"pinned"`
    CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
    UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
```

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).