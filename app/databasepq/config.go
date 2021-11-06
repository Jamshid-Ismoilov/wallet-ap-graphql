package databasepq

import (
	"fmt"
)

var (
	host     = "localhost"
	user     = "jamshid"
	password = "1111"
	dbname   = "task_1"
	port     = 5432
)

var DB_CONFIG = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d",
	host, user, password, dbname, port,
)
