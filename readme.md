## Start
`chmod +x obsgo_amd64`

`./obsgo_amd64 -h`
## Put
```bash
AccessKeyID=ASDFGH12345 SecretAccessKey=Nhbn7sDKJBfsfjsnbsd77Dfsjn endPoint="https://obs.ap-xxx.myhuaweicloud.com" ./obsgo_amd64 -op=put -bucket=name -key=path/to/name.txt -file=/local/test.txt
```
## delete
```bash
AccessKeyID=ASDFGH12345 SecretAccessKey=Nhbn7sDKJBfsfjsnbsd77Dfsjn endPoint="https://obs.ap-xxx.myhuaweicloud.com" ./obsgo_amd64 -op=del -bucket=name -key=path/to/name.txt
```

## development
```bash
# TELEGRAM_BOT_TOKEN=1234:ABCD \
# TELEGRAM_CHAT_ID=111222333 \
AccessKeyID=ASDFGH12345 \
SecretAccessKey=Nhbn7sDKJBfsfjsnbsd77Dfsjn \
endPoint="https://obs.ap-xxx.myhuaweicloud.com" \
go run *.go -op=put -bucket=name -key=path/to/name.txt -file=/local/test.txt
```