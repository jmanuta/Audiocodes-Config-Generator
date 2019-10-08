package main

import "strconv"

func (s SIPSlice) build(c Customer, i *string) {

	// Inside
	s[0].Index = setIndex(s, i) //parse current index value from live ini file; increment
	s[0].InterfaceName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-SIPif"
	s[0].NetworkInterface = "InsideIP"
	s[0].UDPPort = setPort("UDP", i) // parse current port number from live ini file; increment
	s[0].TLSContext = "default"
	s[0].TLSMutualAuthentication = -1
	s[0].MediaRealm = "MR-Inside"
	s[0].ApplicationType = 2
	s[0].SRDName = "DefaultSRD"
	s[0].ClassificationFailureResponseType = 500
	s[0].PreClassificationManSet = -1
	s[0].BlockUnRegUsers = -1
	s[0].MaxNumOfRegUsers = -1
	s[0].EnableUnAuthenticatedRegistrations = -1
	s[0].CallSetupRulesSetId = -1

	// Outside
	s[1].Index = s[0].Index + 1
	s[1].InterfaceName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-SIPif"
	s[1].NetworkInterface = "OutsideIP"
	s[1].TLSPort = setPort("TLS", i)
	s[1].TLSContext = "TeamsCerts"
	s[1].TCPKeepAliveEnable = 1
	s[1].MediaRealm = "MR-Outside"
	s[1].ApplicationType = 2
	s[1].SRDName = "DefaultSRD"
	s[1].ClassificationFailureResponseType = 500
	s[1].PreClassificationManSet = -1
	s[1].BlockUnRegUsers = -1
	s[1].MaxNumOfRegUsers = -1
	s[1].EnableUnAuthenticatedRegistrations = -1
	s[1].CallSetupRulesSetId = -1

}

func (p ProxySlice) build(c Customer, i *string) {

	// Inside
	p[0].Index = setIndex(p, i)
	p[0].ProxyName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-PS"
	p[0].ProxyKeepAliveTime = 60
	p[0].SRDName = "DefaultSRD"
	p[0].ProxyRedundancyMode = -1
	p[0].DNSResolveMethod = 1
	p[0].SBCIPv4SIPInterfaceName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-SIPif"
	p[0].MinActiveServersLB = 1
	p[0].SuccessDetectionRetries = 1
	p[0].SuccessDetectionInterval = 10
	p[0].FailureDetectionRetransmissions = -1

	// Outside
	p[1].Index = p[0].Index + 1
	p[1].ProxyName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-PS"
	p[1].EnableProxyKeepAlive = 1
	p[1].ProxyKeepAliveTime = 15
	p[1].ProxyLoadBalancingMethod = 2
	p[1].IsProxyHotSwap = 1
	p[1].SRDName = "DefaultSRD"
	p[1].TLSContextName = "TeamCerts"
	p[1].ProxyRedundancyMode = -1
	p[1].DNSResolveMethod = 1
	p[1].SBCIPv4SIPInterfaceName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-SIPif"
	p[1].MinActiveServersLB = 1
	p[1].SuccessDetectionRetries = 1
	p[1].SuccessDetectionInterval = 10
	p[1].FailureDetectionRetransmissions = -1

}

func (g IPGSlice) build(c Customer, i *string) {

	// Inside
	g[0].Index = setIndex(g, i)
	g[0].Name = "BWE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	g[0].ProxySetName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-PS"
	g[0].SIPGroupName = c.BW.Domain
	g[0].SipReRoutingMode = -1
	g[0].SRDName = "DefaultSRD"
	g[0].MediaRealm = "MR-Inside"
	g[0].ClassifyByProxySet = 1
	g[0].ProfileName = "IPP-Inside"
	g[0].MaxNumOfRegUsers = -1
	g[0].InboundManSet = -1
	g[0].OutboundManSet = setOutManID(i) // This needs to be calculated
	g[0].SourceUriInput = -1
	g[0].DestUriInput = -1
	g[0].Password = "$1$gQ=="
	g[0].DTLSContext = "default"
	g[0].SBCOperationMode = -1
	g[0].CallSetupRulesSetId = -1

	// Outside
	g[1].Index = g[0].Index + 1
	g[1].Name = "MSE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	g[1].ProxySetName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-PS"
	g[1].SIPGroupName = c.AC.Domain
	g[1].SipReRoutingMode = -1
	g[1].SRDName = "DefaultSRD"
	g[1].MediaRealm = "MR-Outside"
	g[1].ProfileName = "IPP-Outside"
	g[1].MaxNumOfRegUsers = -1
	g[1].InboundManSet = -1
	g[1].OutboundManSet = -1
	g[1].SourceUriInput = -1
	g[1].DestUriInput = -1
	g[1].ContactName = c.AC.Domain
	g[1].Password = "$1$gQ=="
	g[1].DTLSContext = "TeamsCerts"
	g[1].SBCOperationMode = -1
	g[1].TopologyLocation = 1
	g[1].CallSetupRulesSetId = -1

}

