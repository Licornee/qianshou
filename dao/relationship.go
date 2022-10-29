package dao

import (
    "Qianshou/model"
    "errors"
    "gorm.io/gorm"
    "sync"
)

type RelationshipDao struct {
}

var (
    relationshipDao  *RelationshipDao
    relationshipOnce sync.Once // 保证 dao 对象单例
)

func NewRelationshipDaoInstance() *RelationshipDao {
    relationshipOnce.Do(
        func() {
            relationshipDao = &RelationshipDao{}
        })
    return relationshipDao
}

// 查询某用户的所有关系
func (dao *RelationshipDao) SelectAllRelationshipsByUserId(userId string) (rs []model.Relationship, err error) {
    if err = Db.Where("user_id = ?", userId).Find(&rs).Error; err != nil {
        return
    }
    return
}

// 更新关系
func (dao *RelationshipDao) UpdateRelationship(r model.Relationship) (state model.State, err error) {
    userId := r.UserId
    otherUserId := r.OtherUserId
    or := &model.Relationship{}

    // 开启事务
    err = Db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Where("user_id = ? AND other_user_id = ?", otherUserId, userId).Find(&or).Error; err != nil {
            if !errors.Is(err, gorm.ErrRecordNotFound) {
                return err
            }
            // or 不存在，直接创建/更新 r
            tx.Set("gorm:insert_option", "ON CONFLICT").Create(&r)
        } else { // or 存在，根据 or 的 state 创建 / 更新 r
            if or.State == model.LIKED && r.State == model.LIKED {
                r.State = model.MATCHED
                tx.Model(&or).Update("state", model.MATCHED)
            } else if or.State == model.MATCHED && r.State == model.DISLIKED {
                tx.Model(&or).Update("state", model.LIKED)
            }

            tx.Set("gorm:insert_option", "ON CONFLICT").Create(&r)
        }

        // 返回 nil 提交事务
        return nil
    })

    if err != nil {
        return "", err
    }

    return r.State, nil
}
