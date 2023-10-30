package services

import (
	"testing"
)

func TestNewItemService(t *testing.T) {
	is, err := NewItemService(WithMysqlItem("root:@tcp(127.0.0.1:3306)/hmshop"))
	if err != nil {
		t.Errorf(err.Error())
	}
	_, err = is.GetAllItems()
	if err != nil {
		t.Errorf(err.Error())
	}
}
