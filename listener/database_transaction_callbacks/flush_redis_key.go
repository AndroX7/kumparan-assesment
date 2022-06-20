package database_transaction_callbacks

import (
	"gorm.io/gorm"
)

func (c *Callback) FlushRedisKey(db *gorm.DB) {
	if db.Statement.Schema != nil {
		switch db.Statement.Schema.Table {

		}
	}
}
