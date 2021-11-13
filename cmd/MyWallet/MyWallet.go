package MyWallet

import (
	. "MyWallet/pkg/Usuario"
	"time"
)

func NuevoUsuario(Nombre string, Apellido string, fnac time.Time) Usuario {
	return NewUsuario()
}

func añadirCuentaUsuario(user Usuario, nombreCuenta string, saldo float64, descripcion string) {
	user.AddCuenta(nombreCuenta, saldo, descripcion)
}

func añadirCuentaBancariaUsuario(user Usuario, id string, nombreCuenta string, saldo float64, descripcion string) {
	user.AddCuentaBancaria(id, nombreCuenta, saldo, descripcion)
}
