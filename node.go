package nodego

import (
    "github.com/robertkrimen/otto"
    "fmt"
)

type Environment struct {
    vm *otto.Otto
}

func (env *Environment) Run() {
    value, err := env.vm.Call(SOURCE, nil, NewProcess(env.vm))
    fmt.Println(value, err)
}

func NewEnvironment() *Environment {
    return &Environment{
        otto.New(),
    }
}

type ModuleLoader func (vm *otto.Otto) otto.Value

var modules = make(map[string]ModuleLoader)

// NODE_MODULE(module_name, Initialize)
func NodeModule(name string, loader ModuleLoader) {
    modules[name] = loader
}

func ThrowError(vm *otto.Otto, errorType, msg string) {
    value, _ := vm.Call("new " + errorType, nil, msg)
    panic(value)
}

func ThrowTypeError(vm *otto.Otto, msg string) {
    ThrowError(vm, "TypeError", msg)
}
