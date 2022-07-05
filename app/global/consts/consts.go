package consts

//全局常量
const (
	ValidatorPrefix      string = "validator_prefix_" //验证器前缀
	CacheUserTokenPrefix string = "UserToken:"        //用户登陆token缓存Key前缀
	CacheUserTokenTTL    int    = 7200                //用户登陆token缓存Key前缀

	RequestErrorParams int = 400001 //请求参数错误
	RequestErrorHeader int = 400002

	BusinessErrorUserSave      int = 500001
	BusinessErrorUserNotFund   int = 500002
	BusinessErrorUserPassword  int = 500003
	BusinessErrorUserLoginFail int = 500004
	BusinessErrorUserNotLogin  int = 500005
)

// BusinessErrorMap 因GO目前无法定义常量map 故定义为全局包中的map
var BusinessErrorMap = map[int]string{
	BusinessErrorUserSave:     "用户信息保存失败！",
	BusinessErrorUserNotFund:  "用户信息不存在！",
	BusinessErrorUserPassword: "用户密码不正确！",
	BusinessErrorUserNotLogin: "用户未登录！",

	RequestErrorHeader: "请求中Header头参数错误！",
}
