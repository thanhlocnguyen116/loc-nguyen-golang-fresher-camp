package ginrestaurant

import (
	"locnguyen/common"
	"locnguyen/component"
	"locnguyen/modules/restaurants/biz"
	"locnguyen/modules/restaurants/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := biz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
