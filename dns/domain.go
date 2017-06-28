package dns

type DnsServerType struct {
	DnsServer string
}

type DomainType struct {
	DomainId    string
	DomainName  string
	AliDomain   bool
	GroupId     string
	GroupName   string
	InstanceId  string
	VersionCode string
	PunyCode    string
	DnsServers struct {
		DnsServerType []DnsServerType
	}
}

type DomainGroupType struct {
	GroupId   string
	GroupName string
}
