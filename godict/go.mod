module github.com/gcottom/refract/godict

go 1.22.3

replace github.com/gcottom/refract/safereflect => ../safereflect

replace github.com/gcottom/refract/refractutils => ../refractutils

require (
	github.com/gcottom/refract/refractutils v0.0.0-20241013020713-c30bda7cedd9
	github.com/gcottom/refract/safereflect v0.0.0-20241013021005-260a7adec38c
)
