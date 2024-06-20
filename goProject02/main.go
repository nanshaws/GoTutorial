package main

import "nanshawsCache/cache"

func main() {
	cache := cache.NewMenCach()
	cache.SetMaxMemory("300B")
}
