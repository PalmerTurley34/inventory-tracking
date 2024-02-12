# TUI App Docs

This app is built using the bubbletea framework. For more info on the bubbletea framework, check out the [github page](https://github.com/charmbracelet/bubbletea). There's lots of great tutorials and examples!

This app also depends on [lipgloss](https://github.com/charmbracelet/lipgloss) for styling and [huh](https://github.com/charmbracelet/huh) for using forms.

## Code structure

All code for the TUI app can be found in the `cmd/inventory_tracker_app/` directory.

I used some prefixes on the files to organize things a little better: cmd, lists, page etc.

### main.go

Starts the TUI program as well as a local HTTP server to make requests to.

### model.go

Contains all the logic and state for the main bubbletea model. The `model` struct contains some state regarding the application, such as:

* which page the app is currently showing
* what items the user has selected
* messages to display to the user

### forms.go

Contains some functions for generating new forms for the user to complete, such as:

* logging in
* creating an account
* creating new items

### page Prefix

All files starting with "page" contain the `Update` method for each page. This method is how each page handles the different messages that the bubbletea application creates. Depending on the message, the page may update the model state, switch to a new page, or execute an I/O command.

### cmd Prefix

All files starting with "cmd" contain implementations for commands. Commands are the way bubbletea performs I/O operations. Commands are functions that execute and return a message, which is given to the `Update` method to update the UI. All the "cmd" files will contain the function to execute (typically making an HTTP request) and also the messages the command might return.

### lists Prefix

Files with the "lists" prefix contain the structs the implement the `list.Item` interface. These are used in the app's main page for displaying inventory items and actions the user can take.

### styles.go

This file defines all of the `lipgloss` styles that are used to render the UI.
