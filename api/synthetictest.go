package api

import (
	"encoding/json"
	"fmt"
)

// ResourcePath is the path to the Synthetic Tests resource in the Instana API
const SyntheticTestResourcePath = "/api/synthetics/settings/tests"

// ApiTag represents an RBAC tag
type ApiTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// MultipleScriptsConfiguration for Jest-based scripts
type MultipleScriptsConfiguration struct {
	Bundle     *string `json:"bundle,omitempty"`
	ScriptFile *string `json:"scriptFile,omitempty"`
}

// DNSFilterQueryTime represents DNS query time filter
type DNSFilterQueryTime struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    int64  `json:"value"`
}

// DNSFilterTargetValue represents DNS target value filter
type DNSFilterTargetValue struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// SSLCertificateValidation represents SSL certificate validation rule
type SSLCertificateValidation struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// ICMPValidation represents a validation rule for ICMP synthetic tests.
// The Operator must be one of: CONTAINS, EQUALS, GREATER_THAN,
// GREATER_THAN_OR_EQUALS, IS, LESS_THAN, LESS_THAN_OR_EQUALS, MATCHES,
// NOT_MATCHES.
type ICMPValidation struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    int64  `json:"value"`
}

// SyntheticTypeICMP is the syntheticType value for ICMP ping tests.
const SyntheticTypeICMP = "ICMP"

// SyntheticTypeSSLCertificate is the syntheticType value for SSL certificate tests.
const SyntheticTypeSSLCertificate = "SSLCertificate"

type SyntheticTestConfig struct {
	MarkSyntheticCall bool    `json:"markSyntheticCall"`
	Retries           int32   `json:"retries,omitempty"`
	RetryInterval     int32   `json:"retryInterval,omitempty"`
	SyntheticType     string  `json:"syntheticType"`
	Timeout           *string `json:"timeout,omitempty"`

	// HttpAction fields
	URL              *string           `json:"url,omitempty"`
	Operation        *string           `json:"operation,omitempty"`
	Headers          map[string]string `json:"headers,omitempty"`
	Body             *string           `json:"body,omitempty"`
	ValidationString *string           `json:"validationString,omitempty"`
	FollowRedirect   *bool             `json:"followRedirect,omitempty"`
	AllowInsecure    *bool             `json:"allowInsecure,omitempty"`
	ExpectStatus     *int32            `json:"expectStatus,omitempty"`
	ExpectMatch      *string           `json:"expectMatch,omitempty"`
	ExpectExists     []string          `json:"expectExists,omitempty"`
	ExpectNotEmpty   []string          `json:"expectNotEmpty,omitempty"`
	ExpectJson       json.RawMessage   `json:"expectJson,omitempty"`

	// HttpScript fields
	Script     *string                       `json:"script,omitempty"`
	ScriptType *string                       `json:"scriptType,omitempty"`
	FileName   *string                       `json:"fileName,omitempty"`
	Scripts    *MultipleScriptsConfiguration `json:"scripts,omitempty"`

	// BrowserScript fields (shares Script, ScriptType, FileName, Scripts from HttpScript)
	Browser     *string `json:"browser,omitempty"`
	RecordVideo *bool   `json:"recordVideo,omitempty"`

	// DNS fields
	Lookup           *string                `json:"lookup,omitempty"`
	Server           *string                `json:"server,omitempty"`
	QueryType        *string                `json:"queryType,omitempty"`
	Port             *int32                 `json:"port,omitempty"`
	Transport        *string                `json:"transport,omitempty"`
	AcceptCNAME      *bool                  `json:"acceptCNAME,omitempty"`
	LookupServerName *bool                  `json:"lookupServerName,omitempty"`
	RecursiveLookups *bool                  `json:"recursiveLookups,omitempty"`
	ServerRetries    *int32                 `json:"serverRetries,omitempty"`
	QueryTime        *DNSFilterQueryTime    `json:"queryTime,omitempty"`
	TargetValues     []DNSFilterTargetValue `json:"targetValues,omitempty"`

	// SSLCertificate fields
	Hostname             *string                    `json:"hostname,omitempty"`
	DaysRemainingCheck   *int32                     `json:"daysRemainingCheck,omitempty"`
	SSLPort              *int32                     `json:"sslPort,omitempty"`
	AcceptSelfSignedCert *bool                      `json:"acceptSelfSignedCertificate,omitempty"`
	// ValidationRules holds SSL certificate validation rules (syntheticType=SSLCertificate).
	// For ICMP validation rules see ICMPValidationRules.
	ValidationRules []SSLCertificateValidation `json:"-"`

	// WebpageAction fields (shares URL from HttpAction, Browser and RecordVideo from BrowserScript)
	// No additional fields needed

	// WebpageScript fields (shares Script, Browser, RecordVideo, FileName from BrowserScript)
	// No additional fields needed

	// ICMP fields
	PacketCount    *int32  `json:"packetCount,omitempty"`
	PacketInterval *string `json:"packetInterval,omitempty"`
	PacketSize     *int32  `json:"packetSize,omitempty"`
	PacketTimeout  *string `json:"packetTimeout,omitempty"`
	TargetHost     *string `json:"targetHost,omitempty"`
	UseDNS         *bool   `json:"useDNS,omitempty"`
	UseIPv6        *bool   `json:"useIPv6,omitempty"`
	// ICMPValidationRules holds ICMP ping validation rules (syntheticType=ICMP).
	// For SSL certificate validation rules see ValidationRules.
	ICMPValidationRules []ICMPValidation `json:"-"`
}

