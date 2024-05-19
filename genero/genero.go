package genero

type Genero struct {
	GeneroID   int
	TipoGenero string
}

func (g *Genero) ObtenerGeneroID() int {
	return g.GeneroID
}

func (g *Genero) ObtenerTipoGenero() string {
	return g.TipoGenero
}
