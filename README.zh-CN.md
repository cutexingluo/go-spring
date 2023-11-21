# go-spring 介绍

一个基于 Golang 的简易容器框架，类似于Java的Spring容器。它提供了容器、标签初始化、bean注入等操作。由于go语言的特殊性质，它不能被代理，因此它提供了一个生命周期接口来提供服务。

里面除了提供容器，还提供了各种解析函数，也提供了一些类似Java的一些基类和数据结构，例如 Object接口(实现equals)，Iterator 接口 , LinkList , ArrayDeque 等。 



[English document 英文文档](./README.md)

[Chinese document 中文文档](./README.zh-CN.md)

## 1.模块介绍

开放的包 包括 **application , core , frame** 等。当然其他包也可以使用。

1. 你可以新建生命周期结构体，然后实现 core.BeanLifeCycle 或者 core.InBeanLifeCycle 来立马得到需要实现的方法。或者只实现里面的部分方法也可以。(包含 Create , AfterInstantiation , AfterInitialization , OnMain )
2. 然后使该结构体的 **Create** 的返回值返回 SingleBeans 或者 MultiBeans 。
3. 然后可以实现 **AfterInstantiation** 方法，该方法执行时机在 **tag 初始化后 ，bean注入之前**。
4. 然后可以实现 **AfterInitialization** 方法，该方法执行时机在 **bean注入之后**。
5. 然后可以实现 **OnMain** 方法，可以选的两种参数/返回值方法。在里面可以执行主要流程代码。
6. 以上四个方法实现是**可选**的，也就是说你可以一个也不实现。

## 2.注意事项

### 1.名词解释

**Context** : BeanFactory类型的全局单例，是该项目的核心。可以重新赋值自己实现的对象，也可以仅重新赋值里面的Container接口。通过 core.Context 或者 core.AppContext() 获取该对象，然后进行一系列编程式的对bean操作。也可以通过  core.Context.BeanContainer 获取更多操作的方法。

**MultiBean** : 多例 bean ,  可以通过生命周期添加或获取，或者通过 core.Context 添加或获取。添加时需要自行设置BeanName，然后会自动获取 BeanType (它的类型)，然后会添加到容器里面。此时若添加同类型 SingleBean 则会添加失败。

**SingleBean** : 单例 bean , 可以通过生命周期添加或获取，或者通过 core.Context 添加或获取。当一个结构体对象添加进去，它仅会保留一个单例，并且该对象的 BeanName 为该 **包名+结构体名**。此时添加同类型 MultiBean 则会添加失败。如果是在 Create 里面创建，则优先扫描 MultiBean 。

所以 MultiBean 和 SingleBean 是**互斥的**。如果仅创建一个单例，则推荐SingleBean，如果是多例则只能使用 MultiBean 。

**application.Run** : 参数为any, 会按照生命周期依次执行对象的方法。如果有多个对象参数，则会执行所有对象的一个阶段，再执行所有对象的下一阶段。

### 2.tag 解释

**default** : 如果添加的bean没有值，则默认会添加 tag(string) 转化后的值(如果转化失败则不赋值)。

**value** : 和 default 不同的是，value会强制赋值，无论bean里面是否有值。

两种方式仅支持简易的数据。包括除Complex的其他所有**基本类型**，以及简单的基本类型（除）对应的 slice 和 map 这两种**一级聚合类型**，例如 int , string, uint64, map[int]int, []int，[]string 等 ，以下是**不支持的** 例如 map[int] []int , []Human, [] [] int 等

**cap** : 为切片提供的一个 tag , 可以设置容量，其他设置则不会生效。

**bean** : 填写bean的名称，通过生命周期会自动添加到该位置，需要注意**防止循环依赖**。1.如果需要手动注入多例的某个bean又不想影响其他相同类型的bean，可以在注入后手动修改。2.如果目标字段是 ptr 类型则会直接赋值（修改该对象会影响容器里面的bean），如果是 struct 类型，则会**复制**一份赋给该字段

SingleBean : 单例 bean , 可以通过生命周期添加或获取，或者通过 core.Context 添加或获取，当一个结构体对象添加进去，它仅会保留一个单例，并且该对象的 BeanName 为该 **包名+结构体名**。此时添加其他

## 3.快速开始

### 1.普通实例

1.新建一个管理bean的结构体 和 两个结构体对象

```go
type BeanManager struct {
}

type Task struct {
	Name string `default:"task_you_must_do"` // default name
}

type Human struct {
	Name    string            `default:"name_value"`
	Age     int               `json:"age" default:"0"` // more tag
	Sex     int8              `default:"0"` // int8
	Address string            `default:"address_value"`
	CanSwim bool              `value:"true"`
	Bags    []int             `value:"[1,2,3,4,5,6]" default:"[2,4]"` // value override default
	Chs     []int             `value:"[1,2]"  cap:"6"`                //set  cap
	Mp      map[string]string `value:" { \"a\":\"b\",\"c\":\"d\",\"e\":\"d\" } "`
	Task    *Task             `bean:"task1"` // if it is single bean, it must be domain.Task
	Task2   Task              `bean:"task2"` // struct type not ptr, it will copy
}
```

2.实现声明周期 core.BeanLifeCycle 方法

