
# Terratest

- References
    - https://terratest.gruntwork.io/examples/
    - https://github.com/gruntwork-io/terratest/blob/master/test/docker_hello_world_example_test.go

## Docker

Terratest로 간단한 Hello World를 테스트하는 테스트 코드가 있습니다.

- 성공 사례

```shell
go test success_test.go
```

- 실패 사례

```shell
go test failure_test.go
```