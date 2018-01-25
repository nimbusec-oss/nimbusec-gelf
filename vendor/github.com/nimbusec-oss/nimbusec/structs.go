package nimbusec

import (
	"time"
)

type DomainID string
type Domain struct {
	ID     DomainID `json:"id"`
	Bundle BundleID `json:"bundle"`
	Name   string   `json:"name"`
	URL    string   `json:"url"`
}

type BundleID string
type Bundle struct {
	ID           BundleID `json:"id"`
	Name         string   `json:"name"`
	StartDate    int64    `json:"startDate"`
	EndDate      int64    `json:"endDate"`
	Price        int      `json:"price"`
	Currency     string   `json:"currency"`
	Active       int      `json:"active"`
	Capacity     int      `json:"capacity"`
	ChargingType string   `json:"chargingType"`

	Features struct {
		Defacement struct {
			Available bool `json:"available"`
			Nimbusec  bool `json:"nimbusec"`
			ZoneH     bool `json:"zoneh"`
		} `json:"defacement"`
		UnwantedContent struct {
			Available bool `json:"available"`
		} `json:"unwantedContent"`
		Malware struct {
			Available bool `json:"available"`
			Nimbusec  bool `json:"nimbusec"`
			Ikarus    bool `json:"ikarus"`
			Avira     bool `json:"avira"`
			LastLine  bool `json:"lastline"`
			ClamAV    bool `json:"clamav"`
		} `json:"malware"`
		Reputation struct {
			Available bool `json:"available"`
		} `json:"reputation"`
		TLS struct {
			Available bool `json:"available"`
		} `json:"tls"`
		Webshell struct {
			Available bool `json:"available"`
		} `json:"webshell"`
		Application struct {
			Available bool `json:"available"`
		} `json:"application"`
		Scanning struct {
			Available        bool `json:"available"`
			FastScanInterval int  `json:"fastScanInterval"`
			DeepScanInterval int  `json:"deepScanInterval"`
			Quota            int  `json:"quota"`
			FromEU           bool `json:"fromEU"`
			FromUS           bool `json:"fromUS"`
			FromASIA         bool `json:"fromAsia"`
			Mobile           bool `json:"mobile"`
		} `json:"scanning"`
		Notification struct {
			Available   bool `json:"available"`
			EMail       bool `json:"email"`
			TextMessage bool `json:"textMessage"`
		} `json:"notification"`
	} `json:"features"`
}

type NotificationID string
type Notification struct {
	ID         NotificationID `json:"id"`
	Domain     DomainID       `json:"domain"`
	User       UserID         `json:"user"`
	Transport  string         `json:"transport"`
	Blacklist  int            `json:"blacklist"`
	Defacement int            `json:"defacement"`
	Malware    int            `json:"malware"`
}

type NotificationUpdate struct {
	Transport  string `json:"transport"`
	Blacklist  int    `json:"blacklist"`
	Defacement int    `json:"defacement"`
	Malware    int    `json:"malware"`
}

type UserID string
type User struct{}

type IssueID string
type Issue struct {
	ID        IssueID     `json:"id"`
	Domain    DomainID    `json:"domain"`
	Status    IssueStatus `json:"status"`
	Event     string      `json:"event"`
	Category  string      `json:"category"`
	Severity  int         `json:"severity"`
	FirstSeen time.Time   `json:"firstSeen"`
	LastSeen  time.Time   `json:"lastSeen"`
	Details   interface{} `json:"details,omitempty"`
}

type IssueUpdate struct {
	Status  IssueStatus `json:"status"`
	Comment string      `json:"comment"`
}

type IssueStatus string

const (
	IssueStatusPending       IssueStatus = "pending"
	IssueStatusAcknowledged  IssueStatus = "acknowledged"
	IssueStatusIgnored       IssueStatus = "ignored"
	IssueStatusFalsePositive IssueStatus = "falsepositive"
)

type ApplicationOutdatedDetails struct {
	Name          string `json:"name"`
	URL           string `json:"url,omitempty"`
	Path          string `json:"path,omitempty"`
	Version       string `json:"version"`
	LatestVersion string `json:"latestVersion"`
}

type ApplicationVulnerableDetails struct {
	Name            string `json:"name"`
	URL             string `json:"url,omitempty"`
	Path            string `json:"path,omitempty"`
	Version         string `json:"version"`
	Vulnerabilities []struct {
		CVE         string  `json:"cve"`
		Score       float64 `json:"score"`
		Description string  `json:"description"`
		Link        string  `json:"link"`
	} `json:"vulnerabilities"`
}

type DefacementDetails struct {
	URL    string `json:"url"`
	Threat string `json:"threat"`
}

type ZoneHDetails struct {
	URL    string `json:"url"`
	Threat string `json:"threat"`
}

type MalwareDetails struct {
	URL    string `json:"url"`
	Threat string `json:"threat"`
	AV     string `json:"av"`
}

type LastlineDetails struct {
	Score           int `json:"score"`
	AnalysisSubject struct {
		URL     string `json:"url"`
		Referer string `json:"referer"`
	} `json:"analysis_subject"`
	Threat            string   `json:"threat"`
	ThreatClass       string   `json:"threat_class"`
	MaliciousActivity []string `json:"malicious_activity"`
}

type TLSCertificateDetails struct {
	Order     int      `json:"order"`
	NotAfter  int64    `json:"notAfter"`
	NotBefore int64    `json:"notBefore"`
	Issuer    string   `json:"issuer"`
	CName     string   `json:"cName"`
	AltNames  []string `json:"altNames"`
}

type TLSConfigurationDetails struct {
	Protocol string `json:"protocol,omitempty"`
	Cipher   string `json:"cipher,omitempty"`
}

type BlacklistDetails struct {
	Blacklist    string   `json:"blacklist"`
	BlacklistURL string   `json:"blacklistURL"`
	Reasons      []string `json:"reasons"`
}

type SuspiciousLinkDetails struct {
	URL   string `json:"url"`
	Links []struct {
		Link       string             `json:"link"`
		Blacklists []BlacklistDetails `json:"blacklists"`
	} `json:"links"`
}

type SuspiciousRequestDetails struct {
	Entity     string             `json:"entity"`
	URLs       []string           `json:"urls"`
	Blacklists []BlacklistDetails `json:"blacklists"`
}

type ChangedFileDetails struct {
	URL    string  `json:"url"`
	DiffID int     `json:"diff"`
	Score  float64 `json:"score"`
	Votes  map[string]struct {
		Votes       int         `json:"votes"`
		Probability float64     `json:"probability"`
		Extra       interface{} `json:"extra"`
	} `json:"votes"`
}

type WebshellDetails struct {
	AV          string `json:"av"`
	MD5         string `json:"md5"`
	Path        string `json:"path"`
	Size        int    `json:"size"`
	Owner       string `json:"owner"`
	Group       string `json:"group"`
	MTime       int    `json:"mtime"`
	Threat      string `json:"threat"`
	Permissions string `json:"permissions"`
}
