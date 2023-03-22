package Controllers

import (
	"net/http"

	"github.com/zxc10110/mvc_14/Models"

	"github.com/gin-gonic/gin"
)

//GetAllHospitals ... Get all hospitals
func GetAllProduct(c *gin.Context) {
	data, err := Models.GetAllProducts()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
