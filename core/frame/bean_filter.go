package frame

import (
	"github.com/cutexingluo/go-spring/core"
)

// AddBeanFilterFunc you can add bean filter function here
func AddBeanFilterFunc(executionTime int, beanFilterFunc *core.BeanFilterFunction) {
	tagInitialized := core.Context.BeanChains[executionTime]
	tagInitialized = append(tagInitialized, beanFilterFunc)
	core.Context.BeanChains[executionTime] = tagInitialized
}

// InitBeanFunc Add BeanFilterFunction InitializeValue and BeanInject
func InitBeanFunc() {
	// parse.Initialize
	AddBeanFilterFunc(core.TagInitialized, &core.BeanFilterFunction{
		Mode:   core.AllBeanMode,
		Filter: InitializeValue,
	})

	// parse.BeanInject
	AddBeanFilterFunc(core.BeanInjected, &core.BeanFilterFunction{
		Mode:   core.AllBeanMode,
		Filter: BeanInject,
	})
}
