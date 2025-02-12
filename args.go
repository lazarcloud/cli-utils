package cli_utils

type RuntimeArgs map[string]interface{}

func (r RuntimeArgs) Get(name string) string {
	return r[name].(string)
}

type Arg struct {
	Name        string
	Description string
	Required    bool
	Default     interface{}
	Check       func(arg *Arg, value string) error
}

type Args []Arg

func (args Args) has(name string) bool {
	for _, a := range args {
		if a.Name == name {
			return true
		}
	}
	return false
}

func (args Args) isMissing(name string) bool {
	return !args.has(name)
}
