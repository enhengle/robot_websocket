package pojo

/**
 * @Author:lingwang
 * @Description:websocket的相关返回、请求结构
 * @Version: 1.0.0
 * @Date :2022/2/26 17:36
 */

var WebUrl WebSocketUrl

//wss连接地址
type WebSocketUrl struct {
	Url string `json:"url"`
}

//服务端返回信息
type Timtout struct {
	Op int64 `json:"op"`
	D  int64 `json:"d"`
}

//服务端返回信息-@机器人信息
type ReadyEvent struct {
	OP int64  `json:"op"`
	S  int64  `json:"s"`
	T  string `json:"t"`
	D  struct {
		Author struct {
			Avatar   string `json:"avatar"`
			Bot      string `json:"bot"`
			Id       string `json:"id"`
			Username string `json:"username"`
		} `json:"author"`
		ChannelId string `json:"channel_id"`
		Content   string `json:"content"`
		GuildId   string `json:"guild_id"`
		Id        string `json:"id"`
		Member    struct {
			JoinedAt string  `json:"joined_at"`
			Nick     string  `json:"nick"`
			Roles    []int64 `json:"roles"`
		} `json:"member"`
	} `json:"d"`
}
