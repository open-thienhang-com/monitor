<div align="center">
  <a href="thienhang.com">
    <img src="https://creative-hub.thienhang.com/5592b9573e8d329d8ece.png" alt="Logo">
  </a>

  <h1 align="center">Event Proxy API</h1>

  <p align="center">
    <a href="thienhang.com">View Demo</a> ·
    <a href="https://thienhang.com">Report Bug</a> ·
    <a href="https://thienhang.com">Request Feature</a>
  </p>
</div>

[![View Demo](https://img.shields.io/badge/View-Demo-blue)](thienhang.com)
[![Report Bug](https://img.shields.io/badge/Report-Bug-red)](https://thienhang.com)
[![Request Feature](https://img.shields.io/badge/Request-Feature-green)](https://thienhang.com)

## Table of Contents

- [Introduction](#introduction)
- [Demo](#demo)
- [Features](#features)
- [Installation](#installation)
  - [Docker and Go Application Management](#docker-and-go-application-management)
    - [Getting Started](#getting-started)
      - [Development Environment](#development-environment)
        - 👨‍💻 For local development:
        ```bash
        make dev
        ```
        - 🚀 For development environment:
        ```bash
        make dev
        ```
- [Helper links](#helper-links)
  - [Jaeger UI](http://localhost:16686)
  - [Prometheus UI](http://localhost:9090)
  - [Grafana UI](http://localhost:3000)
  - [Swagger UI (if exist)](http://localhost:5001/swagger/index.html)
- [Benchmarks](#-benchmarks)
- [Performance Summary](#-performance-summary)
- [Contributing](#contributing)
- [Bug Reports and Feature Requests](#bug-reports-and-feature-requests)
- [License](#license)

## Introduction

Briefly introduce your project. Describe what The Exchange is and what problem it solves. Provide any relevant context or background information.

## Demo

Include a link to the live demo of your project. Screenshots or GIFs demonstrating the key features can also be added here.

## Features

List the key features of The Exchange. You can use bullet points for a concise overview.

- Feature 1
- Feature 2
- ...

## Installation

## Docker and Go Application Management

This sessions provides a set of Docker and Go commands to streamline the development and management of a Go application using Docker. It includes commands for starting local development environments, managing Docker containers, handling Go module dependencies, and utilizing Go's pprof for profiling.

### Getting Started

#### Development Environment

To start the local development environment with Docker Compose:

#### 👨‍💻 For local development:

```bash
make dev
```

#### 🚀 For develop environment :

```bash
make dev
```

## Helper link
### PGAdmin

Using email and password in .env file to login to pgAdmin. After logging in with your credentials of the .env file, you can add your database to pgAdmin.

Right-click "Servers" in the top-left corner and select "Create" -> "Server..."
Name your connection
Change to the "Connection" tab and add the connection details:
Hostname: "host.docker.internal" (this would normally be your IP address of the postgres database - however, docker can resolve this container ip by its name)
> host.docker.internal
Port: "5432"
Maintenance Database: $POSTGRES_DB (see .env)
Username: $POSTGRES_USER (see .env)
Password: $POSTGRES_PW (see .env)

### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090

### Grafana UI:

http://localhost:3000

### Swagger UI (if exist):

http://localhost:5001/swagger/index.html

# 🧲 Benchmarks


[![View Report Duration](https://img.shields.io/badge/View-Demo-blue)](https://c.thienhang.com/display/AA/Report+02+Feb+2024+-+Event+Tracking+Proxy)



# 📊 Performance Summary:

Please check out the report here: https://c.thienhang.com/pages/viewpage.action?pageId=42140287


