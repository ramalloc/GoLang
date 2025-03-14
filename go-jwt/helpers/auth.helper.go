package helpers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error)  {
	typeOfUser := c.GetString("user_type")
	log.Println("User type:", typeOfUser)
	err = nil
	if typeOfUser != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}

// Returning created error if user is not admin
func MatchUserTypeToUid(c *gin.Context, userId string) (err error)  {
	userType := c.GetString("user_type")
	uid := c.GetString("user_id")
	err = nil
	if userType == "USER" && uid != userId{
		err = errors.New("unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}