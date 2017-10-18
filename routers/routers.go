/**********************************************
** @Des: 路由
** @Author: haodaquan
** @Date:   2017-10-13 10:58:26
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-18 10:59:35
***********************************************/

package routers

import (
	"net/http"

	"github.com/george518/PPGo_Api_Demo_Gin/apps"
	"github.com/george518/PPGo_Api_Demo_Gin/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	// router.Static("/static", "/static")
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", apps.IndexApi)
	v0 := router.Group("/v0")
	v0.Use(middlewares.Auth())
	{
		//新增
		//curl -X POST http://127.0.0.1:8000/v0/member -d "login_name=hell31&password=g2223"
		v0.POST("/member", apps.MemberAdd)
		// curl -X GET http://127.0.0.1:8000/v0/member
		v0.GET("/member", apps.MemberList)
		// curl -X GET http://127.0.0.1:8000/v0/member/1
		v0.GET("/member/:id", apps.MemberGet)
		//curl -X PUT http://127.0.0.1:8000/v0/member/1 -d "login_name=haodaquan&password=1234"
		v0.PUT("/member/:id", apps.MemberEdit)
		// curl -X DELETE http://127.0.0.1:8000/v0/member/2
		v0.DELETE("/member/:id", apps.MemberDelete)
	}

	return router
}
