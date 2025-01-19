package main



type User struct {
	noCopy noCopy
}
func main() {
	u1 := User{}
	_ = u1
	testFunc(u1)
}
func testFunc(u User)  {
	
}


// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}