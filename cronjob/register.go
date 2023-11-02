package cronjob

import (
	"account/config"
	"account/model"
	"time"
)

func DeleteExpiredTemporaryEmail() {
	db := config.GetDB()

	db.Where("time_end < ?", time.Now()).Unscoped().Delete(&model.TemporaryCredential{})
}
