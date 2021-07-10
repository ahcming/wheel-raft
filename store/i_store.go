package store

//IStore 存储层抽象接口
type IStore interface {
	Get(key string) (interface{}, bool)

	Set(key string, val interface{}) error
}