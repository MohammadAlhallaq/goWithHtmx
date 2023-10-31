package services

import (
	"goWithHtmx/model"
	"goWithHtmx/model/mysql"
)

type ItemConfiguration func(service *ItemService) error

type ItemService struct {
	item model.ItemRepo
}

func NewItemService(configurations ...ItemConfiguration) (*ItemService, error) {
	service := &ItemService{}
	for _, cfg := range configurations {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithMysqlItem(connectionString string) ItemConfiguration {
	return func(is *ItemService) error {
		itemRepo, err := mysql.New(connectionString)
		if err != nil {
			return err
		}
		is.item = itemRepo
		return nil
	}
}

func (is *ItemService) GetAllItems() ([]model.Item, error) {
	all, err := is.item.GetAll()
	if err != nil {
		return all, err
	}

	return all, err
}
