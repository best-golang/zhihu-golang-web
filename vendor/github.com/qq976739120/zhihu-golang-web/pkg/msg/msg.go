package msg

const (
	TIME_FORMAT  = "2006-01-02 15:04:05"

	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_PRODUCT = 10001
	//ERROR_NOT_EXIST_PRODUCT = 10002
	PRODUCT_ONLINE = 1
	PRODUCT_OFFLINE = 2
	PRODUCT_STORE = 3

	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN = 20003
	ERROR_AUTH = 20004
)

var Msg = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_EXIST_PRODUCT : "已存在该商品名称",
	//ERROR_NOT_EXIST_PRODUCT : "该商品不存在",
	PRODUCT_ONLINE:"商品正常出手",
	PRODUCT_OFFLINE:"商品下线",
	PRODUCT_STORE:"商品仓库中",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}

	return Msg[ERROR]
}