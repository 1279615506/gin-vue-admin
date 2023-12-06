package service

import (
	"errors"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthorityService struct{}

//
//func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
//	if err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
//		return auth, ErrRoleExistence
//	}
//	e := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
//		if err = tx.Create(&auth).Error; err != nil {
//			return err
//		}
//
//	})
//}
