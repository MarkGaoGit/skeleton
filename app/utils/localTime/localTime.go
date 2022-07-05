package localTime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"skeleton/app/global/variable"
	"strings"
	"time"
)

//LocalTime 自定义时间
// 原理 再数据库查询后 会根据绑定的数据结构中 使用反射来执行相应的结构转换
type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(variable.DateFormat, timeStr)
	*t = LocalTime(t1)
	return err
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(variable.DateFormat))

	//当时间为空则转换为空字符串
	if formatted == "\"0001-01-01 00:00:00\"" {
		formatted = "\"\""
	}

	return []byte(formatted), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(variable.DateFormat), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = LocalTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *LocalTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
