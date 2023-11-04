package initDB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Db *gorm.DB

// 链接数据库
func init() {
	var err error
	//dataSourceName: 数据库的连接信息, 这个连接包括了数据库的用户名、密码、数据库主机以及连接的数据库名等信息
	//用户名:密码@协议(地址:端口)/数据库?参数=参数值
	Db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ginstablediffusion")
	if err != nil {
		log.Panicln("err: ", err.Error())
	}
	//防止表后面自动加上了一个s，如article会变成articles
	Db.SingularTable(true)
}
