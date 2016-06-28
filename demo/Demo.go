package demo

type Runner interface {
	Run()
}

type RunnerFunc func()

func (f RunnerFunc) Run() {
	f()
	return
}

type Decorator func(Runner) Runner

func Decorate(r Runner, ds ...Decorator) Runner {
	decorated := r
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
