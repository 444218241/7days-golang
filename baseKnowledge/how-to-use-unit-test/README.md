### Go Test 单元测试简明教程

#### 一个简单例子
- 测试用例示范
  - `TestAdd(t *testing.T)` 
  - `*testing.B` 基准测试(benchmark)的参数。
  - `*testing.M` TestMain的参数类型。
- 运行示范
  - `go test`
  - `go test -v` # -v 参数会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率。
  - `go test -run TestAdd -v` # 如果只想运行其中的一个用例，
   例如 TestAdd，可以用 -run 参数指定，该参数支持通配符 *，和部分正则表达式，例如 ^、$。

#### 子测试(Subtests)