func (a *Account) build(c Customer, i *string) {

	// BW SIP Account
	a.Index = setIndex(a, i)
	a.ServedTrunkGroup = -1
	a.ServedIPGroupName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	a.ServingIPGroupName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	a.Username = c.BW.User
	a.Password = c.BW.Password
	a.HostName = c.BW.Domain
	a.ContactUser = c.BW.User
	a.Register = 1
	a.RegistrarStickiness = 0
	a.RegistrarSearchMode = 0
	a.RegEventPackageSubscription = 0
	a.ApplicationType = 2
	a.RegByServedIPG = 0
	a.UDPPortAssignment = 0

}

func (r IP2IPSlice) build(c Customer, i *string) {

	// Inside to Outside
	r[0].Index = setIndex(r, i)
	r[0].RouteName = "BWE" + c.BW.Enterprise + "-to-" + "MSE" + c.BW.Enterprise + "-" + c.Name + "-IPRouting"
	r[0].RoutingPolicyName = "Default_SBCRoutingPolicy"
	r[0].SrcIPGroupName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	r[0].SrcUsernamePrefix = "*"
	r[0].SrcHost = "*"
	r[0].DestUsernamePrefix = "*"
	r[0].DestHost = "*"
	r[0].RequestType = 0
	r[0].MessageConditionName = ""
	r[0].ReRouteIPGroupName = "Any"
	r[0].Trigger = 0
	r[0].CallSetupRulesSetId = -1
	r[0].DestType = 0
	r[0].DestIPGroupName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	r[0].DestSIPInterfaceName = ""
	r[0].DestAddress = ""
	r[0].DestPort = 0
	r[0].DestTransportType = -1
	r[0].AltRouteOptions = 0
	r[0].GroupPolicy = 0
	r[0].CostGroup = ""
	r[0].DestTags = ""
	r[0].SrcTags = ""
	r[0].IPGroupSetName = ""
	r[0].RoutingTagName = "default"
	r[0].InternalAction = ""

	// Outside to Inside
	r[1].Index = r[0].Index + 1
	r[1].RouteName = "MSE" + c.BW.Enterprise + "-to-" + "BWE" + c.BW.Enterprise + "-" + c.Name + "-IPRouting"
	r[1].RoutingPolicyName = "Default_SBCRoutingPolicy"
	r[1].SrcIPGroupName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	r[1].SrcUsernamePrefix = "*"
	r[1].SrcHost = "*"
	r[1].DestUsernamePrefix = "*"
	r[1].DestHost = "*"
	r[1].RequestType = 0
	r[1].MessageConditionName = ""
	r[1].ReRouteIPGroupName = "Any"
	r[1].Trigger = 0
	r[1].CallSetupRulesSetId = -1
	r[1].DestType = 0
	r[1].DestIPGroupName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	r[1].DestSIPInterfaceName = ""
	r[1].DestAddress = ""
	r[1].DestPort = 0
	r[1].DestTransportType = -1
	r[1].AltRouteOptions = 0
	r[1].GroupPolicy = 0
	r[1].CostGroup = ""
	r[1].DestTags = ""
	r[1].SrcTags = ""
	r[1].IPGroupSetName = ""
	r[1].RoutingTagName = "default"
	r[1].InternalAction = ""

}

func (f *Classification) build(c Customer, i *string) {

	// Classification Outside to Inside
	f.Index = setIndex(f, i)
	f.ClassificationName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-Classification"
	f.MessageConditionName = "Teams Contact"
	f.SRDName = "DefaultSRD"
	f.SrcSIPInterfaceName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-SIPif"
	f.SrcAddress = ""
	f.SrcPort = 0
	f.SrcTransportType = -1
	f.SrcUsernamePrefix = "*"
	f.SrcHost = "*"
	f.DestUsernamePrefix = "*"
	f.DestHost = c.AC.Domain
	f.ActionType = 1
	f.SrcIPGroupName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	f.DestRoutingPolicy = ""
	f.IpProfileName = ""
	f.IPGroupSelection = 0
	f.IpGroupTagName = "default"
}

