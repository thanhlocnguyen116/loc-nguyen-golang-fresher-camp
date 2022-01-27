package ginuser

import (
	"locnguyen/common"
	"locnguyen/component"
	"locnguyen/modules/users/usermodel"
	"locnguyen/modules/users/usersbiz"
	"locnguyen/modules/users/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := usersbiz.NewCreateUserBiz(store)

		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
