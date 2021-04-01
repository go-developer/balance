// Package define ...
//
// Description : define ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 5:58 下午
package define

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
