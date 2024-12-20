package venueStateTable

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"github.com/jinzhu/gorm"
)

func ResetTableTennisVS(db *gorm.DB, venueNum int) {
	for i := 1; i <= venueNum; i++ {
		todayVS := venueModel.TableTennisVenueState{ID: i, Date: "today"}
		tomorrowVS := venueModel.TableTennisVenueState{ID: i, Date: "tomorrow"}
		db.Save(&todayVS)
		db.Save(&tomorrowVS)
	}
}
