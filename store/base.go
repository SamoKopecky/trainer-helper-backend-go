package store

type StoreBase[T any] interface {
	Update(model *T) error
	Insert(model *T) error
	Delete(modelId int) error
	Get() ([]T, error)
	GetById(modelId int) (model T, err error)
	InsertMany(models *[]T) error
	UndeleteMany(modelIds []int) error
	DeleteMany(modelIds []int) error
}
