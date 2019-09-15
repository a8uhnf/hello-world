FROM rabbitmq:3.7-management

RUN apt-get update
RUN apt-get install wget unzip -y
RUN wget https://dl.bintray.com/rabbitmq/community-plugins/3.7.x/rabbitmq_delayed_message_exchange/rabbitmq_delayed_message_exchange-20171201-3.7.x.zip
RUN unzip rabbitmq_delayed_message_exchange-20171201-3.7.x.zip
RUN mv rabbitmq_delayed_message_exchange-20171201-3.7.x.ez plugins

RUN rabbitmq-plugins enable rabbitmq_delayed_message_exchange

#USER root
#RUN ls -la /etc/rabbitmq/
#RUN cat /etc/rabbitmq/rabbitmq.conf
ADD ./rabbitmq.cfg rabbitmq.cfg
RUN cat rabbitmq.cfg >> /etc/rabbitmq/rabbitmq.conf

RUN cat /etc/rabbitmq/rabbitmq.conf

RUN rabbitmq-plugins enable rabbitmq_peer_discovery_k8s
