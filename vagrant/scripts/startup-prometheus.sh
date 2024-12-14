#!/bin/bash


ARCHIVE="prometheus.tar.gz"
FOLDER="prometheus"

sudo apt update
sudo apt install -y curl wget vim

curl -L -o "$ARCHIVE" https://github.com/prometheus/prometheus/releases/download/v3.0.1/prometheus-3.0.1.linux-amd64.tar.gz

tar -xvf "$ARCHIVE"
EXTRACTED_FOLDER=$(tar -tf "$ARCHIVE" | head -1 | cut -f1 -d"/")

if [ -d "$EXTRACTED_FOLDER" ]; then
    mv "$EXTRACTED_FOLDER" "$FOLDER"
else
    exit 1
fi

sudo mv $FOLDER/prometheus /usr/local/bin/
sudo mv $FOLDER/promtool   /usr/local/bin/

sudo addgroup --system prometheus
sudo adduser --shell /sbin/nologin --system --group prometheus

sudo mkdir /etc/prometheus
sudo mkdir /var/lib/prometheus
sudo chown -R prometheus:prometheus /var/lib/prometheus/
sudo chown -R prometheus:prometheus /etc/prometheus
sudo chown -R prometheus:prometheus /usr/local/bin/prometheus
sudo chown -R prometheus:prometheus /usr/local/bin/promtool

sudo tee /etc/systemd/system/prometheus.service > /dev/null <<EOF
[Unit]
Description=Prometheus Monitoring System
Documentation=https://prometheus.io/docs/
Wants=network-online.target
After=network-online.target

[Service]
User=prometheus
Group=prometheus
Type=simple
ExecReload=/bin/kill -HUP \$MAINPID
ExecStart=/usr/local/bin/prometheus \
  --config.file=/etc/prometheus/prometheus.yml \
  --storage.tsdb.path=/var/lib/prometheus/ \
  --web.listen-address=0.0.0.0:9090 \
  --web.external-url= 

SyslogIdentifier=prometheus
Restart=always

[Install]
WantedBy=multi-user.target
EOF


sudo systemctl daemon-reload
sudo systemctl start prometheus
sudo systemctl enable prometheus

rm prometheus.tar.gz
