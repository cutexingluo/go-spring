package core

import (
	"fmt"
	"github.com/cutexingluo/go-spring/common/se/slice_util"
	"reflect"
	"strings"
)

// Container the bean container interface , Use reflect.Value prevents the loss of type information for the target object
type Container interface {
	HasBean(beanNameOrType string) bool                                        // HasBean checks whether the beanName in the container
	HasBeanType(beanType string) bool                                          // HasBeanByType checks whether the beanType in the container
	IsSingleBean(beanNameOrType string) bool                                   // IsSingleBean checks whether the bean is single
	IsMultiBean(beanName string) bool                                          // IsMultiBean checks whether the bean is multiple
	AddSingleBean(bean *reflect.Value) (isAdd bool, err error)                 // AddSingleBean add a single bean to the collection
	AddMultiBean(beanName string, bean *reflect.Value) (isAdd bool, err error) // AddMultiBean add a multi bean to the collection

	UpdateSingleBean(beanType string, bean *reflect.Value) (bool, error) // UpdateSingleBean update the single bean to the collection
	// UpdateSingleBeanFilter update the single bean to the collection by filter, the filter return nil is unchanged
	UpdateSingleBeanFilter(beanType string, filter func(bean *reflect.Value) (ret *reflect.Value, err error)) (bool, error)

	UpdateMultiBean(beanName string, beanType string, bean *reflect.Value) (bool, error) // UpdateMultiBean update the multi bean to the collection
	// UpdateMultiBeanFilter update a multi bean to the collection by filter, the filter return nil is unchanged
	UpdateMultiBeanFilter(beanName string, beanType string, filter func(bean *reflect.Value) (ret *reflect.Value, err error)) (bool, error)

	// UpdateBeanFilter update a bean to the collection by filter, the filter return nil is unchanged
	UpdateBeanFilter(beanName string, beanType string, filter func(isSingle bool, bean *reflect.Value) (ret *reflect.Value, err error)) (bool, error)

	GetAllBeanNames() []string                             // GetAllBeanNames get all bean names
	GetSingleBean(beanType string) *reflect.Value          // GetSingleBean get the single bean
	GetMultiBean(beanName string) *reflect.Value           // GetMultiBean get the multi bean
	GetBean(beanNameOrType string) (*reflect.Value, error) // GetBean get the bean
	GetMultiBeanNames(beanType string) []string            //  GetMultiBeanNames get all bean names
	GetType(beanName string) string                        // GetType get bean type, if it not found return ""

	RemoveSingleBean(beanType string) (bool, error)                 // RemoveSingleBean remove a single bean by beanType, if nil do nothing return false,nil
	RemoveMultiBean(beanName string, beanType string) (bool, error) // RemoveMultiBean remove a multi bean by beanName and beanType, if nil do nothing return false,nil
}

type BeanContainer struct {
	singleBeans    map[string]*reflect.Value // beanType -> bean , single
	multiBeans     map[string]*reflect.Value // beanName -> bean , multi
	types          map[string]string         // beanName -> beanType , single and multi
	beanCollection map[string][]string       // beanType -> beanNames , multi
	typeSet        map[string]int            // beanType 's  Set -> single is 1, multi is 2
}

// NewBeanContainer creates a new Container
func NewBeanContainer() *BeanContainer {
	return &BeanContainer{
		singleBeans:    make(map[string]*reflect.Value),
		multiBeans:     make(map[string]*reflect.Value),
		types:          make(map[string]string),
		beanCollection: make(map[string][]string),
		typeSet:        make(map[string]int),
	}
}

// HasBean checks whether the beanName in the container
func (_this *BeanContainer) HasBean(beanNameOrType string) bool {
	beanNameOrType = BeanNameFilter(beanNameOrType)
	if beanNameOrType == "" {
		return false
	}
	return _this.checkSingleBean(beanNameOrType) || _this.checkMultiBean(beanNameOrType)
}

// HasBeanType checks whether the beanType in the container
func (_this *BeanContainer) HasBeanType(beanType string) bool {
	beanType = BeanNameFilter(beanType)
	if beanType == "" {
		return false
	}
	return _this.hasBeanType(beanType)
}
func (_this *BeanContainer) hasBeanType(beanType string) bool {
	return _this.typeSet[beanType] != 0
}

