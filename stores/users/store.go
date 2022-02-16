package users

import "database/sql"

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {

}
