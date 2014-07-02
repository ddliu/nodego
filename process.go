package nodego

import(
    "github.com/robertkrimen/otto"
    "os"
    "strings"
)

// Create a new process with below properties:
//  - env
//  - nextTick x
//  - argv
//  - _eval
//  - execArgv
//  - _forceRepl
//  - exit x
//  - stdin x
//  - binding
//  - emit
//  - _exiting x
//  - tickCallback x
//  - cwd
//  - _print_eval
//  - on x
//  - reallyExit
//  - send
//  - _rawDebug
//  - platform
//  - moduleLoadList
func NewProcess(vm *otto.Otto) *otto.Object {
    process, _ := vm.Object(`({})`)
    process.Set("env", getEnv())
    process.Set("argv", os.Args)
    process.Set("binding", binding)

    return process
}

func getEnv() map[string]string {
    var env = make(map[string]string)
    for _, v := range os.Environ() {
        parts := strings.SplitN(v, ":", 2)
        var key, value string
        if len(parts) > 0 {
            key = parts[0]
        }
        if len(parts) > 1 {
            value = parts[1]
        }
        env[key] = value
    }

    return env
}

func binding(call otto.FunctionCall) otto.Value {
    name := call.Argument(0).String()
    loader, ok := modules[name]
    if !ok {
        ThrowError(call.Otto, "Error", "Binding not found: " + name)
    }

    return loader(call.Otto)
}