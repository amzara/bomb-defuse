# Bomb Defuse

Small go project I randomly started on a Sunday evening to practice channels, worker pool and channel-based rate limiting. Idea is to have a c4 (I miss counter strike) that takes in commands that needs to be defused. Worker pool spams the goroutine for the right pin and rate-limiting ensures that only 1 command can come in per 100ms. 

TODO: implement bucket based rate limiting for practice purposes.

## Commands

Commands are sent as strings over a channel.

| Command            | Effect                                      |
|--------------------|---------------------------------------------|
| `PLANT`            | Arms the bomb, starts timer and ticker      |
| `STATUS`           | Prints current state                        |
| `PIN`              | Prints the current PIN (cheat/debug)        |
| `BEEP`             | Toggles the beeper on/off                   |
| `DEFUSE d d d d`   | Attempts defusal with a 4-digit PIN         |
| `TIMEOUT`          | Internal signal; triggers explosion         |


## File layout

```
.
├── main.go                      # Entry point, worker pool
└── internal/
    └── services/
        └── bomb/
            └── bomb.go          # Bomb struct, state machine, rate limiter
```

## Running

```bash
go run main.go
```

