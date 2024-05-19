package editorial

type Editorial struct {
	EditorialID     int
	NombreEditorial string
}

func (e *Editorial) ObtenerEditorialID() int {
	return e.EditorialID
}

func (e *Editorial) ObtenerNombreEditorial() string {
	return e.NombreEditorial
}
