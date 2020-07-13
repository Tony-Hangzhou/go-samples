package grammar

import (
	"encoding/json"
	"fmt"
)

//可以选择的控制字段有三种：
// -：不要解析这个字段
// omitempty：当字段为空（默认值）时，不要解析这个字段。比如 false、0、nil、长度为 0 的 array，map，slice，string
// FieldName：当解析 json 的时候，使用这个名字
type StudentWithOption struct {
	StudentId      string //默认使用原定义中的值
	StudentName    string `json:"sname"`           // 解析（encode/decode） 的时候，使用 `sname`，而不是 `Field`
	StudentClass   string `json:"class,omitempty"` // 解析的时候使用 `class`，如果struct 中这个值为空，就忽略它
	StudentTeacher string `json:"-"`               // 解析的时候忽略该字段。默认情况下会解析这个字段，因为它是大写字母开头的
}

func JsonMain() {

	student := &StudentWithOption{StudentId: "1", StudentName: "fengxm", StudentClass: "0903", StudentTeacher: "feng"}
	jsonByte, _ := json.Marshal(student)

	fmt.Println(string(jsonByte))

	newStudent := new(StudentWithOption)
	json.Unmarshal(jsonByte, newStudent)
	fmt.Println(newStudent)
}
