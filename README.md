### What
Local development DNS server that replies "127.0.0.1" to all type A queries, and NXDOMAIN to any other query.

### Why
It's often useful during development to access local services using a local domain. Existing options are:

1. Add them all to `/etc/hosts` (quickly becomes a mess, have to list all subdomains)
2. Run a DNS server like BIND (complex configuration)
3. Run a DNS proxy like [Dnsmasq](http://passingcuriosity.com/2013/dnsmasq-dev-osx/) (reasonable option but still needs configuration)

Using devdns you just need to download a binary and run it. It works best with the OS X `resolver` system (see below).

### How

Build (or [download the OS X binary](https://github.com/robbiev/devdns/releases/download/v1.0/devdns)) and then run `./devdns`. By default it listens on `127.0.0.1:5300` (UDP), you can specify an alternative address as follows: `./devdns -addr="127.0.0.1:6300"`.

On OS X you can use the [resolver system](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man5/resolver.5.html) (`man 5 resolver`) to resolve only a chosen few domains to this local server:

```
sudo mkdir -p /etc/resolver

# all domains ending in ".dev"
sudo vi /etc/resolver/dev
```

Contents of /etc/resolver/dev:

```
nameserver 127.0.0.1
port 5300
```

### Building

Build using the standard go tools:

```
go get .
go build
```

Use `go build -ldflags "-w"` to build a version without debug symbols (smaller binary).
