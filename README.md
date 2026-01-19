# tcpScanner

TCP port scanner CLI tool. Scans ports and shows what's open.

## Build

```bash
cd script
go build -o tcpScanner ./cmd/app
```

## Usage

```bash
# Scan popular ports
./tcpScanner -d example.com -p

# Scan all 65535 ports
./tcpScanner -d 192.168.1.1 -a

# Show help
./tcpScanner -h
```

## Flags

- `-d` - Domain or IP address (required)
- `-a` - Scan all 65535 ports
- `-p` - Scan popular ports only
- `-h` - Show help

## Output

Shows open ports with service names (like nmap):

```
Starting scan for example.com (58 popular ports)
====================================================
Scanning... | [58/58] 100.0%
====================================================

Found 3 open port(s):
PORT     STATE  SERVICE
----------------------
80/tcp   open   http
443/tcp  open   https
22/tcp   open   ssh
```

## Popular Ports

Scans common ports: FTP, SSH, HTTP, HTTPS, databases (MySQL, PostgreSQL, MongoDB), remote access (RDP, VNC), and more.
