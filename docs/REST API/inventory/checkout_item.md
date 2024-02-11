# Item Check Out

Checks out the item based off the item ID in the URL. Returns the new item record.

**Endpoint** : `POST /v1/inventory_items/checkout/{ID}`

**Authenticated Endpoint.** Header: `Authorization: ApiKey <token>` is requried.

Expects item UUID in the URL:

`v1/inventory_items/checkout/12345678-abcd-1234-abcd-123456abcdef`

## Success Response

**Code** : `200 OK`

### Example Response Body

```json
{
    "id": "<item UUID>",
    "created_at": "2024-02-10T01:50:12.89671Z",
    "updated_at": "2024-02-10T01:50:12.896711Z",
    "name": "<item name>",
    "checked_out_at": "2024-02-11T00:43:09.885614Z",
    "checked_in_at": "2024-02-11T00:38:13.179905Z",
    "due_at": "2024-02-12T00:43:09.885614Z",
    "user_id": "<user UUID>"
}
```

## Error Response

**Code** : `401 Unauthorized`

**Condition** : Header `Authorization: ApiKey <token>` is not present

**Code** : `400 Bad Request`

**Condition** : Item is already checked out

### Example Response

```json
{
    "error": "{ID} is not a valid uuid"
}
```
