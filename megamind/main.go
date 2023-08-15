package main

import (
	"fmt"
	"io/ioutil"
	"megamind/routes"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	nbProgram, err := LoadNBConfig()
	if err != nil {
		fmt.Println("failed in loading nb program")
		panic(err)
	}

	routes.InitRoutes(*e, nbProgram)

	e.Logger.Fatal(e.Start(":9812"))
}

func LoadNBConfig() (goja.Callable, error) {
	vm := goja.New()

	new(require.Registry).Enable(vm)

	// script := `
	// 	function myFunction(param)
	// 	{
	// 		console.log("myFunction running ...")
	// 		console.log("Param = ", param)
	// 		return "Nice meeting you, Go"
	// 	}
	// `
	scriptBytes, err := ioutil.ReadFile(os.Getenv("NB_SCRIPT"))
	if err != nil {
		fmt.Println("File not found")
		panic(err)
	}
	prog, err := goja.Compile("", string(scriptBytes), true)
	if err != nil {
		fmt.Printf("Error compiling the script %v ", err)
		return nil, err
	}
	_, err = vm.RunProgram(prog)

	var myJSFunc goja.Callable
	err = vm.ExportTo(vm.Get("nbEvalScript"), &myJSFunc)
	if err != nil {
		fmt.Printf("Error exporting the function %v", err)
		return nil, err
	}

	return myJSFunc, nil
}
