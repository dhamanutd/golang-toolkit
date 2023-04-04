package healthcheck

import (
	toolkit "github.com/dhamanutd/golang-toolkit"
)

var MySqlCheck = func(rt *toolkit.Runtime) error {
	return rt.DB().Raw("select version()").Error
}
