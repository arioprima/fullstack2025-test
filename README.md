# Client API Documentation

## Base URL

```
http://localhost:8080
```

## Endpoints

### 1. Create a Client

**Request:**

```
POST /clients
```

**Headers:**

```
Content-Type: application/json
```

**Body:**

```json
{
  "name": "Client A",
  "slug": "client-a",
  "is_project": "0",
  "self_capture": "1",
  "client_prefix": "CA",
  "client_logo": "https://example.com/logo.jpg",
  "address": "123 Main St",
  "phone_number": "123-456-7890",
  "city": "New York"
}
```

### 2. Get a Client

**Request:**

```
GET /clients/{slug}
```

**Example:**

```
GET /clients/client-a
```

### 3. Update a Client

**Request:**

```
PUT /clients/{slug}
```

**Headers:**

```
Content-Type: application/json
```

**Body:**

```json
{
  "name": "Client B",
  "is_project": "1",
  "self_capture": "0",
  "client_prefix": "CB",
  "client_logo": "https://example.com/new-logo.jpg",
  "address": "456 Elm St",
  "phone_number": "987-654-3210",
  "city": "Los Angeles"
}
```

### 4. Delete a Client

**Request:**

```
DELETE /clients/{slug}
```

**Example:**

```
DELETE /clients/client-a
```
