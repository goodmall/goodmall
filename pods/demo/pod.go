package demo

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
