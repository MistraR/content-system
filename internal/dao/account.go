package dao

import (
	"content-system/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

func (a *AccountDao) IsExist(UserId string) (bool, error) {
	var account = model.Account{}
	err := a.db.Where("user_id=?", UserId).First(&account).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("AccountDao isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (a *AccountDao) Create(account model.Account) error {
	if err := a.db.Create(&account).Error; err != nil {
		fmt.Printf("AccountDao Create = [%v]", err)
		return err
	}
	return nil
}

func (a *AccountDao) FirstByUserId(userId string) (*model.Account, error) {
	var account model.Account
	err := a.db.Where("user_id=?", userId).First(&account).Error
	if err != nil {
		fmt.Printf("FirstByUserId error = [%v]", err)
		return nil, err
	}
	return &account, nil
}
