<div align="center">
    <h1> Docktor</h1>
    <h2> Analyze your Dev environment with surgical precision </h2>
<div>

![](https://img.shields.io/static/v1.svg?label=DevSecOps&message=Docktor&color=f15f86)
![](https://img.shields.io/static/v1.svg?label=Fast&message=Lightweight&color=a8db75)
![](https://img.shields.io/static/v1.svg?label=Scan&message=Docker&color=fdd765)
![](https://img.shields.io/static/v1.svg?label=Scan&message=FileSystem&color=f69666)
![](https://img.shields.io/static/v1.svg?label=Scan&message=RepoGIT&color=aa9cf2)
![](https://img.shields.io/static/v1.svg?label=Scan&message=ConfigFiles&color=78dbe7)

<div align="left">

## About

Docktor is a Web App that deploys an easy-to-use kit of analysis and scanning tools.

Today, developers use a variety of resources. It is more and more difficult to ensure the security of our artifacts.
Especially Docker environments which are an obvious source of vulnerability.

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/report.png)

> **The objective is to have a simple, fast, lightweight and everywhere approach to ensure the security of our productions.**

<div align="center">

### The Docktor Kit Composition

<img src="https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/logo/svelte.png" height="60">
<img src="https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/logo/gin.png" height="60">
<img src="https://d1q6f0aelx0por.cloudfront.net/product-logos/library-alpine-logo.png?" height="60">
<img src="https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/logo/trivy.svg" height="60">
</div>

## Features

> **The user interacts through the simple and pure web application, to select the elements he wants to analyze.**

Docktor takes care of the remaining work thanks to its 3 main components:

- **Sources**: are UIs to select the sources/artifacts/input folders that will be analyzed by the Scanners
- **Scanners**: are external tools that perform the processing
- **Vizualizer**: are UIs that are in charge of enhancing the data for the user

```mermaid
graph TD
    A[Docktor Frontend]

    B(SOURCE - Docker Images)
    C(SOURCE - File System)
    D(SOURCE - Repo GIT)
    E(SOURCE - Config Files)

    F(VIZUALIZER - Report - Scan / Vulenarabilities)

    W[Docktor Backend]

    Y{Computer}

    Z(((SCANNER - Trivy)))

    W --> |report.json| F
    F --> A

    A --> B
    A --> C
    A --> D
    A --> E

    B --> |local docker images| W
    B --> |remote docker images| W
    C --> |local path| W
    D --> |remote URI| W
    E --> |local path| W

    W --> |exec.Command| Y
    Y --> |report.json|W

    Y --> |./trivy ... |Z
    Z --> |report.json| Y
```

### [SOURCE] - Docker Images - SCA (Software Compisotion Analysis)

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/docker.png)

- Trivy detects:

  - Vulnerabilities

    - OS packages (Alpine, Red Hat Universal Base Image, Red Hat Enterprise Linux, CentOS, AlmaLinux, Rocky Linux, CBL-Mariner, Oracle Linux, Debian, Ubuntu, Amazon Linux, openSUSE Leap, SUSE Enterprise Linux, Photon OS and Distroless)

    - Language-specific packages (Bundler, Composer, Pipenv, Poetry, npm, yarn, Cargo, NuGet, Maven, and Go)

<br>
<br>
<br>

<details><summary>CLICK HERE TO VIEW RESULTS</summary>

<br>

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/report.png)

</details>

<br>
<br>
<br>

### [SOURCE] - File System - SCA (Software Compisotion Analysis)

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/file.png)

- Trivy detects:

  - Vulnerabilities

    - OS packages (Alpine, Red Hat Universal Base Image, Red Hat Enterprise Linux, CentOS, AlmaLinux, Rocky Linux, CBL-Mariner, Oracle Linux, Debian, Ubuntu, Amazon Linux, openSUSE Leap, SUSE Enterprise Linux, Photon OS and Distroless)

    - Language-specific packages (Bundler, Composer, Pipenv, Poetry, npm, yarn, Cargo, NuGet, Maven, and Go)

<br>
<br>
<br>

<details><summary>CLICK HERE TO VIEW RESULTS</summary>

<br>

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/file-res.png)

</details>

<br>
<br>
<br>

### [SOURCE] - Config Files - Misconfigurations Analysis

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/conf.png)

- Trivy detects:

  - Misconfigurations
    - Kubernetes
    - Docker
    - Terraform
    - CloudFormation
    - etc.
  - Secrets
    - AWS access key
    - GCP service account
    - GitHub personal access token
    - etc.

<br>
<br>
<br>

<details><summary>CLICK HERE TO VIEW RESULTS</summary>

<br>

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/conf-res.png)

</details>

<br>
<br>
<br>

### [SOURCE] - Repo GIT - SCA (Software Compisotion Analysis)

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/repo.png)

- Trivy detects:

  - Vulnerabilities

    - OS packages (Alpine, Red Hat Universal Base Image, Red Hat Enterprise Linux, CentOS, AlmaLinux, Rocky Linux, CBL-Mariner, Oracle Linux, Debian, Ubuntu, Amazon Linux, openSUSE Leap, SUSE Enterprise Linux, Photon OS and Distroless)

    - Language-specific packages (Bundler, Composer, Pipenv, Poetry, npm, yarn, Cargo, NuGet, Maven, and Go)

<br>
<br>
<br>

<details><summary>CLICK HERE TO VIEW RESULTS</summary>

<br>

![](https://raw.githubusercontent.com/Matbabs/Docktor/main/assets/img/repo-res.png)

</details>

<br>
<br>
<br>

## Getting started

### Run with - Docker compose (Preferred method)

```
curl -LJO https://github.com/Matbabs/Docktor/blob/main/docker-compose.yml
```

Inside the `docker-compose.yml` containing folder.

```
docker-compose up
```

> Note: your `/home` path is map with the `/home` container path, especially to scan and access your local files.

<br>

<details><summary>Run with - Docker run</summary>

<br>

### Run with - Docker run

```
docker run \
    -d \
    -p 3030:80 \
    -p 4040:4040 \
    -v "/var/run/docker.sock:/var/run/docker.sock:rw" \
    -v /home:/home \
    docktor
```

</details>

<br>

### Access UI

**Connect on: http://localhost:3030**
