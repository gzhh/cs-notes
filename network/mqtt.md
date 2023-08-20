# MQTT protocol
Application layer protocol
- https://en.wikipedia.org/wiki/MQTT
- https://mqtt.org/

The MQTT protocol defines two types of network entities: a message broker and a number of clients.
- An MQTT broker is a server that receives all messages from the clients and then routes the messages to the appropriate destination clients.
- An MQTT client is any device (from a micro controller up to a fully-fledged server) that runs an MQTT library and connects to an MQTT broker over a network.

publish/subscribe
- broker
  - Information is organized in a hierarchy of topics.
- client
  - publisher
  - subscriber


## MQTT broker
### EMQX
- https://github.com/emqx/emqx

example
- https://www.emqx.io/docs/en/v5.1/deploy/install-docker.html#use-docker-compose-to-build-an-emqx-cluster
- https://iotassistant.io/home-assistant/install-mqtt-websockets-on-nginx/
- https://blog.51cto.com/u_13045706/3833784

### Mosquitto
- https://github.com/eclipse/mosquitto

example
- https://github.com/eclipse/paho.mqtt.golang/tree/master/cmd/docker
- https://hub.docker.com/_/eclipse-mosquitto
- https://www.donskytech.com/how-to-enable-websockets-in-mosquitto/


## MQTT client
Client toolbox
- https://mqttx.app/

Client lib
- https://github.com/eclipse/paho.mqtt.python
- https://github.com/eclipse/paho.mqtt.golang


## Ref
- https://en.wikipedia.org/wiki/Comparison_of_MQTT_implementations
