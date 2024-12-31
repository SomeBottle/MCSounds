/*
Http 路由

返回 JSON 格式，必有字段：
  - "msg": string, 响应信息
  - "data": interface{}, 数据
*/
package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":  MsgWelcome,
			"data": []any{},
		})
	})
	return r
}
