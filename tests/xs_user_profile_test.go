package tests

import (
	"fmt"
	"testing"

	"github.com/olaola-chat/rbp-proto/dao/xianshi"
)

func TestXsUserProfile(t *testing.T) {
	row, err := xianshi.XsUserProfile.One("uid = ?", 100000000)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(row)
}
