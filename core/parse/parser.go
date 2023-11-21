package parse

import (
	"reflect"
	"strconv"
)

// IsSupport 支持的类型
func IsSupport(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.String, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.Slice, reflect.Map:
		return true
	case reflect.Complex64, reflect.Complex128: //不支持复数
		return false
	case reflect.Array, reflect.Chan: // 不支持数组和通道
		return false
	}
	return false
}

// IsSupportBasic 支持的基本类型
func IsSupportBasic(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.String, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		return true
	case reflect.Complex64, reflect.Complex128: //不支持复数
		return false
	}
	return false
}

// IsSupportComposite 支持的聚合类型
func IsSupportComposite(kind reflect.Kind) bool {
	switch kind {
	case reflect.Slice, reflect.Map:
		return true
	case reflect.Array, reflect.Chan: // 不支持数组和通道
		return false
	}
	return false
}

// IsSupportBean 支持的 bean 类型
func IsSupportBean(kind reflect.Kind) bool {
	return kind == reflect.Ptr || kind == reflect.Struct
}

// Parse 解析字符串为指定类型 (基本类型) 需自行转型
func Parse(kind reflect.Kind, strValue string) (value interface{}, err error) {
	switch kind {
	case reflect.Int:
		value, err = strconv.Atoi(strValue)
	case reflect.String:
		value = strValue
	case reflect.Bool:
		value, err = strconv.ParseBool(strValue)
	case reflect.Uint:
		parseUint, err := strconv.ParseUint(strValue, 10, 64)
		if err != nil {
			return nil, err
		}
		value = uint(parseUint)
	case reflect.Uint8:
		value, err = strconv.ParseUint(strValue, 10, 8)
	case reflect.Uint16:
		value, err = strconv.ParseUint(strValue, 10, 16)
	case reflect.Uint32:
		value, err = strconv.ParseUint(strValue, 10, 32)
	case reflect.Uint64:
		value, err = strconv.ParseUint(strValue, 10, 64)
	case reflect.Int8:
		value, err = strconv.ParseInt(strValue, 10, 8)
	case reflect.Int16:
		value, err = strconv.ParseInt(strValue, 10, 16)
	case reflect.Int32:
		value, err = strconv.ParseInt(strValue, 10, 32)
	case reflect.Int64:
		value, err = strconv.ParseInt(strValue, 10, 64)
	case reflect.Float32:
		value, err = strconv.ParseFloat(strValue, 32)
	case reflect.Float64:
		value, err = strconv.ParseFloat(strValue, 64)
	case reflect.Complex64:
		value, err = strconv.ParseComplex(strValue, 64)
	case reflect.Complex128:
		value, err = strconv.ParseComplex(strValue, 128)
	case reflect.Array: // ["a", "b", "c"]
		fallthrough
	case reflect.Map: // ["a":"hello", "b":"world"]
		fallthrough
	case reflect.Chan: // ["a", "b", "c"]
		fallthrough
	case reflect.Slice: // ["a", "b", "c"]
		fallthrough
	case reflect.Func, reflect.Struct, reflect.Ptr, reflect.Interface:
		fallthrough
	default:
		return
	}
	return
}

// GetKeyKind 获取key的类型
func GetKeyKind(kind *reflect.Kind, v *reflect.Value) (value reflect.Kind) {
	switch *kind {
	case reflect.Slice: // ["a", "b", "c"]
		return v.Type().Elem().Kind()
	case reflect.Array: // ["a", "b", "c"]
		return v.Type().Elem().Kind()
	case reflect.Map: // ["a":"hello", "b":"world"]
		return v.Type().Key().Kind()
	case reflect.Chan: // ["a", "b", "c"]
		return v.Type().Elem().Kind()
	case reflect.Func, reflect.Struct, reflect.Ptr, reflect.Interface:
		fallthrough
	default:
		return v.Kind()
	}
}

// GetValueKind  获取value的类型
func GetValueKind(kind *reflect.Kind, v *reflect.Value) (value reflect.Kind) {
	switch *kind {
	case reflect.Slice: // ["a", "b", "c"]
		return v.Type().Elem().Kind()
	case reflect.Array: // ["a", "b", "c"]
		return v.Type().Elem().Kind()
	case reflect.Map: // ["a":"hello", "b":"world"]
		return v.Type().Elem().Kind()
	case reflect.Chan: // ["a", "b", "c"]
		return v.Type().Elem().Kind()
	case reflect.Func, reflect.Struct, reflect.Ptr, reflect.Interface:
		fallthrough
	default:
		return v.Kind()
	}
}

// CreateKV 创建指定类型 (Array, Map, Chan, Slice) Array 和 Slice 一样 valueKind - map 的 key 类型  chanSize - chan长度，小于0则改为0
func CreateKV(kind *reflect.Kind, v *reflect.Value, chanSize int) (value interface{}) {
	switch *kind {
	case reflect.Array:
		fallthrough
	case reflect.Slice: // ["a", "b", "c"]
		keyKind := GetKeyKind(kind, v)
		return CreateSlice(keyKind)
	case reflect.Map: // ["a":"hello", "b":"world"]
		keyKind := GetKeyKind(kind, v)
		valueKind := GetValueKind(kind, v)
		return CreateMap(keyKind, valueKind)
	case reflect.Chan: // ["a", "b", "c"]
		keyKind := GetKeyKind(kind, v)
		return CreateChan(keyKind, chanSize)
	default:
		return
	}
}

// SetValueKV 设置指定类型 (Array, Map, Slice)
func SetValueKV(kind *reflect.Kind, tar *reflect.Value, strValue string, sliceCap int) (err error) {
	switch *kind {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		err = ParseStrSetSlice(kind, tar, strValue, sliceCap)
	case reflect.Map:
		err = ParseStrSetMap(kind, tar, strValue)
	}
	return
}
