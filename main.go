package main

import (
	"fmt"
	"time"

	"github.com/kokizzu/gotro/L"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func main() {
	const host = `localhost`
	const port = 4222
	const caFile = `ca.pem`

	// server
	serverTlsConfig, err := server.GenTLSConfig(&server.TLSConfigOpts{
		CertFile: `server-cert.pem`,
		KeyFile:  `server-key.pem`,
		CaFile:   caFile,
		Verify:   true,
		Timeout:  2,
	})
	L.PanicIf(err, `server.GenTLSConfig`)

	srv, err := server.NewServer(&server.Options{
		Host:      host,
		Port:      port,
		TLSConfig: serverTlsConfig,
	})
	L.PanicIf(err, `server.NewServer`)

	go srv.Start()
	defer srv.Shutdown()

	time.Sleep(2 * time.Second) // wait server started

	// client
	nc, err := nats.Connect(
		fmt.Sprintf("tls://%s:%d", host, port),
		nats.RootCAs(caFile),
		nats.ClientCert(`client-cert.pem`, `client-key.pem`),
	)
	L.PanicIf(err, `nats.Connect`)
	defer nc.Drain()

	sub, err := nc.SubscribeSync("hello")
	L.PanicIf(err, `nc.SubscribeSync`)
	defer sub.Drain()

	err = nc.Publish("hello", []byte("world"))
	L.PanicIf(err, `nc.Publish`)

	msg, err := sub.NextMsg(time.Second)
	L.PanicIf(err, `sub.NextMsg`)
	fmt.Println(string(msg.Data))
}
