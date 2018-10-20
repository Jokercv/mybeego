package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ItcastCms/models"
)

type ArticleClassController struct {
	beego.Controller
}

func (this *ArticleClassController) Index() {
	this.TplName="ArticleClass/Index.html"
}
func (this *ArticleClassController) ShowArticleClass() {
	o:=orm.NewOrm()
	var articleClasses []models.ArticelClass
	o.QueryTable("article_class").Filter("parent_id").All(&articleClasses)
	this.Data["json"]=map[string]interface{}{"rows":articleClasses}
	this.ServeJSON()
}