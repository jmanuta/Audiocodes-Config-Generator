package main

// Customer contains BW and AC details
type Customer struct {
	Name string
	BW   struct {
		Enterprise string
		Domain     string
		User       string
		Password   string
	}
	AC struct {
		Domain   string
		User     string
		Password string
		URL      string
	}
}

// PortType for UDP or TLS
type PortType string

// SIPSlice is a collection of SIPInterfaces
type SIPSlice []SIPInterface

// SIPInterface for inside and outside interfaces
type SIPInterface struct {
	Index                              int
	InterfaceName                      string
	NetworkInterface                   string
	ApplicationType                    int
	UDPPort                            int
	TCPPort                            int
	TLSPort                            int
	AdditionalUDPPorts                 string
	AdditionalUDPPortsMode             int
	SRDName                            string
	MessagePolicyName                  string
	TLSContext                         string
	TLSMutualAuthentication            int
	TCPKeepAliveEnable                 int
	ClassificationFailureResponseType  int
	PreClassificationManSet            int
	EncapsulatingProtocol              int
	MediaRealm                         string
	SBCDirectMedia                     int
	BlockUnRegUsers                    int
	MaxNumOfRegUsers                   int
	EnableUnAuthenticatedRegistrations int
	UsedByRoutingServer                int
	TopologyLocation                   int
	PreParsingManSetName               string
	AdmissionProfile                   string
	CallSetupRulesSetId                int
}

// ProxySlice is a collection of ProxySets
type ProxySlice []ProxySet

// ProxySet for inside and outside
type ProxySet struct {
	Index                           int
	ProxyName                       string
	EnableProxyKeepAlive            int
	ProxyKeepAliveTime              int
	ProxyLoadBalancingMethod        int
	IsProxyHotSwap                  int
	SRDName                         string
	ClassicicationInput             int
	TLSContextName                  string
	ProxyRedundancyMode             int
	DNSResolveMethod                int
	KeepAliveFailureResp            string
	GWIPv4SIPInterfaceName          string
	SBCIPv4SIPInterfaceName         string
	GWIPv6SIPInterfaceName          string
	SBCIPv6SIPInterfaceName         string
	MinActiveServersLB              int
	SuccessDetectionRetries         int
	SuccessDetectionInterval        int
	FailureDetectionRetransmissions int
}

// IPGSlice is a collection of IPGroups
type IPGSlice []IPGroup

// IPGroup settings
type IPGroup struct {
	Index                       int
	Type                        int
	Name                        string
	ProxySetName                string
	SIPGroupName                string
	ContactUser                 string
	SipReRoutingMode            int
	AlwaysUseRouteTable         int
	SRDName                     string
	MediaRealm                  string
	ClassifyByProxySet          int
	ProfileName                 string
	MaxNumOfRegUsers            int
	InboundManSet               int
	OutboundManSet              int
	RegistrationMode            int
	AuthenticationMode          int
	MethodList                  string
	EnableSBCClientForking      int
	SourceUriInput              int
	DestUriInput                int
	ContactName                 string
	Username                    string
	Password                    string
	UUIFormat                   int
	QOEProfile                  string
	BWProfile                   string
	AlwaysUseSourceAddr         int
	MsgManUserDef1              string
	MsgManUserDef2              string
	SIPConnect                  int
	SBCPSAPMode                 int
	DTLSContext                 string
	CreatedByRoutingServer      int
	UsedByRoutingServer         int
	SBCOperationMode            int
	SBCRouteUsingRequestURIPort int
	SBCKeepOriginalCallID       int
	TopologyLocation            int
	SBCDialPlanName             string
	CallSetupRulesSetId         int
	Tags                        string
	SBCUserStickiness           int
	UserUDPPortAssignment       int
	AdmissionProfile            string
	ProxyKeepAliveUsingIPG      int
}

// Account settings
type Account struct {
	Index                       int
	ServedTrunkGroup            int
	ServedIPGroupName           string
	ServingIPGroupName          string
	Username                    string
	Password                    string
	HostName                    string
	ContactUser                 string
	Register                    int
	RegistrarStickiness         int
	RegistrarSearchMode         int
	RegEventPackageSubscription int
	ApplicationType             int
	RegByServedIPG              int
	UDPPortAssignment           int
}

// IP2IPSlice is a collection of IP2IPRoutings
type IP2IPSlice []IP2IPRouting

// IP2IPRouting settings
type IP2IPRouting struct {
	Index                int
	RouteName            string
	RoutingPolicyName    string
	SrcIPGroupName       string
	SrcUsernamePrefix    string
	SrcHost              string
	DestUsernamePrefix   string
	DestHost             string
	RequestType          int
	MessageConditionName string
	ReRouteIPGroupName   string
	Trigger              int
	CallSetupRulesSetId  int
	DestType             int
	DestIPGroupName      string
	DestSIPInterfaceName string
	DestAddress          string
	DestPort             int
	DestTransportType    int
	AltRouteOptions      int
	GroupPolicy          int
	CostGroup            string
	DestTags             string
	SrcTags              string
	IPGroupSetName       string
	RoutingTagName       string
	InternalAction       string
}

// Classification settings
type Classification struct {
	Index                int
	ClassificationName   string
	MessageConditionName string
	SRDName              string
	SrcSIPInterfaceName  string
	SrcAddress           string
	SrcPort              int
	SrcTransportType     int
	SrcUsernamePrefix    string
	SrcHost              string
	DestUsernamePrefix   string
	DestHost             string
	ActionType           int
	SrcIPGroupName       string
	DestRoutingPolicy    string
	IpProfileName        string
	IPGroupSelection     int
	IpGroupTagName       string
}

// IPOutboundManipulation settings
type IPOutboundManipulation struct {
	Index                    int
	ManipulationName         string
	RoutingPolicyName        string
	IsAdditionalManipulation int
	SrcIPGroupName           string
	DestIPGroupName          string
	SrcUsernamePrefix        string
	SrcHost                  string
	DestUsernamePrefix       string
	DestHost                 string
	CallingNamePrefix        string
	MessageConditionName     string
	RequestType              int
	ReRouteIPGroupName       string
	Trigger                  int
	ManipulatedURI           int
	RemoveFromLeft           int
	RemoveFromRight          int
	LeaveFromRight           int
	Prefix2Add               string
	Suffix2Add               string
	PrivacyRestrictionMode   int
	DestTags                 string
	SrcTags                  string
}

// MMSlice is a collection of MessageManipulations
type MMSlice []MessageManipulations

// MessageManipulations settings
type MessageManipulations struct {
	Index            int
	ManipulationName string
	ManSetID         int
	MessageType      string
	Condition        string
	ActionSubject    string
	ActionType       int
	ActionValue      string
	RowRole          int
}

// ProxyIPSlice is a collection of ProxyIPs
type ProxyIPSlice []ProxyIP

// ProxyIP settings
type ProxyIP struct {
	Index         int
	ProxySetId    string
	ProxyIpIndex  int
	IpAddress     string
	TransportType int
	Priority      int
	Weight        int
}
