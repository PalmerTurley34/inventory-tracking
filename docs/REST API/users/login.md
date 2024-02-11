# Login

Authenticates user's password and returns user data, including and `ApiKey` for authenticated endpoints.

**Endpoint** : `POST /v1/login`

## Request Body

Expects `username` and `password` fields.

### Example Request Body

```json
{
    "username": "your_username_here",
    "password": "password123"
}
```

## Success Response

**Code** : `200 OK`

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

**Code** : `401 Unauthorized`

**Condition** : if `username` and `password` combination does not match user record

### Example Response

```json
{
    "error": "incorrect password, try again"
}
```
