package dbUtils

import (
	"fmt"
	"ginDemo/src/main/entiy"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/*
import后的下划线的作用即：当导入一个包时，该包下的文件里所有init()函数都会被执行，
		然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候
		就可以使用 import _ " "引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用
		init()函数，所以无法通过包名来调用包中的其他函数。
*/

type ymlConfig struct {
	Ip      string `yaml:"ip"`      //服务器ip
	User    string `yaml:"user"`    //用户
	Pwd     string `yaml:"pwd"`     //用户密码
	Db_name string `yaml:"db_name"` //数据库名字
	Db_port string `yaml:"db_port"` //数据库端口
}

func (c *ymlConfig) getYml() *ymlConfig {
	yamlFile, err := ioutil.ReadFile("src/config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

var AdminT, UserT, Db *gorm.DB

func Init() {
	var c ymlConfig
	var dbConfig string
	config := c.getYml()
	//"root:root123@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dbConfig = config.User + ":" + config.Pwd + "@tcp(" + config.Ip + ":" + config.Db_port + ")/" + config.Db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		panic(err)
	}

	AdminT = db
	UserT = db
	//绑定数据库表名
	UserT.AutoMigrate(&entiy.User{})

}
