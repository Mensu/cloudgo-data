package entities

//UserInfoOrmAtomicService .
type UserInfoOrmAtomicService struct{}

//UserInfoOrmService .
var UserInfoOrmService = UserInfoOrmAtomicService{}

// Save .
func (*UserInfoOrmAtomicService) Save(u *UserInfo) error {
	tx := gormDb.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// FindAll .
func (*UserInfoOrmAtomicService) FindAll() []UserInfo {
	ulist := make([]UserInfo, 0, 0)
	checkErr(gormDb.Find(&ulist).Error)
	return ulist
}

// FindByID .
func (*UserInfoOrmAtomicService) FindByID(id int) *UserInfo {
	u := UserInfo{}
	checkErr(gormDb.First(&u).Error)
	return &u
}
