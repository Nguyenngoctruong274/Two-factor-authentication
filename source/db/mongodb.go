package db

import (
	"time"

	"github.com/ioVN/database"
)

type DB struct {
	AccountDb *AccountDB
	CompanyDb *JobDecriptionDB
}

func NewConnection(uri string, dbName string, dbkey []byte, timeout time.Duration) *DB {
	var (
		isLoading      = true
		loadingBarExit = func(err error) {
			var (
				msg string
			)
			if err != nil {
				msg = "Unable to connect"
			} else {
				msg = "Connected"
			}

			isLoading = false
			time.Sleep(500 * time.Millisecond)
			print("\r[MONGO-DB-] ", msg)
			println("              \n")
			if err != nil {
				panic(err)
			}
		}
	)

	go func() {
		for dot := ""; isLoading; func() {
			time.Sleep(500 * time.Millisecond)
			if len(dot) < 3 {
				dot += "."
			} else {
				dot = ""
			}
		}() {
			print("\r[MONGO-DB-] connecting", dot, "   ")
		}
	}()
	// init connection
	conn, err := database.MongoConnect(uri, dbName, timeout)
	if err != nil {
		println("\r[MONGO-DB-] MongoConnect to :", uri, ", dbname :", dbName, ", timeout :", timeout)
		loadingBarExit(err)
	} else {
		loadingBarExit(nil)
	}
	return &DB{
		AccountDb: NewAccoutDB(conn),
		CompanyDb: NewJobDecriptionDB(conn),
	}

}
