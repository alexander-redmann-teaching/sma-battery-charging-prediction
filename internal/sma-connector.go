package internal

import (
	"encoding/binary"
	"fmt"
	"github.com/simonvetter/modbus"
	"log"
	"net"
	"time"
)

func ConnectToInverter(inverterUrl string) {
	ips, err := net.LookupIP(inverterUrl)
	inverterIp := ""

	if err != nil {
		log.Fatal("Test: ", err)
		inverterIp = inverterUrl
	} else {
		inverterIp = ips[0].String()
	}

	log.Println("Try to connect to modbus to : " + ips[0].String())

	// for a TCP endpoint
	// (see examples/tls_client.go for TLS usage and options)
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://" + inverterIp + ":502",
		Timeout: 1 * time.Second,
	})

	if err != nil {
		log.Fatal("Error on creating connection to inverter via modbus: ", err)
		return
	}

	err = client.SetUnitId(3)
	//client.SetEncoding(modbus.BIG_ENDIAN, modbus.HIGH_WORD_FIRST)

	if err != nil {
		log.Fatal("Error on set unit id: ", err)
		return
	}

	err = client.Open()

	if err != nil {
		log.Fatal("Error on open connection to inverter via modbus: ", err)
		return
	}

	susyId, err := client.ReadRawBytes(30003, 32, modbus.INPUT_REGISTER)

	if err != nil {
		log.Fatal("Error on reading data from inverter via modbus: ", err)
		return
	}
	data := binary.BigEndian.Uint32(susyId)
	fmt.Println(data)

	log.Println("result: " + fmt.Sprintf("%v", data))

	err = client.Close()
	if err != nil {
		log.Fatal("Error on closing connection to inverter via modbus: ", err)
		return
	}

}
