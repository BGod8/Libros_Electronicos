package libro

type Libro struct {
	Titulo      string
	AutorID     int
	GeneroID    int
	Year        int
	Descripcion string
	EditorialID int
}

func (l *Libro) ObtenerTitulo() string {
	return l.Titulo
}

func (l *Libro) ObtenerAutorID() int {
	return l.AutorID
}

func (l *Libro) ObtenerGeneroID() int {
	return l.GeneroID
}

func (l *Libro) ObtenerYear() int {
	return l.Year
}

func (l *Libro) ObtenerDescripcion() string {
	return l.Descripcion
}

func (l *Libro) ObtenerEditorialID() int {
	return l.EditorialID
}
