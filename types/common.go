package types

// ResponseBody 固定的返回格式，成功是唯一的，Code 固定为 0，失败不唯一，Code 不为 0
type ResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
