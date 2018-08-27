package main

import "image"

// START1 OMIT

// Concurrency-safe: loaded at package initialization
// before main function.

var icons = map[string]image.Image{
	"spades.png": loadIcon("spades.png"),
	"hearts.png": loadIcon("hearts.png"),
	// ...
}

func Icon(name string) image.Image { return icons[name] }

// END1 OMIT
