module github.com/Allegan/wit

go 1.15

replace github.com/Allegan/wit/pkg/config => ./pkg/config

replace github.com/Allegan/wit/pkg/mediawiki => ./pkg/mediawiki

require (
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20201002202402-0a1ea396d57c // indirect
	golang.org/x/tools v0.0.0-20201002184944-ecd9fd270d5d // indirect
)
