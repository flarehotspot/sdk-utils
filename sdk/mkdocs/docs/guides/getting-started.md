# Getting Started

## 1. Install Docker Desktop {#install-docker}

All we need to run Flare Hotspot development runtime is [Docker](https://www.docker.com/). You can download it from [here](https://www.docker.com/products/docker-desktop).

For a detailed instruction to install Docker Desktop, please refer to the [official documentation](https://docs.docker.com/desktop/) or follow our beginner-friendly tutorial for [windows](./install-docker-windows.md).

!!!note
    For Linux/Mac users, you can install Docker using the package manager of your distribution.

## 2. Download Flare Hotspot SDK {#download-sdk}

Download the latest **devkit-x.x.xzip** file from [https://github.com/flarehotspot/sdk/releases](https://github.com/flarehotspot/sdk/releases) repository. Select the appropriate zip file that's compatible with your CPU architecture. Windows computers are most likely be running Intel or AMD x86 CPUs, so just select `devkit-0.0.13-pre-amd64.zip` (whatever is the latest release file).

![Download Flare Hotspot SDK](./img/01-select-latest-release.png)

After downloading, extract the zip file to your desired location.

![Extract Flare Hotspot Sdk](./img/02-extract-devkit.png)


![Extract Flare Hotspot Sdk](./img/03-extract-devkit.png)

## 3. Start The SDK Runtime {#start-sdk}

To start the SDK runtime, open windows `CMD` or `PowerShell` and navigate to the extracted file's root directory then run:
```sh
docker compose up --build
```

For [VSCode](https://code.visualstudio.com/) users, you can also do this in the terminal.

![Run docker compose up](./img/04-docker-compose-up.png)

Docker may take sometime to download and install the container and its dependencies. Wait for the message `Listening on port :3000` which indicates that the server is already running and ready to accept connections.

![Server is running](./img/05-server-is-running.png)

Now you can access the Flare Hotspot web interface:

- Captive Portal: [http://localhost:3000](http://localhost:3000)
- Admin Dashboard: [http://localhost:3000/admin](http://localhost:3000/admin)
- Database Management: [http://localhost:8080](http://localhost:8080)

To restart the server and recompile the code, you can run:
```sh
docker compose restart app
```

!!!note
    The default login for the admin dashboard is:
    ```
    username: admin
    password: admin
    ```

## 4. Upgrading The SDK {#upgrade-sdk}

To upgrade the SDK, download the latest release from the [SDK releases page](https://github.com/flarehotspot/sdk/releases) and extract the contents to your desired destination. Then copy your plugin directory into the `plugins` directory inside the new SDK.

Now you can [start](#start-sdk) the new SDK runtime.

