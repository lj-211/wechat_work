package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println("magnifier - ", string(buf[0:n]))
}
