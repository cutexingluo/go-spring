package parse

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// ParseStrSetValue parse string to value
func ParseStrSetValue(fieldKind *reflect.Kind, field *reflect.Value, strValue string, sliceCap int) (err error) {
	switch *fieldKind {
	case reflect.String:
		field.SetString(strValue)
	case reflect.Bool:
		val, err := strconv.ParseBool(strValue)
		field.SetBool(val)
		return err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value, err := strconv.ParseUint(strValue, 10, 64)
		field.SetUint(value)
		return err
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value, err := strconv.ParseInt(strValue, 10, 64)
		field.SetInt(value)
		return err
	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(strValue, 64)
		field.SetFloat(value)
		return err
	case reflect.Complex64, reflect.Complex128:
		value, err := strconv.ParseComplex(strValue, 128)
		field.SetComplex(value)
		return err
	case reflect.Slice: // ["a", "b", "c"]
		err = ParseStrSetSlice(fieldKind, field, strValue, sliceCap)
	case reflect.Map: // ["a":"hello", "b":"world"]
		err = ParseStrSetMap(fieldKind, field, strValue)
	case reflect.Array: // ["a", "b", "c"]
		fallthrough
	case reflect.Chan: // ["a", "b", "c"]
		fallthrough
	default:
		return
	}
	return
}

// ParseStrSetSlice set the slice
func ParseStrSetSlice(fieldKind *reflect.Kind, field *reflect.Value, strValue string, cap int) (err error) {
	if cap < 0 {
		cap = 0
	}
	keyKind := GetKeyKind(fieldKind, field)
	switch keyKind {
	case reflect.String:
		slice := make([]string, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Bool:
		slice := make([]bool, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Int:
		slice := make([]int, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Int8:
		slice := make([]int8, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Int16:
		slice := make([]int16, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Int32:
		slice := make([]int32, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Int64:
		slice := make([]int64, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Uint:
		slice := make([]uint, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Uint8:
		slice := make([]uint8, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Uint16:
		slice := make([]uint16, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Uint32:
		slice := make([]uint32, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Uint64:
		slice := make([]uint64, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Float32:
		slice := make([]float32, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	case reflect.Float64:
		slice := make([]float64, cap)
		err := json.Unmarshal([]byte(strValue), &slice)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(slice))
	default:
		return
	}
	return
}

// ParseStrSetMap set map field
func ParseStrSetMap(fieldKind *reflect.Kind, field *reflect.Value, strValue string) (err error) {
	keyKind := GetKeyKind(fieldKind, field)
	valueKind := GetValueKind(fieldKind, field)
	switch keyKind {
	case reflect.String:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[string]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[string]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[string]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[string]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[string]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[string]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[string]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[string]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[string]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[string]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[string]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[string]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[string]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[string]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[string]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	case reflect.Bool:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[bool]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[bool]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[bool]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[bool]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[bool]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[bool]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[bool]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[bool]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[bool]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[bool]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[bool]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[bool]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[bool]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[bool]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[bool]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	case reflect.Int:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[int]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[int]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[int]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[int]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[int]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[int]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[int]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[int]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[int]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[int]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[int]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[int]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[int]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[int]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[int]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	case reflect.Int8:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[int8]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[int8]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[int8]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[int8]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[int8]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[int8]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[int8]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[int8]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[int8]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[int8]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[int8]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[int8]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[int8]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[int8]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[int8]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}

	case reflect.Int16:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[int16]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[int16]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[int16]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[int16]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[int16]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[int16]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[int16]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[int16]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[int16]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[int16]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[int16]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[int16]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[int16]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[int16]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[int16]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	case reflect.Int32:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[int32]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[int32]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[int32]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[int32]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[int32]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[int32]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[int32]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[int32]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[int32]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[int32]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[int32]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[int32]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[int32]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[int32]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	default:
		mapValue := make(map[int32]any)
		err := json.Unmarshal([]byte(strValue), &mapValue)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(mapValue))
	case reflect.Int64:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[int64]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[int64]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[int64]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[int64]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[int64]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[int64]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[int64]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[int64]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[int64]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[int64]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[int64]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[int64]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[int64]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[int64]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[int64]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}

	case reflect.Uint:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[uint]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[uint]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[uint]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[uint]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[uint]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[uint]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[uint]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[uint]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[uint]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[uint]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[uint]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[uint]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[uint]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[uint]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[uint]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	case reflect.Uint8:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[uint8]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[uint8]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[uint8]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[uint8]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[uint8]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[uint8]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[uint8]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[uint8]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[uint8]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[uint8]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[uint8]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[uint8]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[uint8]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[uint8]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[uint8]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	case reflect.Uint16:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[uint16]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[uint16]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[uint16]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[uint16]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[uint16]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[uint16]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[uint16]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[uint16]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[uint16]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[uint16]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[uint16]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[uint16]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[uint16]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[uint16]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[uint16]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}

	case reflect.Uint32:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[uint32]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[uint32]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[uint32]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[uint32]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[uint32]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[uint32]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[uint32]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[uint32]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[uint32]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[uint32]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[uint32]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[uint32]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[uint32]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[uint32]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[uint32]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}

	case reflect.Uint64:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[uint64]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[uint64]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[uint64]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[uint64]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[uint64]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[uint64]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[uint64]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[uint64]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[uint64]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[uint64]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[uint64]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[uint64]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[uint64]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[uint64]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	case reflect.Float32:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[float32]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Bool:
			mapValue := make(map[float32]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[float32]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[float32]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[float32]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[float32]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[float32]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[float32]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[float32]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint16:
			mapValue := make(map[float32]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[float32]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[float32]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[float32]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[float32]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[float32]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}

	case reflect.Float64:
		switch valueKind {
		case reflect.String:
			mapValue := make(map[float64]string)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))

		case reflect.Bool:
			mapValue := make(map[float64]bool)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int:
			mapValue := make(map[float64]int)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int8:
			mapValue := make(map[float64]int8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int16:
			mapValue := make(map[float64]int16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int32:
			mapValue := make(map[float64]int32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Int64:
			mapValue := make(map[float64]int64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint:
			mapValue := make(map[float64]uint)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint8:
			mapValue := make(map[float64]uint8)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))

		case reflect.Uint16:
			mapValue := make(map[float64]uint16)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint32:
			mapValue := make(map[float64]uint32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Uint64:
			mapValue := make(map[float64]uint64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float32:
			mapValue := make(map[float64]float32)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		case reflect.Float64:
			mapValue := make(map[float64]float64)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		default:
			mapValue := make(map[float64]any)
			err := json.Unmarshal([]byte(strValue), &mapValue)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(mapValue))
		}
	}
	return
}
