package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Funciones y estructuras para manejar usuarios, libros, etc.

type Genero struct {
	ID     int
	Nombre string
}

type Autor struct {
	ID           int
	Nombre       string
	Nacionalidad string
}

type Editorial struct {
	ID     int
	Nombre string
}

type Libro struct {
	Titulo      string
	GeneroID    int
	AutorID     int
	Year        int
	Descripcion string
	EditorialID int
}

// Libros registrados

var generos = []Genero{
	{1, "Fantasía"},
	{2, "Realismo Mágico"},
	{3, "Clásicos"},
	{4, "Ciencia Ficción"},
	{5, "Terror"},
	{6, "Comedia"},
	{7, "Drama"},
}

var autores = []Autor{
	{1, "J.R.R. Tolkien", "Inglés"},
	{2, "Gabriel García Márquez", "Colombiano"},
	{3, "Miguel de Cervantes", "Español"},
	{4, "Isaac Asimov", "Estadounidense"},
	{5, "Stephen King", "Estadounidense"},
	{6, "Aristófanes", "Griego"},
	{7, "Khaled Hosseini", "Persa"},
}

var editoriales = []Editorial{
	{1, "Allen & Unwin"},
	{2, "Editorial Sudamericana"},
	{3, "Francisco de Robles"},
	{4, "Alianza Editorial"},
	{5, "On Writing"},
	{6, "Losada"},
	{7, "Riverhead Books"},
}

var libros = []Libro{
	{"El señor de los anillos", 1, 1, 1954, "Un libro de fantasía épica.", 1},
	{"Cien años de soledad", 2, 2, 1967, "Una obra maestra del realismo mágico.", 2},
	{"Don Quijote de la Mancha", 3, 3, 1605, "Una novela clásica de la literatura española.", 3},
	{"Yo Robot", 4, 4, 1950, "Es una de las obras más populares de Asimov.", 4},
	{"El resplandor", 5, 5, 1977, "La maldad acecha las mentes débiles.", 5},
	{"Lisístrata", 6, 6, 2003, "Busco la comedia en los momentos más oscuros.", 6},
	{"Cometas en el cielo", 7, 7, 2003, "Narra el ir y venir en la vida de dos niños.", 7},
}

// Funcion para lectura de biblioteca

func MostrarGenerosDisponibles() {
	fmt.Println("Géneros disponibles:")
	for _, genero := range generos {
		fmt.Printf("%d. %s\n", genero.ID, genero.Nombre)
	}
	fmt.Println("Ingrese 's' para salir del programa.")
}

func MostrarLibrosPorGenero(generoID int) bool {
	encontrado := false
	fmt.Println("Libros disponibles:")
	for _, libro := range libros {
		if libro.GeneroID == generoID {
			encontrado = true
			autor := autores[libro.AutorID-1]
			editorial := editoriales[libro.EditorialID-1]
			genero := generos[libro.GeneroID-1]
			fmt.Printf("Título: %s\nAño: %d\nDescripción: %s\nGénero: %s\nAutor: %s\nNacionalidad del Autor: %s\nEditorial: %s\n----\n",
				libro.Titulo, libro.Year, libro.Descripcion, genero.Nombre, autor.Nombre, autor.Nacionalidad, editorial.Nombre)
		}
	}
	if !encontrado {
		fmt.Println("No hay libros disponibles en este género.")
	}
	if encontrado {
		fmt.Println("Ingrese 'r' para regresar a la lista de géneros o 's' para salir del programa.")
	}
	return encontrado
}

// Funciones para los clientes
type Usuario struct {
	Correo     string
	Contrasena string
}

var usuariosRegistrados = make(map[string]*Usuario)

func GuardarUsuarios() {
	file, err := os.Create("usuarios.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	for _, usuario := range usuariosRegistrados {
		fmt.Fprintf(file, "%s,%s\n", usuario.Correo, usuario.Contrasena)
	}
}

