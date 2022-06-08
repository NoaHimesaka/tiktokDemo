package repository

import (
	"douyinProject/entity"

	"gorm.io/gorm"
)

func FollowOthers(currentUserId int64, toUserId int64) error {
	//db.AutoMigrate()
	follow := DbFollow{Id: currentUserId, FollowId: toUserId}
	follower := DbFollower{Id: toUserId, FollowerId: currentUserId}
	curUser := DbUser{}
	toUser := DbUser{}
	db.Where("id", currentUserId).First(&curUser)
	db.Where("id", toUserId).First(&toUser)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&follow).Error; err != nil {
			return err
		}
		if err := tx.Create(&follower).Error; err != nil {
			return err
		}
		if err := tx.Model(&curUser).Update("followcount", curUser.Followcount+1).Error; err != nil {
			return err
		}
		if err := tx.Model(&toUser).Update("followercount", toUser.Followercount+1).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func UnFollowOthers(currentUserId int64, toUserId int64) error {
	follow := DbFollow{Id: currentUserId, FollowId: toUserId}
	follower := DbFollower{Id: toUserId, FollowerId: currentUserId}
	curUser := DbUser{}
	toUser := DbUser{}
	db.Where("id", currentUserId).First(&curUser)
	db.Where("id", toUserId).First(&toUser)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&curUser).Update("followcount", curUser.Followcount-1).Error; err != nil {
			return err
		}
		if err := tx.Model(&toUser).Update("followercount", toUser.Followercount-1).Error; err != nil {
			return err
		}
		if err := tx.Delete(&follow).Error; err != nil {
			return err
		}
		if err := tx.Delete(&follower).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func QueryFollowListById(userId int64) ([]entity.User, error) {
	sub := []DbFollow{}
	db.Model(&DbFollow{}).Where("id = ?", userId).Find(&sub)
	results := make([]entity.User, len(sub))
	for i, v := range sub {
		user := DbUser{}
		err := db.Model(&DbUser{}).Where("id = ?", v.FollowId).Find(&user).Error
		if err != nil {
			return nil, err
		}
		results[i] = copyUser(user)
	}
	return results, nil
}
func QueryFollowerListById(userId int64) ([]entity.User, error) {
	sub := []DbFollower{}
	db.Model(&DbFollower{}).Where("id = ?", userId).Find(&sub)
	results := make([]entity.User, len(sub))
	for i, v := range sub {
		user := DbUser{}
		err := db.Model(&DbUser{}).Where("id = ?", v.FollowerId).Find(&user).Error
		if err != nil {
			return nil, err
		}
		results[i] = copyUser(user)
	}
	return results, nil
}
