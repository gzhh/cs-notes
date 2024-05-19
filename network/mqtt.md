# MQTT protocol
Application layer protocol
- https://en.wikipedia.org/wiki/MQTT
- https://zh.wikipedia.org/wiki/MQTT
- https://mqtt.org

Docs
- https://mqtt.org/getting-started/
- What is MQTT? https://aws.amazon.com/what-is/mqtt/
- What Is the MQTT Protocol: A Beginner's Guide https://www.emqx.com/en/blog/the-easiest-guide-to-getting-started-with-mqtt
- MQTT Essentials https://www.hivemq.com/mqtt/
- MQTT协议中文版 https://github.com/mcxiaoke/mqtt

The MQTT protocol defines two types of network entities: a message broker and a number of clients.
- An MQTT broker is a server that receives all messages from the clients and then routes the messages to the appropriate destination clients.
- An MQTT client is any device (from a micro controller up to a fully-fledged server) that runs an MQTT library and connects to an MQTT broker over a network.

publish/subscribe
- broker
  - Information is organized in a hierarchy of topics.
- client
  - publisher
  - subscriber


## Best Pracetice
- 用 redis 直接替换 app/物联网的 mq 服务器的可行性大吗？怎么对客户端进行限制 https://v2ex.com/t/965630


## MQTT broker
- https://en.wikipedia.org/wiki/Comparison_of_MQTT_implementations

### EMQX
- https://github.com/emqx/emqx
- https://www.emqx.io/zh
- https://www.emqx.com/zh

Docs
- https://www.emqx.io/docs/en/v5.1/
- https://www.emqx.io/docs/en/v5.1/tutorial/mqtt-programming.html

Install & Config
- https://www.emqx.io/docs/zh/v3.0/config.html
- https://www.emqx.io/docs/en/v5.1/deploy/install-docker.html
- https://iotassistant.io/home-assistant/install-mqtt-websockets-on-nginx/
- EMQ配置通过nginx反向代理wss和ws https://blog.51cto.com/u_13045706/3833784

client sdk
- https://www.emqx.com/en/mqtt-client-sdk
- https://www.emqx.com/en/blog/category/mqtt-programming

client
- https://mqttx.app/
- http://www.emqx.io/online-mqtt-client
- https://www.emqx.io/docs/zh/latest/admin/cli.html
- https://www.emqx.com/zh/blog/mqtt-client-tools

usage
- http://git.malu.me/MQTT/

example
- https://www.emqx.io/docs/en/v5.1/deploy/install-docker.html#use-docker-compose-to-build-an-emqx-cluster
- https://iotassistant.io/home-assistant/install-mqtt-websockets-on-nginx/
- https://blog.51cto.com/u_13045706/3833784

性能
- EMQX 单节点消息吞吐性能测试报告 https://www.emqx.com/zh/resources/emqx-message-throughput-performance-test-report

### Mosquitto
- https://github.com/eclipse/mosquitto
- https://www.mosquitto.org

Docs
- https://www.mosquitto.org/documentation/

example
- https://github.com/eclipse/paho.mqtt.golang/tree/master/cmd/docker
- https://github.com/vvatelot/mosquitto-docker-compose/blob/main/docker-compose.yaml
- https://hub.docker.com/_/eclipse-mosquitto
- https://www.donskytech.com/how-to-enable-websockets-in-mosquitto/


## MQTT client
Client toolbox
- https://mqttx.app/

Client lib
- https://github.com/eclipse/paho.mqtt.golang
- https://github.com/eclipse/paho.mqtt.python
- https://github.com/eclipse/paho.mqtt.javascript

