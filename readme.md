# Monolith solution for task

## Description

you should build a server that sends n number of random messages to a telegram channel/group using telegram bot token with a rate of 1 message per second. Also, a bunch of messages can have one of these priorities Low, Medium, and High. Highest priority message should go first no matter how busy the server is. You are required to build an API gateway and at least a microservices and connecting those services using rabbitmq, activemq or kafka pup/sub, where API gateway receive post request (n, priority) and generates n random messages and publishes to a message broker where on the other side a service consume it with a constant rate (1msg/s)and sends to a telegram channel/group. It would be better if you add ack/nack too

##

- http handler
- bundle generator
- sender
