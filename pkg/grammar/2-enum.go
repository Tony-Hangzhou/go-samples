package grammar

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

/**
go 没有枚举类型，通过自增 iota + const组 达成
iota 可插队
*/

const y = 0x200

func Enum() {
	// println(&y) 不能获取常量的内存地址，因为常量在编译器预处理阶段直接展开，作为指令数据使用
	println(Sunday, Monday, Tuesday)
}
