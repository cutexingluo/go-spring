package bean_init

import (
	"github.com/cutexingluo/go-spring/common/reflect_util"
	"github.com/cutexingluo/go-spring/core"
	"github.com/cutexingluo/go-spring/core/parse"
	"reflect"
)

// BeanInject autowired the bean . tag `bean:""` 利用 tag 注入
func BeanInject(val *reflect.Value) (ret *reflect.Value, err error) {
	return parse.InitByTag(val, &parse.TagParser{ParseFunc: ParseBean})
}

// ParseBean  解析某个对象的Bean
func ParseBean(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) (err error) {
	if parse.IsSupportBean(*kind) { // Bean 类型
		tagValue := structField.Tag.Get("bean")
		if tagValue == "" {
			return nil
		}
		bean, err := core.Context.BeanContainer.GetBean(tagValue) // 获取bean
		if err != nil || bean == nil {
			return err
		}
		if *kind == reflect.Ptr { //指针则赋值
			srcField.Set(*reflect_util.GetPtrByValue(bean))
		} else { // copy 一份
			srcField.Set(*reflect_util.GetStructByValue(bean))
		}
	}
	return nil
}
