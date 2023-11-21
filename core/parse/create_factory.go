package parse

import "reflect"

// CreateSlice 创建切片(Array,  Slice)
func CreateSlice(keyKind reflect.Kind) (value interface{}) {
	switch keyKind {
	case reflect.String:
		value = make([]string, 0)
	case reflect.Bool:
		value = make([]bool, 0)
	case reflect.Int:
		value = make([]int, 0)
	case reflect.Int8:
		value = make([]int8, 0)
	case reflect.Int16:
		value = make([]int16, 0)
	case reflect.Int32:
		value = make([]int32, 0)
	case reflect.Int64:
		value = make([]int64, 0)
	case reflect.Uint:
		value = make([]uint, 0)
	case reflect.Uint8:
		value = make([]uint8, 0)
	case reflect.Uint16:
		value = make([]uint16, 0)
	case reflect.Uint32:
		value = make([]uint32, 0)
	case reflect.Uint64:
		value = make([]uint64, 0)
	case reflect.Float32:
		value = make([]float32, 0)
	case reflect.Float64:
		value = make([]float64, 0)
	case reflect.Complex64:
		value = make([]complex64, 0)
	case reflect.Complex128:
		value = make([]complex128, 0)
	}
	return
}

// CreateChan 创建通道(Chan)
func CreateChan(keyKind reflect.Kind, chanSize int) (value interface{}) {
	if chanSize < 0 {
		chanSize = 0
	}
	switch keyKind {
	case reflect.String:
		value = make(chan string, chanSize)
	case reflect.Bool:
		value = make(chan bool, chanSize)
	case reflect.Int:
		value = make(chan int, chanSize)
	case reflect.Int8:
		value = make(chan int8, chanSize)
	case reflect.Int16:
		value = make(chan int16, chanSize)
	case reflect.Int32:
		value = make(chan int32, chanSize)
	case reflect.Int64:
		value = make(chan int64, chanSize)
	case reflect.Uint:
		value = make(chan uint, chanSize)
	case reflect.Uint8:
		value = make(chan uint8, chanSize)
	case reflect.Uint16:
		value = make(chan uint16, chanSize)
	case reflect.Uint32:
		value = make(chan uint32, chanSize)
	case reflect.Uint64:
		value = make(chan uint64, chanSize)
	case reflect.Float32:
		value = make(chan float32, chanSize)
	case reflect.Float64:
		value = make(chan float64, chanSize)
	case reflect.Complex64:
		value = make(chan complex64, chanSize)
	case reflect.Complex128:
		value = make(chan complex128, chanSize)
	}
	return
}

