# optparams

用于设置关键字参数的管理模块.

思路参考自grpc项目.得力于1.18增加了泛型支持,现在可以大量减少重复代码了.

## 使用例子

```golang
package main

import (
    "fmt"

    "github.com/Golang-Tools/optparams"
)

type NameOpt struct { //<- 1.定义一个关键字参数的结构体
    Name string
}

func WithName(name string) optparams.Option[NameOpt] {//<- 2.定义可用的关键字参数项,一般命名上使用`with`开头
    return optparams.NewFuncOption(
        func(o *NameOpt) {
            o.Name = name
        })
}

func expfunc(opts ...optparams.Option[NameOpt]) {//<- 3.像这样定义函数的关键字参数
    opt := optparams.GetOption(&NameOpt{Name: "a"}, opts...) //<- 4.使用函数GetOption解析关键字参数
    fmt.Println("opt:", opt)
}

func main() {
    expfunc(WithName("b"))
}

```

感慨下泛型真方便
