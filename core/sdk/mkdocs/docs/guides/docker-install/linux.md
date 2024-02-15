# Install Docker Desktop for Linux

We are going to use the [Docker Desktop](https://docs.docker.com/desktop/install/linux-install/) for convenience. Although we can also use docker cli to run the Flare Hotspot development runtime, Docker Desktop provides a more user-friendly interface and is easier to manage. Follow the instructions below to install Docker Desktop on your Linux machine. For windows users, please refer to the [instructions for windows](./windows.md).

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

Now, you can proceed to [downloading and installing](../../getting-started.md#installing-flare-hotspot-sdk) the Flare Hotspot SDK.
