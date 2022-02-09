package storage

type Store struct {
	DB map[string]string
}

//func NewStore() *Store {
//	store := &Store{db: map[string]string{}}
//	return store
//}

func (k Store) Get(key string) string {
	return k.DB[key]
}

func (k Store) Put(key string, value string) {
	k.DB[key] = value
}
