/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-10-13 21:09:02
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-18 09:40:39
***********************************************/
package dbs

import (
	"database/sql"
	"log"

	"github.com/george518/PPGo_Api_Demo_Gin/libs"
	_ "github.com/go-sql-driver/mysql"
)

var Conns *sql.DB

func init() {
	var err error
	username := libs.Conf.Read("mysql", "username")
	password := libs.Conf.Read("mysql", "password")
	dataname := libs.Conf.Read("mysql", "dataname")
	port := libs.Conf.Read("mysql", "port")
	host := libs.Conf.Read("mysql", "host")

	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=true"
	// fmt.Println(dns)
	Conns, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = Conns.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)
}
