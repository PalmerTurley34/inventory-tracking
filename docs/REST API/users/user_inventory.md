# User Inventory

Return all items that are currently checked out by the user.

**Endpoint** : `GET /v1/users/inventory`

**Authenticated Endpoint.** Header: `Authorization: ApiKey <token>` is requried.

## Success Response

**Code** : `200 OK`

### Example Response Body

```json
[
  {
    "id": "<item_UUID",
    "created_at": "2024-02-10T01:50:20.160704Z",
    "updated_at": "2024-02-10T01:50:20.160705Z",
    "name": "<item_name>",
    "checked_out_at": "2024-02-10T23:07:55.664875Z",
    "checked_in_at": null,
    "due_at": "2024-02-11T23:07:55.664875Z",
    "user_id": "<user_UUID>"
  },
  {
    "id": "<item_UUID>",
    "created_at": "2024-02-10T01:50:12.89671Z",
    "updated_at": "2024-02-10T01:50:12.896711Z",
    "name": "<item_name>",
    "checked_out_at": "2024-02-10T23:07:48.601377Z",
    "checked_in_at": null,
    "due_at": "2024-02-11T23:07:48.601377Z",
    "user_id": "<user_UUID>"
  },
  ...
]
```

## Error Response

**Code** : `401 Unauthorized`

**Condition** : Header `Authorization: ApiKey <token>` is not present

### Example Response

```json
{
    "error": "Error with ApiKey: authorization header is not found"
}
```

## Note

If the user has no items checked out, the response body will be `null`.
