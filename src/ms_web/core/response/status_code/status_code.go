package status_code

const (
	SUCCESS       = 200  //成功
	ERROR         = 100  //标准返回给客户的错误
	FORMAT_ERROR  = 1001 //格式不合法
	EXPIRED       = 1002 //过期
	CHINESE_ERROR = 1003 //格式不合法(中文)
	OTHER_ERROR   = 1004 //其它错误
)
