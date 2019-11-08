package factory_method

// interface
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

// class OperatorBase
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// class PlusOperatorFactory
type PlusOperatorFactory struct {}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// class PlusOperator
type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

// class MinusOperatorFactory
type MinusOperatorFactory struct {}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
			&OperatorBase{},
	}
}

// class MinusOperator
type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}
