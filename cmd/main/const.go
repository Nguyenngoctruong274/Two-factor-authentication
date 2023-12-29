package main

import (
	"authentication/source/db"
	"time"
)

func init() {
	db.AuthenConnection = db.NewConnection(AuthenURI, AuthenDB, nil, 15*time.Second)
}
