package core

// BeanLifeCycleExecute execute the beanLifeCycle
func BeanLifeCycleExecute(lifeCycles []interface{}, f func(lifeCycle any) error) (err error) {
	if len(lifeCycles) == 0 {
		return
	}
	for _, lifeCycle := range lifeCycles {
		if lifeCycle == nil {
			continue
		}
		err = f(lifeCycle)
		if err != nil {
			return
		}
	}
	return
}

// CreateHandlerExecute execute the createHandler
func CreateHandlerExecute(lifeCycle interface{}) (err error) {
	if createHandler, ok := lifeCycle.(CreateHandler); ok {
		beanConfig := createHandler.Create()
		err = AddBeanConfig(beanConfig)
		if err != nil {
			return
		}
	}
	return
}

// AfterInstantiationHandlerExecute execute the afterInstantiationHandler
func AfterInstantiationHandlerExecute(lifeCycle interface{}) (err error) {
	if afterInstantiationHandler, ok := lifeCycle.(AfterInstantiationHandler); ok {
		err = afterInstantiationHandler.AfterInstantiation(Context)
		if err != nil {
			return
		}
	}
	return
}

// AfterInitializationHandlerExecute execute the afterInitializationHandler
func AfterInitializationHandlerExecute(lifeCycle interface{}) (err error) {
	if afterInitializationHandler, ok := lifeCycle.(AfterInitializationHandler); ok {
		err = afterInitializationHandler.AfterInitialization(Context)
		if err != nil {
			return
		}
	}
	return
}

// ChooseMainHandlerExecute choose the main handler
func ChooseMainHandlerExecute(lifeCycle interface{}) (err error) {
	if onMainHandler, ok := lifeCycle.(OnMainHandler); ok {
		return onMainAndDestroyHandlerExecute(onMainHandler)
	} else if mainHandler, ok2 := lifeCycle.(MainHandler); ok2 {
		return mainHandlerExecute(mainHandler)
	}
	return
}

// MainHandlerExecute execute the mainHandler
func mainHandlerExecute(mainHandler MainHandler) (err error) {
	return mainHandler.OnMain(Context)
}

// OnMainAndDestroyHandlerExecute execute the OnMainHandler and DestroyHandler
func onMainAndDestroyHandlerExecute(onMainHandler OnMainHandler) (err error) {
	useGo, destroy, err := onMainHandler.OnMain(Context)
	if err != nil {
		return err
	} else if useGo != nil {
		go useGo()
	}
	if destroy {
		if destroyHandler, ok := onMainHandler.(DestroyHandler); ok {
			err = destroyHandler.Destroy(Context)
			if err != nil {
				return
			}
		}
	}
	return
}

// AddBeanConfig adds beans
func AddBeanConfig(beanConfig *BeanConfig) (err error) {
	if beanConfig == nil {
		return
	}
	for beanName, bean := range beanConfig.MultiBeans {
		_, err = Context.AddMultiBean(beanName, bean)
		if err != nil {
			return
		}
	}
	for _, bean := range beanConfig.SingleBeans {
		_, err = Context.AddSingleBean(bean)
		if err != nil {
			return
		}
	}
	return
}
