package service

import (
	"douyinProject/entity"
	"douyinProject/repository"
)

func FollowAction(currentUserId int64, toUserId int64, actionType int) error {

	if actionType == 1 {
		err := repository.FollowOthers(currentUserId, toUserId)
		if err != nil {
			return err
		}
	} else {
		err := repository.UnFollowOthers(currentUserId, toUserId)
		if err != nil {
			return err
		}
	}
	return nil
}

func FollowListAction(toUserId int64) (list []entity.User, err error) {
	list, err = repository.QueryFollowListById(toUserId)
	return
}

func FollowerListAction(toUserId int64) (list []entity.User, err error) {
	list, err = repository.QueryFollowerListById(toUserId)
	return
}
