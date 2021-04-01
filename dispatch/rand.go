// Package dispatch...
//
// Description : dispatch...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 5:58 下午
package dispatch

import (
	"fmt"
	"math/rand"

	"github.com/go-developer/balance/define"
)

// NewRand ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:51 下午 2021/4/1
func NewRand() IDispatch {
	return &Rand{}
}

// Rand 随机选择
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:01 下午 2021/4/1
type Rand struct {
}

// Get 获取 host + 端口
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:01 下午 2021/4/1
func (r Rand) Get(nodeList []*define.SeverNode) (string, *define.Error) {
	if len(nodeList) == 0 {
		return "", define.NewError(define.ErrorTypeNodeListEmpty, "服务器可用节点为空")
	}
	node := nodeList[rand.Intn(len(nodeList))]
	return fmt.Sprintf("%s:%d", node.Host, node.Port), nil
}
