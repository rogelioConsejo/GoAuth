package goAuth

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestUsuario_Password(t *testing.T) {
	u := NewUsuario("spoq@mail.com", "spoqpass")
	err := u.CheckPassword("spoqpass")
	if err != nil {
		t.Errorf("Error al revisar password: %s\n", err)
	} else {
		println("Password verificado")
	}
	err = u.CheckPassword("spoqpas")
	if err == nil {
		t.Errorf("Error al reconocer password incorrecto: %s\n", err)
	} else {
		println("Password incorrecto identificado")
	}

	var i int
	for i = 1; i <= 100; i++ {
		password := String(rand.Intn(20) + 1)
		err = u.CheckPassword(password)
		fmt.Printf("%d Probando password: %s -- ", i, password)
		if err == nil {
			t.Errorf("Error al reconocer password incorrecto: %s\n", err)
		} else {
			println("Password incorrecto identificado")
		}
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

func TestUsuario_Email(t *testing.T){
	mail := "spoq.mail.com"
	emailRegexp, _ := regexp.MatchString("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}", mail)
	println(emailRegexp)
}