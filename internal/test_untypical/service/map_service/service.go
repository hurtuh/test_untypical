package map_service

import (
	"test_untypical/internal/test_untypical/types"
	"test_untypical/pkg/infrastruct"
)

type MapService struct {
	storage map[interface{}]interface{}
}

func NewMapService() *MapService {
	return &MapService{storage: make(map[interface{}]interface{})}
}

func (m *MapService) Delete(key interface{}) error {

	if _, have := m.storage[key]; !have {
		return infrastruct.ErrorNotFound
	}

	delete(m.storage, key)

	return nil
}

func (m *MapService) Upsert(keyValue *types.KeyValue) error {

	m.storage[keyValue.Key] = keyValue.Value

	return nil
}

func (m *MapService) Get(key interface{}) (interface{}, error) {

	if v, ok := m.storage[key]; ok {
		return v, nil
	}

	return nil, infrastruct.ErrorNotFound
}

func (m *MapService) List() (*types.ListValues, error) {

	//В зависимости от поставленной задачи, можно добавить или убрать проверку
	//if len(m.storage) == 0 {
	//	return nil, infrastruct.ErrorNotFound
	//}
	arr := types.ListValues{
		Values: make([]interface{}, 0),
	}

	for _, v := range m.storage {
		arr.Values = append(arr.Values, v)
	}

	return &arr, nil
}
