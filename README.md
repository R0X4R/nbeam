# nbeam

nbeam is a tiny **Go CLI** for one job, turn noisy Nmap output into clean `host:port` lines. I built it after too much time fighting **regexes** and trying to **wrangle raw scan** output into something usable. It exists so I could stop doing that work by hand.

It reads from stdin, detects whether the input looks like Nmap **normal output** or **XML**, and prints live targets in a format that is easy to pipe into the rest of a recon chain.

## Why it exists

If you have ever tried to pull open ports out of Nmap output with regex, you already know the pain. **nbeam** is the result of wanting a fast, no-dependency way to take that mess and get straight to `host:port` pairs.

## Features

- Auto-detects Nmap normal output and XML input from stdin.
- Extracts open ports and prints them as host:port.
- Uses only the Go standard library.
- Works with live scans or saved Nmap output files.
- Plays nicely with other tools like httpx and nuclei.

## Installation

```bash
go install github.com/R0X4R/nbeam@latest
```

**Install from source:**

```bash
git clone https://github.com/R0X4R/nbeam.git && cd nbeam && go install .
```

## Usage

```bash
nmap -sS -sV -T4 -p- scanme.nmap.org -oA nmap

Starting Nmap 7.94SVN ( https://nmap.org ) at 2026-06-09 10:47 CEST

Nmap scan report for scanme.nmap.org (45.33.32.156)

Host is up (0.18s latency).

Other addresses for scanme.nmap.org (not scanned): 2600:3c01::f03c:91ff:fe18:bb2f

Not shown: 65531 closed tcp ports (reset)

PORT      STATE SERVICE    VERSION
22/tcp    open  ssh        OpenSSH 6.6.1p1 Ubuntu 2ubuntu2.13 (Ubuntu Linux; protocol 2.0)
80/tcp    open  http       Apache httpd 2.4.7 ((Ubuntu))
9929/tcp  open  nping-echo Nping echo
31337/tcp open  tcpwrapped
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .

Nmap done: 1 IP address (1 host up) scanned in 47.09 seconds
```

Pipe a live Nmap scan directly into nbeam:

```bash
nmap -sS -sV -T4 -p- scanme.nmap.org -oN - | nbeam

scanme.nmap.org:22
scanme.nmap.org:80
scanme.nmap.org:9929
scanme.nmap.org:31337
```

You can also pipe saved Nmap output files:

```bash
cat nmap.nmap | nbeam
```

```bash
cat nmap.xml | nbeam
```

That makes it easy to keep chaining results into other tools:

```bash
cat nmap.nmap | nbeam | httpx
```

```bash
cat nmap.xml | nbeam | nuclei -tags cves
```

XML works the same way as normal output:

```bash
nmap -sS -sV -T4 -p- scanme.nmap.org -oX - | nbeam
```

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
