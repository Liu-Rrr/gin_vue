package utils

//暂时没用
import (
	"ginDemo/src/main/dbUtils"
	"ginDemo/src/main/entiy"
	"github.com/gin-gonic/gin"
	"log"
)

func AddUser(c *gin.Context) error {
	var user entiy.User
	if err := c.ShouldBindJSON(&user).Error; err != nil {
		log.Fatal("error")
	}
	if err := dbUtils.Db.Create(&user); err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, "成功")

	return nil
}
