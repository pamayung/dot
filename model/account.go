package model

import (
	"dot/core/cache"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Account struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `gorm:"size:255;not null;" json:"first_name"`
	LastName  string    `gorm:"size:255;not null;" json:"last_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ac *Account) GetByID(db *gorm.DB, id uint32) (*Account, error) {
	var err error
	data := cache.GetCache(id)
	if data != nil {
		result := data.(*Account)
		return result, err
	} else {
		err = db.First(ac, id).Error
		if err != nil {
			return ac, err
		}
		if gorm.IsRecordNotFoundError(err) {
			return ac, errors.New("Not Found")
		}
		cache.SetCache(id, ac)
		return ac, err
	}
}

func (ac *Account) Save(db *gorm.DB) (*Account, error) {

	var err error
	err = db.Create(ac).Error
	if err != nil {
		return ac, err
	}
	cache.SetCache(ac.ID, ac)
	return ac, nil
}

func (ac *Account) Update(db *gorm.DB, id uint32) (*Account, error) {

	d := db.Model(&Account{}).Where("id = ?", id).Updates(ac)
	if d.Error != nil {
		return ac, d.Error
	}
	d.Find(ac, id)
	cache.SetCache(id, ac)

	return ac, nil
}

func (ac *Account) Delete(db *gorm.DB, id uint32) (int64, error) {

	d := db.Delete(ac, id)

	if d.Error != nil {
		return 0, d.Error
	}
	cache.DeleteCache(id)
	return d.RowsAffected, nil
}

func (ac *Account) Create(db *gorm.DB, mb *Member) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(ac).Error; err != nil {
		tx.Rollback()
		return db.Error
	}
	cache.SetCache(ac.ID, ac)

	mb.AccountID = ac.ID
	if err := tx.Create(mb).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
