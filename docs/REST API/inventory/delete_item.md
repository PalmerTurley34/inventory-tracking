# New Item

Deletes item record.

**Endpoint** : `DELETE /v1/inventory_items/{ID}`

Expects item UUID in the URL:

`v1/inventory_items/12345678-abcd-1234-abcd-123456abcdef`

## Success Response

**Code** : `200 OK`

## Error Response

**Code** : `400 Bad Request`

**Condition** ID field is not present in the url

### Example Response

```json
{
    "error": "{ID} is not a valid uuid"
}
```