// syntheticTestConfigJSON is used internally for marshaling/unmarshaling the
// shared "validationRules" JSON key whose element type depends on SyntheticType.
type syntheticTestConfigJSON struct {
	MarkSyntheticCall bool    `json:"markSyntheticCall"`
	Retries           int32   `json:"retries,omitempty"`
	RetryInterval     int32   `json:"retryInterval,omitempty"`
	SyntheticType     string  `json:"syntheticType"`
	Timeout           *string `json:"timeout,omitempty"`

	URL              *string           `json:"url,omitempty"`
	Operation        *string           `json:"operation,omitempty"`
	Headers          map[string]string `json:"headers,omitempty"`
	Body             *string           `json:"body,omitempty"`
	ValidationString *string           `json:"validationString,omitempty"`
	FollowRedirect   *bool             `json:"followRedirect,omitempty"`
	AllowInsecure    *bool             `json:"allowInsecure,omitempty"`
	ExpectStatus     *int32            `json:"expectStatus,omitempty"`
	ExpectMatch      *string           `json:"expectMatch,omitempty"`
	ExpectExists     []string          `json:"expectExists,omitempty"`
	ExpectNotEmpty   []string          `json:"expectNotEmpty,omitempty"`
	ExpectJson       json.RawMessage   `json:"expectJson,omitempty"`

	Script     *string                       `json:"script,omitempty"`
	ScriptType *string                       `json:"scriptType,omitempty"`
	FileName   *string                       `json:"fileName,omitempty"`
	Scripts    *MultipleScriptsConfiguration `json:"scripts,omitempty"`

	Browser     *string `json:"browser,omitempty"`
	RecordVideo *bool   `json:"recordVideo,omitempty"`

	Lookup           *string                `json:"lookup,omitempty"`
	Server           *string                `json:"server,omitempty"`
	QueryType        *string                `json:"queryType,omitempty"`
	Port             *int32                 `json:"port,omitempty"`
	Transport        *string                `json:"transport,omitempty"`
	AcceptCNAME      *bool                  `json:"acceptCNAME,omitempty"`
	LookupServerName *bool                  `json:"lookupServerName,omitempty"`
	RecursiveLookups *bool                  `json:"recursiveLookups,omitempty"`
	ServerRetries    *int32                 `json:"serverRetries,omitempty"`
	QueryTime        *DNSFilterQueryTime    `json:"queryTime,omitempty"`
	TargetValues     []DNSFilterTargetValue `json:"targetValues,omitempty"`

	Hostname             *string         `json:"hostname,omitempty"`
	DaysRemainingCheck   *int32          `json:"daysRemainingCheck,omitempty"`
	SSLPort              *int32          `json:"sslPort,omitempty"`
	AcceptSelfSignedCert *bool           `json:"acceptSelfSignedCertificate,omitempty"`
	ValidationRules      json.RawMessage `json:"validationRules,omitempty"`

	PacketCount    *int32  `json:"packetCount,omitempty"`
	PacketInterval *string `json:"packetInterval,omitempty"`
	PacketSize     *int32  `json:"packetSize,omitempty"`
	PacketTimeout  *string `json:"packetTimeout,omitempty"`
	TargetHost     *string `json:"targetHost,omitempty"`
	UseDNS         *bool   `json:"useDNS,omitempty"`
	UseIPv6        *bool   `json:"useIPv6,omitempty"`
}

