package optfunc

type OptionsParams map[string]interface{}
type Options func(o *OptionsStruct)

type OptionsStruct struct {
	optionalParams OptionsParams
}

func (o *OptionsStruct) Get(name string) interface{} {
	v, ok := o.optionalParams[name]
	if !ok {
		return nil
	} else {
		return v
	}
}

func (o *OptionsStruct) Apply(opts ...Options) *OptionsStruct {
	for _, f := range opts {
		f(o)
	}
	return o
}

func With(name string, value interface{}) Options {
	return func(o *OptionsStruct) {
		o.optionalParams[name] = value
	}
}

func NewOptions(op OptionsParams) *OptionsStruct {
	return &OptionsStruct{optionalParams: op}
}
