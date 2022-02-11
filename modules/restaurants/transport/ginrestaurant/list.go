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

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := biz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
