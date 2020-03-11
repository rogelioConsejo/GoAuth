package auth

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestUsuario_Password(t *testing.T) {
	mail := "spoq@mail.com"
	u, err := NewUsuario(mail, "spoqpass")
	if u != nil {
		if u.GetEmail() != mail {
			t.Error("email guardado incorrectamente")
		}
		err = u.CheckPassword("spoqpass")
		if err != nil {
			t.Errorf("Error al revisar passwordHash: %s\n", err)
		} else {
			println("Password verificado")
		}
		err = u.CheckPassword("spoqpas")
		if err == nil {
			t.Errorf("Error al reconocer passwordHash incorrecto: %s\n", err)
		} else {
			println("Password incorrecto identificado")
		}

		var i int
		for i = 1; i <= 10; i++ {
			password := String(rand.Intn(20) + 1)
			err = u.CheckPassword(password)
			fmt.Printf("%d Probando passwordHash: %s -- ", i, password)
			if err == nil {
				t.Errorf("Error al reconocer passwordHash incorrecto: %s\n", err)
			} else {
				println("Password incorrecto identificado")
			}
		}

		err = u.SetPassword("passwordHash inválido")
		if err == nil {
			t.Error("No se detectó passwordHash inválido al usar SetPassword")
		}
		err = u.SetPassword("    ")
		if err == nil {
			t.Error("No se detectó passwordHash inválido al usar SetPassword")
		}
		err = u.SetPassword("  dfw  ")
		if err == nil {
			t.Error("No se detectó passwordHash inválido al usar SetPassword")
		}
		err = u.SetPassword("passwordvalido")
		if err != nil {
			t.Error("No se detectó passwordHash válido al usar SetPassword")
		}
	} else {
		t.Error("usuario nulo")
	}

	u, err = NewUsuario("spoq@mailcom", "spoqpass")
	if u != nil || err == nil {
		t.Error("no se detectó email incorrecto al crear usuario")
	}
	u, err = NewUsuario("spoq@mail.com", "spoqp ass")
	if u != nil || err == nil {
		t.Error("no se detectó passwordHash incorrecto al crear usuario")
	}

}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func TestUsuario_Email(t *testing.T) {
	mailInvalido := "spoq.mail.com"
	mailValido := "spoq@mail.com"
	esValido, err := validarEmail(mailInvalido)
	println(esValido)
	esValido, err = validarEmail(mailValido)
	println(esValido)
	if err != nil {
		t.Errorf("Error al revisar email: %s\n", err)
	}
}

func TestUsuario_RevisarPassword(t *testing.T) {
	const passwordValido = "erD21*dw#$"
	const passwordInvalido = "qwoe'sdwf"
	const passwordCorto = "1doe"
	const passwordLargo  = "1pe9*e93iw1oqiwedfj#*1poeadaaai"
	esValido, err := validarPassword(passwordValido)
	if !esValido {
		t.Error("No se detectó passwordHash Válido")
	}
	if err != nil {
		t.Errorf("Error al revisar Password: %s\n", err)
	}
	if esValido && err == nil{
		println("Password validado correctamente")
	}
	esValido, err = validarPassword(passwordInvalido)
	if esValido {
		t.Error("No se detectó passwordHash inVálido")
	}
	if err != nil {
		t.Errorf("Error al revisar Password: %s\n", err)
	}
	if !esValido && err == nil {
		println("Password inválido detectado correctamente")
	}
	esValido, err = validarPassword(passwordCorto)
	if esValido {
		t.Error("No se detectó passwordHash muy corto")
	}
	if err != nil {
		t.Errorf("Error al revisar Password: %s\n", err)
	}
	if !esValido && err == nil{
		println("Password corto detectado correctamente")
	}
	esValido, err = validarPassword(passwordLargo)
	if esValido {
		t.Error("No se detectó passwordHash muy largo")
	}
	if err != nil {
		t.Errorf("Error al revisar Password: %s\n", err)
	}
	if !esValido && err == nil{
		println("Password Largo detectado correctamente")
	}
}