func CargarUsuarios() {
	file, err := os.Open("usuarios.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, ",")
		if len(partes) == 2 {
			usuariosRegistrados[partes[0]] = &Usuario{Correo: partes[0], Contrasena: partes[1]}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}

func CrearCliente(correo, contrasena, nombreTitular, numeroTarjeta, fechaCaducidad, claveSeguridad string) {
	usuario := &Usuario{
		Correo:     correo,
		Contrasena: contrasena,
	}
	usuariosRegistrados[correo] = usuario
	GuardarUsuarios()
}

func AutenticarUsuario(correo, contrasena string) (*Usuario, bool) {
	usuario, existe := usuariosRegistrados[correo]
	if !existe || usuario.Contrasena != contrasena {
		return nil, false
	}
	return usuario, true
}

// Validaciones
func validarCorreo(correo string) bool {
	if strings.Contains(correo, "@") && strings.HasSuffix(correo, ".com") {
		return true
	}
	return false
}

func validarContrasena(contrasena string) bool {
	if len(contrasena) >= 8 {
		for _, char := range contrasena {
			if unicode.IsUpper(char) {
				return true
			}
		}
	}
	return false
}

func main() {
	CargarUsuarios()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("¿Qué desea hacer?")
		fmt.Println("1. Crear una nueva cuenta de cliente")
		fmt.Println("2. Iniciar sesión")
		fmt.Println("3. Salir")
		fmt.Print("Ingrese su opción: ")
		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			var correo, contrasena string
			for {
				fmt.Print("Ingrese su correo: ")
				correo, _ = reader.ReadString('\n')
				correo = strings.TrimSpace(correo)
				if validarCorreo(correo) {
					break
				} else {
					fmt.Println("Correo inválido. Debe contener un arroba y terminar con '.com'.")
				}
			}

			for {
				fmt.Print("Ingrese la contraseña: ")
				contrasena, _ = reader.ReadString('\n')
				contrasena = strings.TrimSpace(contrasena)
				if validarContrasena(contrasena) {
					break
				} else {
					fmt.Println("Contraseña inválida. Debe tener al menos 8 caracteres y contener al menos una letra mayúscula.")
				}
			}

			fmt.Print("Ingrese el nombre del titular de la tarjeta: ")
			nombreTitular, _ := reader.ReadString('\n')
			nombreTitular = strings.TrimSpace(nombreTitular)

			fmt.Print("Ingrese el número de tarjeta: ")
			numeroTarjeta, _ := reader.ReadString('\n')
			numeroTarjeta = strings.TrimSpace(numeroTarjeta)

			fmt.Print("Ingrese la fecha de caducidad: ")
			fechaCaducidad, _ := reader.ReadString('\n')
			fechaCaducidad = strings.TrimSpace(fechaCaducidad)

			fmt.Print("Ingrese la clave de seguridad: ")
			claveSeguridad, _ := reader.ReadString('\n')
			claveSeguridad = strings.TrimSpace(claveSeguridad)

			CrearCliente(correo, contrasena, nombreTitular, numeroTarjeta, fechaCaducidad, claveSeguridad)
			fmt.Println("Cliente creado exitosamente!")

		case "2":
			fmt.Print("Ingrese su correo: ")
			correo, _ := reader.ReadString('\n')
			correo = strings.TrimSpace(correo)

			fmt.Print("Ingrese la contraseña: ")
			contrasena, _ := reader.ReadString('\n')
			contrasena = strings.TrimSpace(contrasena)

			usuario, autenticado := AutenticarUsuario(correo, contrasena)
			if autenticado {
				fmt.Printf("¡Bienvenido, %s!\n", usuario.Correo)
				for {
					MostrarGenerosDisponibles()
					fmt.Print("Seleccione un género por su número o 's' para salir: ")
					generoIDStr, _ := reader.ReadString('\n')
					generoIDStr = strings.TrimSpace(generoIDStr)
					if generoIDStr == "s" {
						fmt.Println("Saliendo del programa...")
						os.Exit(0)
					}

					var generoID int
					fmt.Sscanf(generoIDStr, "%d", &generoID)

					if MostrarLibrosPorGenero(generoID) {
						fmt.Print("Ingrese 'r' para regresar a la lista de géneros o 's' para salir: ")
						opcion, _ := reader.ReadString('\n')
						opcion = strings.TrimSpace(opcion)
						if opcion == "r" {
							continue
						} else if opcion == "s" {
							fmt.Println("Saliendo del programa...")
							os.Exit(0)
						}
						continue
					}
				}
			} else {
				fmt.Println("Correo o contraseña incorrecta.")
			}

		case "3":
			fmt.Println("Saliendo del programa...")
			os.Exit(0)

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
