# Installing Node Exporter (for port 9100) (LINUX)

### Installing Node Exporter on the nodes is done as follows:

##### Create a Node Exporter user for running the exporter:
```
    sudo useradd --no-create-home --shell /bin/false node_exporter
```
##### Download and extract the latest version of Node Exporter:
```
wget https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz
tar -xzvf node_exporter-0.18.1.linux-amd64.tar.gz
```
##### Send the executable to /usr/local/bin:
```
sudo cp node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/
```
##### Give the Node Exporter user ownership of the executable:
```
sudo chown node_exporter:node_exporter /usr/local/bin/node_exporter
```
##### Perform some cleanup and create and save a Node Exporter service with the below contents:
```
sudo rm node_exporter-0.18.1.linux-amd64 -rf
sudo nano /etc/systemd/system/node_exporter.service
```
```
[Unit]
Description=Node Exporter
Wants=network-online.target
After=network-online.target

[Service]
User=node_exporter
Group=node_exporter
Type=simple
ExecStart=/usr/local/bin/node_exporter
 
[Install]
WantedBy=multi-user.target
```
##### Reload systemctl services list, start the service and enable it to have it start on system restart:

```
sudo systemctl daemon-reload
sudo systemctl start node_exporter
sudo systemctl enable node_exporter
sudo systemctl status node_exporter
```

##### Check if the installation was successful by checking if {NODE_IP}:{PORT}/metrics is accessible from a web browser.