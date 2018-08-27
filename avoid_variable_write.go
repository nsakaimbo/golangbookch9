package main

import "image"

// START1 OMIT

var icons = make(map[string]image.Image)

func loadIcon(name string) image.Image

// Not concurrency-safe!
func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}

// END1 OMIT
