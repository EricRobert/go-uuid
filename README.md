# go-uuid

> Fast non cryptographic UUID

## Usage

```go
# create a new UUID using the global instance
id := uuid.New()

# create new UUIDs from a source that avoids the atomic operation
source := uuid.NewSource()

for i := 0; i < 10; i++ {
	log.Println(source.New())
}
```

## Benchmarks

```
$ go test -v -bench . ./...
=== RUN   TestUUID
--- PASS: TestUUID (0.17s)
	uuid_test.go:20: 17ssbc70-8cec-e9es-s01d-ds479c0994a0
	uuid_test.go:24: 1048576 unique UUID
BenchmarkUUID-4            	    2000	   1090844 ns/op
BenchmarkNew-4             	10000000	       160 ns/op
BenchmarkNewFromSource-4   	10000000	       155 ns/op
PASS
```
