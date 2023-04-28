package main

import (
	"flag"
	"pet/iternal/s_synchoronizatio/web"
)

var (
	addr = flag.String("addr", ":7000", "addres server")
)

func main() {
	web.Startserver(addr)
}
