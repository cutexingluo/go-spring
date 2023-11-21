package core

import (
	"github.com/cutexingluo/go-spring/common/reflect_util"
	"github.com/cutexingluo/go-spring/core/parse"
)

// FactoryContainer the bean container interface. Encapsulation of Container methods
type FactoryContainer interface {
	HasBean(beanNameOrType string) bool                                       // HasBean checks whether the beanName in the container
	HasBeanType(beanType string) bool                                         // HasBeanByType checks whether the beanType in the container
	IsSingleBean(beanNameOrType string) bool                                  // IsSingleBean checks whether the bean is single
	IsMultiBean(beanName string) bool                                         // IsMultiBean checks whether the bean is multiple
	AddSingleBean(bean any) (isAdd bool, err error)                           // AddSingleBean add a single bean to the collection
	AddMultiBean(beanName string, bean any) (isAdd bool, err error)           // AddMultiBean add a multi bean to the collection
	UpdateSingleBean(beanType string, bean any) (bool, error)                 // UpdateSingleBean update the single bean to the collection
	UpdateMultiBean(beanName string, beanType string, bean any) (bool, error) // UpdateMultiBean update the multi bean to the collection
	GetAllBeanNames() []string                                                // GetAllBeanNames get all bean names
	GetSingleBean(beanType string) any                                        // GetSingleBean get the single bean
	GetMultiBean(beanName string) any                                         // GetMultiBean get the multi bean
	GetBean(beanNameOrType string) (any, error)                               // GetBean get the bean
	RemoveSingleBean(beanType string) (bool, error)                           // RemoveSingleBean remove a single bean by beanType, if nil do nothing return false,nil
	RemoveMultiBean(beanName string, beanType string) (bool, error)           // RemoveMultiBean remove a multi bean by beanName and beanType, if nil do nothing return false,nil
}

// BeanFactory bean factory, you can implement your own Container
type BeanFactory struct {
	BeanContainer Container
	BeanTopo      *parse.BeanTopo
	//BeanInitQueue []string                     // after create initialize
	BeanChains map[int][]*BeanFilterFunction // bean filter function, can change the bean
}

func NewBeanFactoryByContainer(container Container) *BeanFactory {
	return &BeanFactory{
		BeanContainer: container,
		BeanTopo:      parse.NewBeanTopo(),
		BeanChains:    make(map[int][]*BeanFilterFunction),
	}
}

func NewBeanFactory() *BeanFactory {
	return NewBeanFactoryByContainer(NewBeanContainer())
}

// HasBean checks whether the beanName in the container
func (_this *BeanFactory) HasBean(beanNameOrType string) bool {
	return _this.BeanContainer.HasBean(beanNameOrType)
}

// HasBeanType checks whether the beanType in the container
func (_this *BeanFactory) HasBeanType(beanType string) bool {
	return _this.BeanContainer.HasBeanType(beanType)
}

// IsSingleBean checks whether the bean is single
func (_this *BeanFactory) IsSingleBean(beanNameOrType string) bool {
	return _this.BeanContainer.IsSingleBean(beanNameOrType)
}

// IsMultiBean checks whether the bean is multiple
func (_this *BeanFactory) IsMultiBean(beanName string) bool {
	return _this.BeanContainer.IsMultiBean(beanName)
}

// AddSingleBean add a single bean to the collection
func (_this *BeanFactory) AddSingleBean(bean any) (isAdd bool, err error) {
	ptr := reflect_util.GetPtr(bean)
	return _this.BeanContainer.AddSingleBean(&ptr)
}

// AddMultiBean add a multi bean to the collection
func (_this *BeanFactory) AddMultiBean(beanName string, bean any) (isAdd bool, err error) {
	ptr := reflect_util.GetPtr(bean)
	return _this.BeanContainer.AddMultiBean(beanName, &ptr)
}

// UpdateSingleBean update the single bean to the collection
func (_this *BeanFactory) UpdateSingleBean(beanType string, bean any) (bool, error) {
	ptr := reflect_util.GetPtr(bean)
	return _this.BeanContainer.UpdateSingleBean(beanType, &ptr)
}

// UpdateMultiBean update the multi bean to the collection
func (_this *BeanFactory) UpdateMultiBean(beanName string, beanType string, bean any) (bool, error) {
	ptr := reflect_util.GetPtr(bean)
	return _this.BeanContainer.UpdateMultiBean(beanName, beanType, &ptr)
}

// GetAllBeanNames get all bean names
func (_this *BeanFactory) GetAllBeanNames() []string {
	return _this.BeanContainer.GetAllBeanNames()
}

// GetSingleBean get the single bean, you must set beanType pkg+structName
func (_this *BeanFactory) GetSingleBean(beanType string) any {
	ret := _this.BeanContainer.GetSingleBean(beanType)
	if ret == nil {
		return nil
	}
	return ret.Interface()
}

// GetMultiBean get the multi bean
func (_this *BeanFactory) GetMultiBean(beanName string) any {
	ret := _this.BeanContainer.GetMultiBean(beanName)
	if ret == nil {
		return nil
	}
	return ret.Interface()
}

// GetBean get the bean, you must set  typeName pkg+structName if it is beanType
func (_this *BeanFactory) GetBean(beanNameOrType string) (any, error) {
	ret, err := _this.BeanContainer.GetBean(beanNameOrType)
	if ret == nil {
		return nil, err
	}
	return ret.Interface(), err
}

// RemoveSingleBean remove a single bean by beanType, if nil do nothing return false,nil
func (_this *BeanFactory) RemoveSingleBean(beanType string) (bool, error) {
	return _this.BeanContainer.RemoveSingleBean(beanType)
}

// RemoveMultiBean remove a multi bean by beanName and beanType, if nil do nothing return false,nil
func (_this *BeanFactory) RemoveMultiBean(beanName string, beanType string) (bool, error) {
	return _this.BeanContainer.RemoveMultiBean(beanName, beanType)
}
