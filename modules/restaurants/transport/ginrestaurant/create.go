package ginrestaurant

import (
	"locnguyen/common"
	"locnguyen/component"
	"locnguyen/modules/restaurants/biz"
	"locnguyen/modules/restaurants/model"
	"locnguyen/modules/restaurants/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		//get data, logic
		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := biz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
