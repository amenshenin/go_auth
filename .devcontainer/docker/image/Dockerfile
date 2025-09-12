FROM ubuntu:22.04

ARG UID
ENV UID $UID
ARG GID
ENV GID $GID
ARG GO_ARCHIVE
ENV GO_ARCHIVE $GO_ARCHIVE

RUN apt-get update -y
RUN apt install curl ca-certificates openssl tar wget git -y

RUN mkdir -p /usr/local/go
RUN mkdir -p /usr/local/go/{bin,src}
RUN chown -R root:root /usr/local/go
RUN wget -O arch.tar.gz https://go.dev/dl/$GO_ARCHIVE
RUN  rm -rf /usr/local/go && tar -C /usr/local -xzf arch.tar.gz

RUN echo 'export GOPATH=/home/root/work' >> ~/.profile
RUN echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.profile

WORKDIR /home/root/work

EXPOSE 8888

CMD tail -f /dev/null
