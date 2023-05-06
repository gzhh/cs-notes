## curl

### HTTP request 传参方式

get

curl https://httpbin.org/anything?name=gzhh

post form

curl -X POST --data "access_token='abcdefghijklmnopqrstuvwxyz'" https://httpbin.org/anything

post data

curl -X POST -H "Content-Type:application/json" --data "access_token='abcdefghijklmnopqrstuvwxyz'" https://httpbin.org/anything

post json data

curl -X POST -H "Content-Type:application/json" --data '{"name":"gzhh","gender":"male"}' https://httpbin.org/anything
