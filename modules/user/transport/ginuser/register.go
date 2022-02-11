package ginuser

import (
	"locnguyen/common"
	"locnguyen/component"
	"locnguyen/component/hasher"
	"locnguyen/modules/user/biz"
	"locnguyen/modules/user/model"
	"locnguyen/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data model.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := biz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
