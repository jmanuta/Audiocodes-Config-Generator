package main

import "fmt"

func (s SIPInterface) String() string {
	return fmt.Sprintf("SIPInterface %v = \"%v\", \"%v\", %v, %v, %v, %v, \"%v\", %v, \"%v\", \"%v\", \"%v\", %v, %v, %v, %v, %v, \"%v\", %v, %v, %v, %v, %v, %v, \"%v\", \"%v\", %v;",
		s.Index,
		s.InterfaceName,
		s.NetworkInterface,
		s.ApplicationType,
		s.UDPPort,
		s.TCPPort,
		s.TLSPort,
		s.AdditionalUDPPorts,
		s.AdditionalUDPPortsMode,
		s.SRDName,
		s.MessagePolicyName,
		s.TLSContext,
		s.TLSMutualAuthentication,
		s.TCPKeepAliveEnable,
		s.ClassificationFailureResponseType,
		s.PreClassificationManSet,
		s.EncapsulatingProtocol,
		s.MediaRealm,
		s.SBCDirectMedia,
		s.BlockUnRegUsers,
		s.MaxNumOfRegUsers,
		s.EnableUnAuthenticatedRegistrations,
		s.UsedByRoutingServer,
		s.TopologyLocation,
		s.PreParsingManSetName,
		s.AdmissionProfile,
		s.CallSetupRulesSetId,
	)
}

func (s SIPSlice) String() string {
	return fmt.Sprintf("[ SIPInterface ]\n\nFORMAT SIPInterface_Index = SIPInterface_InterfaceName, SIPInterface_NetworkInterface, SIPInterface_ApplicationType, SIPInterface_UDPPort, SIPInterface_TCPPort, "+
		"SIPInterface_TLSPort, SIPInterface_AdditionalUDPPorts, SIPInterface_AdditionalUDPPortsMode, SIPInterface_SRDName, SIPInterface_MessagePolicyName, SIPInterface_TLSContext, SIPInterface_TLSMutualAuthentication, "+
		"SIPInterface_TCPKeepaliveEnable, SIPInterface_ClassificationFailureResponseType, SIPInterface_PreClassificationManSet, SIPInterface_EncapsulatingProtocol, SIPInterface_MediaRealm, SIPInterface_SBCDirectMedia, "+
		"SIPInterface_BlockUnRegUsers, SIPInterface_MaxNumOfRegUsers, SIPInterface_EnableUnAuthenticatedRegistrations, SIPInterface_UsedByRoutingServer, SIPInterface_TopologyLocation, SIPInterface_PreParsingManSetName, "+
		"SIPInterface_AdmissionProfile, SIPInterface_CallSetupRulesSetId;\n%v\n%v\n\n[ \\SIPInterface ]\n\n\n",
		s[0], s[1],
	)
}

func (p ProxySet) String() string {
	return fmt.Sprintf("ProxySet %v = \"%v\", %v, %v, %v, %v, \"%v\", %v, \"%v\", %v, %v, \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", %v, %v, %v, %v;",
		p.Index,
		p.ProxyName,
		p.EnableProxyKeepAlive,
		p.ProxyKeepAliveTime,
		p.ProxyLoadBalancingMethod,
		p.IsProxyHotSwap,
		p.SRDName,
		p.ClassicicationInput,
		p.TLSContextName,
		p.ProxyRedundancyMode,
		p.DNSResolveMethod,
		p.KeepAliveFailureResp,
		p.GWIPv4SIPInterfaceName,
		p.SBCIPv4SIPInterfaceName,
		p.GWIPv6SIPInterfaceName,
		p.SBCIPv6SIPInterfaceName,
		p.MinActiveServersLB,
		p.SuccessDetectionRetries,
		p.SuccessDetectionInterval,
		p.FailureDetectionRetransmissions,
	)
}

func (p ProxySlice) String() string {
	return fmt.Sprintf("[ ProxySet ]\n\nFORMAT ProxySet_Index = ProxySet_ProxyName, ProxySet_EnableProxyKeepAlive, ProxySet_ProxyKeepAliveTime, ProxySet_ProxyLoadBalancingMethod, ProxySet_IsProxyHotSwap, ProxySet_SRDName, "+
		"ProxySet_ClassificationInput, ProxySet_TLSContextName, ProxySet_ProxyRedundancyMode, ProxySet_DNSResolveMethod, ProxySet_KeepAliveFailureResp, ProxySet_GWIPv4SIPInterfaceName, ProxySet_SBCIPv4SIPInterfaceName, "+
		"ProxySet_GWIPv6SIPInterfaceName, ProxySet_SBCIPv6SIPInterfaceName, ProxySet_MinActiveServersLB, ProxySet_SuccessDetectionRetries, ProxySet_SuccessDetectionInterval, ProxySet_FailureDetectionRetransmissions;\n"+
		"%v\n%v\n\n[ \\ProxySet ]\n\n\n",
		p[0], p[1],
	)
}

