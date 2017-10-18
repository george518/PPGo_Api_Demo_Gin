/**********************************************
** @Des: 入口文件
** @Author: haodaquan
** @Date:   2017-10-13 10:46:12
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-18 10:57:07
***********************************************/
package main

import (
	db "github.com/george518/PPGo_Api_Demo_Gin/dbs"
	"github.com/george518/PPGo_Api_Demo_Gin/libs"
	"github.com/george518/PPGo_Api_Demo_Gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	defer db.Conns.Close()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	router.Run(":" + libs.Conf.Read("site", "httpport"))
}
