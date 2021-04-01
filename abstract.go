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
	ID     string  `json:"id"`     // 机器编号
	Host   string  `json:"host"`   // ip
	Port   int     `json:"port"`   // 端口
	Weight float64 `json:"weight"` // 权重
	Status int     `json:"status"` // 状态
}

// Server server 的具体配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:59 下午 2021/4/1
type Server struct {
	lock     easylock.EasyLock
	NodeList []*SeverNode
	Balance  IBalance
}

// Add 添加一个Server
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:00 下午 2021/4/1
func (s *Server) Add(node *SeverNode) {
	_ = s.lock.Lock()
	defer s.lock.Unlock()
	s.NodeList = append(s.NodeList, node)
}

// Remove 移除一个server
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:09 下午 2021/4/1
func (s *Server) Remove(nodeID string) {
	_ = s.lock.Lock()
	defer s.lock.Unlock()
	for nodeIndex, item := range s.NodeList {
		if item.ID == nodeID {
			s.NodeList = append(s.NodeList[0:nodeIndex], s.NodeList[nodeIndex:]...)
			break
		}
	}
}

// Get 按照指定策略获取一台机器
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:17 下午 2021/4/1
func (s *Server) Get() string {
	return ""
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
