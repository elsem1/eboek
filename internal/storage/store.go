package storage

func NewStore(path string) *Store { // ontvangt een path en geeft een store pointer terug
	return &Store{}
}

type Store struct{}
