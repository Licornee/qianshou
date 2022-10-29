package router

import (
    "Qianshou/api"
    "github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/users", api.SelectAllUsers)

    router.POST("/users", api.CreateUser)

    router.GET("/users/:user_id/relationships", api.SelectAllRelationshipsByUserId)

    router.PUT("/users/:user_id/relationships/:other_user_id", api.UpdateRelationship)

    return router
}
