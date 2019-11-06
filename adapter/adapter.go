package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type adapter struct {
	Adaptee
}

type adapteeImpl struct {}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}


func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
