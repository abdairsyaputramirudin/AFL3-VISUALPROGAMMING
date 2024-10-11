## AFL3 Visual Programming

### Base URL

http://localhost:8080

### Endpoints

### 1. Get All Books

- **URL**: `/books`
- **Method**: `GET`
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    [
      {
        "id": 1,
        "title": "Clean Code",
        "author": "Robert C. Martin",
        "price": 29.99
      },
      {
        "id": 2,
        "title": "The Pragmatic Programmer",
        "author": "Andrew Hunt",
        "price": 34.99
      }
    ]
    ```

### 2. Get Books By Id

- **URL**: `/books/:id`
- **Method**: `GET`
- **URL Params**:
  - `id` (required): The ID of the book
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "id": 1,
      "title": "Clean Code",
      "author": "Robert C. Martin",
      "price": 29.99
    }
    ```
  - **Status**: `404 Not Found`
  - **Body**:
    ```json
    {
      "message": "book not found"
    }
    ```

### 3. Create a New Book

- **URL**: `/books`
- **Method**: `POST`
- **Request Body**:
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "title": "The Art of Computer Programming",
      "author": "Donald Knuth",
      "price": 59.99
    }
    ```
- **Response**:
  - **Status**: `201 Created`
  - **Body**:
    ```json
    {
      "id": 5,
      "title": "The Art of Computer Programming",
      "author": "Donald Knuth",
      "price": 59.99
    }
    ```
  - **Status**: `400 Bad Request`
  - **Body**:
    ```json
    {
      "error": "Invalid request body"
    }
    ```

### 4. Update a Book

- **URL**: `/books/:id`
- **Method**: `PUT`
- **URL Params**:
  - `id` (required): The ID of the book
- **Request Body**:
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "title": "The Art of Computer Programming - Volume 1",
      "author": "Donald Knuth",
      "price": 69.99
    }
    ```
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "message": "book updated successfully"
    }
    ```
  - **Status**: `404 Not Found`
  - **Body**:
    ```json
    {
      "message": "book not found"
    }
    ```

### 5. Delete a Book

- **URL**: `/books/:id`
- **Method**: `DELETE`
- **URL Params**:
  - `id` (required): The ID of the book
- **Response**:
  - **Status**: `200 OK`
  - **Body**:
    ```json
    {
      "message": "book deleted successfully"
    }
    ```
  - **Status**: `404 Not Found`
  - **Body**:
    ```json
    {
      "message": "book not found"
    }
    ```
