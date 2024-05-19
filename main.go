package main

import (
	"fmt"
	"libro_electronico/autor"
	"libro_electronico/editorial"
	"libro_electronico/genero"
	"libro_electronico/libro"
)

func imprimirLibrosPorGenero(libros []libro.Libro, generoID int) {
	fmt.Printf("%-5s %-30s %-10s %-5s\n", "ID", "Título", "AutorID", "Year")
	fmt.Println("-------------------------------------------------------")
	for i, libro := range libros {
		if libro.ObtenerGeneroID() == generoID {
			fmt.Printf("%-5d %-30s %-10d %-5d\n", i, libro.ObtenerTitulo(), libro.ObtenerAutorID(), libro.ObtenerYear())
		}
	}
}

func imprimirDetallesLibro(libro libro.Libro, autores []autor.Autor, editoriales []editorial.Editorial) {
	var autorNombre, editorialNombre string

	for _, autor := range autores {
		if autor.ObtenerID() == libro.ObtenerAutorID() {
			autorNombre = autor.ObtenerNombre()
			break
		}
	}

	for _, editorial := range editoriales {
		if editorial.ObtenerEditorialID() == libro.ObtenerEditorialID() {
			editorialNombre = editorial.ObtenerNombreEditorial()
			break
		}
	}

	fmt.Println("Detalles del libro seleccionado:")
	fmt.Println("Título:", libro.ObtenerTitulo())
	fmt.Println("Autor:", autorNombre)
	fmt.Println("Año:", libro.ObtenerYear())
	fmt.Println("Descripción:", libro.ObtenerDescripcion())
	fmt.Println("Editorial:", editorialNombre)
}

func main() {
	generos := []genero.Genero{
		{1, "Fantasía"},
		{2, "Realismo Mágico"},
		{3, "Clásicos"},
		{4, "Ciencia Ficción"},
		{5, "Terror"},
		{6, "Comedia"},
		{7, "Drama"},
	}

	autores := []autor.Autor{
		{1, "J.R.R. Tolkien", "Inglés"},
		{2, "Gabriel García Márquez", "Colombiano"},
		{3, "Miguel de Cervantes", "Español"},
		{4, "Isaac Asimov", "Estadounidense"},
		{5, "Stephen King", "Estadounidense"},
		{6, "Aristófanes", "Griego"},
		{7, "Khaled Hosseini", "Persa"},
	}

	editoriales := []editorial.Editorial{
		{1, "Allen & Unwin"},
		{2, "Editorial Sudamericana"},
		{3, "Francisco de Robles"},
		{4, "Alianza Editorial"},
		{5, "On Writing"},
		{6, "Losada"},
		{7, "Riverhead Books"},
	}

	libros := []libro.Libro{
		{"El señor de los anillos", 1, 1, 1954, "Un libro de fantasía épica.", 1},
		{"Cien años de soledad", 2, 2, 1967, "Una obra maestra del realismo mágico.", 2},
		{"Don Quijote de la Mancha", 3, 3, 1605, "Una novela clásica de la literatura española.", 3},
		{"Yo Robot", 4, 4, 1950, "Es una de las obras más populares de Asimov.", 4},
		{"El resplandor", 5, 5, 1977, "La maldad acecha las mentes débiles.", 5},
		{"Lisístrata", 6, 6, 2003, "Busco la comedia en los momentos más oscuros.", 6},
		{"Cometas en el cielo", 7, 7, 2003, "Narra el ir y venir en la vida de dos niños.", 7},
	}

	for {
		fmt.Println("Géneros disponibles:")
		for _, genero := range generos {
			fmt.Printf("%d: %s\n", genero.ObtenerGeneroID(), genero.ObtenerTipoGenero())
		}

		var generoID int
		fmt.Println("Ingresa el ID del género que deseas visualizar:")
		fmt.Scan(&generoID)

		fmt.Println("Estos son los Libros disponibles:")
		imprimirLibrosPorGenero(libros, generoID)

		var libroID int
		fmt.Println("Ingrese el número del libro que deseas ver:")
		fmt.Scan(&libroID)

		if libroID >= 0 && libroID < len(libros) && libros[libroID].ObtenerGeneroID() == generoID {
			imprimirDetallesLibro(libros[libroID], autores, editoriales)
		} else {
			fmt.Println("No se encontró la opción que seleccionaste, inténtalo nuevamente.")
		}

		var opcion string
		fmt.Println("Presiona 1 si deseas regresar al menú principal o cualquier tecla si deseas salir:")
		fmt.Scan(&opcion)

		if opcion != "1" {
			fmt.Println("Saliendo del sistema. ¡Adiós!")
			break
		}
	}
}
