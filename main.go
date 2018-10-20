package main

import (
	_ "ItcastCms/routers"
	"github.com/astaxie/beego"
	_"ItcastCms/models"
	"ItcastCms/models"
)

//判断用户是否有权限。
func CheckUserAction(userExtActions []models.UserAction,roleId int)(b bool)  {
	b=false
	for i:=0;i<len(userExtActions) ;i++  {
		if userExtActions[i].Actions.Id==roleId {
			b=true
			break
		}
	}
	return
}
//判断是否禁止还是允许
func CheckPassUserAction(userExtActions []models.UserAction,roleId int)(b bool)  {
	b=false
	for i:=0;i<len(userExtActions) ;i++  {
		if userExtActions[i].Actions.Id==roleId {
			if userExtActions[i].IsPass==1 {
				b=true
			}
			break
		}
	}
	return

}

//判断用户的角色
func CheckUserRole(userExtRoles []*models.RoleInfo,roleId int)(b bool)  {
 //判断传递过来的roleId是否在userExtRoles中出现。
 	b=false
	for i:=0;i<len(userExtRoles) ;i++  {
		if userExtRoles[i].Id==roleId {
			b=true
			break
		}
	}
	return
}
//判断角色对应的权限
func CheckRoleAction(roleExtActions []*models.ActionInfo,actionId int)(b bool)  {
	b=false
	for i:=0; i<len(roleExtActions);i++  {
		if roleExtActions[i].Id==actionId{
			b=true
			break
		}
	}
	return
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true//启用SESSION

	beego.AddFuncMap("checkUserAction",CheckUserAction)
	beego.AddFuncMap("checkPassUserAction",CheckPassUserAction)
	beego.AddFuncMap("checkUserRole",CheckUserRole)
	beego.AddFuncMap("checkRoleAction",CheckRoleAction)
	beego.Run()
}