// MarshalJSON serialises SyntheticTestConfig, routing the "validationRules"
// JSON key to either SSLCertificateValidation or ICMPValidation depending on
// SyntheticType.
func (c SyntheticTestConfig) MarshalJSON() ([]byte, error) {
	j := syntheticTestConfigJSON{
		MarkSyntheticCall:    c.MarkSyntheticCall,
		Retries:              c.Retries,
		RetryInterval:        c.RetryInterval,
		SyntheticType:        c.SyntheticType,
		Timeout:              c.Timeout,
		URL:                  c.URL,
		Operation:            c.Operation,
		Headers:              c.Headers,
		Body:                 c.Body,
		ValidationString:     c.ValidationString,
		FollowRedirect:       c.FollowRedirect,
		AllowInsecure:        c.AllowInsecure,
		ExpectStatus:         c.ExpectStatus,
		ExpectMatch:          c.ExpectMatch,
		ExpectExists:         c.ExpectExists,
		ExpectNotEmpty:       c.ExpectNotEmpty,
		ExpectJson:           c.ExpectJson,
		Script:               c.Script,
		ScriptType:           c.ScriptType,
		FileName:             c.FileName,
		Scripts:              c.Scripts,
		Browser:              c.Browser,
		RecordVideo:          c.RecordVideo,
		Lookup:               c.Lookup,
		Server:               c.Server,
		QueryType:            c.QueryType,
		Port:                 c.Port,
		Transport:            c.Transport,
		AcceptCNAME:          c.AcceptCNAME,
		LookupServerName:     c.LookupServerName,
		RecursiveLookups:     c.RecursiveLookups,
		ServerRetries:        c.ServerRetries,
		QueryTime:            c.QueryTime,
		TargetValues:         c.TargetValues,
		Hostname:             c.Hostname,
		DaysRemainingCheck:   c.DaysRemainingCheck,
		SSLPort:              c.SSLPort,
		AcceptSelfSignedCert: c.AcceptSelfSignedCert,
		PacketCount:          c.PacketCount,
		PacketInterval:       c.PacketInterval,
		PacketSize:           c.PacketSize,
		PacketTimeout:        c.PacketTimeout,
		TargetHost:           c.TargetHost,
		UseDNS:               c.UseDNS,
		UseIPv6:              c.UseIPv6,
	}

	switch c.SyntheticType {
	case SyntheticTypeICMP:
		if len(c.ICMPValidationRules) > 0 {
			raw, err := json.Marshal(c.ICMPValidationRules)
			if err != nil {
				return nil, fmt.Errorf("marshaling ICMP validationRules: %w", err)
			}
			j.ValidationRules = raw
		}
	default:
		if len(c.ValidationRules) > 0 {
			raw, err := json.Marshal(c.ValidationRules)
			if err != nil {
				return nil, fmt.Errorf("marshaling SSL validationRules: %w", err)
			}
			j.ValidationRules = raw
		}
	}

	return json.Marshal(j)
}

