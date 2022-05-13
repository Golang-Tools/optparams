//optparams 用于管理关键字参数的模块
package optparams

//Option 作为关键字参数的形参类型
//@generic T any 关键字参数对应的结构
type Option[T any] interface {
	Apply(*T)
}

//FuncOption 作为关键字参数项的返回值
//@generic T any 关键字参数对应的结构
type FuncOption[T any] struct {
	f func(*T)
}

func (fo *FuncOption[T]) Apply(do *T) {
	fo.f(do)
}

// NewFuncOption 用于构造关键字参数项目
//@generic T any 关键字参数对应的结构
//例子:
// func WithName(name string) optparams.Option[NameOpt] {
// 	return optparams.NewFuncOption(
// 		func(o *NameOpt) {
// 			o.Name = name
// 		})
// }
func NewFuncOption[T any](f func(*T)) *FuncOption[T] {
	return &FuncOption[T]{
		f: f,
	}
}

// GetOption 用于获取一个完整的关键字参数结构实例
func GetOption[T any](defaultopt *T, opts ...Option[T]) *T {
	if len(opts) > 0 {
		for _, opt := range opts {
			opt.Apply(defaultopt)
		}
	}
	return defaultopt
}
