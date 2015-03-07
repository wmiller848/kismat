//
//

package main

import (
	// Standard
	"fmt"
	"net"
	_ "time"
	// 3rd Party
	_ "code.google.com/p/gopacket"
	"code.google.com/p/gopacket/pcap"
	// Kismat
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\r\n", ifaces)
	handle, err := pcap.OpenLive("lo0", 1024, true, pcap.BlockForever)
        if err != nil {
		fmt.Println(err.Error())
		return
        }
	err = handle.SetBPFFilter("tcp")
	if err != nil {
		fmt.Println(err.Error())
		return
        }
	//packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
  	//fmt.Println(packetSource)
	//for packet := range packetSource.Packets() {
  	//	fmt.Printf("%+v\r\n", packet)
	//}
	data, info, err := handle.ReadPacketData()
	fmt.Println(data, info, err)
}
