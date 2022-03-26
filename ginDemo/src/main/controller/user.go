package controller

import (
	"ginDemo/src/main/dbUtils"
	"ginDemo/src/main/entiy"
	"ginDemo/src/main/utils"
	"github.com/gin-gonic/gin"
	"log"
)

/*
	User & Admin
	gorm.Model
	user string
	age  int
*/

func Register(c *gin.Context) {

	var user entiy.User
	//user.Username = c.GetPostForm()
	//user.Password = c.GetPostForm("password")
	//dbUtils.UserT.Create(&user)
	//utils.Success(c)

	//user.Username, _ = c.GetPostForm("username")
	//user.Password, _ = c.GetPostForm("password")
	//username := c.PostForm("username") //绑定form的值
	//dbUtils.UserT.Create(&user)
	//utils.Success(c)
	//log.Println(username)

	//c.
	//c.BindJSON("username")
	type SU struct {
		message string
		code    string
	}

	//Infor := &SU{
	//	message: "success",
	//	code:    "0",
	//}
	if err := c.ShouldBindJSON(&user); err == nil {
		dbUtils.UserT.Create(&user)
		utils.Success(c)

	} else {
		utils.FailedErr(c, err.Error())
		return
	}

}

func Login(c *gin.Context) {
	var user entiy.User
	var userTemp entiy.User
	if err := c.ShouldBindJSON(&user); err == nil {
		userTemp = user
		//dbUtils.UserT.Create(&user)
		log.Println(userTemp)
		log.Println(user)
		if dbUtils.UserT.Where("username = ?", user.Username).First(&user); userTemp.Username == user.Username && userTemp.Password == user.Password {
			utils.Success(c)
		} else {
			utils.Failed(c)
		}

	} else {
		utils.FailedErr(c, err.Error())
	}

}
