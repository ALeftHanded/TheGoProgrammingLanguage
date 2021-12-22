package main

import "github.com/pkg/profile"

func MemExample() {
	var stack [1024]byte
	heap := make([]byte, 64*1024)
	_, _ = stack, heap
}

func main() {
	//p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	//p := profile.Start(profile.MemProfileAllocs, profile.ProfilePath("."), profile.NoShutdownHook)
	defer profile.Start(profile.MemProfile).Stop()
	MemExample()
}
