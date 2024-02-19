package handle

import "github.com/gin-gonic/gin"

func GetAllCash(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Gabriel",
	})
}

func GetIdCash(c *gin.Context) {
	movimentacao := c.Params.ByName("movimentacao")
	c.JSON(200, gin.H{
		"movimentacao": movimentacao,
	})
}