// IsSingleBean checks whether the bean is single
func (_this *BeanContainer) IsSingleBean(beanNameOrType string) bool {
	beanNameOrType = BeanNameFilter(beanNameOrType)
	if beanNameOrType == "" {
		return false
	}
	return _this.checkSingleBean(beanNameOrType)
}
func (_this *BeanContainer) checkSingleBean(beanNameOrType string) bool {
	return (_this.typeSet[beanNameOrType] == 1 || _this.singleBeans[beanNameOrType] != nil) && _this.types[beanNameOrType] == beanNameOrType
}

// IsMultiBean checks whether the bean is multiple
func (_this *BeanContainer) IsMultiBean(beanName string) bool {
	beanName = BeanNameFilter(beanName)
	if beanName == "" {
		return false
	}
	return _this.checkMultiBean(beanName)
}

// IsMultiBean checks whether the bean is multiple
func (_this *BeanContainer) checkMultiBean(beanName string) bool {
	return _this.types != nil && _this.multiBeans[beanName] != nil
}

// IsMultiBean checks whether the type is multiple
func (_this *BeanContainer) checkMultiType(beanType string) bool {
	return _this.types != nil && (_this.typeSet[beanType] == 2 || _this.multiBeans[beanType] != nil || len(_this.beanCollection) > 0)
}

//------------- add--------------

// addSingleBean add a single bean to the collection, unsafe
func (_this *BeanContainer) addSingleBean(beanType string, bean *reflect.Value) {
	_this.typeSet[beanType] = 1 // single
	_this.singleBeans[beanType] = bean
	_this.types[beanType] = beanType
}

// AddSingleBean add a single bean to the collection
func (_this *BeanContainer) AddSingleBean(bean *reflect.Value) (isAdd bool, err error) {
	beanType := bean.Type().String()
	beanType = BeanNameFilter(beanType)                                 // 统一类型
	if _this.hasBeanType(beanType) || _this.checkSingleBean(beanType) { // 已经存在
		return false, nil
	} else if _this.checkMultiBean(beanType) { // 如果是多实例，则注册失败
		return false, fmt.Errorf("this bean '%s' is not a single bean , because it is in the multiBeans list", beanType)
	}
	_this.addSingleBean(beanType, bean)
	return true, nil
}

// addTypeNameMapping adds the bean type name mapping, multiple
func (_this *BeanContainer) addTypeNameMapping(beanType string, beanName string) bool {
	if _, ok := _this.types[beanName]; ok {
		return false
	}
	_this.types[beanName] = beanType
	_this.beanCollection[beanType] = append(_this.beanCollection[beanType], beanName)
	return true
}

// addMultiBean add a multi bean to the collection, unsafe
func (_this *BeanContainer) addMultiBean(beanName string, beanType string, bean *reflect.Value) {
	_this.typeSet[beanType] = 2 // multi
	_this.multiBeans[beanName] = bean
	_this.addTypeNameMapping(beanType, beanName)
}

// AddMultiBean add a multi bean to the collection
func (_this *BeanContainer) AddMultiBean(beanName string, bean *reflect.Value) (isAdd bool, err error) {
	beanName = BeanNameFilter(beanName)
	if beanName == "" {
		return false, nil
	}
	beanType := bean.Type().String()
	beanType = BeanNameFilter(beanType)                                     // 统一类型
	if _this.checkSingleBean(beanType) || _this.checkSingleBean(beanName) { // 已经存在
		return false, fmt.Errorf("this bean '%s' is not a multi bean , because it's typeBean is in the singleBeans list", beanName)
	} else if _this.checkMultiBean(beanName) { // 已经存在
		return false, nil
	}
	_this.addMultiBean(beanName, beanType, bean)
	return true, nil
}

// ------------- update--------------
// updateSingleBean update a single bean to the collection
func (_this *BeanContainer) updateSingleBean(beanType string, bean *reflect.Value) {
	_this.singleBeans[beanType] = bean
}

// UpdateSingleBean update the single bean to the collection
func (_this *BeanContainer) UpdateSingleBean(beanType string, bean *reflect.Value) (bool, error) {
	beanType = BeanNameFilter(beanType)
	if beanType == "" {
		return false, nil
	}
	if _this.checkMultiType(beanType) { // 如果是多实例，则更改失败
		return false, fmt.Errorf("this bean '%s' is not a single bean , because it has multiBeans ", beanType)
	}
	if _this.checkSingleBean(beanType) { // 如果存在
		_this.updateSingleBean(beanType, bean)
		return true, nil
	} else {
		return false, nil
	}
}

