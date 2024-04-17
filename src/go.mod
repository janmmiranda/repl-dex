module github.com/janmmiranda/repl-dex

go 1.22.2

replace github.com/janmmiranda/repl-dex/internal/pokeapi => ./internal/pokeapi

require github.com/janmmiranda/repl-dex/internal/pokeapi v0.0.0-00010101000000-000000000000

replace github.com/janmmiranda/repl-dex/internal/pokecache => ./internal/pokecache

require github.com/janmmiranda/repl-dex/internal/pokecache v0.0.0 // indirect
