package models

// UserFriend 表示用户之间的好友关系及亲密度
type UserFriend struct {
	ID        int64   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID    int64   `gorm:"not null;column:user_id" json:"user_id"`
	FriendID  int64   `gorm:"not null;column:friend_id" json:"friend_id"`
	Closeness float64 `gorm:"type:decimal(5,4);not null;column:closeness;comment:亲密度，0-1" json:"closeness"`
}

// TableName 指定表名
func (UserFriend) TableName() string {
	return "t_user_friend"
}
