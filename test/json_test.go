package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TestJson(t *testing.T) {
	result := &Result{}
	//调用k8s接口获取集群信息
	err := json.Unmarshal([]byte(""), result)
	if err != nil {
		fmt.Println("occur error => ", err)
	}
	fmt.Println("success => ", result)
}
