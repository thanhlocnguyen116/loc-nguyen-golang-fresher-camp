package ginuser

import (
	"locnguyen/common"
	"locnguyen/component"
	"locnguyen/modules/users/usermodel"
	"locnguyen/modules/users/usersbiz"
	"locnguyen/modules/users/userstorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewUpdateUserBiz(store)

		if err := biz.UpdateUser(c.Request.Context(), id, &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
