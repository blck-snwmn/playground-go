## Test(unit test)
```bash
$ go test -v .
=== RUN   TestEnv
    sample_env_test.go:10: Skipping no CLASS environment variable set
--- SKIP: TestEnv (0.00s)
=== RUN   TestShort
--- PASS: TestShort (0.00s)
=== RUN   TestUnit
--- PASS: TestUnit (0.00s)
PASS
ok      github.com/blck-snwmn/playground-go/testclassification  0.002s
```

## Test(unit test) w/ CLASS=true
```bash
$ CLASS=true go test -v .
=== RUN   TestEnv
--- PASS: TestEnv (0.00s)
=== RUN   TestShort
--- PASS: TestShort (0.00s)
=== RUN   TestUnit
--- PASS: TestUnit (0.00s)
PASS
ok      github.com/blck-snwmn/playground-go/testclassification  0.003s
```

## Test(unit test) w/ CLASS=true and short
```bash
$ CLASS=true go test -v --short .
=== RUN   TestEnv
--- PASS: TestEnv (0.00s)
=== RUN   TestShort
    sample_short_test.go:7: Skipping short test
--- SKIP: TestShort (0.00s)
=== RUN   TestUnit
--- PASS: TestUnit (0.00s)
PASS
ok      github.com/blck-snwmn/playground-go/testclassification  0.002s
```

## Test(integration test)
```bash
$ go test -tags=integ -v .
=== RUN   TestEnv
    sample_env_test.go:10: Skipping no CLASS environment variable set
--- SKIP: TestEnv (0.00s)
=== RUN   TestInteg
--- PASS: TestInteg (0.00s)
=== RUN   TestShort
--- PASS: TestShort (0.00s)
=== RUN   TestUnit
--- PASS: TestUnit (0.00s)
PASS
ok      github.com/blck-snwmn/playground-go/testclassification  0.003s
```