package driver

type Driver interface {
	Run()
	Write([]byte)
	Read() []byte
	Stop()
}
