#!/bin/bash
wget https://dev.mysql.com/get/mysql-apt-config_0.8.16-1_all.deb
dpkg -i mysql-apt-config*
apt-get update
apt-get install mysql-community-server
