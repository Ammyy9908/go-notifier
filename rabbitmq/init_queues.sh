#!/bin/bash

# Wait for RabbitMQ to start
sleep 10

# Log in to the RabbitMQ container
rabbitmqadmin -u user -p password declare queue name=Email_High_Priority durable=true arguments='{"x-max-priority": 10}'
rabbitmqadmin -u user -p password declare queue name=Email_Standard durable=true
rabbitmqadmin -u user -p password declare queue name=Email_Retry durable=true arguments='{"x-message-ttl":60000,"x-dead-letter-exchange":"","x-dead-letter-routing-key":"Email_Standard"}'

rabbitmqadmin -u user -p password declare queue name=SMS_High_Priority durable=true arguments='{"x-max-priority": 10}'
rabbitmqadmin -u user -p password declare queue name=SMS_Standard durable=true
rabbitmqadmin -u user -p password declare queue name=SMS_Retry durable=true arguments='{"x-message-ttl":60000,"x-dead-letter-exchange":"","x-dead-letter-routing-key":"SMS_Standard"}'

rabbitmqadmin -u user -p password declare queue name=WebPush_High_Priority durable=true arguments='{"x-max-priority": 10}'
rabbitmqadmin -u user -p password declare queue name=WebPush_Standard durable=true
rabbitmqadmin -u user -p password declare queue name=WebPush_Retry durable=true arguments='{"x-message-ttl":60000,"x-dead-letter-exchange":"","x-dead-letter-routing-key":"WebPush_Standard"}'
