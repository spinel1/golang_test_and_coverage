# golang_test_and_coverage
Go 언어 테스트 소스 작성 및 커버리지 측정 by Ginkgo + Gomega


```
$ ginkgo
Running Suite: Calculator Suite
===============================
Random Seed: 1619763088
Will run 7 of 7 specs

•••••••
Ran 7 of 7 Specs in 0.000 seconds
SUCCESS! -- 7 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 525.815135ms
Test Suite Passed

$ ginkgo -cover
Running Suite: Calculator Suite
===============================
Random Seed: 1619763093
Will run 7 of 7 specs

•••••••
Ran 7 of 7 Specs in 0.000 seconds
SUCCESS! -- 7 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
coverage: 100.0% of statements

Ginkgo ran 1 suite in 540.784032ms
Test Suite Passed

```

### Coverprofile format

line.column,line.column numberOfStatements count

```
mode: atomic
calc/calculator/calculator.go:3.25,5.2 1 1
calc/calculator/calculator.go:7.26,9.2 1 2
calc/calculator/calculator.go:11.27,13.2 1 2
calc/calculator/calculator.go:15.29,17.2 1 1
calc/calculator/calculator.go:19.24,22.34 2 1
calc/calculator/calculator.go:25.2,25.12 1 1
calc/calculator/calculator.go:22.34,24.3 1 20

```

### Ginkgo 사용

```
### generate calculator_suite_test.go 
$ ginkgo bootstrap

### generate calculator_test.go 
$ ginkgo generate calculator

```
