package controllers

import (
	"github.com/astaxie/beego"
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
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	if err != nil {
		// fmt.Println("getfile err ", err)
	} else {
		//fmt.Println(h.Filename)
		c.SaveToFile("uploadname", "static/img/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
		c.Data["Filename"] = h.Filename
		c.TplName = "file.html"
	}
}
