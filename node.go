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

type ModuleInitialize func (vm *otto.Otto) otto.Value

// NODE_MODULE(module_name, Initialize)
func NodeModule(name string, init ModuleInitialize) {

}