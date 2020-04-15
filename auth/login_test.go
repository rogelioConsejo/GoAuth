package auth

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	token,err:=Login("rserrano@sysandweb.com","123456")
	if err ==nil {
		fmt.Println("Inicio de sesion exitoso: ",token.codigo)
	}else {
		fmt.Print(err)
	}
}

func TestLogout(t *testing.T) {
	err :=Logout("sdfsdfsdfsdf")
	if err!= nil {
		fmt.Print(err)
	}
}