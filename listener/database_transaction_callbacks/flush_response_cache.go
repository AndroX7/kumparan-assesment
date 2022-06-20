package database_transaction_callbacks

import (
	"log"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"

	"github.com/AndroX7/kumparan-assesment/app/api/middleware"
	"github.com/AndroX7/kumparan-assesment/models"
)

func (c *Callback) FlushResponseCache(db *gorm.DB) {
	if db.Statement.Schema != nil {
		go c.responseCacheUsecase.FlushGeneralSet(db.Statement.Schema.Table)

		switch db.Statement.Schema.Table {
		case "product_brands":
			var artistM models.Articles
			err := copier.Copy(&artistM, db.Statement.Model)
			if err != nil {
				log.Println("error flush product_brands related response cache data: ", err)
			} else {
				go c.responseCacheUsecase.FlushFromArtist(&artistM)
			}
			go c.responseCacheUsecase.FlushAllFromSet(middleware.RedisResponseArtistSet)
		}
	}
}
