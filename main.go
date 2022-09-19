package main

import (
	"fmt"

	dbs "example.com/database"
	vs "example.com/resetpassword"
)

func main() {
	db := dbs.Connect()
	vs.HandleFunc()

	// dbb := vs.ResetPassword()
	fmt.Println(db.Exec("select * from user_table"))
	// fmt.Println(dbb)
	// gonfig.Connect()
}
