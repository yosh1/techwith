
package common

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func AuthReq(token string) gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println(token)
        c.Next()
    }
}
