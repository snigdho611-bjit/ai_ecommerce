# E-commerce API Documentation

## Table of Contents

-   [Product Endpoints](#product-endpoints)
-   [User Endpoints](#user-endpoints)
-   [Cart Endpoints](#cart-endpoints)
-   [Error Formats](#error-formats)

## Product Endpoints

### Create Product

-   **Endpoint**: `/api/products`
-   **Method**: `POST`
-   **Authentication**: Required
-   **Request Body**:
    ```json
    {
        "name": "string",
        "description": "string",
        "price": "float",
        "stock": "int"
    }
    ```

### Get All Products

-   **Endpoint**: `/api/products`
-   **Method**: `GET`
-   **Authentication**: Not Required

### Get Product By ID

-   **Endpoint**: `/api/products/:id`
-   **Method**: `GET`
-   **Authentication**: Not Required

### Update Product

-   **Endpoint**: `/api/products/:id`
-   **Method**: `PUT`
-   **Authentication**: Required
-   **Request Body**:
    ```json
    {
        "name": "string",
        "description": "string",
        "price": "float",
        "stock": "int"
    }
    ```

### Delete Product

-   **Endpoint**: `/api/products/:id`
-   **Method**: `DELETE`
-   **Authentication**: Required

## User Endpoints

### Register

-   **Endpoint**: `/api/users/register`
-   **Method**: `POST`
-   **Authentication**: Not Required
-   **Request Body**:
    ```json
    {
        "name": "string",
        "email": "string",
        "password": "string"
    }
    ```

### Login

-   **Endpoint**: `/api/users/login`
-   **Method**: `POST`
-   **Authentication**: Not Required
-   **Request Body**:
    ```json
    {
        "email": "string",
        "password": "string"
    }
    ```

### Get User Profile

-   **Endpoint**: `/api/users/profile`
-   **Method**: `GET`
-   **Authentication**: Required

### Update User Profile

-   **Endpoint**: `/api/users/profile`
-   **Method**: `PUT`
-   **Authentication**: Required
-   **Request Body**:
    ```json
    {
        "name": "string",
        "email": "string",
        "password": "string"
    }
    ```

## Cart Endpoints

### Add Item to Cart

-   **Endpoint**: `/api/cart`
-   **Method**: `POST`
-   **Authentication**: Required
-   **Request Body**:
    ```json
    {
        "product": "string",
        "quantity": "int"
    }
    ```

### Get Cart

-   **Endpoint**: `/api/cart`
-   **Method**: `GET`
-   **Authentication**: Required

### Update Cart Item

-   **Endpoint**: `/api/cart/:id`
-   **Method**: `PUT`
-   **Authentication**: Required
-   **Request Body**:
    ```json
    {
        "quantity": "int"
    }
    ```

### Delete Cart Item

-   **Endpoint**: `/api/cart/:id`
-   **Method**: `DELETE`
-   **Authentication**: Required

## Error Formats

### Error Response

```json
{
    "message": "string",
    "errors": ["string"]
}
```

### Validation Error

```json
{
    "message": "Validation Error",
    "errors": ["string"]
}
```

### Unauthorized Error

```json
{
    "message": "Unauthorized"
}
```

### Not Found Error

```json
{
    "message": "Not Found"
}
```

### Internal Server Error

```json
{
    "message": "Internal Server Error"
}
```

### Not Implemented Error

```json
{
    "message": "Not Implemented"
}
```

### Service Unavailable Error

```json
{
    "message": "Service Unavailable"
}
```

### Bad Request Error

```json
{
    "message": "Bad Request"
}
```

### Conflict Error

```json
{
    "message": "Conflict"
}
```

### Forbidden Error

```json
{
    "message": "Forbidden"
}
```

### Too Many Requests Error

```json
{
    "message": "Too Many Requests"
}
```

### Unauthorized Error

```json
{
    "message": "Unauthorized"
}
```