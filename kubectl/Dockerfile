FROM alpine
ARG VCS_REF
ARG BUILD_DATE

# Metadata
ENV KUBE_LATEST_VERSION="v1.9.0"
ENV HELM_VERSION="helm-v2.10.0-linux-amd64.tar.gz"

RUN apk add --update ca-certificates \
 && apk add --update -t deps curl \
 && curl -L https://storage.googleapis.com/kubernetes-release/release/${KUBE_LATEST_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
 && chmod +x /usr/local/bin/kubectl \
 && apk del --purge deps \
 && rm /var/cache/apk/*

RUN curl -L https://storage.googleapis.com/kubernetes-helm/${HELM_VERSION} -o . \ 
 && tar -zxvf HELM_VERSION \
 && mv linux-amd64/helm /usr/local/bin/helm \
 && chmod +x /usr/local/bin/helm

WORKDIR /root