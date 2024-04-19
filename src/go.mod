module github.com/janmmiranda/repl-dex

go 1.22.2

replace github.com/janmmiranda/repl-dex/internal/pokeapi => ./internal/pokeapi

require (
	github.com/google/uuid v1.6.0
	github.com/janmmiranda/repl-dex/internal/pokeapi v0.0.0-00010101000000-000000000000
	github.com/peterh/liner v1.2.2
)

replace github.com/janmmiranda/repl-dex/internal/pokecache => ./internal/pokecache

require (
	github.com/janmmiranda/repl-dex/internal/pokecache v0.0.0 // indirect
	github.com/mattn/go-runewidth v0.0.3 // indirect
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
)
