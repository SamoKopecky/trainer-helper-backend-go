package store

type StoreBase[T any] interface {
	Update(model *T) error
	Insert(model *T) error
	Get() ([]T, error)
	InsertMany(models *[]T) error
	Undelete(modelId int) error
}
