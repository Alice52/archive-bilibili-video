package api

import (
	"fmt"
	"github.com/alice52/archive/bili/util"
	"github.com/micro-services-roadmap/kit-common/viperx"
	"testing"
)

func TestClient_IsLogin(t *testing.T) {
	client := &BClient{}
	if client.isLogin() {
		fmt.Println("client has already login")
	} else {
		fmt.Println("client has not login")
	}
}

func TestClient_MySpaceInfo(t *testing.T) {
	viperx.InitViper("../config-local.yaml")

	info, err := logonFunc().MySpaceInfo()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
