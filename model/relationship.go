package model

type State string
type RelationshipType string

const (
    LIKED        State            = "liked"
    DISLIKED     State            = "disliked"
    MATCHED      State            = "matched"
    RELATIONSHIP RelationshipType = "relationship"
)

type Relationship struct {
    UserId      string           `json:"-"`
    OtherUserId string           `json:"user_id"`
    State       State            `json:"state"`
    Type        RelationshipType `json:"type"`
}

func (Relationship) TableName() string {
    return "qianshou.relationship"
}
