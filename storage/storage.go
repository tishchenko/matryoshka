package storage

type Storage interface {
	Save()
	Load()
}