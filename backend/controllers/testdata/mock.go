package testdata

type serviceMock struct{}

func NewServiceMock()*serviceMock{
	return &serviceMock{}
}