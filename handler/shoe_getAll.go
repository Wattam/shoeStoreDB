package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	shoe "github.com/wattam/shoeStoreDB/models"
)

func ShoeGetAll(store shoe.AllGetter) gin.HandlerFunc {

	return func(c *gin.Context) {
		results := store.GetAll()

		c.JSON(http.StatusOK, results)
	}
}
