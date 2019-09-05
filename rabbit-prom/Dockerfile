FROM rabbitmq:3.7.2-management

RUN apt-get update

RUN apt-get install wget unzip -y

#RUN apt-get install -y unzip

RUN wget https://dl.bintray.com/rabbitmq/community-plugins/3.7.x/rabbitmq_delayed_message_exchange/rabbitmq_delayed_message_exchange-20171201-3.7.x.zip
RUN unzip rabbitmq_delayed_message_exchange-20171201-3.7.x.zip
RUN mv rabbitmq_delayed_message_exchange-20171201-3.7.x.ez plugins

RUN wget https://github.com/deadtrickster/prometheus_rabbitmq_exporter/releases/download/v3.7.2.3/prometheus_cowboy-0.1.4.ez

RUN mv prometheus_cowboy-0.1.4.ez plugins

RUN wget https://github.com/deadtrickster/prometheus_rabbitmq_exporter/releases/download/v3.7.2.3/prometheus_httpd-2.1.8.ez

RUN mv prometheus_httpd-2.1.8.ez plugins 

RUN wget https://github.com/deadtrickster/prometheus_rabbitmq_exporter/releases/download/v3.7.2.3/prometheus-3.5.1.ez

RUN mv prometheus-3.5.1.ez plugins 

RUN wget https://github.com/deadtrickster/prometheus_rabbitmq_exporter/releases/download/v3.7.2.3/accept-0.3.3.ez

RUN mv accept-0.3.3.ez plugins

RUN rabbitmq-plugins enable rabbitmq_delayed_message_exchange



RUN wget https://github.com/deadtrickster/prometheus_rabbitmq_exporter/releases/download/v3.7.2.3/prometheus_rabbitmq_exporter-3.7.2.3.ez

RUN mv prometheus_rabbitmq_exporter-*.ez plugins/

RUN ls -la plugins/ 
RUN pwd


#RUN unzip $RABBITMQ_HOME/plugins/rabbitmq_delayed_message_exchange-20171201-3.7.x.zip -d $RABBITMQ_HOME/plugins

#RUN ls -la $RABBITMQ_HOME/plugins/ 

#RUN rm -rf $RABBITMQ_HOME/plugins/rabbitmq_delayed_message_exchange-20171201-3.7.x.zip


RUN rabbitmq-plugins enable --offline rabbitmq_consistent_hash_exchange

RUN rabbitmq-plugins enable prometheus_rabbitmq_exporter
