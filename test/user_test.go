package test

import (
	"im/driver"
	"im/model"
	"testing"
)

func TestUser(t *testing.T) {
	driver.DB.Find(&model.User{})
}
