package fix

import (
	"fmt"

	"github.com/quickfixgo/quickfix"
)

// Application implements the quickfix.Application interface
type Application struct{ respChan chan string }

// handleResp provide interface to respond to a component
func (t Application) handleResp(val string) { t.respChan <- val }

// OnCreate is called when quickfix creates a new session.
// A session comes into and remains in existence for the life of application.
// Sessions exist whether or not a counter party is connected to it.
// As soon as a session is created, you can begin sending messages to it.
func (t Application) OnCreate(sessionID quickfix.SessionID) {
	t.handleResp(fmt.Sprintf("OnCreate. SessionID: %s\n", sessionID))
	return
}

// OnLogon notifies when a valid logon has been established.
// This is called when a connection has been established and the FIX logon process
// has completed with both parties exchanging valid logon messages.
func (t Application) OnLogon(sessionID quickfix.SessionID) {
	t.handleResp(fmt.Sprintf("OnLogon. SessionID: %s\n", sessionID))
	return
}

// OnLogout notifies when a FIX session is no longer online.
// This could happen during a normal logout exhange or because of a forced
// termination or a loss of network connection.
func (t Application) OnLogout(sessionID quickfix.SessionID) {
	t.handleResp(fmt.Sprintf("OnLogout. SessionID: %s\n", sessionID))
	return
}

// FromAdmin notifies you when an administrative message is sent from a counterparty to your FIX engine.
// This can be useful for doing extra validation on logon messages.
func (t Application) FromAdmin(
	msg *quickfix.Message,
	sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {

	t.handleResp(fmt.Sprintf(
		"FromAdmin. SessionID:%s  Msg:%s \n", sessionID, msg.String()))
	return
}

// ToAdmin notifies administrative messages that are being sent from your FIX engine to the counter party.
// This is useful for logging purposes.
func (t Application) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	t.respChan <- fmt.Sprintf("ToAdmin. SessionID: %s\n", sessionID)
	ModifyMsg(msg)
	return
}

// ToApp is a callback for application messages that are being sent to a counterparty.
// If you throw a DoNotSend error in this function, the application will not send the message.
// This is useful if the application has been asked to resend a message such as an order that is no longer
// relevant for the current market.
func (t Application) ToApp(
	msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	t.respChan <- fmt.Sprintf(
		"ToApp. SessionID: %s Sending: %s\n", sessionID, msg.String())
	return
}

// FromApp receives application level requests.
// If your application is a sell-side, this is where you will get your new order requests.
// If you're buy-side, this is where you will get your execution reports.
func (t Application) FromApp(msg *quickfix.Message,
	sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	t.respChan <- msg.String()
	return
}
