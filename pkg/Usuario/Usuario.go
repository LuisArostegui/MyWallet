package Usuario

import (
	. "MyWallet/pkg/MoneyAccount"
	"time"
)

type Usuario struct {
	Nombre   string
	Apellido string
	fnac     time.Time
	cuentas  []Cuenta
}

func NewUsuario(opciones ...Opcion) Usuario {
	user := Usuario{Nombre: "N/A", Apellido: "N/A",
		fnac: time.Date(1997, time.January, 1, 0, 0, 0, 0, time.Local)}
	for _, opcion := range opciones {
		opcion(&user)
	}
	return user
}

func (user *Usuario) AddCuenta(nombreAccount string, saldo float64, descripcion string) {
	a := NewAccount(NombreAccount(nombreAccount), Saldo(saldo), Fcreate(time.Now()), Descripcion(descripcion))
	user.cuentas = append(user.cuentas, a)
}

func (user *Usuario) AddCuentaBancaria(id string, nombreAccount string, saldo float64, descripcion string) {
	a := NewAccount(NombreAccount(nombreAccount), Saldo(saldo), Fcreate(time.Now()), Descripcion(descripcion))
	user.cuentas = append(user.cuentas, a)
}

func (user Usuario) GetEdad() uint8 {
	return uint8(time.Now().Year() - user.fnac.Year())
}