func (g IPGroup) String() string {
	return fmt.Sprintf("IPGroup %v = %v, \"%v\", \"%v\", \"%v\", \"%v\", %v, %v, \"%v\", \"%v\", %v, \"%v\", %v, %v, %v, %v, %v, \"%v\", %v, %v, %v, \"%v\", \"%v\", \"%v\", %v, \"%v\", \"%v\", %v, \"%v\", \"%v\", %v, %v, "+
		"\"%v\", %v, %v, %v, %v, %v, %v, \"%v\", %v, \"%v\", %v, %v, \"%v\", %v;",
		g.Index,
		g.Type,
		g.Name,
		g.ProxySetName,
		g.SIPGroupName,
		g.ContactUser,
		g.SipReRoutingMode,
		g.AlwaysUseRouteTable,
		g.SRDName,
		g.MediaRealm,
		g.ClassifyByProxySet,
		g.ProfileName,
		g.MaxNumOfRegUsers,
		g.InboundManSet,
		g.OutboundManSet,
		g.RegistrationMode,
		g.AuthenticationMode,
		g.MethodList,
		g.EnableSBCClientForking,
		g.SourceUriInput,
		g.DestUriInput,
		g.ContactName,
		g.Username,
		g.Password,
		g.UUIFormat,
		g.QOEProfile,
		g.BWProfile,
		g.AlwaysUseSourceAddr,
		g.MsgManUserDef1,
		g.MsgManUserDef2,
		g.SIPConnect,
		g.SBCPSAPMode,
		g.DTLSContext,
		g.CreatedByRoutingServer,
		g.UsedByRoutingServer,
		g.SBCOperationMode,
		g.SBCRouteUsingRequestURIPort,
		g.SBCKeepOriginalCallID,
		g.TopologyLocation,
		g.SBCDialPlanName,
		g.CallSetupRulesSetId, g.Tags,
		g.SBCUserStickiness,
		g.UserUDPPortAssignment,
		g.AdmissionProfile,
		g.ProxyKeepAliveUsingIPG,
	)
}

func (g IPGSlice) String() string {
	return fmt.Sprintf("[ IPGroup ]\n\nFORMAT IPGroup_Index = IPGroup_Type, IPGroup_Name, IPGroup_ProxySetName, IPGroup_SIPGroupName, IPGroup_ContactUser, IPGroup_SipReRoutingMode, IPGroup_AlwaysUseRouteTable, "+
		"IPGroup_SRDName, IPGroup_MediaRealm, IPGroup_ClassifyByProxySet, IPGroup_ProfileName, IPGroup_MaxNumOfRegUsers, IPGroup_InboundManSet, IPGroup_OutboundManSet, IPGroup_RegistrationMode, IPGroup_AuthenticationMode, "+
		"IPGroup_MethodList, IPGroup_EnableSBCClientForking, IPGroup_SourceUriInput, IPGroup_DestUriInput, IPGroup_ContactName, IPGroup_Username, IPGroup_Password, IPGroup_UUIFormat, IPGroup_QOEProfile, IPGroup_BWProfile, "+
		"IPGroup_AlwaysUseSourceAddr, IPGroup_MsgManUserDef1, IPGroup_MsgManUserDef2, IPGroup_SIPConnect, IPGroup_SBCPSAPMode, IPGroup_DTLSContext, IPGroup_CreatedByRoutingServer, IPGroup_UsedByRoutingServer, "+
		"IPGroup_SBCOperationMode, IPGroup_SBCRouteUsingRequestURIPort, IPGroup_SBCKeepOriginalCallID, IPGroup_TopologyLocation, IPGroup_SBCDialPlanName, IPGroup_CallSetupRulesSetId, IPGroup_Tags, IPGroup_SBCUserStickiness, "+
		"IPGroup_UserUDPPortAssignment, IPGroup_AdmissionProfile, IPGroup_ProxyKeepAliveUsingIPG;\n%v\n%v\n\n[ \\IPGroup ]\n\n\n",
		g[0], g[1],
	)
}

