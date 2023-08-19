package db

import "github.com/edr3x/chi-explore/db/schemas"

func DbSync() {
	DB.AutoMigrate(&schemas.User{})
}
