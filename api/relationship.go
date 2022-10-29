package api

import (
    "Qianshou/model"
    "Qianshou/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

func SelectAllRelationshipsByUserId(c *gin.Context) {
    userId := c.Param("user_id")
    rs, err := service.SelectAllRelationshipsByUserId(userId)

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
        Data: rs,
    })
}

func UpdateRelationship(c *gin.Context) {
    userId := c.Param("user_id")
    otherUserId := c.Param("other_user_id")
    json := make(map[string]interface{})
    err := c.Bind(&json)
    if err != nil {
        c.JSON(http.StatusOK, &model.Data{
            Code: -1,
            Msg:  err.Error(),
        })
        return
    }

    r, err := service.UpdateRelationship(userId, otherUserId, model.State(json["state"].(string)))

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
        Data: r,
    })
}