// CreateMap 创建Map
func CreateMap(keyKind reflect.Kind, valueKind reflect.Kind) (value interface{}) {
	switch keyKind {
	case reflect.String:
		switch valueKind {
		case reflect.String:
			value = make(map[string]string)
		case reflect.Bool:
			value = make(map[string]bool)
		case reflect.Int:
			value = make(map[string]int)
		case reflect.Int8:
			value = make(map[string]int8)
		case reflect.Int16:
			value = make(map[string]int16)
		case reflect.Int32:
			value = make(map[string]int32)
		case reflect.Int64:
			value = make(map[string]int64)
		case reflect.Uint:
			value = make(map[string]uint)
		case reflect.Uint8:
			value = make(map[string]uint8)
		case reflect.Uint16:
			value = make(map[string]uint16)
		case reflect.Uint32:
			value = make(map[string]uint32)
		case reflect.Uint64:
			value = make(map[string]uint64)
		case reflect.Float32:
			value = make(map[string]float32)
		case reflect.Float64:
			value = make(map[string]float64)
		case reflect.Complex64:
			value = make(map[string]complex64)
		case reflect.Complex128:
			value = make(map[string]complex128)
		default:
			value = make(map[string]any)
		}
	case reflect.Bool:
		switch valueKind {
		case reflect.String:
			value = make(map[bool]string)
		case reflect.Bool:
			value = make(map[bool]bool)
		case reflect.Int:
			value = make(map[bool]int)
		case reflect.Int8:
			value = make(map[bool]int8)
		case reflect.Int16:
			value = make(map[bool]int16)
		case reflect.Int32:
			value = make(map[bool]int32)
		case reflect.Int64:
			value = make(map[bool]int64)
		case reflect.Uint:
			value = make(map[bool]uint)
		case reflect.Uint8:
			value = make(map[bool]uint8)
		case reflect.Uint16:
			value = make(map[bool]uint16)
		case reflect.Uint32:
			value = make(map[bool]uint32)
		case reflect.Uint64:
			value = make(map[bool]uint64)
		case reflect.Float32:
			value = make(map[bool]float32)
		case reflect.Float64:
			value = make(map[bool]float64)
		case reflect.Complex64:
			value = make(map[bool]complex64)
		case reflect.Complex128:
			value = make(map[bool]complex128)
		default:
			value = make(map[bool]any)
		}
	case reflect.Int:
		switch valueKind {
		case reflect.String:
			value = make(map[int]string)
		case reflect.Bool:
			value = make(map[int]bool)
		case reflect.Int:
			value = make(map[int]int)
		case reflect.Int8:
			value = make(map[int]int8)
		case reflect.Int16:
			value = make(map[int]int16)
		case reflect.Int32:
			value = make(map[int]int32)
		case reflect.Int64:
			value = make(map[int]int64)
		case reflect.Uint:
			value = make(map[int]uint)
		case reflect.Uint8:
			value = make(map[int]uint8)
		case reflect.Uint16:
			value = make(map[int]uint16)
		case reflect.Uint32:
			value = make(map[int]uint32)
		case reflect.Uint64:
			value = make(map[int]uint64)
		case reflect.Float32:
			value = make(map[int]float32)
		case reflect.Float64:
			value = make(map[int]float64)
		case reflect.Complex64:
			value = make(map[int]complex64)
		case reflect.Complex128:
			value = make(map[int]complex128)
		default:
			value = make(map[int]any)
		}
	case reflect.Int8:
		switch valueKind {
		case reflect.String:
			value = make(map[int8]string)
		case reflect.Bool:
			value = make(map[int8]bool)
		case reflect.Int:
			value = make(map[int8]int)
		case reflect.Int8:
			value = make(map[int8]int8)
		case reflect.Int16:
			value = make(map[int8]int16)
		case reflect.Int32:
			value = make(map[int8]int32)
		case reflect.Int64:
			value = make(map[int8]int64)
		case reflect.Uint:
			value = make(map[int8]uint)
		case reflect.Uint8:
			value = make(map[int8]uint8)
		case reflect.Uint16:
			value = make(map[int8]uint16)
		case reflect.Uint32:
			value = make(map[int8]uint32)
		case reflect.Uint64:
			value = make(map[int8]uint64)
		case reflect.Float32:
			value = make(map[int8]float32)
		case reflect.Float64:
			value = make(map[int8]float64)
		case reflect.Complex64:
			value = make(map[int8]complex64)
		case reflect.Complex128:
			value = make(map[int8]complex128)
		default:
			value = make(map[int8]any)
		}
	case reflect.Int16:
		switch valueKind {
		case reflect.String:
			value = make(map[int16]string)
		case reflect.Bool:
			value = make(map[int16]bool)
		case reflect.Int:
			value = make(map[int16]int)
		case reflect.Int8:
			value = make(map[int16]int8)
		case reflect.Int16:
			value = make(map[int16]int16)
		case reflect.Int32:
			value = make(map[int16]int32)
		case reflect.Int64:
			value = make(map[int16]int64)
		case reflect.Uint:
			value = make(map[int16]uint)
		case reflect.Uint8:
			value = make(map[int16]uint8)
		case reflect.Uint16:
			value = make(map[int16]uint16)
		case reflect.Uint32:
			value = make(map[int16]uint32)
		case reflect.Uint64:
			value = make(map[int16]uint64)
		case reflect.Float32:
			value = make(map[int16]float32)
		case reflect.Float64:
			value = make(map[int16]float64)
		case reflect.Complex64:
			value = make(map[int16]complex64)
		case reflect.Complex128:
			value = make(map[int16]complex128)
		default:
			value = make(map[int16]any)
		}

	case reflect.Int32:
		switch valueKind {
		case reflect.String:
			value = make(map[int32]string)
		case reflect.Bool:
			value = make(map[int32]bool)
		case reflect.Int:
			value = make(map[int32]int)
		case reflect.Int8:
			value = make(map[int32]int8)
		case reflect.Int16:
			value = make(map[int32]int16)
		case reflect.Int32:
			value = make(map[int32]int32)
		case reflect.Int64:
			value = make(map[int32]int64)
		case reflect.Uint:
			value = make(map[int32]uint)
		case reflect.Uint8:
			value = make(map[int32]uint8)
		case reflect.Uint16:
			value = make(map[int32]uint16)
		case reflect.Uint32:
			value = make(map[int32]uint32)
		case reflect.Uint64:
			value = make(map[int32]uint64)
		case reflect.Float32:
			value = make(map[int32]float32)
		case reflect.Float64:
			value = make(map[int32]float64)
		case reflect.Complex64:
			value = make(map[int32]complex64)
		case reflect.Complex128:
			value = make(map[int32]complex128)
		default:
			value = make(map[int32]any)
		}

	case reflect.Int64:
		switch valueKind {
		case reflect.String:
			value = make(map[int64]string)
		case reflect.Bool:
			value = make(map[int64]bool)
		case reflect.Int:
			value = make(map[int64]int)
		case reflect.Int8:
			value = make(map[int64]int8)
		case reflect.Int16:
			value = make(map[int64]int16)
		case reflect.Int32:
			value = make(map[int64]int32)
		case reflect.Int64:
			value = make(map[int64]int64)
		case reflect.Uint:
			value = make(map[int64]uint)
		case reflect.Uint8:
			value = make(map[int64]uint8)
		case reflect.Uint16:
			value = make(map[int64]uint16)
		case reflect.Uint32:
			value = make(map[int64]uint32)
		case reflect.Uint64:
			value = make(map[int64]uint64)
		case reflect.Float32:
			value = make(map[int64]float32)
		case reflect.Float64:
			value = make(map[int64]float64)
		case reflect.Complex64:
			value = make(map[int64]complex64)
		case reflect.Complex128:
			value = make(map[int64]complex128)
		default:
			value = make(map[int64]any)
		}

	case reflect.Uint:
		switch valueKind {
		case reflect.String:
			value = make(map[uint]string)
		case reflect.Bool:
			value = make(map[uint]bool)
		case reflect.Int:
			value = make(map[uint]int)
		case reflect.Int8:
			value = make(map[uint]int8)
		case reflect.Int16:
			value = make(map[uint]int16)
		case reflect.Int32:
			value = make(map[uint]int32)
		case reflect.Int64:
			value = make(map[uint]int64)
		case reflect.Uint:
			value = make(map[uint]uint)
		case reflect.Uint8:
			value = make(map[uint]uint8)
		case reflect.Uint16:
			value = make(map[uint]uint16)
		case reflect.Uint32:
			value = make(map[uint]uint32)
		case reflect.Uint64:
			value = make(map[uint]uint64)
		case reflect.Float32:
			value = make(map[uint]float32)
		case reflect.Float64:
			value = make(map[uint]float64)
		case reflect.Complex64:
			value = make(map[uint]complex64)
		case reflect.Complex128:
			value = make(map[uint]complex128)
		default:
			value = make(map[uint]any)
		}

	case reflect.Uint8:
		switch valueKind {
		case reflect.String:
			value = make(map[uint8]string)
		case reflect.Bool:
			value = make(map[uint8]bool)
		case reflect.Int:
			value = make(map[uint8]int)
		case reflect.Int8:
			value = make(map[uint8]int8)
		case reflect.Int16:
			value = make(map[uint8]int16)
		case reflect.Int32:
			value = make(map[uint8]int32)
		case reflect.Int64:
			value = make(map[uint8]int64)
		case reflect.Uint:
			value = make(map[uint8]uint)
		case reflect.Uint8:
			value = make(map[uint8]uint8)
		case reflect.Uint16:
			value = make(map[uint8]uint16)
		case reflect.Uint32:
			value = make(map[uint8]uint32)
		case reflect.Uint64:
			value = make(map[uint8]uint64)
		case reflect.Float32:
			value = make(map[uint8]float32)
		case reflect.Float64:
			value = make(map[uint8]float64)
		case reflect.Complex64:
			value = make(map[uint8]complex64)
		case reflect.Complex128:
			value = make(map[uint8]complex128)
		default:
			value = make(map[uint8]any)
		}

	case reflect.Uint16:
		switch valueKind {
		case reflect.String:
			value = make(map[uint16]string)
		case reflect.Bool:
			value = make(map[uint16]bool)
		case reflect.Int:
			value = make(map[uint16]int)
		case reflect.Int8:
			value = make(map[uint16]int8)
		case reflect.Int16:
			value = make(map[uint16]int16)
		case reflect.Int32:
			value = make(map[uint16]int32)
		case reflect.Int64:
			value = make(map[uint16]int64)
		case reflect.Uint:
			value = make(map[uint16]uint)
		case reflect.Uint8:
			value = make(map[uint16]uint8)
		case reflect.Uint16:
			value = make(map[uint16]uint16)
		case reflect.Uint32:
			value = make(map[uint16]uint32)
		case reflect.Uint64:
			value = make(map[uint16]uint64)
		case reflect.Float32:
			value = make(map[uint16]float32)
		case reflect.Float64:
			value = make(map[uint16]float64)
		case reflect.Complex64:
			value = make(map[uint16]complex64)
		case reflect.Complex128:
			value = make(map[uint16]complex128)
		default:
			value = make(map[uint16]any)
		}

	case reflect.Uint32:
		switch valueKind {
		case reflect.String:
			value = make(map[uint32]string)
		case reflect.Bool:
			value = make(map[uint32]bool)
		case reflect.Int:
			value = make(map[uint32]int)
		case reflect.Int8:
			value = make(map[uint32]int8)
		case reflect.Int16:
			value = make(map[uint32]int16)
		case reflect.Int32:
			value = make(map[uint32]int32)
		case reflect.Int64:
			value = make(map[uint32]int64)
		case reflect.Uint:
			value = make(map[uint32]uint)
		case reflect.Uint8:
			value = make(map[uint32]uint8)
		case reflect.Uint16:
			value = make(map[uint32]uint16)
		case reflect.Uint32:
			value = make(map[uint32]uint32)
		case reflect.Uint64:
			value = make(map[uint32]uint64)
		case reflect.Float32:
			value = make(map[uint32]float32)
		case reflect.Float64:
			value = make(map[uint32]float64)
		case reflect.Complex64:
			value = make(map[uint32]complex64)
		case reflect.Complex128:
			value = make(map[uint32]complex128)
		default:
			value = make(map[uint32]any)
		}

	case reflect.Uint64:
		switch valueKind {
		case reflect.String:
			value = make(map[uint64]string)
		case reflect.Bool:
			value = make(map[uint64]bool)
		case reflect.Int:
			value = make(map[uint64]int)
		case reflect.Int8:
			value = make(map[uint64]int8)
		case reflect.Int16:
			value = make(map[uint64]int16)
		case reflect.Int32:
			value = make(map[uint64]int32)
		case reflect.Int64:
			value = make(map[uint64]int64)
		case reflect.Uint:
			value = make(map[uint64]uint)
		case reflect.Uint8:
			value = make(map[uint64]uint8)
		case reflect.Uint16:
			value = make(map[uint64]uint16)
		case reflect.Uint32:
			value = make(map[uint64]uint32)
		case reflect.Uint64:
			value = make(map[uint64]uint64)
		case reflect.Float32:
			value = make(map[uint64]float32)
		case reflect.Float64:
			value = make(map[uint64]float64)
		case reflect.Complex64:
			value = make(map[uint64]complex64)
		case reflect.Complex128:
			value = make(map[uint64]complex128)
		default:
			value = make(map[uint64]any)
		}

	case reflect.Float32:
		switch valueKind {
		case reflect.String:
			value = make(map[float32]string)
		case reflect.Bool:
			value = make(map[float32]bool)
		case reflect.Int:
			value = make(map[float32]int)
		case reflect.Int8:
			value = make(map[float32]int8)
		case reflect.Int16:
			value = make(map[float32]int16)
		case reflect.Int32:
			value = make(map[float32]int32)
		case reflect.Int64:
			value = make(map[float32]int64)
		case reflect.Uint:
			value = make(map[float32]uint)
		case reflect.Uint8:
			value = make(map[float32]uint8)
		case reflect.Uint16:
			value = make(map[float32]uint16)
		case reflect.Uint32:
			value = make(map[float32]uint32)
		case reflect.Uint64:
			value = make(map[float32]uint64)
		case reflect.Float32:
			value = make(map[float32]float32)
		case reflect.Float64:
			value = make(map[float32]float64)
		case reflect.Complex64:
			value = make(map[float32]complex64)
		case reflect.Complex128:
			value = make(map[float32]complex128)
		default:
			value = make(map[float32]any)
		}

	case reflect.Float64:
		switch valueKind {
		case reflect.String:
			value = make(map[float64]string)
		case reflect.Bool:
			value = make(map[float64]bool)
		case reflect.Int:
			value = make(map[float64]int)
		case reflect.Int8:
			value = make(map[float64]int8)
		case reflect.Int16:
			value = make(map[float64]int16)
		case reflect.Int32:
			value = make(map[float64]int32)
		case reflect.Int64:
			value = make(map[float64]int64)
		case reflect.Uint:
			value = make(map[float64]uint)
		case reflect.Uint8:
			value = make(map[float64]uint8)
		case reflect.Uint16:
			value = make(map[float64]uint16)
		case reflect.Uint32:
			value = make(map[float64]uint32)
		case reflect.Uint64:
			value = make(map[float64]uint64)
		case reflect.Float32:
			value = make(map[float64]float32)
		case reflect.Float64:
			value = make(map[float64]float64)
		case reflect.Complex64:
			value = make(map[float64]complex64)
		case reflect.Complex128:
			value = make(map[float64]complex128)
		default:
			value = make(map[float64]any)
		}

	case reflect.Complex64:
		switch valueKind {
		case reflect.String:
			value = make(map[complex64]string)
		case reflect.Bool:
			value = make(map[complex64]bool)
		case reflect.Int:
			value = make(map[complex64]int)
		case reflect.Int8:
			value = make(map[complex64]int8)
		case reflect.Int16:
			value = make(map[complex64]int16)
		case reflect.Int32:
			value = make(map[complex64]int32)
		case reflect.Int64:
			value = make(map[complex64]int64)
		case reflect.Uint:
			value = make(map[complex64]uint)
		case reflect.Uint8:
			value = make(map[complex64]uint8)
		case reflect.Uint16:
			value = make(map[complex64]uint16)
		case reflect.Uint32:
			value = make(map[complex64]uint32)
		case reflect.Uint64:
			value = make(map[complex64]uint64)
		case reflect.Float32:
			value = make(map[complex64]float32)
		case reflect.Float64:
			value = make(map[complex64]float64)
		case reflect.Complex64:
			value = make(map[complex64]complex64)
		case reflect.Complex128:
			value = make(map[complex64]complex128)
		default:
			value = make(map[complex64]any)
		}

	case reflect.Complex128:
		switch valueKind {
		case reflect.String:
			value = make(map[complex128]string)
		case reflect.Bool:
			value = make(map[complex128]bool)
		case reflect.Int:
			value = make(map[complex128]int)
		case reflect.Int8:
			value = make(map[complex128]int8)
		case reflect.Int16:
			value = make(map[complex128]int16)
		case reflect.Int32:
			value = make(map[complex128]int32)
		case reflect.Int64:
			value = make(map[complex128]int64)
		case reflect.Uint:
			value = make(map[complex128]uint)
		case reflect.Uint8:
			value = make(map[complex128]uint8)
		case reflect.Uint16:
			value = make(map[complex128]uint16)
		case reflect.Uint32:
			value = make(map[complex128]uint32)
		case reflect.Uint64:
			value = make(map[complex128]uint64)
		case reflect.Float32:
			value = make(map[complex128]float32)
		case reflect.Float64:
			value = make(map[complex128]float64)
		case reflect.Complex64:
			value = make(map[complex128]complex64)
		case reflect.Complex128:
			value = make(map[complex128]complex128)
		default:
			value = make(map[complex128]any)
		}
	default:
		value = make(map[any]any)
	}
	return
}
