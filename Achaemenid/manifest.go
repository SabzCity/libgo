/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"time"

	etime "../earth-time"
	lang "../language"
)

// Manifest store server manifest data
// All string slice is multi language and in order by ManifestLanguages order
type Manifest struct {
	SocietyID  uint32
	DomainName string
	Email      string
	Icon       string

	Organization   map[lang.Language]string
	Name           map[lang.Language]string
	Description    map[lang.Language]string
	TermsOfService map[lang.Language]string
	Licence        map[lang.Language]string
	TAGS           []string // Use to categorized apps e.g. Music, GPS, ...

	RequestedPermission []uint32 // ServiceIDs from PersiaOS services e.g. InternetInBackground, Notification, ...
	TechnicalInfo       TechnicalInfo
}

// TechnicalInfo store some technical information but may different from really server condition!
type TechnicalInfo struct {
	// Shutdown settings
	ShutdownDelay time.Duration // the server will wait for at least this amount of time for active streams to finish!

	// Server Overal rate limit
	MaxOpenConnection     uint64         // The maximum number of concurrent connections the app may serve.
	ConnectionIdleTimeout etime.Duration // In seconds
	// MaxStreamHeaderSize   uint64 // For stream protocols with variable header size like HTTP
	// MaxStreamPayloadSize  uint64 // For stream protocols with variable payload size like sRPC, HTTP, ...

	// Guest rete limit - Connection.OwnerType==0
	GuestMaxConnections            uint64 // 0 means App not accept guest connection.
	GuestMaxUserConnectionsPerAddr uint64
	GuestMaxConcurrentStreams      uint32
	GuestMaxStreamConnectionDaily  uint32 // Max open stream per day for a guest connection. overflow will drop on creation!
	GuestMaxServiceCallDaily       uint64 // 0 means no limit and good for PayAsGo strategy!
	GuestMaxBytesSendDaily         uint64
	GuestMaxBytesReceiveDaily      uint64
	GuestMaxPacketsSendDaily       uint64
	GuestMaxPacketsReceiveDaily    uint64

	// Registered rate limit - Connection.OwnerType==1
	RegisteredMaxConnections            uint64
	RegisteredMaxUserConnectionsPerAddr uint64
	RegisteredMaxConcurrentStreams      uint32
	RegisteredMaxStreamConnectionDaily  uint32 // Max open stream per day for a Registered user connection. overflow will drop on creation!
	RegisteredMaxServiceCallDaily       uint64 // 0 means no limit and good for PayAsGo strategy!
	RegisteredMaxBytesSendDaily         uint64
	RegisteredMaxBytesReceiveDaily      uint64
	RegisteredMaxPacketsSendDaily       uint64
	RegisteredMaxPacketsReceiveDaily    uint64

	// If you want to know Connection.OwnerType>1 rate limit strategy, You must read server codes!!

	// Minimum hardware specification for each instance of application.
	CPUCores uint8  // Number
	CPUSpeed uint64 // Hz
	RAM      uint64 // Byte
	GPU      uint64 // Hz
	Network  uint64 // Byte per second
	Storage  uint64 // Byte, HHD||SSD||... indicate by DataCentersClassForDataStore

	// Distribution
	DistributeOutOfSociety bool          // Allow to run service-only instance of app out of original society belong to.
	DataCentersClass       uint8         // 0:FirstClass 256:Low-Quality default:5
	MaxNodeNumber          uint32        // default:3
	NodeFailureTimeOut     time.Duration // Max suggestion is 6 hour, other service only node replace failed node! not use in network failure, it is handy proccess!
}
