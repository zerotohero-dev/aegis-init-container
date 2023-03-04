/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"fmt"
	"github.com/zerotohero-dev/aegis-core/log"
	"github.com/zerotohero-dev/aegis-sdk-go/startup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.InfoLn("Starting Aegis Init Container")
	go startup.Watch()

	// Block the process from exiting, but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	select {
	case e := <-s:
		fmt.Println(e)
		panic("bye cruel world!")
	}
}
