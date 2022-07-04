package consts

//全局常量
const (
	ValidatorPrefix    string = "validator_prefix_" //验证器前缀
	RequestParamsError int    = 400001              //请求参数错误

	BusinessErrorUserSave int = 500001 //用户信息保存错误
)

// BusinessErrorMap 因GO目前无法定义常量map 故定义为全局包中的map
var BusinessErrorMap = map[int]string{
	BusinessErrorUserSave: "用户信息保存失败！",
}