// UpdateSingleBeanFilter update the single bean to the collection by filter
func (_this *BeanContainer) UpdateSingleBeanFilter(beanType string, filter func(bean *reflect.Value) (ret *reflect.Value, err error)) (bool, error) {
	beanType = BeanNameFilter(beanType)
	if beanType == "" {
		return false, nil
	}
	if _this.checkMultiType(beanType) { // 如果是多实例 则更改失败
		return false, fmt.Errorf("this bean '%s' is not a single bean , because it has multiBeans ", beanType)
	}
	if _this.checkSingleBean(beanType) { // 如果存在
		bean, err := filter(_this.singleBeans[beanType])
		if err != nil {
			return false, err
		}
		if bean != nil {
			_this.updateSingleBean(beanType, bean)
		}
		return true, nil
	} else {
		return false, nil
	}

}

// updateMultiBean update a multi bean to the collection
func (_this *BeanContainer) updateMultiBean(beanName string, bean *reflect.Value) {
	_this.multiBeans[beanName] = bean
}

// UpdateMultiBean update a multi bean to the collection
func (_this *BeanContainer) UpdateMultiBean(beanName string, beanType string, bean *reflect.Value) (bool, error) {
	beanName = BeanNameFilter(beanName) // 统一类型
	beanType = BeanNameFilter(beanType) // 统一类型
	if beanName == "" || beanType == "" {
		return false, nil
	}
	if _this.checkSingleBean(beanType) || _this.checkSingleBean(beanName) { // 如果是单实例，则更改失败
		return false, fmt.Errorf("this bean 'type: '%s' , bean: '%s' ' is not a multi bean, because it has a singleBean ", beanType, beanName)
	} else if _this.checkMultiType(beanType) { // 如果存在
		_this.updateMultiBean(beanName, bean)
		return true, nil
	} else {
		return false, nil
	}
}

// UpdateMultiBeanFilter update a multi bean to the collection by filter
func (_this *BeanContainer) UpdateMultiBeanFilter(beanName string, beanType string, filter func(bean *reflect.Value) (ret *reflect.Value, err error)) (bool, error) {
	beanName = BeanNameFilter(beanName) // 统一类型
	beanType = BeanNameFilter(beanType) // 统一类型
	if beanName == "" || beanType == "" {
		return false, nil
	}
	if _this.checkSingleBean(beanType) || _this.checkSingleBean(beanName) { // 如果是单实例 则更改失败
		return false, fmt.Errorf("this bean 'type: '%s' , bean: '%s' ' is not a multi bean, because it has a singleBean ", beanType, beanName)
	} else if _this.checkMultiBean(beanName) { // 如果存在
		bean, err := filter(_this.multiBeans[beanName])
		if err != nil {
			return false, err
		}
		if bean != nil {
			_this.updateMultiBean(beanName, bean)
		}
		return true, nil
	} else {
		return false, nil
	}
}

// UpdateBeanFilter update a bean to the collection by filter
func (_this *BeanContainer) UpdateBeanFilter(beanName string, beanType string, filter func(isSingle bool, bean *reflect.Value) (ret *reflect.Value, err error)) (bool, error) {
	beanName = BeanNameFilter(beanName) // 统一类型
	beanType = BeanNameFilter(beanType) // 统一类型
	if beanName == "" || beanType == "" {
		return false, nil
	}
	if _this.checkSingleBean(beanType) || _this.checkSingleBean(beanName) { // 如果是单实例
		bean, err := filter(true, _this.singleBeans[beanType])
		if err != nil {
			return false, err
		}
		if bean != nil {
			_this.updateSingleBean(beanType, bean)
		}
		return true, nil
	} else if _this.checkMultiBean(beanName) { // 如果存在多例
		bean, err := filter(false, _this.multiBeans[beanName])
		if err != nil {
			return false, err
		}
		if bean != nil {
			_this.updateMultiBean(beanName, bean)
		}
		return true, nil
	} else {
		return false, nil
	}
}

//------------- get--------------

// GetAllBeanNames get all bean names
func (_this *BeanContainer) GetAllBeanNames() []string {
	typeSlice := make([]string, len(_this.types))
	var index = 0
	for k := range _this.types {
		typeSlice[index] = k
		index++
	}
	return typeSlice
}

