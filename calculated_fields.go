package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func setIndex(c interface{}, i *string) int {
	switch c := c.(type) {
	case SIPSlice:
		r, _ := regexp.Compile("SIPInterface [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case ProxySlice:
		r, _ := regexp.Compile("ProxySet [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case IPGSlice:
		r, _ := regexp.Compile("IPGroup [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case *Account:
		r, _ := regexp.Compile("Account [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case IP2IPSlice:
		r, _ := regexp.Compile("IP2IPRouting [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case *Classification:
		r, _ := regexp.Compile("Classification [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case *IPOutboundManipulation:
		r, _ := regexp.Compile("IPOutboundManipulation [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case MMSlice:
		r, _ := regexp.Compile("MessageManipulations [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	case ProxyIPSlice:
		r, _ := regexp.Compile("ProxyIp [0-9].*")
		d := r.FindAllString(*i, -1)
		split := strings.Split(d[len(d)-1], " ")
		last, _ := strconv.Atoi(split[1])
		return (last + 1)
	default:
		fmt.Printf("Case: Not Found (%T)", c)
		return 0
	}
}

func setPort(p PortType, i *string) int {

	r, _ := regexp.Compile("SIPInterface [0-9].*\n")
	d := r.FindAllString(*i, -1)
	ports := []int{}
	for _, v := range d {

		pos := 0
		if p == "UDP" {
			pos = 3 // position after "=" in the SIPInterface config, starting from 0
		}
		if p == "TLS" {
			pos = 5
		}

		// whitespace in split
		m := strings.Split(v, ", ")
		i, _ := strconv.Atoi(m[pos])
		ports = append(ports, i)
	}

	// Find the highest currently used port number and return the increment by 1
	max := 0
	for _, v := range ports {
		if v > max {
			max = v
		}
	}
	return max + 1
}

func setOutManID(i *string) int {

	// search for IPGroup entries in the entire ini; -1 means return all
	r, _ := regexp.Compile("IPGroup [0-9].*\n")
	d := r.FindAllString(*i, -1)

	// store the extracted manip IDs from current config
	ids := []int{}

	// loop over all regex results values (discard key)
	for _, v := range d {

		// split values by whitespace
		m := strings.Split(v, ", ")

		// position 13 is the outbound manip ID
		i, _ := strconv.Atoi(m[13])

		// add to the ids slice
		ids = append(ids, i)
	}

	// find the largest number in ids
	max := 0

	// loop over all ids, find the max
	for _, v := range ids {
		if v > max {
			max = v
		}
	}

	// return largest number + 1
	return max + 1
}
