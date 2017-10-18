/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-10-13 13:47:59
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-17 17:06:46
***********************************************/
package apps

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/george518/PPGo_Api_Demo_Gin/models"
	"github.com/gin-gonic/gin"
)

//获取列表
func MemberList(c *gin.Context) {
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")
	page := 1
	pageSize := 4

	list, n, err := models.ListMember(page, pageSize, filters...)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "SUCCESS",
			"data":      list,
			"count":     n,
			"page_size": pageSize,
			"current:":  page,
		})
	}
}

//获取一个会员
func MemberGet(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))

	mem, err := models.OneMember(mid)
	fmt.Println(mem, err)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		//log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    mem,
		})
	}

}

//修改会员信息
func MemberEdit(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))
	m := new(models.Member)
	m.Id = mid
	m.LoginName = c.Request.FormValue("login_name")
	m.Password = c.Request.FormValue("password")

	if n, err := m.UpdateMember(mid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}

//新增
func MemberAdd(c *gin.Context) {
	m := new(models.Member)
	m.LoginName = c.Request.FormValue("login_name")
	m.Password = c.Request.FormValue("password")

	if id, err := m.AddMember(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		m.Id = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    m,
		})
	}

}

//删除
func MemberDelete(c *gin.Context) {
	mid, _ := strconv.Atoi(c.Param("id"))

	if n, err := models.DeleteMember(mid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}
