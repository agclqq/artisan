package prowjob

type Commander interface {
	GetCommand() string
	Usage() string
	Handle(*Context)
}
