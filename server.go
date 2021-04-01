// Package balance ...
//
// Description : 接口定义
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 2:43 下午
package balance

import (
	"github.com/go-developer/balance/define"
	"github.com/go-developer/balance/dispatch"
	"github.com/go-developer/gopkg/easylock"
)

//  ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:51 下午 2021/4/1
func NewServer(nodeList []*define.SeverNode, d dispatch.IDispatch) *Server {
	return &Server{
		lock:     easylock.NewLock(),
		NodeList: nodeList,
		Balance:  d,
	}
}

// Server server 的具体配置
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:59 下午 2021/4/1
type Server struct {
	lock     easylock.EasyLock
	NodeList []*define.SeverNode
	Balance  dispatch.IDispatch
}

// Add 添加一个Server
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:00 下午 2021/4/1
func (s *Server) Add(node *define.SeverNode) {
	_ = s.lock.Lock()
	defer func() {
		_ = s.lock.Unlock()
	}()
	s.NodeList = append(s.NodeList, node)
}

// Remove 移除一个server
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:09 下午 2021/4/1
func (s *Server) Remove(nodeID string) {
	_ = s.lock.Lock()
	defer func() {
		_ = s.lock.Unlock()
	}()
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
func (s *Server) Get() (string, *define.Error) {
	_ = s.lock.RLock()
	defer func() {
		_ = s.lock.RUnlock()
	}()
	return s.Balance.Get(s.NodeList)
}
