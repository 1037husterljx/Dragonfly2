ARG BASE_IMAGE=alpine:3.14

FROM golang:1.17.3-alpine3.14 as server-builder

ARG GOPROXY
ARG GOTAGS
ARG GOGCFLAGS

WORKDIR /go/src/d7y.io/dragonfly/v2

RUN apk --no-cache add bash make gcc libc-dev git

COPY . /go/src/d7y.io/dragonfly/v2

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN make build-scheduler && make install-scheduler


FROM reg.docker.alibaba-inc.com/amwp/kubeone-base:latest
ENV APPNAME dragonfly2-scheduler
ENV SCHEDULER_CONFIG /home/admin/$APPNAME/config/scheduler.yaml

WORKDIR /home/admin/
RUN echo "ps -eo pid,command|grep 'dragonfly2-scheduler/target/scheduler'|grep -v 'grep' | awk '{print \$1}'|xargs kill -2" > /home/admin/stop.sh && \
    echo "/home/admin/dragonfly2-scheduler/target/scheduler  > /dev/null 2>&1 &" > /home/admin/start.sh
RUN chown -R admin:admin /home/admin

COPY --from=server-builder /opt/dragonfly/bin/scheduler /home/admin/$APPNAME/target/
COPY --from=server-builder /go/bin/dlv /bin/
