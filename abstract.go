// Package balance ...
//
// Description : 接口定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 2:43 下午
package balance

import (
	"github.com/go-developer/gopkg/easylock"
)

// SeverNode 服务器节点的数据结构
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:46 下午 2021/4/1
type SeverNode struct {
	Host   string  `json:"host"`
	Port   int     `json:"port"`
	Weight float64 `json:"weight"`
	Status int     `json:"status"`
}

// Server server 的具体配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:59 下午 2021/4/1
type Server struct {
	lock     easylock.EasyLock
	NodeList []*SeverNode
}

// Add 添加一个Server
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:00 下午 2021/4/1
func (s *Server) Add(node *SeverNode) {
	s.lock.Lock("")
	defer s.lock.Unlock("")
	s.NodeList = append(s.NodeList, node)
}

// IBalance 负载均衡的接口定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:44 下午 2021/4/1
type IBalance interface {
	Get() string
	Add(node string)
	Remove(node string)
}
