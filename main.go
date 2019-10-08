package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	custNamePtr := flag.String("name", "NAME NOT SET", "Customer name")
	entPtr := flag.String("ent", "ENT NOT SET", "Enterprise number")
	bwDomainPtr := flag.String("bwdomain", "BWDOMAIN NOT SET", "Broadworks domain")
	bwUserPtr := flag.String("bwuser", "BWUSER NOT SET", "Broadworks user (Pilot Number)")
	bwPassPtr := flag.String("bwpass", "BWPASS NOT SET", "Broadworks auth password")
	acDomain := flag.String("acdomain", "ACDOMAIN NOT SET", "Audiocodes SBC customer FQDN")
	sbcIP := flag.String("sbcip", "192.168.0.1", "Audiocodes SBC IP address") // replace with SBC IP for API

	flag.Parse()

	c := Customer{Name: *custNamePtr}
	c.BW.Enterprise = *entPtr
	c.BW.Domain = *bwDomainPtr
	c.BW.User = *bwUserPtr
	c.BW.Password = *bwPassPtr
	c.AC.User = "Username" // Audiocodes username
	c.AC.Password = "Password" // Audiocodes password
	c.AC.Domain = *acDomain
	c.AC.URL = fmt.Sprintf("http://%s/api/v1/files/ini", *sbcIP)

	// Character limit of 9
	if len(*custNamePtr) > 9 {
		log.Fatal("Error: name must be 9 characters or less!")
	}

	// Retrieve current INI
	i := c.INI()

	// SIPInterface
	s := make(SIPSlice, 2)
	s.build(c, i)
	fmt.Printf("%v", s)

	// ProxySet
	p := make(ProxySlice, 2)
	p.build(c, i)
	fmt.Printf("%v", p)

	// IPGroup
	g := make(IPGSlice, 2)
	g.build(c, i)
	fmt.Printf("%v", g)

	// Account
	a := &Account{}
	a.build(c, i)
	fmt.Printf("%v", a)

	// IP2IPRouting
	r := make(IP2IPSlice, 2)
	r.build(c, i)
	fmt.Printf("%v", r)

	// Classification
	f := &Classification{}
	f.build(c, i)
	fmt.Printf("%v", f)

	// IPOutboundManipulation
	n := &IPOutboundManipulation{}
	n.build(c, i)
	fmt.Printf("%v", n)

	// MessageManipulations
	m := make(MMSlice, 6)
	m.build(c, i, s, g)
	fmt.Printf("%v", m)

	//ProxyIp
	x := make(ProxyIPSlice, 2)
	x.build(sbcIP, i, p)
	fmt.Printf("%v", x)
}
