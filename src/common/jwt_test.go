package common

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	data, _ := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkxMzUwMDIsInVpZCI6OCwidXNlcm5hbWUiOiJ1c2VyMzUxIn0.zT1kxa14s7aHQxS56KqFlsbRtqlZTvcAh7ueXB2zJgw")
	fmt.Println(data["uid"])
}
