package Models

import (
	"fmt"

	"github.com/zxc10110/mvc_14/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllHospitals
func GetAllProducts() (product []Product, err error) {
	if err = Config.DB.Select("*").
		Table("Test.Product").
		Find(&product).Error; err != nil {
		return
	}
	fmt.Println(product)
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
