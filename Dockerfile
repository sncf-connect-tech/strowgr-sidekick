FROM centos:6.8

LABEL application.name "StrowgrSidekick"
LABEL application.desc "Strowgr docker sidekick"

COPY sidekick /sidekick
COPY entrypoint.sh /entrypoint.sh

# Install several versions of hap
RUN yum install -y gcc \
    && curl http://www.haproxy.org/download/1.4/src/haproxy-1.4.22.tar.gz |tar -xzC /tmp/ \
    && curl http://www.haproxy.org/download/1.4/src/haproxy-1.4.27.tar.gz |tar -xzC /tmp/ \
    && curl http://www.haproxy.org/download/1.5/src/haproxy-1.5.18.tar.gz |tar -xzC /tmp/ \
    && cd /tmp/haproxy-1.4.22 && make TARGET=linux26 && mkdir -p /opt/haproxy-1.4.22/bin && cp haproxy /opt/haproxy-1.4.22/bin \
    && cd /tmp/haproxy-1.4.27 && make TARGET=linux26 && mkdir -p /opt/haproxy-1.4.27/bin && cp haproxy /opt/haproxy-1.4.27/bin \
    && cd /tmp/haproxy-1.5.18 && make TARGET=linux26 && mkdir -p /opt/haproxy-1.5.18/bin && cp haproxy /opt/haproxy-1.5.18/bin \
    && rm -rf /tmp/haproxy* \
    && chmod +x /sidekick && chmod +x /entrypoint.sh \
    && yum remove -y gcc \
    && yum clean all

ENV LOOKUP_ADDR="nsqlookupd:4161"
ENV PRODUCER_ADDR="nsqd:4150"
ENV PRODUCER_REST_ADDR="http://nsqd:4151"
ENV CLUSTER_ID="cluster"
ENV VIP="sidekick"
ENV HTTP_PORT="50000"
ENV HAP_HOME="/data"
ENV STATUS="master"
ENV ID="sidekick"
ENV ARGS=""

VOLUME /data

EXPOSE 50000

ENTRYPOINT ["/entrypoint.sh"]
