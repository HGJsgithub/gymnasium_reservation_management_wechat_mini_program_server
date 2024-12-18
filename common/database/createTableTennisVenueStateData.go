package database

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venue"
	"fmt"
	"github.com/jinzhu/gorm"
)

func CreateTableTennisVenueStateData(db *gorm.DB, venueNum int) {
	for i := 1; i <= venueNum; i++ {
		oneVs1 := venue.TableTennisVenueState1{ID: i}
		oneVs2 := venue.TableTennisVenueState2{ID: i}
		if db.First(&oneVs1).RecordNotFound() == true {
			db.Create(&venue.TableTennisVenueState1{
				ID:  i,
				T9:  false,
				T10: false,
				T11: false,
				T12: false,
				T13: false,
				T14: false,
				T15: false,
				T16: false,
				T17: false,
				T18: false,
				T19: false,
				T20: false,
				T21: false,
			})
		} else {
			fmt.Println(i, "号场地今天的乒乓球场地已经存在！")
			fmt.Println()
		}
		if db.First(&oneVs2).RecordNotFound() == true {
			db.Create(&venue.TableTennisVenueState2{
				ID:  i,
				T9:  false,
				T10: false,
				T11: false,
				T12: false,
				T13: false,
				T14: false,
				T15: false,
				T16: false,
				T17: false,
				T18: false,
				T19: false,
				T20: false,
				T21: false,
			})
		} else {
			fmt.Println(i, "号场地明天的乒乓球场地已经存在！")
			fmt.Println()
		}
	}
}
