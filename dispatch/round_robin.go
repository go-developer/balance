// Package dispatch ...
//
// Description : Round Robin 轮询选择服务器
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 8:04 下午
package dispatch

import (
	"fmt"

	"github.com/go-developer/balance/define"
	"github.com/go-developer/gopkg/easylock"
)

// NewRoundRobin 轮询调度
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:07 下午 2021/4/1
func NewRoundRobin() IDispatch {
	return &RoundRobin{
		lock:          easylock.NewLock(),
		nextNodeIndex: 0,
	}
}

// RoundRobin 轮询选择机器
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:06 下午 2021/4/1
type RoundRobin struct {
	lock          easylock.EasyLock
	nextNodeIndex int
}

// Get ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 8:05 下午 2021/4/1
func (r *RoundRobin) Get(nodeList []*define.SeverNode) (string, *define.Error) {
	if len(nodeList) == 0 {
		return "", define.NewError(define.ErrorTypeNodeListEmpty, "服务器可用节点为空")
	}
	_ = r.lock.Lock()
	defer func() {
		_ = r.lock.Unlock()
	}()
	if r.nextNodeIndex >= len(nodeList) {
		// 记录过索引之后, 在下次访问之前, 可能移除了某些节点, 所以要检测越界
		r.nextNodeIndex = len(nodeList) - 1
	}
	node := fmt.Sprintf("%s:%d", nodeList[r.nextNodeIndex].Host, nodeList[r.nextNodeIndex].Port)
	r.nextNodeIndex = (r.nextNodeIndex + 1) % len(nodeList)
	return node, nil
}
