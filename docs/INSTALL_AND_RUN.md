# Install and Run the Oasis API Server

## Configuring the Nodes

In order for the API to be able to run correctly, Prometheus metrics should be enabled in the Oasis node's configuration file `config.yml` in the `/serverdir/etc/` directory which was set during the installation of the Oasis node. To enable the metrics append this to the end of the file :

```
metrics:
    mode: pull
    address: 0.0.0.0:9090 
```

Change `:9090` to the port you want prometheus to be exposed at.

## Installing the Node Exporter

Please follow the [installation guide](INSTALL_NODE_EXPORTER.md) to install Node Exporter that will be running along side your node.

## Configuring the API

Configuring the API involves setting up three config files. This is a strict requirement for the API to be able to run.

Start off by cloning this repository and navigating into it:
```bash
git clone https://github.com/SimplyVC/oasis_api_server
cd oasis_api_server
```

The API can be configured in either of two ways, by [using the setup script](#using-the-setup-script), or [manually](#manually).

### Using the Setup Script

Please note that you will need to have installed Python before using this script.
To install Python, pip and pipenv follow [this guide](INSTALL_PYTHON.md).

You can run the Python script `run_setup.py` which will guide you through the configuration process, and ask you to enter the values one by one. This can be done by running the following, starting from the root project directory, `oasis_api_server`:
```bash
pipenv sync
pipenv run python run_setup.py
```
### Manually

Alternatively, for advanced users, you can make a copy of the `example_*.ini` files inside the `config` folder without the `example_` prefix, and manually change the values as required.

## Installing the API and Dependencies

This section will guide you through the installation of the API and any of its dependencies.

We recommend that the API is installed on a Linux system, given the simpler installation and running process. 
Nevertheless, instructions on how to install the API on a Windows system are also provided.

You can either [run the API from source](#running-from-source) or [run the API using Docker](#run-using-docker).
This affects what dependencies you will need to install.

### Running from Source

Running the API from source only requires you to install Golang.

#### Install Golang

It is assumed that since this API needs to be run on the same machine as the Oasis node then Golang is already installed,
therefore no documentation is provided for it's installation,

#### Running the API

After having installed golang you can now run the API as follows from the project directory:
```bash
bash run_api.sh
```

#### Running the API as a Linux Service

Running the API as a service means that it starts up automatically on boot and restarts automatically if it runs into some issue and stops running. To do so, we recommend the following steps:
```bash
# Add a new user to run the API
sudo adduser <USER>

# Grant permissions
sudo chown -R <USER>:<USER> <API_DIR>/  # ownership of api
sudo chmod +x <API_DIR>/run_api.sh      # execute permission for runner
```

The service file will now be created:

```bash
# Create the service file
sudo nano /etc/systemd/system/oasis_api_server.service
```

It should contain the following, replacing `<USER>` with the created user's name and the two `<API_DIR>` with the API Server's installation directory:

```bash
[Unit]
Description=Oasis API Server
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
User=<USER>
TimeoutStopSec=90s
WorkingDirectory=<API_DIR>/
ExecStart=/bin/bash <API_DIR>/run_api.sh

[Install]
WantedBy=multi-user.target
```

Lastly, we will *enable* and *start* the alerter service:
```bash
sudo systemctl enable oasis_api_server.service
sudo systemctl start oasis_api_server.service
```

Check out `systemctl status oasis_api_server` to confirm that the API is running.


### Run using Docker

To run the API using Docker, you shouldn't be surprised to find out that you need to install Docker.

You will then obtain the Docker image, make sure that the config files are where they should be, and run everything.

#### Installing Docker on your Machine

To install Docker on your machine, follow [this guide](INSTALL_DOCKER.md)

#### Obtaining the Docker Image

This part can be done in either of two ways, either by building the Docker image yourself, or by downloading it from Docker Hub.

##### Building the Docker Image

First start off by cloning this repository:
```bash
git clone https://github.com/SimplyVC/oasis_api_server
```

Then run the following commands to build the image:
```bash
cd oasis_api_server
docker build -t simplyvc/oasis_api_server:1.0.6 .
```


##### Downloading the Pre-Built Docker Image from DockerHub

The pre-built Docker image can simply be downloaded by running the following command:
```bash
docker pull simplyvc/oasis_api_server:1.0.6
```

#### Config Files Directory and Permissions

The config files needed by the Docker image are the same as those generated in the `Configuring the API` section above.\
These config files can be moved to any directory of your choosing `<CONFIG_DIR>`.

##### On Ubuntu

If you created a new user `<USER>` earlier on, set the permissions as follows:
```bash
sudo chown -R <USER>:<USER> <CONFIG_DIR>
```

##### On Windows

No further steps are required.

#### Running the Docker Image

Now that the Docker image is on your machine, and you have written configurations for it, you can run it as follows, where `<CONFIG_DIR>` is the **full path** to the folder containing the previously created config files. Suppose that <CONFIG_DIR> is the **full path** to the folder containing the previously created config files, <INTERNAL_SOCK_DIR> is the **full path** on your machine where the internal socket file can be found for the Oasis node (e.g: serverdir/node/), and <PATH_IN_NODE_CONFIG> is the specified location in the API setup where the API will look for the internal socket (e.g: serverdir/node/). The API Server must also be able to find the **full path** of the Sentry `tls_identity_cert.pem` files. Therefore the path <INTERNAL_TLS_DIR> is the **full path** on your machine where the `tls_identity_cert.pem` file can be found, and <PATH_IN_SENTRY_CONFIG> is the specified location in the API setup where the API will look for the `tls_identity_cert.pem` file (e.g: serverdir/etc/). Now that the Docker image is on your machine and you have written configuration for it, you can run it as follows:
```bash
docker run --network="host" -p 127.0.0.1:8686:8686 \
    --mount type=bind,source=<CONFIG_DIR>,target=/app/config/ \
    --mount type=bind,source=<INTERNAL_SOCK_DIR>,target=<PATH_IN_NODE_CONFIG> \
    --mount type=bind,source=<INTERNAL_TLS_DIR>,target=<PATH_IN_SENTRY_CONFIG> \
    -d simplyvc/oasis_api_server:1.0.6
```

Note: The port after `-p` and before the `:` is used to route a port from the machine to the internal port of the Docker. If this is changed, any program which refers to the API Docker container must refer to this port.\
Example: with `8686`:`5367`, the API URL must look like `http://1.2.3.4:8686`, but the configured port inside the files must be `5367`, it is suggested to leave them them both as `8686`.

## Confirming the API Works

If you wish to make sure that the API is running, the following should return `{"result":"pong"}`:
```bash
curl -X GET http://localhost:8686/api/ping
```

If you wish to check the API's connection to a node, you can run the following for some node `<NODE>`:
```bash
curl -X GET http://localhost:8686/api/pingnode?name=<NODE>
```

**Note**

Using curl with multiple queries requires the entire url to be inside double quotation marks as seen in the example below:

```bash
curl -X GET "http://localhost:8686/api/registry/nodes?name=<NODE>&height=300"
```
---
[Back to API front page](../README.md)