func (a Account) String() string {
	return fmt.Sprintf("[ Account ]\n\nFORMAT Account_Index = Account_ServedTrunkGroup, Account_ServedIPGroupName, Account_ServingIPGroupName, Account_Username, Account_Password, Account_HostName, Account_ContactUser, "+
		"Account_Register, Account_RegistrarStickiness, Account_RegistrarSearchMode, Account_RegEventPackageSubscription, Account_ApplicationType, Account_RegByServedIPG, Account_UDPPortAssignment;\nAccount %v = %v, \"%v\", "+
		"\"%v\", \"%v\", \"%v\", \"%v\", \"%v\", %v, %v, %v, %v, %v, %v, %v;\n\n[ \\Account ]\n\n\n",
		a.Index,
		a.ServedTrunkGroup,
		a.ServedIPGroupName,
		a.ServingIPGroupName,
		a.Username,
		a.Password,
		a.HostName,
		a.ContactUser,
		a.Register,
		a.RegistrarStickiness,
		a.RegistrarSearchMode,
		a.RegEventPackageSubscription,
		a.ApplicationType,
		a.RegByServedIPG,
		a.UDPPortAssignment,
	)
}

func (r IP2IPRouting) String() string {
	return fmt.Sprintf("IP2IPRouting %v = \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", %v, \"%v\", \"%v\", %v, %v, %v, \"%v\", \"%v\", \"%v\", %v, %v, %v, %v, \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\";",
		r.Index,
		r.RouteName,
		r.RoutingPolicyName,
		r.SrcIPGroupName,
		r.SrcUsernamePrefix,
		r.SrcHost,
		r.DestUsernamePrefix,
		r.DestHost,
		r.RequestType,
		r.MessageConditionName,
		r.ReRouteIPGroupName,
		r.Trigger,
		r.CallSetupRulesSetId,
		r.DestType,
		r.DestIPGroupName,
		r.DestSIPInterfaceName,
		r.DestAddress,
		r.DestPort,
		r.DestTransportType,
		r.AltRouteOptions,
		r.GroupPolicy,
		r.CostGroup,
		r.DestTags,
		r.SrcTags,
		r.IPGroupSetName,
		r.RoutingTagName,
		r.InternalAction,
	)
}

func (g IP2IPSlice) String() string {
	return fmt.Sprintf("[ IP2IPRouting ]\n\nFORMAT IP2IPRouting_Index = IP2IPRouting_RouteName, IP2IPRouting_RoutingPolicyName, IP2IPRouting_SrcIPGroupName, IP2IPRouting_SrcUsernamePrefix, IP2IPRouting_SrcHost, "+
		"IP2IPRouting_DestUsernamePrefix, IP2IPRouting_DestHost, IP2IPRouting_RequestType, IP2IPRouting_MessageConditionName, IP2IPRouting_ReRouteIPGroupName, IP2IPRouting_Trigger, IP2IPRouting_CallSetupRulesSetId, "+
		"IP2IPRouting_DestType, IP2IPRouting_DestIPGroupName, IP2IPRouting_DestSIPInterfaceName, IP2IPRouting_DestAddress, IP2IPRouting_DestPort, IP2IPRouting_DestTransportType, IP2IPRouting_AltRouteOptions, "+
		"IP2IPRouting_GroupPolicy, IP2IPRouting_CostGroup, IP2IPRouting_DestTags, IP2IPRouting_SrcTags, IP2IPRouting_IPGroupSetName, IP2IPRouting_RoutingTagName, IP2IPRouting_InternalAction;\n%v\n%v\n\n[ \\IP2IPRouting ]\n\n\n",
		g[0], g[1],
	)
}

func (f Classification) String() string {
	return fmt.Sprintf("[ Classification ]\n\nFORMAT Classification_Index = Classification_ClassificationName, Classification_MessageConditionName, Classification_SRDName, Classification_SrcSIPInterfaceName, "+
		"Classification_SrcAddress, Classification_SrcPort, Classification_SrcTransportType, Classification_SrcUsernamePrefix, Classification_SrcHost, Classification_DestUsernamePrefix, Classification_DestHost, "+
		"Classification_ActionType, Classification_SrcIPGroupName, Classification_DestRoutingPolicy, Classification_IpProfileName, Classification_IPGroupSelection, Classification_IpGroupTagName;\nClassification %v "+
		"= \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", %v, %v, \"%v\", \"%v\", \"%v\", \"%v\", %v, \"%v\", \"%v\", \"%v\", %v, \"%v\";\n\n[ \\Classification ]\n\n\n",
		f.Index,
		f.ClassificationName,
		f.MessageConditionName,
		f.SRDName,
		f.SrcSIPInterfaceName,
		f.SrcAddress,
		f.SrcPort,
		f.SrcTransportType,
		f.SrcUsernamePrefix,
		f.SrcHost,
		f.DestUsernamePrefix,
		f.DestHost,
		f.ActionType,
		f.SrcIPGroupName,
		f.DestRoutingPolicy,
		f.IpProfileName,
		f.IPGroupSelection,
		f.IpGroupTagName,
	)
}