// UnmarshalJSON deserialises SyntheticTestConfig, routing the "validationRules"
// JSON key to either ICMPValidationRules or ValidationRules based on SyntheticType.
func (c *SyntheticTestConfig) UnmarshalJSON(data []byte) error {
	var j syntheticTestConfigJSON
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}

	c.MarkSyntheticCall = j.MarkSyntheticCall
	c.Retries = j.Retries
	c.RetryInterval = j.RetryInterval
	c.SyntheticType = j.SyntheticType
	c.Timeout = j.Timeout
	c.URL = j.URL
	c.Operation = j.Operation
	c.Headers = j.Headers
	c.Body = j.Body
	c.ValidationString = j.ValidationString
	c.FollowRedirect = j.FollowRedirect
	c.AllowInsecure = j.AllowInsecure
	c.ExpectStatus = j.ExpectStatus
	c.ExpectMatch = j.ExpectMatch
	c.ExpectExists = j.ExpectExists
	c.ExpectNotEmpty = j.ExpectNotEmpty
	c.ExpectJson = j.ExpectJson
	c.Script = j.Script
	c.ScriptType = j.ScriptType
	c.FileName = j.FileName
	c.Scripts = j.Scripts
	c.Browser = j.Browser
	c.RecordVideo = j.RecordVideo
	c.Lookup = j.Lookup
	c.Server = j.Server
	c.QueryType = j.QueryType
	c.Port = j.Port
	c.Transport = j.Transport
	c.AcceptCNAME = j.AcceptCNAME
	c.LookupServerName = j.LookupServerName
	c.RecursiveLookups = j.RecursiveLookups
	c.ServerRetries = j.ServerRetries
	c.QueryTime = j.QueryTime
	c.TargetValues = j.TargetValues
	c.Hostname = j.Hostname
	c.DaysRemainingCheck = j.DaysRemainingCheck
	c.SSLPort = j.SSLPort
	c.AcceptSelfSignedCert = j.AcceptSelfSignedCert
	c.PacketCount = j.PacketCount
	c.PacketInterval = j.PacketInterval
	c.PacketSize = j.PacketSize
	c.PacketTimeout = j.PacketTimeout
	c.TargetHost = j.TargetHost
	c.UseDNS = j.UseDNS
	c.UseIPv6 = j.UseIPv6

	if len(j.ValidationRules) > 0 && string(j.ValidationRules) != "null" {
		switch j.SyntheticType {
		case SyntheticTypeICMP:
			if err := json.Unmarshal(j.ValidationRules, &c.ICMPValidationRules); err != nil {
				return fmt.Errorf("unmarshaling ICMP validationRules: %w", err)
			}
		default:
			if err := json.Unmarshal(j.ValidationRules, &c.ValidationRules); err != nil {
				return fmt.Errorf("unmarshaling SSL validationRules: %w", err)
			}
		}
	}

	return nil
}

type SyntheticTest struct {
	ID               string              `json:"id"`
	Label            string              `json:"label"`
	Description      *string             `json:"description,omitempty"`
	Active           bool                `json:"active"`
	ApplicationID    *string             `json:"applicationId,omitempty"`
	Applications     []string            `json:"applications,omitempty"`
	MobileApps       []string            `json:"mobileApps,omitempty"`
	Websites         []string            `json:"websites,omitempty"`
	Configuration    SyntheticTestConfig `json:"configuration"`
	CustomProperties map[string]string   `json:"customProperties,omitempty"`
	Locations        []string            `json:"locations"`
	PlaybackMode     string              `json:"playbackMode"`
	TestFrequency    *int32              `json:"testFrequency,omitempty"`
	RbacTags         []RbacTag           `json:"rbacTags,omitempty"`
	TenantId         *string             `json:"tenantId,omitempty"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject for SyntheticTest
func (s *SyntheticTest) GetIDForResourcePath() string {
	return s.ID
}
