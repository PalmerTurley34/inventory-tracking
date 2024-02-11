# New User

Creates new user record and returns user data, including an `ApiKey` for authenticated endpoints.

**Endpoint** : `POST /v1/users`

## Request Body

Expects `name`, `username`, and `password` fields.

`username` field must be unique.

`password` field must be at least 8 characters in length.

### Example Request Body

```json
{
    "name": "your_name_here",
    "username": "your_username_here",
    "password": "password123"
}
```

## Success Response

**Code** : `201 Created`

### Example Response Body

```json
{
    "id": "<UUID String>",
    "created_at": "2024-02-10T01:48:22.545585Z",
    "updated_at": "2024-02-10T01:48:22.545585Z",
    "name": "<name>",
    "username": "<username>",
    "is_admin": false,
    "api_key": "<api key token>"
}
```

## Error Response

**Code** : `400 Bad Request`

**Condition** :

* There is already a record with the same `username` field (must be unique)
* `password` is not at least 8 characters in length

### Example Response

```json
{
    "error": "Username: <username> already exists. Please choose a unique username."
}
```
