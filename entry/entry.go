package entry

import (
	"fmt"
	"runtime"
)

//定义key,value 结构
type Entry struct {
	key   interface{}
	value interface{}
}

func NewEntry(key, value interface{}) *Entry {
	return &Entry{
		key:   key,
		value: value,
	}
}

func (e *Entry) Get() string {
	return fmt.Sprintf("key: %v, value: %v", e.key, e.value)
}

func (e *Entry) GetKey() interface{} {
	return e.key
}

func (e *Entry) GetValue() interface{} {
	return e.value
}

func (e *Entry) SetValue(value interface{}) {
	e.value = value
}

//计算出元素占用内存字节数
func (e *Entry) Len() int {
	return CalcLen(e.value)
}

//计算value占用内存大小
func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	// case Value: // 结构体，数组，切片，map,要求实现 Value 接口，该接口只有1个 Len 方法，返回占用的内存字节数，如果没有实现该接口，则panic
	// 	n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, int8, uint8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.value", value))
	}

	return n
}
