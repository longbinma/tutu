package controllers

import (
	"fmt"
	"os"
	"time"
	"tutu/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

type Upload struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *Upload) Get() {
	/*
		var lists []orm.ParamsList
		num, err = o.Raw("SELECT user_name FROM user WHERE status = ?", 1).ValuesList(&lists)
		if err == nil && num > 0 {
			fmt.Println(lists[0][0]) // slene
		}
	*/
	c.TplName = "upload.html"
}

func (c *Upload) Post() {
	// 获取下拉菜单内容
	var Type_path string
	jsoninfo := c.GetString("type")
	if jsoninfo == "" {
		fmt.Println("jsoninfo is empty")
	}

	desinfo := c.GetString("des")
	fmt.Println(desinfo)
	if desinfo == "" {
		fmt.Println("des is empty")
	}

	// 获取文件名称并保存文件
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	if err != nil {
		// fmt.Println("getfile err ", err)
	} else {
		//fmt.Println(h.Filename)

		Type_path = "static/" + jsoninfo
		err := os.Mkdir(Type_path, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("创建目录" + Type_path + "成功")
		}

		c.SaveToFile("uploadname", Type_path+"/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
		c.Data["Filename"] = h.Filename
		c.TplName = "file.html"
	}

	//创建一个mysql链接
	o := orm.NewOrm()
	images := models.Images{
		Name:        h.Filename,
		Path:        Type_path + "/" + h.Filename,
		Type_name:   jsoninfo,
		Des:         desinfo,
		Upload_time: time.Now(),
	}

	// insert 插入数据到mysql
	o.Begin()
	id, err := models.ImagesAdd(o, images)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err != nil {
		o.Rollback()
		fmt.Println("插入images表出错,事务回滚")
	} else {
		o.Commit()
		fmt.Println("插入images表成功,事务提交")
	}
}
