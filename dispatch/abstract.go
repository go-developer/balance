// Package dispatch...
//
// Description : dispatch...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 5:53 下午
package dispatch

import "github.com/go-developer/balance/define"

// IDispatch 负载均衡的接口定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:44 下午 2021/4/1
type IDispatch interface {
	// Get 获取一个节点
	Get(nodeList []*define.SeverNode) (string, *define.Error)
}
