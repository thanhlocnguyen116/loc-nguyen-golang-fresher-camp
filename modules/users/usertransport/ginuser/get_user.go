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
		user_id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewGetUserBiz(store)

		data, err := biz.GetUser(c.Request.Context(), user_id)

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
