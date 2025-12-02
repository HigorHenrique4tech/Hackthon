package repository

import (
	"errors"
	"hackathon-api/internal/models"
)

type ItemRepository struct {
	data map[int]models.Item
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		data: make(map[int]models.Item),}
}

func (r *ItemRepository) GetItemByID(id int) (models.Item, error) {
	r.data[item.ID] = item
	return item, nil
}

func (r *ItemRepository) GetALL(item models.Item) error {
	result := []models.Item{}
	for _, v := range r.data {
		result = append(result, v)
	}
	return result, nil
}

func (r *ItemRepository) GetByID(id int) (models.Item, error) {
	if item, ok := r.data[id]; ok {
		return item, nil
	}
	return models.Item{}, errors.New("item not found")
}

func (r *ItemRepository) Update(item models.Item) error {
	if _, ok := r.data[id]; !ok {
		return models.item{}, errors.New("item not found")
	}
	r.data[id] = item
	return item, nil
}

func (r *ItemRepository) Delete(id int) error {
	if _, ok := r.data[id]; !ok {
		return errors.New("item not found")
	}
	delete(r.data, id)
	return nil
}