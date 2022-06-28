FROM golang:1.17-stretch as server-builder

ARG GOPROXY
ARG GOTAGS
ARG GOGCFLAGS

WORKDIR /go/src/d7y.io/dragonfly/v2

COPY . /go/src/d7y.io/dragonfly/v2

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN make build-dfget && make install-dfget


FROM reg.docker.alibaba-inc.com/amwp/kubeone-base:latest
ENV APPNAME dragonfly2-seedpeer
ENV DFGET_CONFIG /home/admin/$APPNAME/config/seedpeer.yaml

WORKDIR /home/admin/
RUN echo "ps -eo pid,command|grep 'dragonfly2-seedpeer/target/dfget'|grep -v 'grep' | awk '{print \$1}'|xargs kill -2" > /home/admin/stop.sh && \
    echo "/home/admin/dragonfly2-seedpeer/target/dfget daemon --config /home/admin/dragonfly2-seedpeer/config/seedpeer.yaml > /home/admin/dragonfly2-seedpeer/service.log 2>&1 &" > /home/admin/start.sh && \
    echo "/home/admin/dragonfly2-seedpeer/target/dfget --workhome /home/admin/dragonfly2-seedpeer \"\$@\"" > /home/admin/dfget && \
    chmod +x /home/admin/dfget
RUN chown -R admin:admin /home/admin

COPY --from=server-builder /opt/dragonfly/bin/dfget /home/admin/$APPNAME/target/
COPY --from=server-builder /go/bin/dlv /bin/
