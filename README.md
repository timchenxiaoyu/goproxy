## http or https proxy

this is proxy for http or https with golang 
### build
```$xslt
go build
```


### run
```$xslt
./goproxy --log_dir /tmp -v 10

```

### example
```$xslt
curl --proxy http://127.0.0.1:8080 https://www.baidu.com
```