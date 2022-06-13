package tests

import (
	"douyin/app/model/testm"
	"douyin/boot/global"
)

type TagService struct{}

func (t *TagService) CreateTag(tags testm.Tags) (err error) {

	err = global.DYDB.Create(&tags).Error
	return err
}
