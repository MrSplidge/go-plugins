module go-plugins

go 1.22.0

require (
	github.com/MrSplidge/go-xmldom v1.1.4
	github.com/MrSplidge/go-coutil v0.0.0
	github.com/antchfx/xpath v1.2.5 // indirect
	github.com/mattn/go-isatty v0.0.20
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225
)

replace github.com/MrSplidge/go-coutil v0.0.0 => ..\go-coutil

require golang.org/x/sys v0.17.0 // indirect
