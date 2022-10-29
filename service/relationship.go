package service

import (
    "Qianshou/dao"
    "Qianshou/model"
)

func SelectAllRelationshipsByUserId(userId string) (rs []model.Relationship, err error) {
    return dao.NewRelationshipDaoInstance().SelectAllRelationshipsByUserId(userId)
}

func UpdateRelationship(userId string, otherUserId string, state model.State) (r model.Relationship, err error) {
    r.UserId = userId
    r.OtherUserId = otherUserId
    r.State = state
    r.Type = model.RELATIONSHIP
    r.State, err = dao.NewRelationshipDaoInstance().UpdateRelationship(r)
    return
}
