package service

import (
    "Qianshou/dao"
    "Qianshou/model"
)

func SelectAllUsers() ([]model.User, error) {
    return dao.NewUserDaoInstance().SelectAllUsers()
}

func CreateUser(name string) (user model.User, err error) {
    user.Type = model.USER
    user.Name = name
    user.Id, err = dao.NewUserDaoInstance().CreateUser(user)
    return
}
