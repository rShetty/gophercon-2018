# gophercon-2018
## Code companion for my talk at Gophercon India 2018, slides [here](http://slides.com/rajeevbharshetty/resiliency-in-distributed-systems)

## Usage:
Run `dep ensure` for pulling in dependencies

### Timeout
To start a slow server
Run 
```
go run timeout/leak.go

```

To start a non-leaking server
Run 
```
go run timeout/timeout.go

```

### Retry

To start a retry server
Run 
```
go run retry/retry.go

```

### Circuit Breaker

To start a circuit breaker server
Run 
```
go run circuitbreaker/breaker.go

```