// GetType get bean type, if it not found return ""
func (_this *BeanContainer) GetType(beanName string) string {
	beanName = BeanNameFilter(beanName) // 统一类型
	if beanName == "" {
		return ""
	}
	return _this.types[beanName]
}

// GetMultiBeanNames get all bean names
func (_this *BeanContainer) GetMultiBeanNames(beanType string) []string {
	return _this.beanCollection[beanType]
}

// GetSingleBean get a single bean by beanType, you must set  typeName pkg+structName if it is beanType
func (_this *BeanContainer) GetSingleBean(beanType string) *reflect.Value {
	beanType = BeanNameFilter(beanType) // 统一类型
	if beanType == "" {
		return nil
	}
	if _this.checkSingleBean(beanType) {
		return _this.singleBeans[beanType]
	}
	return nil
}

// GetMultiBean get a multi bean by beanName
func (_this *BeanContainer) GetMultiBean(beanName string) *reflect.Value {
	beanName = BeanNameFilter(beanName) // 统一类型
	if beanName == "" {
		return nil
	}
	if _this.checkMultiBean(beanName) {
		return _this.multiBeans[beanName]
	}
	return nil
}

// GetBean get a bean by beanName or beanType, you must set  typeName pkg+structName if it is beanType
func (_this *BeanContainer) GetBean(beanNameOrType string) (*reflect.Value, error) {
	beanNameOrType = BeanNameFilter(beanNameOrType) // 统一类型
	if beanNameOrType == "" {
		return nil, fmt.Errorf("this bean '%s' is not exist in the container", beanNameOrType)
	}
	if _this.checkSingleBean(beanNameOrType) {
		return _this.GetSingleBean(beanNameOrType), nil
	} else if _this.checkMultiBean(beanNameOrType) {
		return _this.GetMultiBean(beanNameOrType), nil
	} else {
		// 返回错误 拼接错误
		return nil, fmt.Errorf("this bean '%s' is not exist in the container", beanNameOrType)
	}
}

// ------------- remove--------------
func (_this *BeanContainer) removeSingleBean(beanType string) {
	delete(_this.singleBeans, beanType)
	delete(_this.types, beanType)
}

// RemoveSingleBean remove a single bean by beanType, if nil do nothing return false,nil
func (_this *BeanContainer) RemoveSingleBean(beanType string) (bool, error) {
	beanType = BeanNameFilter(beanType) // 统一类型
	if beanType == "" {
		return false, nil
	}
	if _this.checkSingleBean(beanType) {
		_this.removeSingleBean(beanType)
		return true, nil
	} else if _this.checkMultiBean(beanType) {
		return false, fmt.Errorf("this bean '%s' is not a single bean, because it is in the multiBeans list", beanType)
	} else {
		return false, nil
	}
}

// removeMultiBean remove a multi bean by beanName and beanType
func (_this *BeanContainer) removeMultiBean(beanName string, beanType string) {
	delete(_this.multiBeans, beanName)
	delete(_this.types, beanName)
	beans := _this.beanCollection[beanType]
	slice_util.RemoveFirstString(beans, beanName) // 效率低
}

// RemoveMultiBean remove a multi bean by beanName and beanType, if nil do nothing return false,nil
// 不推荐删除，因为删除效率低
func (_this *BeanContainer) RemoveMultiBean(beanName string, beanType string) (bool, error) {
	beanType = BeanNameFilter(beanType) // 统一类型
	beanName = BeanNameFilter(beanName) // 统一类型
	if beanType == "" || beanName == "" {
		return false, nil
	}
	if _this.checkMultiBean(beanName) && _this.checkMultiType(beanType) {
		_this.removeMultiBean(beanName, beanType)
		return true, nil
	} else if _this.IsSingleBean(beanType) || _this.checkSingleBean(beanName) {
		return false, fmt.Errorf("this bean ' type : '%s', bean: '%s' ' is not a multi bean, because it is in the singleBeans list", beanType, beanName)
	} else {
		return false, nil
	}
}

// BeanNameFilter 统一指针和结构体的类型
func BeanNameFilter(beanName string) string {
	if beanName == "" {
		return beanName
	}
	s := strings.TrimSpace(beanName)
	if len(s) > 0 && s[0] == '*' {
		s = strings.TrimSpace(s[1:])
	}
	return s
}
