package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Books       BookModel
	Permissions PermissionModel
	Tokens      TokenModel // Add a new Tokens field.
	Users       UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Books:       BookModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db}, // Initialize a new TokenModel instance.
		Users:       UserModel{DB: db},
	}
}
