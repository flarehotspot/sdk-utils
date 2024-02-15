+++
title = "Getting Started - Windows"
description = "Series of tutorials to get you started with using the Flare Hotspot SDK."
date = 2021-05-01T08:00:00+00:00
updated = 2021-05-01T08:00:00+00:00
draft = false
weight = 1
sort_by = "weight"
template = "docs/page.html"
+++

# Getting Started

## Overview

This tutorial is for Windows users. If you are a linux user, you can follow the [linux tutorial](../getting-started-linux) instead.

The plugin SDK is a set of tools and libraries that allow you to build applications that interact with the [Flare Hotspot System](https://www.flarehotspot.com). It is developed using the [Go Programming Language](https://go.dev). Go is a simple but fast and efficient programming language suitable for building system applications while also being easy to learn for beginners. This series of tutorials will guide you through the process of building a simple plugin using the Flare Hotspot SDK.

---


# Prerequisites
All we need to get started is [Docker](https://www.docker.com). Below are the instructions to install Docker Desktop on your system.

## Installing Docker Desktop

We'll use the [Docker Desktop](https://docs.docker.com/desktop/install/linux-install/) for convenience. Although the docker desktop is not officially supported on Linux, we can still install it using the following steps.

Add your user to the `kvm` group:
```sh
sudo usermod -aG kvm $USER
```

You need to **REBOOT** your computer after adding your user to the `kvm` group. After rebooting, you can proceed to install Docker Desktop.
```sh
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

# Download and install docker-desktop debian package:
wget https://desktop.docker.com/linux/main/amd64/137060/docker-desktop-4.27.2-amd64.deb -O ~/Downloads/docker-desktop.deb
sudo apt install -y ~/Downloads/docker-desktop.deb
```

**NOTE**: At the end of the installation process, `apt` displays an error due to installing a downloaded package. You can ignore this error message.

```sh
# N: Download is performed unsandboxed as root, as file '/home/user/Downloads/docker-desktop.deb' couldn't be accessed by user '_apt'. - pkgAcquire::Run (13: Permission denied)
```

---

# Installing Flare Hotspot SDK

## Download SDK
To install the plugin SDK, download the latest **devkit-x.x.xzip** file from [sdk-releases](https://github.com/flarehotspot/sdk-releases/releases) repository.
After downloading, extract the zip file to your desired location.
```sh
# replace ~/Downloads/devkit-0.0.5.zip with the path to the downloaded zip file
unzip ~/Downloads/devkit-0.0.5.zip -d ~/Documents/devkit-0.0.5
cd ~/Documents/devkit-0.0.5
```

Below is the directory structure of the zip file:
```
|- config
|- core
|- main
|- mock-files
|- plugins
    |-- com.flarego.default-theme
    |-- com.flarego.sample-plugin
```

Notice the `com.flarego.sample-plugin` directory. We are going to use this sample plugin as the base of our project.

## Start SDK Runtime

To start the SDK runtime, you need to run:
```sh
cd ~/Documents/devkit-0.0.5
docker compose up
```

Now you can access the Flare Hotspot web interface:

- Captive Portal: [http://localhost:3000](http://localhost:3000)
- Admin Dashboard: [http://localhost:3000/admin](http://localhost:3000/admin)
- Database Management: [http://localhost:8080](http://localhost:8080)

The default login for the admin dashboard is:
```
username: admin
password: admin
```

To stop the SDK runtime, you need to run:
```sh
cd ~/Documents/devkit-0.0.5
docker compose down
```
