package test

import (
	"encoding/json"
	"fmt"
	"log"
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

func TestJson2(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "Jack"
	m["age"] = 12
	m["contact"] = map[string]interface{}{
		"country": "China",
		"city":    "Beijing",
	}
	//将映射序列化到JSON字符串
	data, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		log.Println("ERROR:", err)
	}
	fmt.Println(string(data))
}