func (n IPOutboundManipulation) String() string {
	return fmt.Sprintf("[ IPOutboundManipulation ]\n\nFORMAT IPOutboundManipulation_Index = IPOutboundManipulation_ManipulationName, IPOutboundManipulation_RoutingPolicyName, IPOutboundManipulation_IsAdditionalManipulation, "+
		"IPOutboundManipulation_SrcIPGroupName, IPOutboundManipulation_DestIPGroupName, IPOutboundManipulation_SrcUsernamePrefix, IPOutboundManipulation_SrcHost, IPOutboundManipulation_DestUsernamePrefix, "+
		"IPOutboundManipulation_DestHost, IPOutboundManipulation_CallingNamePrefix, IPOutboundManipulation_MessageConditionName, IPOutboundManipulation_RequestType, IPOutboundManipulation_ReRouteIPGroupName, "+
		"IPOutboundManipulation_Trigger, IPOutboundManipulation_ManipulatedURI, IPOutboundManipulation_RemoveFromLeft, IPOutboundManipulation_RemoveFromRight, IPOutboundManipulation_LeaveFromRight, "+
		"IPOutboundManipulation_Prefix2Add, IPOutboundManipulation_Suffix2Add, IPOutboundManipulation_PrivacyRestrictionMode, IPOutboundManipulation_DestTags, IPOutboundManipulation_SrcTags;\nIPOutboundManipulation %v = "+
		"\"%v\", \"%v\", %v, \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", \"%v\", %v, \"%v\", %v, %v, %v, %v, %v, \"%v\", \"%v\", %v, \"%v\", \"%v\";\n\n[ \\IPOutboundManipulation ]\n\n\n",
		n.Index,
		n.ManipulationName,
		n.RoutingPolicyName,
		n.IsAdditionalManipulation,
		n.SrcIPGroupName,
		n.DestIPGroupName,
		n.SrcUsernamePrefix,
		n.SrcHost,
		n.DestUsernamePrefix,
		n.DestHost,
		n.CallingNamePrefix,
		n.MessageConditionName,
		n.RequestType,
		n.ReRouteIPGroupName,
		n.Trigger,
		n.ManipulatedURI,
		n.RemoveFromLeft,
		n.RemoveFromRight,
		n.LeaveFromRight,
		n.Prefix2Add,
		n.Suffix2Add,
		n.PrivacyRestrictionMode,
		n.DestTags,
		n.SrcTags,
	)
}

func (m MessageManipulations) String() string {
	return fmt.Sprintf("MessageManipulations %v = \"%v\", %v, \"%v\", \"%v\", \"%v\", %v, \"%v\", %v;",
		m.Index,
		m.ManipulationName,
		m.ManSetID,
		m.MessageType,
		m.Condition,
		m.ActionSubject,
		m.ActionType,
		m.ActionValue,
		m.RowRole,
	)
}

func (m MMSlice) String() string {
	return fmt.Sprintf("[ MessageManipulations ]\n\nFORMAT MessageManipulations_Index = MessageManipulations_ManipulationName, MessageManipulations_ManSetID, MessageManipulations_MessageType, MessageManipulations_Condition, "+
		"MessageManipulations_ActionSubject, MessageManipulations_ActionType, MessageManipulations_ActionValue, MessageManipulations_RowRole;\n%v\n%v\n%v\n%v\n\n[ \\MessageManipulations ]\n\n\n",
		m[0], m[1], m[2], m[3],
	)
}

func (p ProxyIP) String() string {
	return fmt.Sprintf("ProxyIp %v = \"%v\", %v, \"%v\", %v, %v, %v;",
		p.Index,
		p.ProxySetId,
		p.ProxyIpIndex,
		p.IpAddress,
		p.TransportType,
		p.Priority,
		p.Weight,
	)
}

func (p ProxyIPSlice) String() string {
	return fmt.Sprintf("[ ProxyIp ]\n\nFORMAT ProxyIp_Index = ProxyIp_ProxySetId, ProxyIp_ProxyIpIndex, ProxyIp_IpAddress, ProxyIp_TransportType, ProxyIp_Priority, ProxyIp_Weight;\n%v\n%v\n\n[ \\ProxyIp ]\n\n\n",
		p[0], p[1],
	)
}
