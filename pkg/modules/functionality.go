package modules

import (
	"github.com/ropnop/go-windapsearch/pkg/ldapsession"
	"github.com/spf13/pflag"
	"gopkg.in/ldap.v3"
)

type FunctionalityModule struct{}

func init() {
	AllModules = append(AllModules, new(FunctionalityModule))
}

func (FunctionalityModule) Name() string {
	return "functionality"
}

func (FunctionalityModule) Description() string {
	return "Print domain functionality level"
}

func (FunctionalityModule) FlagSet() *pflag.FlagSet {
	return pflag.NewFlagSet("functionality", pflag.ExitOnError)
}

func (FunctionalityModule) DefaultAttrs() []string {
	return []string{
		"defaultNamingContext",
		"domainFunctionality",
		"forestFunctionality",
		"domainControllerFunctionality",
		"dnsHostName",
	}
}

func (FunctionalityModule) Run(session *ldapsession.LDAPSession, attrs []string) (*ldap.SearchResult, error) {
	sr := ldap.NewSearchRequest(
		"",
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0, 0, false,
		"(objectClass=*)",
		attrs,
		nil)
	return session.LConn.Search(sr)
}
