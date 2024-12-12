# API Documentation

This API provides endpoints for managing customer authentication, product management, shopping cart operations, and transaction handling.

## Table of Contents

- [Customer Authentication](#customer-authentication)
  - [Login](#login)
  - [Register](#register)
- [Product Management](#product-management)
  - [Fetch Products](#fetch-products)
  - [Product Details](#product-details)
  - [Caching with Valkey](#caching-with-valkey)
- [Shopping Cart Management](#shopping-cart-management)
  - [View Cart](#view-cart)
  - [Add to Cart](#add-to-cart)
  - [Remove from Cart](#remove-from-cart)
- [Checkout and Transactions](#checkout-and-transactions)
  - [Get Transactions](#get-transactions)
  - [Checkout](#checkout)
  - [Cancel Transaction](#cancel-transaction)
  - [Handle Payment](#handle-payment)
- [User Profile](#user-profile)
- [Error Handling](#error-handling)
- [Security](#security)
- [Dependencies](#dependencies)


## Customer Authentication

### Login

**Endpoint:** `POST /auth/login`

**Description:** Logs in a user and returns a JWT token.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "message": "Login successful"
}
```

### Register

**Endpoint:** `POST /auth/register`

**Description:** Registers a new user.

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "message": "User registered successfully"
}
```

## Product Management

### Fetch Products

**Endpoint:** `GET /products`

**Description:** Fetches a paginated list of products, optionally filtered by category, price range, and search query.

**Query Parameters:**
- `page` (int, default: 1)
- `limit` (int, default: 10)
- `category` (string, e.g., "electronics")
- `min_price` (float, default: 0)
- `max_price` (float, default: 0)
- `search` (string)

**Response:**

```json
{
  "current_page": 1,
  "total_items": 100,
  "data": [
    {
      "id": 1,
      "name": "Product 1",
      "price": 19.99,
      "description": "This is product 1."
    }
  ]
}
```

### Product Details

**Endpoint:** `GET /products/:id`

**Description:** Retrieves details of a specific product by ID.

**Path Parameters:**
- `id` (int)

**Response:**

```json
{
  "data": {
    "id": 1,
    "name": "Product 1",
    "price": 19.99,
    "description": "This is product 1."
  }
}
```

### Caching with Valkey

The product API utilizes Valkey for caching frequently accessed product data.

## Shopping Cart Management

### View Cart

**Endpoint:** `GET /cart`

**Description:** Retrieves the current user's shopping cart items and total.

**Response:**

```json
{
  "message": "Cart retrieved successfully",
  "data": [
    {
      "id": 1,
      "product_name": "Product 1",
      "quantity": 2,
      "price": 19.99
    }
  ],
  "total": 39.98
}
```

### Add to Cart

**Endpoint:** `POST /cart`

**Description:** Adds a product to the user's shopping cart.

**Request Body:**

```json
{
  "product_id": 1,
  "quantity": 2
}
```

**Response:**

```json
{
  "message": "Product added to cart successfully"
}
```

### Remove from Cart

**Endpoint:** `DELETE /cart/:id`

**Description:** Removes a specific product or all items from the cart.

**Path Parameters:**
- `id` (int, product ID)

**Query Parameters:**
- `all` (bool, default: false)

**Response:**

```json
{
  "message": "Cart item(s) removed successfully"
}
```

## Checkout and Transactions

### Get Transactions

**Endpoint:** `GET /transaction`

**Description:** Retrieves the user's transaction history.

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "total_amount": 39.98,
      "status": "completed",
      "created_at": "2023-10-01T12:00:00Z"
    }
  ]
}
```

### Checkout

**Endpoint:** `POST /transaction/checkout`

**Description:** Initiates the checkout process and records a transaction.

**Response:**

```json
{
  "message": "Checkout successful"
}
```

### Cancel Transaction

**Endpoint:** `DELETE /transaction/:id/cancel`

**Description:** Cancels a transaction by ID.

**Path Parameters:**
- `id` (int, transaction ID)

**Response:**

```json
{
  "message": "Transaction canceled successfully"
}
```

### Handle Payment

**Endpoint:** `POST /transaction/:id/pay`

**Description:** Processes payment for a transaction.

**Path Parameters:**
- `id` (int, transaction ID)

**Response:**

```json
{
  "message": "Payment successful"
}
```

## User Profile

### Profile

**Endpoint:** `GET /user/me`

**Description:** Retrieves the authenticated user's profile.

**Response:**

```json
{
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "created_at": "2023-09-01T10:00:00Z"
  }
}
```

## Error Handling

- `400 Bad Request`: Invalid request data or parameters.
- `401 Unauthorized`: Missing or invalid JWT token.
- `404 Not Found`: Resource not found.
- `500 Internal Server Error`: Server-side error.

## Security

- **Authentication:** JWT-based authentication is used for securing endpoints.
- **Password Hashing:** Passwords are securely hashed using bcrypt.

## Dependencies

- Go Fiber for the web framework.
- JWT for token-based authentication.
- Valkey for caching.
- Logging with Logrus.


## Deployment

