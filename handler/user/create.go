package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create creates a new user account.
// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user [post]
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	t, _ := time.Parse("2006-01-02", r.Birthday)
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		Sex:      r.Sex,
		Age:      r.Age,
		Birthday: t,
		Address:  r.Address,
	}
	// Validate the data.
	if err := u.Validate(); err != nil {

		log.Error("Validate the data.", err)
		SendResponse(c, &errno.Errno{Code: errno.ErrValidation.Code, Message: err.Error()}, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
