// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Author struct {
	ID   int32          `json:"id"`
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}

type Person struct {
	ID sql.NullInt32 `json:"id"`
}