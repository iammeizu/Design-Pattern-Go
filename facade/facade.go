package facade

import "fmt"

// interface
type API interface {
	Test() string
}

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

// struct
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

type aModuleImpl struct {}

type bModuleImpl struct {}

// New
func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// func
func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
	}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
