package user

import (
	"strconv"

	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	//var u model.UserModel
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	t, _ := time.Parse("2006-01-02", r.Birthday)

	log.Info(r.Username)
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		Sex:      r.Sex,
		Age:      r.Age,
		Birthday: t,
		Address:  r.Address,
	}
	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		log.Info(err.Error())
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
