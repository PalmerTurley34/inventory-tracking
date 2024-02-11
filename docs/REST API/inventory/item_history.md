# Item History

Shows the check out history of the given item.

**Endpoint** : `GET /v1/inventory_items/history/{ID}`

Expects item UUID in the URL:

`v1/inventory_items/history/12345678-abcd-1234-abcd-123456abcdef`

## Success Response

**Code** : `200 OK`

### Example Response Body

```json
[
  {
    "name": "<item name>",
    "username": "<username>",
    "checked_out_at": "2024-02-10T04:46:21.979753Z",
    "checked_in_at": "2024-02-10T21:59:13.79227Z"
  },
  {
    "name": "<item name>",
    "username": "<username>",
    "checked_out_at": "2024-02-10T04:45:48.351725Z",
    "checked_in_at": "2024-02-10T04:46:11.508561Z"
  },
  {
    "name": "<item name>",
    "username": "<username>",
    "checked_out_at": "2024-02-09T20:07:31.862244Z",
    "checked_in_at": "2024-02-10T04:45:41.528091Z"
  }
]
```

## Error Response

**Code** : `400 Bad Request`

**Condition** `ID` is not present in the url

### Example Response

```json
{
    "error": "{ID} is not a valid uuid"
}
```
