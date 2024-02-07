package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh"
)

func NewInitialForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("Welcome to the Toy Box!").Description("Select an option to begin:"),
			huh.NewSelect[string]().
				Key("option").
				Options(huh.NewOptions("Login", "Create Account")...),
		),
	)
}

func NewLoginForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("Enter username and password."),
			huh.NewInput().
				Key("username").
				Prompt("Username: "),

			huh.NewInput().
				Key("password").
				Prompt("Password: ").
				Password(true).
				Validate(validatePassword),

			huh.NewConfirm().
				Key("done").
				Title("Confirm").
				Validate(func(v bool) error {
					if !v {

						return fmt.Errorf("Shift+Tab to go back")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("No"),
		),
	)
}

func NewCreateUserForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Enter you information:"),

			huh.NewInput().
				Key("name").
				Prompt("Name: ").
				Validate(func(s string) error {
					if len(s) > 0 {
						return nil
					}
					return fmt.Errorf("name cannot be blank")
				}),

			huh.NewInput().
				Key("username").
				Prompt("Username: ").
				Validate(validateUsername),

			huh.NewInput().
				Key("password").
				Prompt("Password: ").
				Password(true).
				Placeholder("Must be at least 8 characters").
				Validate(validatePassword),

			huh.NewConfirm().
				Key("done").
				Title("Confirm").
				Validate(func(v bool) error {
					if !v {

						return fmt.Errorf("Shift+Tab to go back")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("No"),
		),
	)
}

func NewCreateItemForm() *huh.Form {
	return nil
}

func validatePassword(s string) error {
	if len(s) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

func validateUsername(s string) error {
	resp, err := http.Post(
		"http://localhost:8080/v1/valid_username",
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"username": "%s"}`, s))),
	)
	if err != nil {
		return fmt.Errorf("couldn't contact server: %v", err)
	}
	defer resp.Body.Close()
	respBody := struct {
		Valid bool `json:"valid"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return fmt.Errorf("couldnt validate username: %v", err)
	}
	if respBody.Valid {
		return nil
	}
	return fmt.Errorf("username already exists")
}
