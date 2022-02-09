package ginuser

import (
	"locnguyen/common"
	"locnguyen/component"
	"locnguyen/modules/users/usersbiz"
	"locnguyen/modules/users/userstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewGetUserBiz(store)

		data, err := biz.GetUser(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
