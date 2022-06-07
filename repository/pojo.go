package repository

type dbVideo struct {
	Id            int64
	Authorid      int64
	Playurl       string
	Coverurl      string
	Favoritecount int64
	Commentcount  int64
	Title         string
	Created       int64 `gorm:"column:createtime"`
}

func (dbVideo) TableName() string {
	return "dyvideo"
}

type UserFavorite struct {
	Id     int64
	User   DbUser    `gorm:"ForeignKey:Id"`
	Videos []dbVideo `gorm:"many2many:userFavoriteVideo"`
}

func (UserFavorite) TableName() string {
	return "favoriteList"
}

type DbUser struct {
	Id            int64
	Username      string
	Password      string
	Followcount   int64
	Followercount int64
}

func (DbUser) TableName() string {
	return "dyuser"
}

type DbFollow struct {
	Id       int64 `gorm:"column:id"`
	FollowId int64 `gorm:"column:follow_id"`
}

func (DbFollow) TableName() string {
	return "dyfollow"
}

type DbFollower struct {
	Id         int64 `gorm:"column:id"`
	FollowerId int64 `gorm:"column:follower_id"`
}

func (DbFollower) TableName() string {
	return "dyfollower"
}
