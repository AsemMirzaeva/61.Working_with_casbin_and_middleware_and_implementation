package storage

import (
	"database/sql"
	"fmt"
)

func OpenSql(drname, url string) (*sql.DB, error) {
	db, err := sql.Open(drname, url)
	if err != nil {
		return nil, fmt.Errorf("error: db opening %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error: db connection :%v", err)
	}

	return db, nil
}
