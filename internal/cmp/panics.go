package cmp

// Panics returns true if the function passed to it panics.
// Otherwise, it returns false.
func Panics(f func()) (bool, interface{}) {
	panics := false
	var err interface{}
	func() {
		defer func() {
			if err = recover(); err != nil {
				panics = true
			}
		}()
		f()
	}()
	return panics, err
}
