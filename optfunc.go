package optfunc

type OptionsParams map[string]interface{}
type OptionsFunc func(o *Options)

type Options struct {
	optionalParams OptionsParams
}

func (o *Options) Get(name string) interface{} {
	v, ok := o.optionalParams[name]
	if !ok {
		return nil
	} else {
		return v
	}
}

func (o *Options) Apply(opts ...OptionsFunc) {
	for _, f := range opts {
		f(o)
	}
}

func With(name string, value interface{}) OptionsFunc {
	return func(o *Options) {
		o.optionalParams[name] = value
	}
}

func NewOptions(op OptionsParams) *Options {
	return &Options{optionalParams: op}
}
