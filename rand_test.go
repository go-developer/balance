// Package dispatch...
//
// Description : dispatch...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 6:48 下午
package balance

import (
	"fmt"
	"testing"

	"github.com/go-developer/balance/dispatch"

	"github.com/go-developer/balance/define"
)

// TestRand 测试随机选节点
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:48 下午 2021/4/1
func TestRand(t *testing.T) {
	result := make(map[string]int)
	list := []*define.SeverNode{
		{"127.0.0.1:8080", "127.0.0.1", 8080, 0, 1},
		{"127.0.0.1:8081", "127.0.0.1", 8081, 0, 1},
		{"127.0.0.1:8082", "127.0.0.1", 8082, 0, 1},
		{"127.0.0.1:8083", "127.0.0.1", 8083, 0, 1},
	}
	s := NewServer(list, dispatch.NewRand())
	for i := 0; i < 10000; i++ {
		node, _ := s.Get()
		if _, exist := result[node]; !exist {
			result[node] = 0
		}
		result[node]++
	}
	fmt.Printf("%+v", result)
}
