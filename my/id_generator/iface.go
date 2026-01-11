package id_generator

type IdGenIface interface {
	GenSnowFlakeID() int64
}
