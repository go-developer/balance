// Package define...
//
// Description : define...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-04-01 6:38 下午
package define

// NewError ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:40 下午 2021/4/1
func NewError(errType string, message string) *Error {
	return &Error{
		errType: errType,
		message: message,
	}
}

// Error 定义异常
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:38 下午 2021/4/1
type Error struct {
	errType string
	message string
}

// Error 获取错误信息, 兼容内置 error 类型
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:39 下午 2021/4/1
func (e Error) Error() string {
	return e.message
}

// GetErrType 获取错误类型
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:42 下午 2021/4/1
func (e Error) GetErrType() string {
	return e.errType
}

const (
	// ErrorTypeNodeListEmpty 服务器节点列表为空
	ErrorTypeNodeListEmpty = "NODE_LIST_EMPTY"
)
