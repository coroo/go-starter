package middlewares

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Db(conn sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
