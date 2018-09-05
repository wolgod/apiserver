package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service"
	"apiserver/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest

	if err := c.ShouldBind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	conn := util.Get()
	//记得销毁本次链连接
	defer conn.Close()

	//写入数据
	_, err := conn.Do("SET", "go_key", "redigo")
	if err != nil {
		fmt.Println("error while setting")
	}
	//获取value并转成字符串
	account_balance, err := redis.String(conn.Do("GET", "go_key"))
	if err != nil {
		fmt.Println("error while getting")
	}
	fmt.Printf("redisGet:%s \n", account_balance)

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
