// A service on gin 1.9. gin 1.10 changed Default() and New() to take
// variadic options: every existing call still compiles, so this is the
// change the risk gate should call medium and point at, not block.
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) { c.String(200, "ok") })
	_ = r.Run(":8080")
}
