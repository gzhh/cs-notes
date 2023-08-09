# HTTP 总结
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
