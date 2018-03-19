package demo

//  可被各个 端 继承 此处作为共享基类 BaseDemoPod ?  然后其他 entry-point处继承（内嵌）此类即可
type DemoPod struct {
	Config struct {
		DSN string //`mapstructure:"dsn"`,
	}
}

func NewDemoPod(options ...func(*DemoPod)) *DemoPod {
	pod := DemoPod{}

	for _, opt := range options {
		opt(&pod)
	}

	return &pod
}
