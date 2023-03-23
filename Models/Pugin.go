package Models

import (
	"fmt"

	"github.com/zxc10110/mvc_14/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllProduct
func GetAllProducts() (product []Product, err error) {
	if err = Config.DB.Table("Test.Product").
		Order("qty").
		Find(&product).Error; err != nil {
		return
	}
	fmt.Println(product)
	return
}

//GetProductID
func GetProductID(req ProductOperator) (product Product, err error) {
	if err = Config.DB.Table("Test.Product").
		Where("pid = ?", req.Pid).
		Find(&product).Error; err != nil {
		return
	}
	fmt.Println(product)
	return
}

//getQtyById
func GetQtyId(req ProductOperator) (product Product, err error) {
	//get amount by id
	if err = Config.DB.Select("qty").
		Table("Test.Product").
		Where("pid = ?", "Ticket-01").
		Find(&product).Error; err != nil {
		return
	}

	fmt.Println("Qty : ", product)

	return
}

//getUpdateProductId
func UpdateProductId(req ProductOperator, qty Product) (err error) {
	//get amount by id
	fmt.Println("Current Qty", qty)
	fmt.Println("Input Qty", req.Qty)

	var amount int
	if req.Operator == "decrease" {
		amount = qty.Qty - req.Qty
	}

	if req.Operator == "increase" {
		amount = qty.Qty + req.Qty
	}

	fmt.Println("Operator : ", req.Operator)
	fmt.Println("Amount : ", amount)

	if err = Config.DB.Table("Test.Product").
		Where("pid = ?", "Ticket-01").
		Update(map[string]interface{}{
			"qty": amount,
		}).
		Error; err != nil {
		return
	}

	return
}

/*
//CountPatient
func GetCountPatientPerHospital() (count []CountPatient, err error) {
	if err = Config.DB.Select("COUNT(p.hnid) as count, h.hid, h.title").
		Table("Test.hospitals as h").
		Joins("left join patients as p on p.hid = h.hid").
		Group("hid, title").
		Find(&count).Error; err != nil {
		return
	}
	return
}

//OrderCountPatient
func GetOrderCountPatient() (count []CountPatient, err error) {
	if err = Config.DB.Select("COUNT(p.hnid) as count, h.hid, h.title").
		Table("Test.hospitals as h").
		Joins("left join patients as p on p.hid = h.hid").
		Group("hid, title").
		Order("count DESC").
		Find(&count).Error; err != nil {
		return
	}
	return
}
*/
