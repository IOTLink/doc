package models

import (
	postgres "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = "123456"
	DB_NAME     = "fabric"
	DB_PORT     = 5432
	DB_HOST     = "127.0.0.1"
)

type ManageDB struct {
	Dbuser string
	Dbpasswd string
	Dbname string
	Dbport uint32
	Dbhost string
	Database  *postgres.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (db *ManageDB) InitDB(user string, pwd string, dbname string, port uint32, host string) {
	if user != "" {
		db.Dbuser = user
	} else {
		db.Dbuser = DB_USER
	}

	if pwd != "" {
		db.Dbpasswd = pwd
	} else {
		db.Dbpasswd = DB_PASSWORD
	}

	if dbname != "" {
		db.Dbname = dbname
	} else {
		db.Dbname = DB_NAME
	}

	if port != 0 {
		db.Dbport = port
	} else {
		db.Dbport = DB_PORT
	}

	if host != "" {
		db.Dbhost = host
	} else {
		db.Dbhost = DB_HOST
	}

	db.Database = nil
}

func (db *ManageDB) RegisterDB() (error){
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		db.Dbuser, db.Dbpasswd, db.Dbname)
	dbase, err := postgres.Open("postgres", dbinfo)
	if err != nil {
		log.Printf("connecting postgresql abort %s!", error.Error(err))
		return err
	} else {
		log.Printf("connecting postgresql success!")
	}
	db.Database = dbase

	err = db.Database.Ping()
	if err != nil {
		log.Printf("connecting postgresql failt!")
	}
	return nil
}

func (db *ManageDB) UnRegisterDB() {
	log.Printf("close postgresql success!")
	db.Database.Close()
}

func (db *ManageDB) InsertAppInfo(appid string, appkey string, registime string) (int, error){
	var lastInsertId int
	err := db.Database.QueryRow("INSERT INTO app_reg_tab(appid, appkey, registime) VALUES($1,$2,$3) returning id;", appid, appkey, registime).Scan(&lastInsertId)
	if err != nil {
		log.Printf("insert error: %s", error.Error(err))
		return 0, err
	}
	log.Printf("last inserted id = %d", lastInsertId)
	return lastInsertId, nil
}

func (db *ManageDB) QueryAppInfo(appid string) (string, string, error){
	sql := fmt.Sprintf("select appkey,registime from app_reg_tab where appid = '%s'", appid)
	err := db.Database.Ping()
	if err != nil {
		log.Printf("connecting postgresql failt!")
	}
	rows, err := db.Database.Query(sql)
	if err != nil {
		log.Printf("query error: %s", error.Error(err))
		return "", "", err
	}
	for rows.Next() {
		var appkey string
		var timestamp string
		err = rows.Scan(&appkey, &timestamp)
		if err != nil {
			log.Printf("query error: %s", error.Error(err))
			return "", "", err
		}
		return appkey,timestamp, nil
	}
	return "", "", nil
}


func (db *ManageDB) IsExist(appid string) (bool, error) {
	sql := fmt.Sprintf("select id from app_reg_tab where appid = '%s'", appid)
	log.Printf("Query %s",sql)
	if db.Database == nil {
		log.Printf("db.Database is nil")
	}
	rows, err := db.Database.Query(sql)
	if err != nil {
		log.Printf("IsExist query error: %s", error.Error(err))
		return false, err
	}
	if rows == nil {
		log.Printf("appid is not exist")
		return false, err
	}
	return rows.Next(),nil
}

/*
//https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html

 */














