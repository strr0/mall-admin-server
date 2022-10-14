package config

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	casbinUtil "github.com/casbin/casbin/v2/util"
	"github.com/casbin/json-adapter/v2"
	"github.com/gin-gonic/gin"
	"log"
	"mall-admin-server/auth"
	"mall-admin-server/util"
	"net/http"
)

var enforcer *casbin.Enforcer
var exclude = []string{
	"/admin/register",
	"/admin/login",
	"/admin/refreshToken",
	"/admin/logout",
}

// 初始化模型
func getModel() model.Model {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	//m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj")
	return m
}

// 初始化权限配置
func getAdapter() *jsonadapter.Adapter {
	b := []byte(`[]`)
	a := jsonadapter.NewAdapter(&b) // Use b as the data source.
	return a
}

func InitPolicy() {
	db := GetDb()
	var p []map[string]interface{}
	db.Raw("select r.name role_name, m.name menu_name from ums_role_menu_relation rmr "+
		"left join ums_role r on rmr.role_id = r.id "+
		"left join ums_menu m on rmr.menu_id = m.id "+
		"where m.level = ?", 1).Scan(&p)
	for _, policy := range p {
		_, _ = enforcer.AddPolicy(policy["role_name"], policy["menu_name"], "*")
	}
	var g []map[string]interface{}
	db.Raw("select a.username admin_name, r.name role_name from ums_admin_role_relation arr " +
		"left join ums_admin a on arr.admin_id = a.id " +
		"left join ums_role r on arr.role_id = r.id").Scan(&g)
	for _, policy := range g {
		_, _ = enforcer.AddRoleForUser(policy["admin_name"].(string), policy["role_name"].(string))
	}
}

func init() {
	var err error
	m := getModel()
	a := getAdapter()
	enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("casbin failed: %v", err)
	}
}

func AuthCheckRole(menu string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 忽略认证
		for _, uri := range exclude {
			if casbinUtil.KeyMatch2(ctx.Request.URL.Path, uri) {
				ctx.Next()
				return
			}
		}
		// token认证
		tokenString := ctx.Request.Header.Get("Authorization")
		_, user := auth.ParseToken(tokenString)
		if user == "" {
			ctx.JSON(http.StatusOK, util.Failed("权限不足"))
			ctx.Abort()
			return
		}
		result, err := enforcer.Enforce(user, menu, "*")
		if err != nil {
			ctx.JSON(http.StatusOK, util.Failed("系统异常"))
			ctx.Abort()
			return
		}
		if result {
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusOK, util.Failed("权限不足"))
		ctx.Abort()
	}
}
