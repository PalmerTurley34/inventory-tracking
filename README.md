# Toy Box

This is a TUI app, built with the [bubbletea](https://github.com/charmbracelet/bubbletea/tree/master) framework in Go. The backend is an HTTP REST API built in Go, also using a PostgreSQL database. Documentation for the app can be found [here](docs/bubbletea%20app/README.md). And documentation for the REST API can be found [here](docs/REST%20API/README.md).

## Tech Stack

* [Bubbletea](https://github.com/charmbracelet/bubbletea/tree/master) TUI library
* [Chi](https://github.com/go-chi/chi) HTTP library
* [SQLC](https://sqlc.dev/) SQL code generation tool
* [goose](https://github.com/pressly/goose) database migration tool
* [postgreSQL](https://www.postgresql.org/) database

## Description

I created this app as a model for sharing toys for my toddlers who are trying to learn the concept.

This is a sample app, but could easily be translated into other domains such as:

* A library book checkout system
* Tracking employee's use of company resources
* Keeping track of repeated tasks, chores, daily cleaning checklist, etc.
* And more!

The main page of the app is a view of three lists: a list of all the items and their check out status, a list of the actions you can perform, and a list of the current user's checked out items:

![App Main Page](docs/images/main_page.png "Main Page")

The app also contains other pages for filling out forms, view an item's check out history, etc.

<img src="docs/images/create_item_page.png" height="300">

<img src="docs/images/create_user_page.png" height="300">

<img src="docs/images/item_history.png" height="200">

## Setup and Installation

### Environment Variables

The only environment variable required is `DB_URL`. This variable is used by the HTTP server to connect to the database. The expected database type is postgres. I use a `.env` file in the root directory of the project for this, with this line:

`DB_URL=postgres://palmer:@localhost:5432/my_database`

### Database

The database schema can be found in `sql/shema/`. I used [goose](https://github.com/pressly/goose) to handle the migrations. From within the `sql/schema/` directory run:

`goose postgres <DB_URL> up`

to create the tables.

### Build and Run

You can compile the executable from the root directory of the project:

`go build ./cmd/inventory_tracker_app`

run the executable:

`./inventory_tracker_app`