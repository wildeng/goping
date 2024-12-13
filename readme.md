# Goping

Inspired by a simlar script that a friend of mine realised, I've implemented a colorised version of a simple ping to quickly identify slow responses:

1. Green if response time is < 50ms
2. Yellow if response time is < 150 ms
3. Red anything above the 150ms threshold
4. No colour if response time isn't available.

Take it as it is, nothing more than an excercise to keep my go knowledge ready.

### Usage

```
go run goping <your url>
```

or

```
go build goping
./goping <your url>
```
