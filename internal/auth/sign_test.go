package auth

import (
	"YimiTV-go/internal/definition"
	"fmt"
	"testing"
)

func TestSignature(t *testing.T) {

	var p definition.ValidateTokenParameter

	p.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiN2Y2YWQ2MGNmYWU1ZTMwYTFlZTY1YTc0NjE1NjJmMzZkM2NkODU0ZjAyZDQxNmUzMTljZjBkNzE3YTNmNjdjZCIsImV4cCI6MTczOTg3Njc4OSwiaWF0IjoxNzM3Mjg0Nzg5LCJzY29wZSI6InJlYWQgd3JpdGUiLCJ1c2VyX2lkIjoiMlBieEYzcUlia0ZYeWVhUUh5VGs1N2NUd2VvIn0.ned8aBNsruUOsdYGiIcW_ZxmSprnBuCCd4diKQeecGU"
	p.PageToken = "a"
	p.TimeUnix = 1737285676

	fmt.Println(Signature("a", "a1", &p))

}
