package controllers

import (
	"fmt"
	"os"
	"time"
	"tutu/models"

	"reflect"

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
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	var maps []orm.Params
	var Images models.Images
	//res, err := o.QueryTable("Article").Filter("id", "name").All(&Article)
	num, err := o.QueryTable("Images").Values(&maps, "id", "name", "Path", "Type_name", "Des")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(Images.Id, Images.Name)
		fmt.Println("res是:", reflect.TypeOf(maps))
		fmt.Println(num)
		for _, v := range maps {

			fmt.Println(reflect.TypeOf(v))
		}
		c.Data["m"] = maps
	}

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
	fmt.Println(jsoninfo)
	if jsoninfo == "" {
		fmt.Println("jsoninfo is empty")
	}

	desinfo := c.GetString("desname")
	fmt.Println(desinfo)
	if desinfo == "" {
		fmt.Println("des is empty")
	}

	// 获取文件名称并保存文件
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
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
