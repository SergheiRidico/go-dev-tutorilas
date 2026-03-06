# Vintage Jazz Records Repository

A lightweight RESTful web service built with the **Gin Web Framework** in Go.

The service contains the following endpoints:
| Method | Endpoint Path | Description                                               |
| :----- | :------------ | :-------------------------------------------------------- |
| `GET`  | `/albums`     | Get a list of all albums, returned as JSON.               |
| `POST` | `/albums`     | Add a new album from request data sent as JSON.           |
| `GET`  | `/users/{id}` | Get an album by its ID, returning the album data as JSON. |

---

### Prerequisites
* **Go** 1.16 or higher
* Command terminal
* **cURL** tool

---

### How to run

```bash
$ go run .
```
### How to access endpoints

Get all albums
```bash
$ curl -i http://localhost:8080/albums
```

Get single album
```bash
$ curl http://localhost:8080/albums/1
```

Add new album
```bash
$ curl -i http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
