package main

import db "github.com/PalmerTurley34/inventory-tracking/internal/database"

type errMsg struct {
	err error
}

type userLoggedInMsg struct {
	userInfo db.User
}

type userLoggedOutMsg struct{}

type userCreatedMsg struct{}
