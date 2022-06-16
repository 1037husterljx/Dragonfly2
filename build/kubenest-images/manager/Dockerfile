ARG BASE_IMAGE=alpine:3.14

FROM golang:1.17.3-alpine3.14 as server-builder

ARG GOPROXY
ARG GOTAGS
ARG GOGCFLAGS

WORKDIR /go/src/d7y.io/dragonfly/v2

RUN apk --no-cache add bash make gcc libc-dev git

COPY . /go/src/d7y.io/dragonfly/v2

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN make build-manager && make install-manager

FROM node:12-alpine as console-builder
WORKDIR /build
COPY manager/console/package.json /build
RUN npm install --loglevel warn --progress false
COPY manager/console /build
RUN npm run build

FROM reg.docker.alibaba-inc.com/amwp/kubeone-base:latest
ENV APPNAME dragonfly2-manager
ENV MANAGER_CONFIG /home/admin/dragonfly2-manager/config/manager.yaml

WORKDIR /home/admin/
RUN echo "ps -eo pid,command|grep 'dragonfly2-manager/target/manager'|grep -v 'grep' | awk '{print \$1}'|xargs kill -2" > /home/admin/stop.sh && \
    echo "/home/admin/dragonfly2-manager/target/manager  > /dev/null 2>&1 &" > /home/admin/start.sh
RUN chown -R admin:admin /home/admin

COPY --from=server-builder /opt/dragonfly/bin/manager /home/admin/$APPNAME/target/
COPY --from=console-builder /build/dist /home/admin/dragonfly2-manager/target/console
COPY --from=server-builder /go/bin/dlv /bin/
