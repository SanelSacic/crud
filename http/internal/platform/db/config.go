package db

import "time"

// DataSource contains username,pssword and database name.
var (
	DataSource         = "root:sonyss-a3@/made"
	MaxIdleConnections = 200
	TimeoutConnection  = 60 * time.Second
)
