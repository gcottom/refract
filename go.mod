module github.com/gcottom/refract

go 1.22.0

replace (
	github.com/gcottom/refract/gendynamic => ./gendynamic
	github.com/gcottom/refract/godict => ./godict
	github.com/gcottom/refract/refractutils => ./refractutils
	github.com/gcottom/refract/safereflect => ./safereflect
)

require (
	github.com/gcottom/refract/gendynamic v0.0.0-20241013021005-260a7adec38c
	github.com/gcottom/refract/godict v0.0.0-20241013021005-260a7adec38c
	github.com/gcottom/refract/refractutils v0.0.1
)

require github.com/gcottom/refract/safereflect v0.0.1 // indirect
