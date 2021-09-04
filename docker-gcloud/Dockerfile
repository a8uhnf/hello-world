FROM google/cloud-sdk:355.0.0-debian_component_based

RUN apt -y update

RUN apt -y install \
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

RUN apt -y update

RUN apt install -y docker-ce docker-ce-cli containerd.io jq

# RUN docker version

RUN apt -y update
