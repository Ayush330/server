package config

import (
	"fmt"
)

const (
	sql_domain       = "localhost"
	sql_username     = "root"
	sql_password     = "newpassword"
	sql_port         = "3306"
	database         = "splitwise"
	MAXM_CONNECTIONS = 20
)

func GetSqlAddress() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sql_username, sql_password, sql_domain, sql_port, database)
}
