package core

import (
	"fmt"
	"reflect"
)

type RpcReq struct {
	ServiceName string
	MethodName  string
	PayLoad     string
}

type AresMethod struct {
	MethodName string
	CallFun    reflect.Method
	ParamsType reflect.Type
	Param1     interface{}
}

func (aresMethod *AresMethod) Invoke(param interface{}) []reflect.Value {
	reValue := reflect.ValueOf(param)
	parmams := []reflect.Value{reflect.ValueOf(aresMethod.Param1), reValue}
	return aresMethod.CallFun.Func.Call(parmams)
}

type ServiceMethods struct {
	ServiceName    string
	AresMethodsMap map[string]*AresMethod
}

type Core struct {
	ServiceMthodsMap map[string]*ServiceMethods
}

func (core *Core) Init() {
	core.ServiceMthodsMap = map[string]*ServiceMethods{}
}

func (core *Core) RegisterController(contoller interface{}) {
	rtype := reflect.TypeOf(contoller)
	kd := rtype.Kind()
	fmt.Println("kind = " + kd.String())
	methods := rtype.NumMethod()
	fmt.Println("methods count = ", methods)
	for i := 0; i < methods; i++ {
		method := rtype.Method(i)
		var serviceName = rtype.Elem().Name()
		var methodName = method.Name
		fmt.Println("---------- method_name = " + methodName + " service_name = " + serviceName)
		var val *ServiceMethods
		var ok bool
		if val, ok = core.ServiceMthodsMap[serviceName]; !ok {
			val = &ServiceMethods{}
			core.ServiceMthodsMap[serviceName] = val
		}
		aresMethod := AresMethod{}
		aresMethod.CallFun = method
		aresMethod.Param1 = contoller
		aresMethod.MethodName = methodName
		//method params
		mt := method.Type
		numIn := mt.NumIn()
		for j := 0; j < numIn; j++ {
			param := mt.In(j)
			fmt.Println("++++++pramtype = " + param.Elem().Name())
			aresMethod.ParamsType = param.Elem()
		}
		val.addMethod(&aresMethod)
	}
}

func (srviceMethods *ServiceMethods) addMethod(aresMethod *AresMethod) {
	if srviceMethods.AresMethodsMap == nil {
		srviceMethods.AresMethodsMap = map[string]*AresMethod{}
	}
	srviceMethods.AresMethodsMap[aresMethod.MethodName] = aresMethod

}

func (core *Core) GetCallFun(serviceName string, methodName string) *AresMethod {
	var val *ServiceMethods
	var ok bool
	if val, ok = core.ServiceMthodsMap[serviceName]; !ok {
		return nil
	}
	if val.AresMethodsMap == nil {
		return nil
	}
	return val.AresMethodsMap[methodName]
}
