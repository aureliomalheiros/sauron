#!/bin/bash


ARCHIVE="node_exporter.tar.gz"
FOLDER="node_exporter"

sudo apt update

curl -L -o "$ARCHIVE" https://github.com/prometheus/node_exporter/releases/download/v1.8.2/node_exporter-1.8.2.linux-amd64.tar.gz

tar -xvf "$ARCHIVE"
EXTRACTED_FOLDER=$(tar -tf "$ARCHIVE" | head -1 | cut -f1 -d"/")

if [ -d "$EXTRACTED_FOLDER" ]; then
    mv "$EXTRACTED_FOLDER" "$FOLDER"
else
    exit 1
fi

sudo mv $FOLDER/node_exporter /usr/local/bin/

sudo addgroup --system node_exporter
sudo adduser --shell /sbin/nologin --system --group node_exporter


sudo tee /etc/systemd/system/node_exporter.service > /dev/null <<EOF
[Unit]
Description=Node Exporter
Documentation=https://prometheus.io/docs/
Wants=network-online.target
After=network-online.target

[Service]
User=node_exporter
Group=node_exporter
Type=simple
ExecReload=/bin/kill -HUP \$MAINPID
ExecStart=/usr/local/bin/node_exporter 

[Install]
WantedBy=multi-user.target
EOF


sudo systemctl daemon-reload
sudo systemctl start node_exporter
sudo systemctl enable node_exporter

rm node_exporter.tar.gz
