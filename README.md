````markdown
# TCP Scanner

## Service Overview

This service performs TCP port scanning for a specified host. It checks which ports are open and which are closed within the range from **1 to 65,535** (all valid TCP ports).

## Architecture

The project consists of two main components:
- **Backend** - a Go API server built using the standard `net/http` library
- **Frontend** - a static web interface implemented with HTML, CSS, and JavaScript

## Quick Start with Docker

### Requirements
- Docker  
- Docker Compose  

### Run the Entire Project with a Single Command

```bash
# Clone the repository
git clone <repository-url>
cd tcpScanner

# Build and start the application
docker-compose up --build
````

After startup:

* **Frontend** will be available at: [http://localhost:3000](http://localhost:3000)
* **Backend API** will be available at: [http://localhost:8080](http://localhost:8080)

### Stopping the Project

```bash
# Stop containers
docker-compose down

# Stop containers and remove volumes
docker-compose down -v
```

## API Documentation

### [POST] /scan

Scans TCP ports for the specified host.

**Request:**

```json
{
    "host": "example.com"
}
```

**Response:**

```json
{
    "host": "example.com",
    "open_ports": [80, 443, 22, 8080]
}
```

### [GET] /health

Service health check endpoint.

**Response:**

```json
{
    "status": "healthy",
    "service": "tcp-scanner"
}
```

## Performance

Scan duration depends on the target host, network latency, and firewall configuration.

Approximate scan times for popular targets:

* `scanme.nmap.org` - ~25 seconds
* `sberbank.ru` - ~2 minutes 25 seconds
* `outline.com` - ~1 minute 15 seconds

## Security

* Input validation at the API level
* CORS configuration for secure frontend–backend interaction
* Validation of IP address and domain name formats
