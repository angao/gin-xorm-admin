package router

import (
	"html/template"

	"github.com/angao/gin-xorm-admin/controller"
	"github.com/angao/gin-xorm-admin/router/middlewares"
	"github.com/angao/gin-xorm-admin/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Init 路由
func Init() {
	router := gin.New()

	store := sessions.NewCookieStore([]byte("--secret--key--"))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		MaxAge:   3600,
	})
	router.Use(sessions.Sessions("session_id", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandler)

	router.NoRoute(middlewares.NoRoute)

	router.Static("/public", "public")
	// router.HTMLRender = utils.LoadTemplates("views")

	router.SetFuncMap(template.FuncMap{
		"formatAsDate": utils.FormatDate,
	})

	// login authentication
	auth := new(controller.AuthController)
	router.GET("/login", auth.ToLogin)
	router.POST("/login", auth.Login)
	router.GET("/logout", auth.Logout)

	index := new(controller.IndexController)
	router.GET("/", middlewares.Auth(), index.Home)
	router.GET("/blackboard", middlewares.Auth(), index.BlackBoard)

	// user /mgr/
	user := new(controller.UserController)
	userGroup := router.Group("/mgr", middlewares.Auth())
	{
		userGroup.GET("/", user.Home)
		userGroup.POST("/list", user.List)
		userGroup.GET("/user_info", user.Info)
		userGroup.GET("/user_add", user.ToAdd)
		userGroup.GET("/user_edit/:id", user.ToEdit)
		userGroup.GET("/role_assign/:id", user.ToRoleAssign)

		userGroup.POST("/add", user.Add)
		userGroup.POST("/delete", user.Delete)
		userGroup.POST("/reset", user.Reset)
		userGroup.POST("/freeze", user.Freeze)
		userGroup.POST("/unfreeze", user.UnFreeze)
		userGroup.POST("/setRole", user.SetRole)
	}

	// dept /dept/
	dept := new(controller.DeptController)
	deptGroup := router.Group("/dept", middlewares.Auth())
	{
		deptGroup.POST("/tree", dept.List)
	}

	// role /role/
	role := new(controller.RoleController)
	roleGroup := router.Group("/role", middlewares.Auth())
	{
		roleGroup.GET("/", role.Index)
		roleGroup.POST("/list", role.List)
		roleGroup.GET("/role_add", role.ToAdd)
		roleGroup.POST("/roleTreeListByUserId/:userId", role.RoleTreeListByUserID)
	}

	menu := new(controller.MenuController)
	menuGroup := router.Group("/menu", middlewares.Auth())
	{
		menuGroup.GET("/", menu.Index)
		menuGroup.POST("/list", menu.List)
		menuGroup.POST("/remove", menu.Remove)

		menuGroup.GET("/menu_add", menu.ToAdd)
		menuGroup.POST("/selectMenuTreeList", menu.SelectMenuTreeList)

		menuGroup.GET("/menu_edit/:menuId", menu.Edit)
	}
	router.Run(":3000")
}
