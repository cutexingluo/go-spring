package bean_init

import (
	"fmt"
	"github.com/cutexingluo/go-spring/core"
	"github.com/cutexingluo/go-spring/core/parse"
	"reflect"
	"strconv"
	"strings"
)

const (
	DefaultValue  = "default" // default value if it is nil.  默认值
	OverrideValue = "value"   // override value, existing values will also be overwritten.  覆盖值, 就算已存在的值也会被覆盖
	CapValue      = "cap"     // the slice cap value. 切片容量
	BeanValue     = "bean"    // the bean value.  bean 值
)

// Initialize -初始化目标对象(InitializeByPtr 和 InitializeByStruct 整合版本)，返回新对象，
// 使用 .Interface() 获取该对象，使用 .Addr() 获取目标指针的value对象。
// 利用 tag `value="" default=""` 初始化值。如果tag为空 不做处理。自行断言。
func Initialize(srcVal *reflect.Value) (value *reflect.Value, err error) {
	return parse.InitByTag(srcVal, &parse.TagParser{ParseFunc: ParseValue})
}

// InitializeByPtr -初始化指针目标对象，利用 tag `value=""  default=""` 初始化值。如果tag为空 不做处理。无需断言
func InitializeByPtr[T any](src *T) (value T, err error) {
	srcVal := reflect.ValueOf(src)
	dstVal, err := parse.InitPtrByTag(&srcVal, &parse.TagParser{ParseFunc: ParseValue})
	return dstVal.Interface().(T), err
}

// InitializeByStruct -初始化结构体目标对象(新建对象(增加开销))，利用 tag `value=""  default=""` 初始化值。如果tag为空 不做处理。无需断言
func InitializeByStruct[T any](src T) (value T, err error) {
	srcVal := reflect.ValueOf(src)
	dstVal, err := parse.InitStructByTag(&srcVal, &parse.TagParser{ParseFunc: ParseValue})
	return dstVal.Interface().(T), err
}

// ParseValue  解析某个对象的value
func ParseValue(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) (err error) {
	if parse.IsSupportBasic(*kind) { // 基本类型
		tagValue := strings.TrimSpace(structField.Tag.Get(OverrideValue))
		if tagValue == "" {
			defaultValue := strings.TrimSpace(structField.Tag.Get(DefaultValue)) // 默认值
			if defaultValue != "" && srcField.IsZero() {
				err = parse.ParseStrSetValue(kind, srcField, defaultValue, 0)
			}
			return err
		}
		err = parse.ParseStrSetValue(kind, srcField, tagValue, 0)
	} else if parse.IsSupportComposite(*kind) { // 聚合类型
		capStr := strings.TrimSpace(structField.Tag.Get(CapValue))
		var sliceCap = 0
		if *kind == reflect.Slice && capStr != "" {
			sliceCap, err = strconv.Atoi(capStr) // if error pass
		}
		tagValue := strings.TrimSpace(structField.Tag.Get(OverrideValue)) // 覆盖值
		if tagValue == "" {
			defaultValue := strings.TrimSpace(structField.Tag.Get(DefaultValue)) // 默认值
			if defaultValue != "" && srcField.IsNil() {
				err = parse.SetValueKV(kind, srcField, defaultValue, sliceCap)
			}
			return err
		}
		err = parse.SetValueKV(kind, srcField, tagValue, sliceCap)
	} else if parse.IsSupportBean(*kind) { // 如果是bean类型，则需要解析bean的属性
		//fmt.Println(srcVal, kind, srcField, structField, "is bean")
		beanName := core.BeanNameFilter(structField.Tag.Get(BeanValue)) //获取 bean name
		//fmt.Println(beanName)
		if beanName == "" {
			return nil
		}
		if !core.Context.HasBean(beanName) {
			return fmt.Errorf(" the bean '%s' is not in the container", beanName)
		}
		srcType := core.BeanNameFilter(srcVal.Type().String()) // 获取源对象类型

		//fmt.Println(srcType, beanName)
		if core.Context.IsMultiBean(srcType) {
			names := core.Context.BeanContainer.GetMultiBeanNames(srcType)
			for _, name := range names { // 该类型的所有bean全部与目标bean进行链接
				core.Context.BeanTopo.Add(name, beanName) // 添加边
			}
		} else if core.Context.IsSingleBean(srcType) {
			core.Context.BeanTopo.Add(srcType, beanName)
		}
	}
	return err
}
