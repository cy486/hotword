package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	orm2 "github.com/astaxie/beego/orm"
	"hotword/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}
func (c *MainController) ShowTab() {
	c.TplName = "tab.html"
}

//按照分类查询数据
func (c *MainController) FindWord() {
	is_a := c.GetString("is_a")
	fmt.Println(is_a)
	orm := orm2.NewOrm()
	var word []models.HotWord
	_, err := orm.QueryTable("HotWord").Filter("IsA", is_a).All(&word)
	if err != nil {
		beego.Info("查询失败")
		return
	}
	c.Data["words"] = word
	c.TplName = "word.html"
}

//根据word查询url
func (c *MainController) FindUrl() {
	id, err := c.GetInt("id")
	orm := orm2.NewOrm()
	fmt.Println(id)
	var word []models.WordContent
	_, err = orm.QueryTable("WordContent").Filter("Word", id).All(&word)
	if err != nil {
		beego.Info("查询url错误")
		beego.Info(err)
	}
	c.Data["words"] = word
	c.TplName = "urls.html"
}

//查询词云
func (c *MainController) ShowCloud() {
	c.TplName = "cloud.html"
}
func (c *MainController) ShowUrl() {
	url := c.GetString("Url")
	fmt.Println(url)
	c.Redirect(url, 302)
}
func (c *MainController) Search() {
	key := c.GetString("top-search")
	orm := orm2.NewOrm()
	var word models.HotWord
	_, err := orm.QueryTable("HotWord").Filter("Word", key).All(&word)
	if err != nil {
		beego.Info("查询失败")
		return
	}
	c.Data["words"] = word
	c.TplName = "detail.html"
}
func (c *MainController) ShowGuanxi() {
	c.TplName = "guanxi.html"
}
func (c *MainController) ShowAdd() {
	c.TplName = "add.html"
}
func (c *MainController) AddNewWord() {
	word := c.GetString("word")
	fmt.Println(word)
	orm := orm2.NewOrm()
	var hotword models.HotWord
	hotword.Word = word
	content := con(word)
	hotword.Content = orm2.TextField(content)
	hotword.IsA = shouzimu(word)
	_, err := orm.Insert(&hotword)
	if err != nil {
		beego.Info("添加失败")
	}
	_ = orm.Read(&hotword)
	findcontent(hotword.Word /*page, */, hotword.Id)
	c.Data["words"] = hotword
	c.TplName = "detail.html"
}
func (c *MainController) DownLoad() {
	c.Ctx.Output.Download("./file/hotword.docx")
	c.TplName = "index.html"
}
