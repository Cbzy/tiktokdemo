package RegexpLib

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

func TestAlphabet(t *testing.T) {

	buf := "112a"
	str := []string{"123", "abc", "1aa", "33asd"}
	start := 1
	end := 3

	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(`^(\w){`)
	stringBuilder.WriteString(fmt.Sprint(start) + `,` + fmt.Sprint(end))
	stringBuilder.WriteString(`}$`)
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(stringBuilder.String())
	if reg1 == nil {
		t.Error(reg1)
		return
	}
	////根据规则提取关键信息
	//

	result1 := reg1.Match([]byte(buf))
	fmt.Println("result1 = ", result1)
	for _, s := range str {
		result1 := reg1.Match([]byte(s))
		fmt.Println("result1 = ", result1)
	}
}