func (n *IPOutboundManipulation) build(c Customer, i *string) {

	// Add +1 from Inside to Outside
	n.Index = setIndex(n, i)
	n.ManipulationName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-OM Add +1"
	n.RoutingPolicyName = "Default_SBCRoutingPolicy"
	n.IsAdditionalManipulation = 0
	n.SrcIPGroupName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	n.DestIPGroupName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-IPG"
	n.SrcUsernamePrefix = "*"
	n.SrcHost = "*"
	n.DestUsernamePrefix = "*"
	n.DestHost = "*"
	n.CallingNamePrefix = "*"
	n.MessageConditionName = ""
	n.RequestType = 0
	n.ReRouteIPGroupName = "Any"
	n.Trigger = 0
	n.ManipulatedURI = 1
	n.RemoveFromLeft = 0
	n.RemoveFromRight = 0
	n.LeaveFromRight = 255
	n.Prefix2Add = "+1"
	n.Suffix2Add = ""
	n.PrivacyRestrictionMode = 0
	n.DestTags = ""
	n.SrcTags = ""
}

func (m MMSlice) build(c Customer, i *string, s SIPSlice, g IPGSlice) {

	// Contact URL Host
	m[0].Index = setIndex(m, i)
	m[0].ManipulationName = "MSE" + c.BW.Enterprise + "-" + c.Name + "-MM Contact URL Host"
	m[0].ManSetID = 1
	m[0].MessageType = "Options"
	m[0].Condition = "param.Message.Address.dst.SIPInterface=='" + strconv.Itoa(s[1].Index) + "'"
	m[0].ActionSubject = "Header.Contact.URL.Host"
	m[0].ActionType = 2
	m[0].ActionValue = "'" + c.AC.Domain + "'"
	m[0].RowRole = 0

	// Add PAI
	m[1].Index = m[0].Index + 1
	m[1].ManipulationName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-MM Add PAI"
	m[1].ManSetID = g[0].OutboundManSet
	m[1].MessageType = "Invite"
	m[1].Condition = ""
	m[1].ActionSubject = "Header.P-Asserted-Identity"
	m[1].ActionType = 0
	m[1].ActionValue = "'sip:" + c.BW.User + "@" + c.BW.Domain + "'"
	m[1].RowRole = 0

	// Modify Refer-to
	m[2].Index = m[0].Index + 2
	m[2].ManipulationName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-MM Modify Refer-to"
	m[2].ManSetID = g[0].OutboundManSet
	m[2].MessageType = "Refer"
	m[2].Condition = ""
	m[2].ActionSubject = "Header.Refer-To.URL.Host"
	m[2].ActionType = 2
	m[2].ActionValue = "'" + c.BW.Domain + "'"
	m[2].RowRole = 0

	// Modify Host URL
	m[3].Index = m[0].Index + 3
	m[3].ManipulationName = "BWE" + c.BW.Enterprise + "-" + c.Name + "-MM Modify Host URL"
	m[3].ManSetID = g[0].OutboundManSet
	m[3].MessageType = "Invite"
	m[3].Condition = ""
	m[3].ActionSubject = "Header.From.URL.Host"
	m[3].ActionType = 2
	m[3].ActionValue = "'" + c.BW.Domain + "'"
	m[3].RowRole = 0

}

func (p ProxyIPSlice) build(sbc *string, i *string, ps ProxySlice) {

	// Inside
	p[0].Index = setIndex(p, i)
	p[0].ProxySetId = strconv.Itoa(ps[0].Index)
	p[0].ProxyIpIndex = 0
	p[0].IpAddress = "lax-iad3.masteraccess.com"
	if *sbc == "199.168.181.189" {
		p[0].IpAddress = "ord-iad3.masteraccess.com"
	}

	p[0].TransportType = 0
	p[0].Priority = 0
	p[0].Weight = 0

	// Outside
	p[1].Index = p[0].Index + 1
	p[1].ProxySetId = strconv.Itoa(ps[1].Index)
	p[1].ProxyIpIndex = 0
	p[1].IpAddress = "teams.local"
	p[1].TransportType = 2
	p[1].Priority = 0
	p[1].Weight = 0
}
