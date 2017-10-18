/**********************************************
** @Des: 权限检查
** @Author: haodaquan
** @Date:   2017-10-16 21:49:31
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-18 09:47:11
***********************************************/
package middlewares

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/george518/PPGo_Api_Demo_Gin/libs"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.FormValue("app_key")
		sign := c.Request.FormValue("sign")
		ts := c.Request.FormValue("ts")
		method := c.Request.Method
		now := time.Now().Unix()

		if key != libs.Conf.Read("api", "apikey") {
			noAuth(c, "Key not found")
			return
		}
		//时差两秒返回无权 注意修改
		time_check, _ := strconv.Atoi(ts)
		if (now - int64(time_check)) > 100000000000 {
			noAuth(c, "Time out")
			return
		}

		if Sign(key, ts, method, sign) == false {
			noAuth(c, "Unauthorized")
			return
		}
		c.Next()
		return
	}
}

func noAuth(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
	c.Abort()
}

func Sign(key, ts, method, sign string) bool {
	secret := libs.Conf.Read("api", "apisecrect")

	b := bytes.Buffer{}
	b.WriteString("app_key=")
	b.WriteString(key)
	b.WriteString("&app_secret=")
	b.WriteString(secret)
	b.WriteString("&method=")
	b.WriteString(method)
	b.WriteString("&ts=")
	b.WriteString(ts)
	if libs.Md5([]byte(b.String())) == sign {
		return true
	}
	return false
}
