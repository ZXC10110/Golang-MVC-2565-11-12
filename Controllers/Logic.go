package Controllers

import (
	"fmt"
	"net/http"

	"github.com/zxc10110/mvc_14/Models"

	"github.com/gin-gonic/gin"
)

//GetAllProduct ... Get all product
func GetAllProduct(c *gin.Context) {
	data, err := Models.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, data)
		return
	}
}

//GetProductById ... Get product by id
func GetProductById(c *gin.Context) {
	var product Models.ProductOperator
	c.ShouldBindJSON(&product)

	request := Models.ProductOperator{
		Pid: product.Pid,
	}

	fmt.Println(request.Pid)

	data, err := Models.GetProductID(request)
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, data)
		return
	}
}

//GetProductById ... Get product by id
func GetUpdateById(c *gin.Context) {
	var product Models.ProductOperator
	c.ShouldBindJSON(&product)

	request := Models.ProductOperator{
		Pid:      product.Pid,
		Qty:      product.Qty,
		Operator: product.Operator,
	}

	fmt.Println(request.Pid)
	fmt.Println(request.Qty)
	fmt.Println(request.Operator)

	//get amount by id
	qty, err := Models.GetQtyId(request)
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	}

	//checking qty
	if qty.Qty <= 0 && request.Operator == "decrease" {
		c.JSON(http.StatusBadRequest, "สินค้าเหลือ 0 ไม่สามารถแก้ไขได้")
		return
	}

	//checking request qty == string then qty == 0
	if request.Qty == 0 {
		c.JSON(http.StatusBadRequest, "ต้องเป็นตัวเลขเท่านั้น")
		return
	}

	//update amount by id
	er := Models.UpdateProductId(request, qty)
	if er != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	}

	//get updated amount by id
	data, err := Models.GetProductID(request)
	if err != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		c.JSON(http.StatusOK, data)
		return
	}
}
