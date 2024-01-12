FROM alpine:3.18
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
ENV SONOBUOY_VERSION 0.57.0
RUN apk add curl tar gzip
RUN curl -sfL https://github.com/vmware-tanzu/sonobuoy/releases/download/v${SONOBUOY_VERSION}/sonobuoy_${SONOBUOY_VERSION}_linux_amd64.tar.gz | tar xvzf - -C /usr/bin
COPY run-test.sh /usr/bin
CMD ["/usr/bin/run-test.sh"]