```go
type BeanManager struct {
}
// add beans
func (_this *BeanManager) Create() *core.BeanConfig {
    // If it (bean) is not a ptr type, it will be converted to a ptr type
    // 如果添加的 bean 的不是 ptr 类型 ， 它会转为 ptr 类型并放入容器
	return &core.BeanConfig{ 
		SingleBeans: []interface{}{
			&domain.Human{ // ptr type . If it is a ptr type, it will be assigned directly; if it is a struct type, it will be copied
				Name: "xingluo",
				Age:  18,
			},
		},
		MultiBeans: map[string]interface{}{
			"task1": &domain.Task{  //This BeanName will be injected into the Human structure as a basis
				Name: "xingluo_will_do", // override default name
			},
			"task2": &domain.Task{},
		},
	}
}
//Can not be implemented
func (_this *BeanManager) AfterInstantiation(factory *core.BeanFactory) error {
	fmt.Println("------------------")
	fmt.Println("AfterInstantiation -->")
	bean, err := factory.GetBean("domain.Human")
	if err != nil {
		return err
	}
	fmt.Println(bean.(*domain.Human)) // it must be ptr type
	return nil
}
//Can not be implemented
func (_this *BeanManager) AfterInitialization(factory *core.BeanFactory) error {
	fmt.Println("------------------")
	fmt.Println("AfterInitialization -->")
	bean, err := factory.GetBean("domain.Human")
	if err != nil {
		return err
	}
	fmt.Println(bean.(*domain.Human))
	return nil
}

func (_this *BeanManager) OnMain(factory *core.BeanFactory) error {
	fmt.Println("------------------")
	fmt.Println("OnMain -->")
	bean, err := factory.GetBean("domain.Human")
	if err != nil {
		return err
	}
	human := bean.(*domain.Human)
	fmt.Println(human)
	fmt.Println(human.Task)
    fmt.Println(human.Task2)
	return nil
}
```

3.注册到执行链中并执行

```go
func main() {
    // main func 主程序入口
	err := application.Run( //Sequential execution
		&config.BeanManager{},
	)
	if err != nil {
		fmt.Println("err -> ", err)
		return
	}
	time.Sleep(time.Second * 1)
}

```



4.执行结果

```bash
------------------
AfterInstantiation -->
&{xingluo 18 0 address_value true [1 2 3 4 5 6] [1 2] map[a:b c:d e:d] <nil> {}}
------------------
AfterInitialization -->
&{xingluo 18 0 address_value true [1 2 3 4 5 6] [1 2] map[a:b c:d e:d] 0xc0000582f0 {task_you_must_do}}
------------------
OnMain -->
&{xingluo 18 0 address_value true [1 2 3 4 5 6] [1 2] map[a:b c:d e:d] 0xc0000582f0 {task_you_must_do}}
&{xingluo_will_do}
{task_you_must_do}
```

5.自定义添加过滤器

```go
func main{
    // new func filter
	tagFunc := application.NewFuncInitByTag(
		func(srcVal *reflect.Value, kind *reflect.Kind, srcField *reflect.Value, structField *reflect.StructField) error {
			// TODO add your filter func here, you can get tag to do something
			return nil
		})
	// add filter
	application.AddBeanFilterFunc(core.BeanInjected, &core.BeanFilterFunction{
        Mode:   core.AllBeanMode, // all mode (contains single and multi)
		Filter: tagFunc,
	})
	// main func 主程序入口
	application.Run( //Sequential execution
		&config.BeanManager{},
	)
}
```

### 2.简易实例

1.新建一个管理bean的结构体 和 两个结构体对象

```go
type BeanManager struct {
}

type Task struct {
	Name string `default:"task_you_must_do"` // default name
}

type Human struct {
	Name    string            `default:"name_value"`
	Task    *Task             `bean:"task1"` // if it is single bean, it must be domain.Task
	Task2   Task              `bean:"task2"` // struct type not ptr, it will copy
}
```

2.实现声明周期 core.BeanLifeCycle 的两个方法

```go
type BeanManager struct {
}
// add beans
func (_this *BeanManager) Create() *core.BeanConfig {
	return &core.BeanConfig{ 
		SingleBeans: []interface{}{
			&domain.Human{ 
				Name: "xingluo",
				Age:  18,
			},
		},
		MultiBeans: map[string]interface{}{
			"task1": &domain.Task{  
				Name: "xingluo_will_do", // override default name
			},
			"task2": &domain.Task{},
		},
	}
}

func (_this *BeanManager) OnMain(factory *core.BeanFactory) error {
	bean, err := factory.GetBean("domain.Human") // pkg + struct, it is type
	if err != nil {
		return err
	}
	human := bean.(*domain.Human) // Must be converted to ptr type
	fmt.Println(human)
	fmt.Println(human.Task)
    fmt.Println(human.Task2)
	return nil
}
```

3.注册到执行链中并执行

```go
func main() {
    // main func 主程序入口
	application.Run( //Sequential execution
		&config.BeanManager{},
	)
}
```

4.执行结果

```bash
&{xingluo 18 0 address_value true [1 2 3 4 5 6] [1 2] map[a:b c:d e:d] 0xc0000582f0 {task_you_must_do}}
&{xingluo_will_do}
{task_you_must_do}
```

## 4.结语

欢迎大家提意见或者建议，如有bug欢迎反馈。

