# go-spring Introduction

A simple container framework based on **Golang**, similar to Java's Spring container. It provides operations such as
container, label initialization, and bean injection. Due to the special nature of the go language, it cannot be proxied,
so it provides a lifecycle interface to provide services.

In addition to providing containers, it also provides various parsing functions, as well as some Java like base classes
and data structures, such as the Object interface (implementing equals), Iterator interface, LinkList, ArrayDeque, and
so on.



[English document 英文文档](./README.md)

[Chinese document 中文文档](./README.zh-CN.md)



## 1. Module Introduction

Open packages include **application, core, framework, ** etc. Of course, other packages can also be used.

1. You can create a new lifecycle structure and then implement the core.BeanLifeCycle or core.InBeanLifeCycle
   immediately obtains the method that needs to be implemented. Alternatively, implementing only some of the methods
   within it is also possible. (Including Create, AfterInstantiation, AfterInitialization, OnMain)

2. Then return the return value of the **Create** of the structure to SingleBeans or MultiBeans.

3. Then, the **AfterInstantiation** method can be implemented, which is executed after **tag initialization and before**
   bean injection.

4. Then, the **AfterInitialization** method can be implemented, which is executed after **bean injection**.

Then, the **OnMain** method can be implemented, which includes two parameter/return value methods that can be selected.
The main process code can be executed inside.

6. The above four methods are **optional**, which means you can implement none of them.



## 2.Precautions

### 1. Explanation of Terms

**Context**: The global singleton of **BeanFactory** type is the core of this project. You can reassign the objects you
implement, or you can only reassign the Container interface inside. Through **core.Context** or **core.AppContext()**
retrieves the object and performs a series of programmatic operations on the bean. You can also use the methods for
obtaining more operations for **core.Context.BeanContainer**.

**MultiBean**: Multiple instance beans can be added or obtained through the lifecycle, or through the core Add or obtain
Context. When adding, you need to set the BeanName yourself, and then automatically obtain the BeanType (its type),
which will be added to the container. If you add a SingleBean of the same type at this time, the addition will fail.

**SingleBean**: A singleton bean that can be added or obtained through its lifecycle, or through the core Add or obtain
Context. When a structural object is added, it only retains one singleton, and the BeanName of the object is the **
package name+struct name**. Adding the same type of MultiBean at this time will result in a failure to add. If created
in Create, priority is given to scanning MultiBeans.

So MultiBean and SingleBean are **mutually exclusive**. If only one single instance is created, SingleBean is
recommended. If there are multiple instances, only MultiBean can be used.

**Application.Run**: If the parameter is any, the methods of the object will be executed sequentially according to its
lifecycle. If there are multiple object parameters, one stage of all objects will be executed, followed by the next
stage of all objects.



### 2. Tag Explanation

**Default**: If the added bean does not have a value, the converted value of tag (string) will be added by default (if
the conversion fails, no value will be assigned).

**Value**: Unlike default, value will be forcibly assigned regardless of whether there is a value in the bean.

Both methods only support simple data. Including all other **basic types** except Complex, as well as two **first level
aggregation types** corresponding to simple basic types (except), namely **slice** and **map**, such as int, string,
uint64, map[int]int, []int, []string, etc. The following are **unsupported**, such as map[int] []int, [] Human, [] []
int, map[string]map[string]int , etc

**Cap**: A tag provided for slicing, which can set the capacity, but other settings will not take effect.

**Bean**: Fill in the name of the bean, which will be automatically added to this location through the lifecycle. Please
note **to prevent circular dependencies**. If you need to manually inject multiple instances of a bean without affecting
other beans of the same type, you can manually modify it after injection. 2. If the target field is of type ptr, it will
be directly assigned a value (modifying the object will affect the beans in the container). If it is of type struct, it
will be **copied** and assigned to the field



## 3.Quick Start



### 1. Ordinary examples

1.Create a new management bean structure and two structure objects

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

2.Implement the declaration cycle core.BeanLifeCycle method

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

3.Register in the execution chain and execute

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

4.Results of execution

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

5.Custom Add Filter

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

### 2. Simple examples

1.Create a new management bean structure and two structure objects

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

2.Implement the declaration cycle Two methods for core.BeanLifeCycle

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

3.Register in the execution chain and execute

```go
func main() {
    // main func 主程序入口
	application.Run( //Sequential execution
		&config.BeanManager{},
	)
}
```

4.Results of execution

```bash
&{xingluo 18 0 address_value true [1 2 3 4 5 6] [1 2] map[a:b c:d e:d] 0xc0000582f0 {task_you_must_do}}
&{xingluo_will_do}
{task_you_must_do}
```

## 4. Conclusion

Welcome to provide feedback or suggestions, and if there are any bugs, please feel free to provide feedback.
