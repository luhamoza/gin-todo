**Title:** Gin Todo API

**Description:**

This is a simple RESTful API built with Gin for managing a todo list. It allows users to create, retrieve, update, and delete todos using JSON requests.

**Features:**

* Create new todos
* Get a list of all todos
* Get a specific todo by ID
* Update an existing todo
* Delete a todo

**Dependencies:**

* Gin: [https://github.com/gin-gonic](https://github.com/gin-gonic)
* [https://github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) 

**Installation:**

1. Clone this repository:

   ```bash
   git clone https://github.com/luhamoza/gin-todo.git
   ```

2. Install dependencies:

   ```bash
   cd gin-todo
   go mod download
   ```

**Database:**

This project currently uses a Sqlite database connection.  If you prefer a different database, you will need to update the `db` package to use the appropriate driver and configuration for your chosen database.

**Usage:**

1. Start the server:

   ```bash
   go run main.go
   ```

   The server will run on port `:3001` by default.

2. Use a tool like Postman or curl to send HTTP requests:

   * **Get all todos:**

     ```
     GET http://localhost:3001/todos
     ```

   * **Create a new todo:**

     ```
     POST http://localhost:3001/todos Content-Type: application/json

     {
       "title": "Buy groceries"
     }
     ```

   * **Get a todo by ID:**

     ```
     GET http://localhost:3001/todos/1
     ```

   * **Update a todo:**

     ```
     PUT http://localhost:3001/todos/1 Content-Type: application/json

     {
       "title": "Updated title"
     }
     ```

   * **Delete a todo:**

     ```
     DELETE http://localhost:3001/todos/1
     ```

**API Reference:**

| Endpoint        | Method | Description                                                    |
|-----------------|--------|-----------------------------------------------------------------|
| /todos          | GET    | Gets a list of all todos.                                          |
| /todos          | POST   | Creates a new todo.                                                |
| /todos/:id      | GET    | Gets a specific todo by ID.                                      |
| /todos/:id      | PUT    | Updates an existing todo.                                        |
| /todos/:id      | DELETE | Deletes a todo.                                                   |



**License:**

This project is licensed under the MIT License.
