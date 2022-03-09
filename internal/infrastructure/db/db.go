package db

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hizzuu/plate-backend/conf"
)

func NewMysqlDB() (*sql.DB, error) {
	db, err := sql.Open(conf.C.DB.Dbms, MysqlDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MysqlDSN() string {
	c := &mysql.Config{
		User:                 conf.C.DB.User,
		Passwd:               conf.C.DB.Pass,
		Net:                  conf.C.DB.Net,
		Addr:                 conf.C.DB.Host + ":" + conf.C.DB.Port,
		DBName:               conf.C.DB.Name,
		Loc:                  time.Local,
		ParseTime:            conf.C.DB.ParseTime,
		AllowNativePasswords: conf.C.DB.AllowNativePasswords,
	}
	return c.FormatDSN()
}
