package dao

import (
    "Qianshou/model"
    "sync"
)

type UserDao struct {
}

var (
    userDao  *UserDao
    userOnce sync.Once // 保证 dao 对象单例
)

func NewUserDaoInstance() *UserDao {
    userOnce.Do(
        func() {
            userDao = &UserDao{}
        })
    return userDao
}

// 查询所有用户信息
func (dao *UserDao) SelectAllUsers() (users []model.User, err error) {
    if err = Db.Find(&users).Error; err != nil {
        return
    }
    return
}

// 添加用户
func (dao *UserDao) CreateUser(user model.User) (id string, err error) {
    if err = Db.Create(&user).Error; err != nil {
        return "", err
    }

    return user.Id, nil
}
