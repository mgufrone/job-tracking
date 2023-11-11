package facades

import (
	"github.com/goravel/framework/facades"
	"gorm.io/gorm"
)

func DBSource() *gorm.DB {
	db, err := facades.App().Make("db.source")
	if err != nil {
		return nil
	}
	return db.(*gorm.DB)
}
