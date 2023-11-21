package core

// the core.Context is the global context variable, you can rebuild it yourself. you can also implement your own Container in it
// the main function is in application package, you can use application.Run()

// Context - Global context variable, you can rebuild it yourself. you can also implement your own Container in it
// 全局上下文变量，你可以自行重建该变量,  你也可以实现自己的 Container 替换里面的 BeanContainer
var Context = NewBeanFactory()

func AppContext() *BeanFactory {
	return Context
}
