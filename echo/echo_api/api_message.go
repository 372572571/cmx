package echo_api

type ApiMessage struct {
	ApiContent            string // 已经写入的内容
	ApiContentRequest     string // 已经写入的内容
	ApiContentResponse    string // 已经写入的内容
	OpenContentRequest    string // 已经写入的内容
	OpenContentResponse   string
	IsCommentedOutputOnly bool
	IsUnmixed             bool // 是否不包含外部message
	Name                  string
}
