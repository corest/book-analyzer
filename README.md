# bookanalyzer

## How to run

1. With `go run`

```
cat book_analyzer_set.in | go run main.go -target-size 200
```

2. With binary

```
go build
./bookanalyzer -target-size 200 < book_analyzer_set.in
