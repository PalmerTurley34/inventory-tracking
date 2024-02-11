# New Item

Creates new item record and returns item data.

**Endpoint** : `POST /v1/inventory_items`

## Request Body

Expects `name` field.

### Example Request Body

```json
{
    "name": "your_item_name_here"
}
```

## Success Response

**Code** : `201 Created`

### Example Response Body

```json
{
  "id": "<item UUID>",
  "created_at": "2024-02-11T01:00:10.26612Z",
  "updated_at": "2024-02-11T01:00:10.26612Z",
  "name": "<item name>",
  "checked_out_at": null,
  "checked_in_at": null,
  "due_at": null,
  "user_id": null
}
```

## Error Response

**Code** : `400 Bad Request`

**Condition** `name` field is not present in request body

### Example Response

```json
{
    "error": "couldn't find name field in request body"
}
```
