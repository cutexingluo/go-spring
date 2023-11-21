package core

import "reflect"

// LifeCycle  life cycle, in file mode. 生命周期接口, 此处为文件模式
type LifeCycle interface {
	Create()              //  create beans to container.   创建时，在此处添加 bean
	AfterInstantiation()  // after bean tag initialization, 在此处 bean tag 已经被解析，但还未被注入 bean
	AfterInitialization() // after bean tag initialization and bean injection. Here are some starting actions 完成 bean 的注入 之后执行, 在此处进行开始操作
	OnMain()              // the main func of the file, execute main code here.  该文件的主程序，在此处执行主要代码
	Destroy()             // Perform a series of closing operations here.   销毁时，在此处执行一系列的收尾操作
}

// BeanLifeCycle - you can use the interface to implement the life cycle quickly. 生命周期接口, 你可以使用此接口快速实现生命周期
//
// no destroy function , You can use this interface to achieve the ability to apply Avoid executing the Destory method after the Run method.
// The main method is not fully applicable in the OnMain method, which can have more applications.
// 你可以使用该接口，来达到能够在application.Run方法后避免执行Destory方法。而且主方法不完全在OnMain方法里面可以能够有更多的应用
type BeanLifeCycle interface {
	CreateHandler
	AfterInstantiationHandler
	AfterInitializationHandler
	MainHandler
}

// InBeanLifeCycle - you can use the interface to implement the life cycle quickly. 生命周期接口, 你可以使用此接口快速实现生命周期
type InBeanLifeCycle interface {
	CreateHandler
	AfterInstantiationHandler
	AfterInitializationHandler
	OnMainHandler
	DestroyHandler
}

// CreateHandler create bean interface, you should implement this, bean  life cycle, in struct mode. 生命周期接口， 此处为结构体模式
type CreateHandler interface {
	//Create - create beans to container.   创建时，在此处添加 bean
	Create() *BeanConfig
}

// AfterInstantiationHandler update bean interface, you should implement this, bean  life cycle, in struct mode. 生命周期接口， 此处为结构体模式
type AfterInstantiationHandler interface {
	// AfterInstantiation -  after bean tag initialization, update beans to container.  tag 已经被解析，但还未被注入 bean   创建时，在此处更改 bean.
	AfterInstantiation(factory *BeanFactory) error
}

// AfterInitializationHandler update bean interface, you should implement this, bean  life cycle, in struct mode. 生命周期接口, 此处为结构体模式
type AfterInitializationHandler interface {
	// AfterInitialization - after bean tag initialization and bean injection. Here are some starting actions 完成 bean 的注入 之后执行, 在此处进行开始操作
	AfterInitialization(factory *BeanFactory) error
}

// MainHandler  main func of the file, execute main code here. like  OnMainHandler 该文件的主程序，在此处执行主要代码
//
// If both (MainHandler and  OnMainHandler) exist, only OnMainHandler will be executed,  如果和 OnMainHandler 同时存在，则只执行 OnMainHandler
type MainHandler interface {
	// OnMain - the main func of the file, execute main code here. is run only once. and will run Destroy func if it exists.
	//该文件的主程序 在此处执行主要代码
	OnMain(factory *BeanFactory) error
}

// OnMainHandler  main func of the file, execute main code here. like MainHandler  该文件的主程序，在此处执行主要代码
//
// If both (MainHandler and  OnMainHandler) exist, only OnMainHandler will be executed,  如果和 MainHandler 同时存在，则只执行 OnMainHandler
type OnMainHandler interface {
	// OnMain - the main func of the file, execute main code here. if useGo is not nil, the ret-fun(useGo) will be executed in go-routine .if runDestroy is true,  Destroy func will be executed.
	//该文件的主程序 在此处执行主要代码, useGo 返回 true 表示使用go-routine 来执行, runDestroy返回true, Destroy()方法将被执行
	OnMain(factory *BeanFactory) (useGo func(), runDestroy bool, err error)
}

// DestroyHandler Perform a series of closing operations here.   销毁时，在此处执行一系列的收尾操作
type DestroyHandler interface {
	// Destroy - Perform a series of closing operations here.   销毁时，在此处执行一系列的收尾操作
	Destroy(factory *BeanFactory) error
}

// BeanConfig  配置的bean, 只能配一个。如果两个配了，则使用MultiBeans
type BeanConfig struct {
	SingleBeans []interface{}          // single bean 单例Bean
	MultiBeans  map[string]interface{} // multi beans 注入的 bean列表,  map[beanName]bean
}

// BeanFilterFunction bean filter function
type BeanFilterFunction struct {
	// filter mode 1 is for single bean, 2 is for multi bean, 3 is for all bean
	// you can set core.SingleBeanMode , core.MultiBeanMode , core.AllBeanMode
	Mode   int
	Filter func(bean *reflect.Value) (ret *reflect.Value, err error) // filter function, if return nil do nothing
}

// in bean filter function
const (
	SingleBeanMode = iota + 1
	MultiBeanMode
	AllBeanMode
)

// Timing of execution (Execution timing) , it used by bean filter function slice in  BeanChains
const (
	BeanCreated = iota + 1
	TagInitialized
	BeanInjected = iota << 1
)
