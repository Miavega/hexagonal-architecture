package application

type useCases struct {
}

type UseCases interface {
	GetRoute()
	PostWriterRoute()
	PostWorkerRoute()
}

func NewUseCases() UseCases {
	return &useCases{}
}
