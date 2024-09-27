package negate

func Negate(myBoolenan *bool) bool {
	*myBoolenan = !*myBoolenan
	return *myBoolenan
}