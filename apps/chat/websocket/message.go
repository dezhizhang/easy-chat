package websocket

type Message struct {
	Method string      `json:"method"`  // 请求的方法
	FormId string      `json:"form_id"` //
	Data   interface{} `json:"data"`    // 请求消息
}
