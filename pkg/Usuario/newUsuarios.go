package Usuario

import (
	"time"
)

type Opcion func(*Usuario)

func Nombre(nombre string) Opcion {
	return func(user *Usuario) {
		user.Nombre = nombre
	}
}

func Apellido(apellido string) Opcion {
	return func(user *Usuario) {
		user.Apellido = apellido
	}
}

func Fnac(fnac time.Time) Opcion {
	return func(user *Usuario) {
		user.fnac = fnac
	}
}
