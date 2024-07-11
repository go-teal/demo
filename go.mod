module github.com/go-teal/demo

go 1.21.7

// require github.com/go-teal/teal v0.1.2

replace github.com/go-teal/teal => ./../teal

require (
	github.com/go-teal/teal v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.33.0
)

require (
	github.com/go-teal/gota v0.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	gonum.org/v1/gonum v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
