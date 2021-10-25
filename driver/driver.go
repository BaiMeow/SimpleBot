package driver

type Driver interface {
	Run() error
	Write([]byte) error
	Read() ([]byte, error)
	Stop() error
}
