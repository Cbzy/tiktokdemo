package RegexpLib

import (
	"bytes"
	"fmt"
	"regexp"
)

func RegSome(str string, value string, start int, end int) bool {
	var stringBuilder bytes.Buffer
	switch str {
	case "AZ09":
		stringBuilder.WriteString(`^(\w){`)
		stringBuilder.WriteString(fmt.Sprint(start) + `,` + fmt.Sprint(end))
		stringBuilder.WriteString(`}$`)
		//解析正则表达式，如果成功返回解释器
		reg := regexp.MustCompile(stringBuilder.String())
		return reg.Match([]byte(value))

	}
	return false
}
