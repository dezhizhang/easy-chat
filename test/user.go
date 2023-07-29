package test

import (
	"fmt"
	"im/driver"
	"im/model"
	"testing"
)

func Test(t testing.T) {
	tx := driver.DB.Find(&model.User{})
	fmt.Println(tx)
}
