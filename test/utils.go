package test

import "flag"

func IsTesting() bool {
	return flag.Lookup("test.v") != nil
}
