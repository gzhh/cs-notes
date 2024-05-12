# HTTP 总结
- https://pkg.go.dev/net/http


原理
- Go Internals: HTTP request multiplexing in Go https://akshay-kumar.hashnode.dev/go-internals-http-request-multiplexing-in-go-1
- 一文读懂 Go Http Server 原理 https://mp.weixin.qq.com/s/aLiqEuD9T8sERVrfGDSxvw
- Go 号称几行代码开启一个 HTTP Server，底层都做了什么？https://mp.weixin.qq.com/s/n7mSUB6pxoYmr5u575Nqqg
- 面试官：net/http库知道吗？能说说优缺点吗？https://mp.weixin.qq.com/s/AcmL60a-ZYdlFK2t0Et9Qg

Third Party Package
- https://github.com/valyala/fasthttp
  - https://github.com/tsingson/fasthttp-guide
  - Go Fasthttp 实践之 Hello World https://mp.weixin.qq.com/s/LRR_BIQVUx44OHReTJWv9Q
- Resty https://github.com/go-resty/resty


## 常见问题
### io.ReadAll 读取 Response.Body 出现 EOF 错误
原因
- 可能是服务端程序异常退出导致链接中断
- 可能是请求超时问题

Ref
- [https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi](https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi)
- [https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi](https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi)

### io.ReadAll 读取 Response.Body 出现 unexpected EOF 错误
问题复现
```
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getBodyFromURL(service string, clientTimeout int) (string, error) {
	var netClient = &http.Client{
		Timeout: time.Duration(clientTimeout) * time.Millisecond,
	}

	rsp, err := netClient.Get(service)
	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP request error, response code: %d", rsp.StatusCode)
	}


	// 造成 io.ReadAll 出现 unexpected EOF 错误，可能是读取内容的真实长度与 content-length 不一致
	buf, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	fmt.Printf("buf %s\n", string(buf))

	return string(bytes.TrimSpace(buf)), nil
}

func TestReadAll(t *testing.T) {
	bodyErrorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := `{"code": 0}`
		w.Header().Set("Content-Type", "application/json")

        // 错误长度
		w.Header().Set("Content-Length", "100")

        // 正确长度
		// w.Header().Set("Content-Length", strconv.Itoa(len(content)))

		w.Write([]byte(content))
	}))

	_, err := getBodyFromURL(bodyErrorServer.URL, 1000)
	if err != nil {
		panic(err)
	}
}
```

分析
- 可能是读取 response 的真实内容长度和 content-length 长度不一致导致的
- 也可能是请求返回的数据使用 Content-Encoding = gzip 压缩导致的，需要请求头设置 Accept-Encoding = gzip，然后手动解压 response

Ref
- https://github.com/wnanbei/direwolf/issues/1
