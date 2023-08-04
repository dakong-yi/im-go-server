package middleware

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

// 自定义的中间件，用于记录 POST 请求的请求体内容
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 打印请求方法和路径
		fmt.Println("请求:", c.Request.Method, c.Request.URL.Path)

		// 打印请求头
		fmt.Println("请求头:")
		for key, values := range c.Request.Header {
			for _, value := range values {
				fmt.Println(key+":", value)
			}
		}

		// 打印查询参数
		fmt.Println("查询参数:")
		queryParams := c.Request.URL.Query()
		for key, values := range queryParams {
			for _, value := range values {
				fmt.Println(key+":", value)
			}
		}
		// 从请求中获取请求体内容
		body, _ := c.GetRawData()

		// 打印请求体内容
		fmt.Println("请求体内容:", string(body))

		// 将请求体内容重新设置到请求中，以便后续处理程序使用
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// 继续处理其他中间件和请求处理程序
		c.Next()
	}
}
