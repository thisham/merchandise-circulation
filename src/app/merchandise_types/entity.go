package merchandise_types

type Domain struct {
	ID   int
	Name string
}

// main clauses: GET, CREATE, AMEND, REMOVE
type Services interface {
	GetAllMerchandiseTypes() ([]Domain, error)
}

// main clauses: INSERT, SELECT, UPDATE, DELETE
type Data interface {
	SelectAllMerchandiseTypes() (Domain, error)
}
