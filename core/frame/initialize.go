package frame

import (
	"github.com/cutexingluo/go-spring/core/bean_init"
	"reflect"
)

// Initialize -初始化目标对象，利用 tag `value="" default=""` 初始化值。
// 返回新对象。如果tag为空 不做处理。自行断言
func Initialize(src any) (value *reflect.Value, err error) {
	srcVal := reflect.ValueOf(src)
	return bean_init.Initialize(&srcVal)
}

// InitializeNoError -初始化目标对象，利用 tag `value="" default=""` 初始化值。返回新对象。如果tag为空 不做处理。不抛出错误, 自行断言
func InitializeNoError(src any) (value *reflect.Value) {
	initialize, _ := Initialize(src)
	return initialize
}

// InitializeByPtr -初始化指针目标对象，利用 tag `value=""  default=""` 初始化值。如果tag为空 不做处理。无需断言
func InitializeByPtr[T any](src *T) (value T, err error) {
	return bean_init.InitializeByPtr(src)
}

// InitializeByStruct -初始化结构体目标对象(新建对象(增加开销))，
// 返回新对象。利用 tag `value=""  default=""` 初始化值。如果tag为空 不做处理。无需断言
func InitializeByStruct[T any](src T) (value T, err error) {
	return bean_init.InitializeByStruct(src)
}

// InitializeValue -初始化目标对象，利用 tag `value="" default=""` 初始化值。
// 返回新对象。如果tag为空 不做处理。自行断言
func InitializeValue(val *reflect.Value) (value *reflect.Value, err error) {
	return bean_init.Initialize(val)
}

// BeanInject autowired the bean
func BeanInject(val *reflect.Value) (ret *reflect.Value, err error) {
	return bean_init.BeanInject(val)
}
