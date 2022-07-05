package consts

//全局常量
const (
	ValidatorPrefix    string = "validator_prefix_" //验证器前缀
	RequestParamsError int    = 400001              //请求参数错误

	BusinessErrorUserSave      int = 500001
	BusinessErrorUserNotFund   int = 500002
	BusinessErrorUserPassword  int = 500003
	BusinessErrorUserLoginFail int = 500004
)

// BusinessErrorMap 因GO目前无法定义常量map 故定义为全局包中的map
var BusinessErrorMap = map[int]string{
	BusinessErrorUserSave:     "用户信息保存失败！",
	BusinessErrorUserNotFund:  "用户信息不存在！",
	BusinessErrorUserPassword: "用户密码不正确！",
}
