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
