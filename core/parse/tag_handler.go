package parse

import (
	"errors"
	"reflect"
)

type TagHandler interface {
	Parse(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) error // parse 解析函数, 返回解析后的对象
}

type TagParser struct {
	ParseFunc func(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) error // parse 解析函数, 返回解析后的对象
}

func (_this *TagParser) Parse(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) error {
	if _this.ParseFunc != nil {
		return _this.ParseFunc(srcVal, kind, srcField, structField)
	}
	return nil
}

// NewFuncInitByTag - new func init by tag
func NewFuncInitByTag(tagParser *TagParser) func(src *reflect.Value) (dst *reflect.Value, err error) {
	fn := func(src *reflect.Value) (dst *reflect.Value, err error) {
		return InitByTag(src, tagParser)
	}
	return fn
}

// FieldHandler 字段处理器
func FieldHandler(srcVal *reflect.Value, handler func(srcField *reflect.Value, fieldStruct *reflect.StructField)) {
	for i := 0; i < srcVal.NumField(); i++ {
		fieldVal := srcVal.Field(i)
		fieldStruct := srcVal.Type().Field(i)
		if handler != nil {
			handler(&fieldVal, &fieldStruct)
		}
	}
}

// InitByTagParser - 根据 tag 初始化指针对象或结构体对象, 返回对象. 推荐使用指针类型(引用)
func InitByTagParser(src *reflect.Value, tagMap map[reflect.Kind]TagParser) (dst any, err error) {
	return InitByTag(src, &TagParser{
		ParseFunc: func(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) error {
			var parser TagParser
			parser, ok := tagMap[*kind]
			if !ok {
				return nil
			}
			return parser.Parse(srcVal, kind, srcField, structField)
		},
	},
	)
}

// InitByTag - 根据 tag 初始化指针对象或结构体对象, 返回对象. 推荐使用指针类型(引用)
func InitByTag(src *reflect.Value, handler TagHandler) (dst *reflect.Value, err error) {
	t := src
	if t.Kind() != reflect.Struct && t.Kind() != reflect.Ptr {
		return src, errors.New("invalid type, must be  ptr type or struct type. ")
	}
	if t.Kind() == reflect.Ptr {
		ret, err := InitPtrByTag(src, handler)
		return ret, err
	} else if t.Kind() == reflect.Struct {
		ret, err := InitStructByTag(src, handler)
		return ret, err
	}
	return src, err
}

// InitPtrByTag - 根据 tag 初始化指针对象, 返回指针指向的对象
func InitPtrByTag(src *reflect.Value, handler TagHandler) (dst *reflect.Value, err error) {
	if src.Kind() != reflect.Ptr {
		return src, errors.New("invalid type , must be pointer type")
	}
	ret := src           // 新建对象, 返回指针
	dstVal := src.Elem() // 新建对象, 返回结构体
	// 设置 值
	FieldHandler(&dstVal, func(srcField *reflect.Value, fieldStruct *reflect.StructField) {
		kind := srcField.Kind()
		err = handler.Parse(&dstVal, &kind, srcField, fieldStruct)
		return
	})
	return ret, nil
}

// InitStructByTag - 根据 tag 初始化结构体, 返回新的对象
func InitStructByTag(src *reflect.Value, handler TagHandler) (dst *reflect.Value, err error) {
	if src.Kind() != reflect.Struct {
		return src, errors.New("invalid type, must be struct type")
	}
	ret := reflect.New(src.Type()) // 新建对象, 返回指针
	dstVal := ret.Elem()           // 新建对象, 返回结构体
	// 设置 值
	FieldHandler(&dstVal, func(srcField *reflect.Value, fieldStruct *reflect.StructField) {
		kind := srcField.Kind()
		err = handler.Parse(&dstVal, &kind, srcField, fieldStruct)
		return
	})
	return &ret, nil
}

// 		kind := srcField.Kind()
//		var parser TagParser
//		parser, ok := tagMap[kind]
//		if !ok {
//			return
//		}
//		val := parser.parse(kind, srcField, fieldStruct)
//		if val != nil {
//			srcField.Set(reflect.ValueOf(val))
//		}
