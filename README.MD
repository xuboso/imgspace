## Golang实现的图床程序

### TODO
- 使用协程提高并发处理能力
- 做到热部署
- 限制每个IP或每天最多上传多少张，防止被刷

编译命令

for Linux
`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o img_app`

for Mac
`GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o macOS_img_app`

for windows 64位
`GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o windows64_img_app.exe`

for windows 32位
`GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o windows32_img_app.exe`


1. git clone https://github.com/xuboso/imgspace.git

2. 安装caddy, 执行命令 `CADDY_TELEMETRY=on curl https://getcaddy.com | bash -s personal http.cache,http.expires,http.geoip`

3. 修改Caddyfile里面的host改成自己的

4. 启动caddy服务

5. 启动go程序 `./img_app`