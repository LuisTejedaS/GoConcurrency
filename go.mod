module example/concurrency

go 1.17

replace example/channels => ./channels

require example/channels v0.0.0-00010101000000-000000000000

require (
	example/routines v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.3.0 // indirect
)

replace example/routines => ./routines
