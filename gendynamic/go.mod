module github.com/gcottom/refract/gendynamic

go 1.22.3

replace (
	github.com/gcottom/refract/safereflect => ../safereflect
	github.com/gcottom/refract/refractutils => ../refractutils
)

require github.com/gcottom/refract/safereflect v0.0.0-20241013014421-507e0b369ab8

require github.com/gcottom/refract/refractutils v0.0.0-20241013020713-c30bda7cedd9 // indirect
