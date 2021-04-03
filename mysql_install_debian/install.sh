#!/bin/bash
wget https://dev.mysql.com/get/mysql-apt-config_0.8.16-1_all.deb
dpkg -i mysql-apt-config*
apt-get update
apt-get install mysql-community-server

## edit to allow mysql connect from any host, by default it's 127.0.0.1
# sudo vim /etc/mysql/mysql.conf.d/mysqld.cnf