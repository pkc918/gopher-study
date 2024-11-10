### 记录 esbuild 学习过程

#### 拉取项目到本地后，如何初始化项目，为pr做准备
打开项目的 `docs/development.md` 文件
1. *Build*： 使用 `make esbuild` or `go build ./cmd/esbuild`
2. *Test*：使用 `make test-go` or `go test ./internal/...`
3. *Publish*:
    - 修改 version.txt 版本号
    - 将版本号复制到 `CHANGELOG.md`
    - 执行 `make publish-all`

执行在浏览器上
1. 执行 `make platform-wasm` 构建 esbuild 的 WebAssembly 版本
2. 使用 HTTP 服务开启文件 `./esbuild --servedir=.`
3. 点击 `/scripts/try.html` 在你的浏览器上