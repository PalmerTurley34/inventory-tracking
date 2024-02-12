# REST API Docs

All code for the the API can be found in the internal/backend/ directory.

The toy box API has endpoints for users to create and interact with inventory items. Authenticated enpoints must have an `Authorization: ApiKey <token>` header in the request.

## Open Endpoints

* [New User](users/new_user.md) : `POST /v1/users`
* [Login](users/login.md) : `POST /v1/login`
* [Validate Username](users/valid_username.md) : `POST /v1/valid_username`
* [All items](inventory/all_items.md) : `GET /v1/inventory_items`
* [New Item](inventory/create_item.md) : `POST /v1/inventory_items`
* [Delete Item](inventory/delete_item.md) : `DELETE /v1/inventory_items/{ID}`
* [Item History](inventory/item_history.md) : `GET /inventory_items/history/{ID}`

## Authenticated Endpoints

* [User Inventory](users/user_inventory.md) : `GET /v1/users/inventory`
* [Check Out Item](inventory/checkout_item.md) : `POST /v1/inventory_items/checkout/{ID}`
* [Check In Item](inventory/checkin_item.md) : `POST /v1/inventory_items/checkin/{ID}`

## Errors

In the event of an error with any request, the response body will always contain an `error` field with the corresponding error message.

**For Example** :

```json
{
    "error": "incorrect password, try again"
}
```
