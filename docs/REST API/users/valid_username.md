# Validate Username

Checks if the username provided is unique in the database.

**Endpoint** : `POST /v1/valid_username`

Returns a `valid` field that is a boolean

## Success Response

**Code** : `200 OK`

### Example Response Body

```json
{
    "valid": false
}
```
