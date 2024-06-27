package base

type Name string

const (
	SimpleMutexImplCAS     Name = "simple_mutex_impl_cas"
	SimpleMutexImplSema    Name = "simple_mutex_impl_sema"
	SimpleMutexImplSemaCAS Name = "simple_mutex_impl_sema_cas"
	SimpleMutexImplV0      Name = "simple_mutex_impl_v0"
	SimpleMutexImplV1      Name = "simple_mutex_impl_v1"
)
