package models

import (
	"fmt"

	u "github.com/mtjhartley/concerts-api/internal/pkg/utils"

	"github.com/jinzhu/gorm"
)

type Concert struct {
	gorm.Model
	Name         string `json:"name"`
	Date         string `json:"date"`
	FacebookLink string `json:"facebook_link"`
	TicketLink   string `json:"ticket_link"`
	UserId       uint   `json:"user_id"` //The user that this Concert belongs to
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (concert *Concert) Validate() (map[string]interface{}, bool) {

	if concert.Name == "" {
		return u.Message(false, "Concert name should be on the payload"), false
	}

	if concert.Date == "" {
		return u.Message(false, "Date should be on the payload"), false
	}

	if concert.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (concert *Concert) Create() map[string]interface{} {

	if resp, ok := concert.Validate(); !ok {
		return resp
	}

	GetDB().Create(concert)

	resp := u.Message(true, "success")
	resp["concert"] = concert
	return resp
}

func GetConcert(id uint) *Concert {

	concert := &Concert{}
	err := GetDB().Table("concert").Where("id = ?", id).First(concert).Error
	if err != nil {
		return nil
	}
	return concert
}

func GetConcerts(user uint) []*Concert {

	concert := make([]*Concert, 0)
	err := GetDB().Table("concert").Where("user_id = ?", user).Find(&concert).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return concert
}
