//package name should be same with folder
package db

import "shopping/models"

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
		Name:  "Yeezy",
	}
}
