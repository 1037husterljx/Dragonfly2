FROM golang:1.18-stretch as server-builder

ARG GOPROXY
ARG GOTAGS
ARG GOGCFLAGS

WORKDIR /go/src/d7y.io/dragonfly/v2

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
ENV APP_NAME dragonfly2-manager
ENV MANAGER_CONFIG /home/admin/dragonfly2-manager/config/manager.yaml

RUN yum clean all && \
    yum install tengine-proxy-2.5.4 -b current -y && \
    yum install -b current t-security-xagent -y && \
    yum clean all


WORKDIR /home/admin/
COPY build/kubenest-images/manager/admin /home/admin/
RUN chown -R admin:admin /home/admin

COPY --from=server-builder /opt/dragonfly/bin/manager /home/admin/$APP_NAME/target/
COPY --from=server-builder /go/bin/dlv /bin/
COPY --from=console-builder /build/dist /home/admin/dragonfly2-manager/target/console
