package reflect_util

import (
	"fmt"
	"reflect"
)

// Println 输出信息
func Println(val reflect.Value) {
	fmt.Println(" ----------------- reflect info ----------------- ")
	fmt.Println("reflect.Value: ", val)
	fmt.Println("reflect interface(): ", val.Interface())
	fmt.Println("reflect Type(): ", val.Type())
	fmt.Println("reflect Kind(): ", val.Kind())
}

// PrintlnDfs   Dfs输出信息 (如果是指针，则递归)
func PrintlnDfs(val reflect.Value) {
	fmt.Println(" ----------------- reflect info ----------------- ")
	fmt.Println("reflect.Value: ", val)
	fmt.Println("reflect interface(): ", val.Interface())
	fmt.Println("reflect Type(): ", val.Type())
	fmt.Println("reflect Kind(): ", val.Kind())
	if val.Kind() == reflect.Ptr {
		fmt.Print(" -------  in reflect.ptr  ")
		PrintlnDfs(val.Elem())
		fmt.Println(" -------  out reflect.ptr ")
	}
}

// GetType 获取类型
func GetType(tar any) reflect.Type {
	return reflect.TypeOf(tar)
}

// GetStructByValue 统一转为结构体
func GetStructByValue(tar *reflect.Value) *reflect.Value {
	if tar.Kind() == reflect.Ptr {
		el := tar.Elem()
		return &el
	} else {
		return tar
	}
}

// GetStruct 统一转为结构体
func GetStruct[T any](tar T) reflect.Value {
	val := reflect.ValueOf(tar)
	if val.Kind() == reflect.Ptr {
		return val.Elem()
	} else {
		return val
	}
}

// GetPtrByValue 统一转为指针, 结构体会有复制开销 (因为避免被回收)
func GetPtrByValue(tar *reflect.Value) *reflect.Value {
	if tar.Kind() == reflect.Ptr {
		return tar
	}
	ptrValue := reflect.New(tar.Type())
	ptrValue.Elem().Set(*tar)
	return &ptrValue
}

// GetPtr 统一转为指针, 结构体会有复制开销 (因为避免被回收)
func GetPtr[T any](tar T) reflect.Value {
	value := reflect.ValueOf(tar)
	if value.Kind() == reflect.Ptr {
		return value
	}
	ptrValue := reflect.New(value.Type())
	ptrValue.Elem().Set(value)
	return ptrValue
}

// ValueOf 转为Value, 是 value 不转型
func ValueOf(tar any) *reflect.Value {
	if r, ok := tar.(reflect.Value); ok {
		return &r
	} else if r, ok := tar.(*reflect.Value); ok {
		return r
	}
	ret := reflect.ValueOf(tar)
	return &ret
}

// ValuesOf 批量转化为 reflect.Value  切片
func ValuesOf(tar ...any) []reflect.Value {
	values := make([]reflect.Value, len(tar))
	for index, value := range tar {
		values[index] = reflect.ValueOf(value)
	}
	return values
}

// ValuesByLenOf 批量转化, 固定长度 , 少于则nil , 多则舍去
func ValuesByLenOf(length int, tar ...any) []reflect.Value {
	values := make([]reflect.Value, length)
	for i := 0; i < length; i++ {
		values[i] = reflect.ValueOf(tar[i])
	}
	return values
}

// ValuesParse 批量解码
func ValuesParse(tar ...reflect.Value) []any {
	values := make([]any, len(tar))
	for index, value := range tar {
		values[index] = value.Interface()
	}
	return values
}

// ValueTypeEquals 判断类型是否相等, 指针和结构体统一
func ValueTypeEquals(a, b reflect.Value) bool {
	return TypeEquals(a.Type(), b.Type())
}

// TypeEquals 判断类型是否相等, 指针和结构体统一
func TypeEquals(ta, tb reflect.Type) bool {
	return GetStructType(ta) == GetStructType(tb)
}

// GetStructType 获取结构体类型
func GetStructType(ta reflect.Type) reflect.Type {
	if ta.Kind() == reflect.Ptr {
		ta = ta.Elem()
	}
	return ta
}

//--------------------ReflectInfo----------------

// ReflectInfo 反射信息
type ReflectInfo struct {
	Target reflect.Value
}

// NewReflectInfo 新建反射信息
func NewReflectInfo(tar reflect.Value) *ReflectInfo {
	return &ReflectInfo{
		Target: tar,
	}
}

// NewReflectInfoByValue 新建反射信息
func NewReflectInfoByValue(value any) *ReflectInfo {
	val := reflect.ValueOf(value)
	return &ReflectInfo{
		Target: val,
	}
}

// CallMethod 普通形式调用方法
func (_this *ReflectInfo) CallMethod(methodName string, args ...any) (ret []any, err error) {
	t := _this.Target.Type()
	if method, ok := t.MethodByName(methodName); ok {
		newSlice := make([]any, len(args)+1)
		newSlice[0] = _this.Target.Interface()
		copy(newSlice[1:], args)
		value := method.Func.Call(ValuesOf(newSlice...))
		return ValuesParse(value...), nil
	}
	return nil, fmt.Errorf("the method '%s' is not found", methodName)
}

// CallMethodByFixedLen 普通形式调用方法, 固定长度, 少于则赋nil ,多则舍去
func (_this *ReflectInfo) CallMethodByFixedLen(methodName string, args ...any) (ret []any, err error) {
	t := _this.Target.Type()
	if method, ok := t.MethodByName(methodName); ok {
		length := method.Type.NumIn() // 固定长度
		newSlice := make([]any, length+1)
		newSlice[0] = _this.Target.Interface()
		copy(newSlice[1:], args)
		value := method.Func.Call(ValuesOf(newSlice...))
		return ValuesParse(value...), nil
	}
	return nil, fmt.Errorf("the method '%s' is not found", methodName)
}
