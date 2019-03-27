package router

import (
	"github.com/angao/gin-xorm-admin/controller"
	"github.com/angao/gin-xorm-admin/router/middlewares"
	"github.com/angao/gin-xorm-admin/utils"
	ginsessions "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// Init 路由
func Init(port string) {
	router := gin.New()

	cookieStore := sessions.NewCookieStore([]byte("jDIkFg6ju7kEM7DOIWGcXSLwCL6QaMZy"))
	store := utils.NewStore(cookieStore)
	router.Use(ginsessions.Sessions("session_id", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.ErrorHandler)

	router.NoRoute(middlewares.NoRoute)

	router.Static("/public", "public")
	// router.HTMLRender = utils.LoadTemplates("views")

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
		userGroup.POST("/add", user.Add)

		userGroup.GET("/user_edit/:id", user.ToEdit)
		userGroup.POST("/edit", user.Edit)

		userGroup.POST("/delete", user.Delete)

		userGroup.GET("/user_chpwd", user.ToChangePasswd)
		userGroup.POST("/changePwd", user.ChangePwd)
		userGroup.POST("/reset", user.Reset)

		userGroup.GET("/role_assign/:id", user.ToRoleAssign)
		userGroup.POST("/setRole", user.SetRole)

		userGroup.POST("/freeze", user.Freeze)
		userGroup.POST("/unfreeze", user.UnFreeze)
	}

	// dept /dept/
	dept := new(controller.DeptController)
	deptGroup := router.Group("/dept", middlewares.Auth())
	{
		deptGroup.GET("/", dept.Index)
		deptGroup.POST("/list", dept.List)
		deptGroup.POST("/tree", dept.Tree)

		deptGroup.GET("/dept_add", dept.ToAdd)
		deptGroup.POST("/add", dept.Add)

		deptGroup.GET("/dept_update/:deptId", dept.ToEdit)
		deptGroup.POST("/update", dept.Edit)

		deptGroup.POST("/delete", dept.Delete)
	}

	// role /role/
	role := new(controller.RoleController)
	roleGroup := router.Group("/role", middlewares.Auth())
	{
		roleGroup.GET("/", role.Index)
		roleGroup.POST("/list", role.List)

		roleGroup.GET("/role_add", role.ToAdd)
		roleGroup.POST("/add", role.Add)

		roleGroup.GET("/role_edit/:roleId", role.ToEdit)
		roleGroup.POST("/edit", role.Edit)

		roleGroup.POST("/remove", role.Remove)

		roleGroup.GET("/role_assign/:roleId", role.ToAssign)

		roleGroup.POST("/roleTreeList", role.TreeList)
		roleGroup.POST("/roleTreeListByUserId/:userId", role.RoleTreeListByUserID)

		roleGroup.POST("/setAuthority", role.SetAuthority)
	}

	menu := new(controller.MenuController)
	menuGroup := router.Group("/menu", middlewares.Auth())
	{
		menuGroup.GET("/", menu.Index)
		menuGroup.POST("/list", menu.List)

		menuGroup.GET("/menu_add", menu.ToAdd)
		menuGroup.POST("/add", menu.Add)

		menuGroup.GET("/menu_edit/:menuId", menu.ToEdit)
		menuGroup.POST("/edit", menu.Edit)

		menuGroup.POST("/remove", menu.Remove)

		menuGroup.POST("/selectMenuTreeList", menu.SelectMenuTreeList)
		menuGroup.POST("/menuTreeListByRoleId/:roleId", menu.TreeListByRoleID)
	}

	notice := new(controller.NoticeController)
	noticeGroup := router.Group("/notice", middlewares.Auth())
	{
		noticeGroup.GET("/", notice.Index)
		noticeGroup.POST("/list", notice.List)

		noticeGroup.GET("/notice_add", notice.ToAdd)
		noticeGroup.POST("/add", notice.Add)

		noticeGroup.GET("/notice_update/:noticeId", notice.ToEdit)
		noticeGroup.POST("/update", notice.Edit)

		noticeGroup.POST("/delete", notice.Delete)
	}

	router.Run(":" + port)
}
