package application

import (
	"github.com/cutexingluo/go-spring/core"
	"github.com/cutexingluo/go-spring/core/frame"
	"github.com/cutexingluo/go-spring/core/parse"
	"reflect"
)

// AutoConfig -   add tag-func 	parse.Initialize and parse.BeanInject  to  core.Context.BeanChains  自动配置
var AutoConfig = true

func SetAutoConfig(flag bool) {
	AutoConfig = flag
}

// Run starts the application. It will load the file at filePath and start the application.
// 主程序入口， 启动应用
func Run(lifeCycle ...interface{}) (err error) {

	if AutoConfig {
		frame.InitBeanFunc() // add func tag to  core.Context.BeanChains
	}

	// BeanCreated
	if err = core.BeanLifeCycleExecute(lifeCycle, func(lifeCycle any) error {
		return core.CreateHandlerExecute(lifeCycle)
	}); err != nil {
		return
	}
	if err = core.BeanCreatedFunc(); err != nil {
		return
	}

	// TagInitialized -> AfterInstantiation
	if err = core.TagInitializedFunc(); err != nil {
		return
	}
	if err = core.BeanLifeCycleExecute(lifeCycle, func(lifeCycle any) error {
		return core.AfterInstantiationHandlerExecute(lifeCycle)
	}); err != nil {
		return
	}

	// BeanInjected -> AfterInitialization
	if err = core.BeanInjectedFunc(); err != nil {
		return
	}
	if err = core.BeanLifeCycleExecute(lifeCycle, func(lifeCycle any) error {
		return core.AfterInitializationHandlerExecute(lifeCycle)
	}); err != nil {
		return
	}

	// ChooseMainHandler
	if err = core.BeanLifeCycleExecute(lifeCycle, func(lifeCycle any) error {
		return core.ChooseMainHandlerExecute(lifeCycle)
	}); err != nil {
		return
	}
	return
}

// AddBeanFilterFunc you can add bean filter function here.
// executionTime - such as core.BeanCreated, core.TagInitialized, core.BeanInjected
func AddBeanFilterFunc(executionTime int, beanFilterFunc *core.BeanFilterFunction) {
	frame.AddBeanFilterFunc(executionTime, beanFilterFunc)
}

// NewFuncInitByTagParser - new func init by tag, you can use it in AddBeanFilterFunc, like NewFuncInitByTag
func NewFuncInitByTagParser(tagParser *parse.TagParser) func(src *reflect.Value) (dst *reflect.Value, err error) {
	return parse.NewFuncInitByTag(tagParser)
}

// NewFuncInitByTag - new func init by tag, you can use it in AddBeanFilterFunc, like NewFuncInitByTagParser
func NewFuncInitByTag(fn func(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) error) func(src *reflect.Value) (dst *reflect.Value, err error) {
	return parse.NewFuncInitByTag(
		&parse.TagParser{ParseFunc: fn})
}
