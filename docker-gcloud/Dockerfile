FROM google/cloud-sdk:319.0.0-debian_component_based

RUN apt-get -y update

RUN apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -

RUN add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"

RUN apt-get -y update

RUN apt-get install -y docker-ce docker-ce-cli containerd.io

RUN chown 777 /var/run/docker.sock

RUN docker version