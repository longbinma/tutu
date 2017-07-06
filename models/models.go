package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
const (
		_DB_NAME        = "root:123456@/beeblog?charset=utf8"
	_SQLITE3_DRIVER = "mysql"
)
*/

type Images struct {
	Id          int    `orm:"index"`
	Name        string `orm:"index"`
	Path        string
	Type_name   string
	Upload_time time.Time
}

type Utype struct {
	Id          int    `orm:"index"`
	Name        string `orm:"index"`
	Created     time.Time
	Path        string
	Images_path string
}

func RegisterDB() {
	/*
		if !com.IsExist(_DB_NAME) {
			os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
			os.Create(_DB_NAME)
		}
	*/
	orm.RegisterModel(new(Images), new(Utype))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/tutu?charset=utf8", 10, 10)
}
