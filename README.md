
# Example how to use mTLS with NATS

1. Install `mkcert`
```shell
sudo apt install libnss3-tools certutil
curl -JLO "https://dl.filippo.io/mkcert/latest?for=linux/amd64"
chmod +x mkcert-v*-linux-amd64
sudo mv mkcert-v*-linux-amd64 ~/go/bin/mkcert # or /usr/local/bin/mkcert
```

2. generate both certs
```
mkcert \
  -key-file server-key.pem \
  -cert-file server-cert.pem \
  localhost

mkcert \
  -client \
  -key-file client-key.pem \
  -cert-file client-cert.pem \
  localhost

sudo cp ~/.local/share/mkcert/rootCA.pem  ca.pem
```

3. run example:
```
go rum main.go
```