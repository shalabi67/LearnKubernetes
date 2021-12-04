FROM ubuntu

RUN apt-get update && apt-get -y install wget && apt-get -y install systemd

RUN wget https://github.com/prometheus/node_exporter/releases/download/v1.3.0/node_exporter-1.3.0.linux-amd64.tar.gz && \
    tar xvfz node_exporter-1.3.0.linux-amd64.tar.gz

RUN useradd -M -r -s /bin/false node_exporter && \
    cp node_exporter-1.3.0.linux-amd64/node_exporter /usr/local/bin/ && \
    chown node_exporter:node_exporter /usr/local/bin/node_exporter

COPY node-exporter.service /etc/systemd/system/node_exporter.service
COPY exec.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/exec.sh
CMD /usr/local/bin/exec.sh




