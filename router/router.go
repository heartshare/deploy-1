package router

import (
    "deploy/api"
    "deploy/router/middleware"
    "github.com/gin-gonic/gin"
)

func InitRouters() (router *gin.Engine) {
    var (
        routerGroup *gin.RouterGroup
    )
    router = gin.Default()
    router.Use(middleware.Cors()) //增加跨域请求头
    routerGroup = router.Group("")

    authRouter := routerGroup.Group("auth")
    {
        authRouter.POST("login", api.Login)
        authRouter.GET("captcha", api.Captcha)
        authRouter.GET("captcha/:captchaId", api.CaptchaImg)
    }

    sysRouter := routerGroup.Group("sys").Use(middleware.JWTAuth())
    {
        sysRouter.GET("getDeployInfo", api.GetDeployServerInfo)
        sysRouter.POST("modifyPwd", api.ModifyPwd)
        sysRouter.POST("getUserList", api.GetUserList)
        sysRouter.POST("saveUser", api.SaveUser)
        sysRouter.POST("setUserStatus", api.SetUserStatus)
    }

    serverRouter := routerGroup.Group("server").Use(middleware.JWTAuth())
    {
        serverRouter.POST("getServerList", api.GetServerList)
        serverRouter.POST("saveServer", api.SaveServer)
        serverRouter.GET("deleteServer", api.DeleteServer)
    }

    repoRouter := routerGroup.Group("project").Use(middleware.JWTAuth())
    {
        repoRouter.POST("getProjectList", api.GetProjectList)
        repoRouter.POST("saveProject", api.SaveProject)
        repoRouter.GET("initProject", api.InitProject)
        repoRouter.GET("delProject", api.DelProject)
    }

    envCfgRouter := routerGroup.Group("envCfg").Use(middleware.JWTAuth())
    {
        envCfgRouter.POST("getEnvCfgList", api.GetEnvCfgList)
        envCfgRouter.POST("saveEnvCfg", api.SaveEnvCfg)
        envCfgRouter.GET("delEnvCfg", api.DelEnvCfg)
        envCfgRouter.GET("getCfgOptions", api.GetCfgOptions)
    }

    taskRouter := routerGroup.Group("task").Use(middleware.JWTAuth())
    {
        taskRouter.POST("getTaskList", api.GetTaskList)
        taskRouter.POST("saveTask", api.SaveTask)
        taskRouter.GET("deleteTask", api.DeleteTask)
        taskRouter.GET("getBranches", api.GetBranches)
        taskRouter.GET("getVersions", api.GetVersions)
        taskRouter.GET("getEnvOptions", api.GetEnvOptions)
        taskRouter.GET("rollBack", api.RollBack)
        taskRouter.GET("deploy", api.Deploy)
        router.GET("ws", api.DeployInfo) //websocket 发布阶段监听
    }
    return
}
