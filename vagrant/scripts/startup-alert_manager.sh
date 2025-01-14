#!/bin/bash


ARCHIVE="alert_manager.tar.gz"
FOLDER="alert_manager"

sudo apt update

curl -L -o "$ARCHIVE" https://github.com/prometheus/alertmanager/releases/download/v0.27.0/alertmanager-0.27.0.linux-amd64.tar.gz

tar -xvf "$ARCHIVE"
EXTRACTED_FOLDER=$(tar -tf "$ARCHIVE" | head -1 | cut -f1 -d"/")

if [ -d "$EXTRACTED_FOLDER" ]; then
    mv "$EXTRACTED_FOLDER" "$FOLDER"
else
    exit 1
fi

sudo mv $FOLDER/alertmanager /usr/local/bin/
sudo mv $FOLDER/amtool /usr/local/bin

sudo mkdir /etc/alertmanager
sudo mkdir /etc/alertmanager/templates
sudo mkdir /var/lib/alertmanager

sudo addgroup --system alertmanager
sudo adduser --shell /sbin/nologin --system --group alertmanager

sudo chown -R alertmanager:alertmanager /etc/alertmanager
sudo chown -R alertmanager:alertmanager /var/lib/alertmanager

sudo tee /etc/systemd/system/alertmanager.service > /dev/null <<EOF
[Unit]
Description=Prometheus Alert Manager
Documentation=https://prometheus.io/docs/
Wants=network-online.target
After=network-online.target

[Service]
User=alertmanager
Group=alertmanager
Type=simple
ExecStart=/usr/local/bin/alertmanager \
    --config.file=/etc/alertmanager/alertmanager.yml \
    --web.route-prefix=/

[Install]
WantedBy=multi-user.target
EOF


sudo systemctl daemon-reload
sudo systemctl start alertmanager
sudo systemctl enable alertmanager

rm alert_manager.tar.gz
