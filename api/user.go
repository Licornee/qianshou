package api

import (
    "Qianshou/model"
    "Qianshou/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SelectAllUsers(c *gin.Context) {
    users, err := service.SelectAllUsers()
    if err != nil {
        c.JSON(http.StatusOK, &model.Data{
            Code: -1,
            Msg:  err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, &model.Data{
        Code: 0,
        Msg:  "Success!",
        Data: users,
    })
}

func CreateUser(c *gin.Context) {
    json := make(map[string]interface{})
    err := c.Bind(&json)
    if err != nil {
        c.JSON(http.StatusOK, &model.Data{
            Code: -1,
            Msg:  err.Error(),
        })
        return
    }

    user, err := service.CreateUser(json["name"].(string))
    if err != nil {
        c.JSON(http.StatusOK, &model.Data{
            Code: -1,
            Msg:  err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, &model.Data{
        Code: 0,
        Msg:  "Success!",
        Data: user,
    })
}
