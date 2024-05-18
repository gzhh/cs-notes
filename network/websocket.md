# WebSocket
- https://en.wikipedia.org/wiki/WebSocket
- https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API

WebSocket is distinct from HTTP. Both protocols are located at layer 7 in the OSI model and depend on TCP at layer 4. Although they are different, RFC 6455 states that WebSocket "is designed to work over HTTP ports 443 and 80 as well as to support HTTP proxies and intermediaries", thus making it compatible with HTTP. To achieve compatibility, the WebSocket handshake uses the HTTP Upgrade header to change from the HTTP protocol to the WebSocket protocol.

原理
- 看完让你彻底搞懂Websocket原理 https://blog.csdn.net/frank_good/article/details/50856585

Tool
- websocat - Command-line client for WebSockets https://github.com/vi/websocat


## WebSocket handshake
- https://en.wikipedia.org/wiki/WebSocket#Protocol_handshake
- https://en.wikipedia.org/wiki/HTTP/1.1_Upgrade_header
- https://www.nginx.com/blog/websocket-nginx/


## ws 和 wss
ws 和 wss 的区别
- https://stackoverflow.com/questions/46557485/difference-between-ws-and-wss

wss 下证书问题
- 使用自签证书需要手动跳过浏览器安全限制
- 可以使用 letsencrypt 等第三方证书，或付费机构证书


## Go Websocket 使用
- https://github.com/gorilla/websocket

Go websocket cors 问题
- https://stackoverflow.com/questions/22644392/chrome-websockets-cors-policy
- https://stackoverflow.com/questions/65034144/how-to-add-a-trusted-origin-to-gorilla-websockets-checkorigin
- https://stackoverflow.com/questions/33323337/update-send-the-origin-header-with-js-websockets
- https://github.com/kataras/neffos/issues/11#issuecomment-520689681


## HTTP2 vs Websocket
- https://stackoverflow.com/questions/28582935/does-http-2-make-websockets-obsolete
