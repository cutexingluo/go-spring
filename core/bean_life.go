package core

import (
	"github.com/cutexingluo/go-spring/common/base"
	"reflect"
)

// BeanCreatedFunc BeanCreated
func BeanCreatedFunc() (err error) {
	// BeanCreated
	allBeanNames := Context.BeanContainer.GetAllBeanNames() // get all bean names
	for i := BeanCreated; i < (1 << 3); i++ {
		if i&BeanCreated != 0 {
			err = updateFunction(i, allBeanNames)
			if err != nil {
				return err
			}
		}
	}
	return
}

// TagInitializedFunc TagInitialized
func TagInitializedFunc() (err error) {
	// TagInitialized
	allBeanNames := Context.BeanContainer.GetAllBeanNames() // get all bean names
	for i := TagInitialized; i < (1 << 3); i++ {
		if i&TagInitialized != 0 {
			err = updateFunction(i, allBeanNames)
			if err != nil {
				return err
			}
		}
	}
	return
}

// BeanInjectedFunc BeanInjected
func BeanInjectedFunc() error {
	// BeanInjected
	build, err := Context.BeanTopo.Build()
	if err != nil {
		return err
	}
	base.Reverse(build)
	for i := BeanInjected; i < (1 << 3); i++ {
		if i&BeanInjected != 0 {
			err := updateFunction(i, build)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// chain of execution functions
func updateFunction(executionTime int, beanNames []string) error {
	for _, funcStruct := range Context.BeanChains[executionTime] { // get the function struct
		for _, beanName := range beanNames { // iterate all bean names
			beanType := Context.BeanContainer.GetType(beanName)
			_, err := Context.BeanContainer.UpdateBeanFilter(
				beanName, beanType,
				func(isSingle bool, bean *reflect.Value) (ret *reflect.Value, err error) {
					if isSingle && funcStruct.Mode == SingleBeanMode ||
						!isSingle && funcStruct.Mode == MultiBeanMode ||
						funcStruct.Mode == AllBeanMode {
						return funcStruct.Filter(bean)
					}
					return nil, nil
				})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
