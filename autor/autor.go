package autor

type Autor struct {
	AutorID      int
	Nombre       string
	Nacionalidad string
}

func (a *Autor) ObtenerID() int {
	return a.AutorID
}

func (a *Autor) ObtenerNombre() string {
	return a.Nombre
}

func (a *Autor) ObtenerNacionalidad() string {
	return a.Nacionalidad
}
