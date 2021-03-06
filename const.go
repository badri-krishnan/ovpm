package ovpm

// Version defines the version of ovpm.
var Version = "development"

const (
	// DefaultVPNPort is the default OpenVPN port to listen.
	DefaultVPNPort = "1197"

	// DefaultVPNProto is the default OpenVPN protocol to use.
	DefaultVPNProto = UDPProto

	// DefaultVPNNetwork is the default OpenVPN network to use.
	DefaultVPNNetwork = "10.9.0.0/24"

	// DefaultVPNDNS is the default DNS to push to clients.
	DefaultVPNDNS = "8.8.8.8"

	// DefaultDaemonPort is the port OVPMD will listen by default if something else is not specified.
	DefaultDaemonPort = 9090

	// DefaultKeepalivePeriod is the default ping period to check if the remote peer is alive.
	DefaultKeepalivePeriod = "2"

	// DefaultKeepaliveTimeout is the default ping timeout to assume that remote peer is down.
	DefaultKeepaliveTimeout = "4"

	etcBasePath = "/etc/ovpm/"
	varBasePath = "/var/db/ovpm/"

	_DefaultConfigPath    = etcBasePath + "ovpm.ini"
	_DefaultDBPath        = varBasePath + "db.sqlite3"
	_DefaultVPNConfPath   = varBasePath + "server.conf"
	_DefaultVPNCCDPath    = varBasePath + "ccd"
	_DefaultCertPath      = varBasePath + "server.crt"
	_DefaultKeyPath       = varBasePath + "server.key"
	_DefaultCACertPath    = varBasePath + "ca.crt"
	_DefaultCAKeyPath     = varBasePath + "ca.key"
	_DefaultDHParamsPath  = varBasePath + "dh4096.pem"
	_DefaultCRLPath       = varBasePath + "crl.pem"
	_DefaultStatusLogPath = varBasePath + "openvpn-status.log"
)

// Testing is used to determine whether we are testing or running normally.
// Set it to true when testing.
var Testing = false
