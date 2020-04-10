package auth

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	token,err:=Login("rserrano@sysandweb.com","123456")
	if err !=nil {
		fmt.Println(token)
	}
}
