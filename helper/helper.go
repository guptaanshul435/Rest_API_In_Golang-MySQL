package helper

import (
	"log"

	"anshulgithub.com/anshul/usermangement/models"
	"github.com/badoux/checkmail"
)

func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func EmtRequest(user *models.User) bool {
	if user.Name == "" || user.Address == "" || user.PhoneNo == "" {
		return true
	}
	return false
}

func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}

func IsValidNumber(phoneNo string) bool {
	if len(phoneNo) > 10 || len(phoneNo) < 10 {
		return false
	}
	return true
}


