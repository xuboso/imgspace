## Golang实现的图床程序

### TODO
- 使用协程提高并发处理能力
- 做到热部署
- 限制每个IP或每天最多上传多少张，防止被刷

编译命令

for linux
`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o img_app`

for mac
`GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o img_app`