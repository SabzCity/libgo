/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"../assets"
)

// DefaultServer use as default server.
var DefaultServer = &defaultServer
var defaultServer Server

// Server represents needed data to serving as a network server.
type Server struct {
	Status                int // States locate in const of this file.
	Manifest              Manifest
	Network               Network
	ProtocolsHandlers     ProtocolsHandlers
	Services              Services
	PublicKeyCryptography PublicKeyCryptography
	Connections           Connections
	Assets                assets.Folder // Any data files to serve!
}

// Server Status
const (
	// ServerStateStop indicate server had been stopped
	ServerStateStop int = iota
	// ServerStateRunning indicate server is working
	ServerStateRunning
	// ServerStateStopping indicate server want to stop
	ServerStateStopping
	// ServerStateStarting indicate server plan to start and working on it
	ServerStateStarting
)

// Init method use to initialize related object with default data to prevent from panic!
func (s *Server) Init() {
	if s == nil {
		s = DefaultServer
	}
	s.Services.init()
	s.Connections.init()
}

// Start will start the server.
func (s *Server) Start() (err error) {
	s.Status = ServerStateStarting

	// Recover from panics if exist.
	// defer panicHandler(s)

	// watch for SIGTERM and SIGINT from the operating system, and notify the app on the channel
	var sig = make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM)
	signal.Notify(sig, syscall.SIGINT)
	go s.HandleOSSignals(sig)

	// Get UserGivenPermission from OS

	// Make & Register publicKey
	if err = s.PublicKeyCryptography.RegisterPublicKey(); err != nil {
		return err
	}

	// Get GP & MTU from OS router.
	if err = s.Network.RegisterGP(); err != nil {
		return err
	}

	s.Status = ServerStateRunning

	// register s.HandleGP() for income packet handler

	return nil
}

// HandleOSSignals use to handle OS signals!
func (s *Server) HandleOSSignals(sigChannel chan os.Signal) {
	var sig = <-sigChannel
	switch sig {
	case syscall.SIGTERM:
		// wait for our os signal to stop the app
		// on the graceful stop channel
		// this goroutine will block until we get an OS signal
		Log("caught sig: %+v", sig)

		// sleep for 60 seconds to waiting for app to finish,
		Log("Waiting for server to finish, will take 60 seconds")

		s.Shutdown()

		os.Exit(s.Status)
	}
}

// Shutdown use to graceful stop server!!
func (s *Server) Shutdown() {
	// ... Do business Logic for shutdown
	// Shutdown works by:
	// first closing open listener for income packet and refuse all new packet,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down

	// Send signal to DNS & Certificate server to revoke app data.

	// Wait to finish above logic
	time.Sleep(60 * time.Second)

	// it must change to ServerStateStop(0) otherwise it means app can't close normally
	s.Status = ServerStateRunning
}

// SendStream use to register a stream to send pool and automatically send to other side.
func (s *Server) SendStream(st *Stream) (err error) {
	// First Check st.Connection.Status to ability send stream over it

	return nil
}