package service

import "test_untypical/internal/test_untypical/types"

//ServiceI Интерфейс описываюищй все необходимые методы для реализации сервиса
//Обычно я бы не стал завязывать и ключ и значение на пустом интерфейсе,
//но в условиях задач не было уточнения по типам данных для хранилища
type ServiceI interface {
	Delete(key interface{}) error
	Upsert(value *types.KeyValue) error
	Get(key interface{}) (interface{}, error)
	List() (*types.ListValues, error)
}
