# Getting Started With Fuzzing

This is a basic example of fuzzing in Go. The goal of fuzzing is to test corner cases by generating random test data and submitting it to code to see whether it properly handles unexpected input.
The code example contains a simple flowed function, its fixed version, unit and fuzz tests. 

---

### Prerequisites
* **Go** 1.18 or later 
* Command terminal
* CPU with AMD64 or ARM64 architecture

---

### How to run

```bash
$ go run .
```

### How run to the tests
Run unit test and fuzz tests with only seed corpus. Ensure all tests successfully passed. 
```bash
$ go test -v
```

Run only fuzz test with invalid inputs generated automatically. 
Ensure fuzz test failed and invalid input is present in *testdata/fuzz/FuzzReverse* directory.
```bash
$ go test -fuzz=Fuzz
```

Comment flawed function and uncomment fixed one in main.go file. Run fuzz test again and ensure test successfully passed. 
```bash
$ go test -fuzz=Fuzz -fuzztime 20s
```
