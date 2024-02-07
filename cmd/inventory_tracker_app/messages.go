package main

import db "github.com/PalmerTurley34/inventory-tracking/internal/database"

type errMsg struct {
	err error
}

type loginFailMsg struct {
	err error
}

type loginSucessMsg struct {
	userInfo db.User
}

type userLoggedOutMsg struct{}

type userCreateFailMsg struct {
	err error
}

type userCreatedMsg struct {
	userInfo db.User
}
