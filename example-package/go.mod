module example.com/m

go 1.22.5

require (
	github.com/codedByHasan/golang-for-beginners/cryptit v0.0.0
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

// Replacing package with local path
replace github.com/codedByHasan/golang-for-beginners/cryptit => ../cryptit
