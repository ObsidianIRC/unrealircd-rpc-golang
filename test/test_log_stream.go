package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ObsidianIRC/unrealircd-rpc-golang"
)

func main() {
	// Connection parameters
	username := "adminpanel"
	password := "password"
	apiLogin := username + ":" + password
	wsURL := "wss://localhost:8600/api"

	fmt.Println("Connecting to UnrealIRCd RPC server...")
	fmt.Printf("URL: %s\n", wsURL)
	fmt.Printf("Username: %s\n", username)

	// Create connection with TLS verification disabled for localhost testing
	conn, err := unrealircd.NewConnection(wsURL, apiLogin, &unrealircd.Options{
		TLSVerify: false,
	})
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	fmt.Println("Connected successfully!")

	// Subscribe to log events
	fmt.Println("\nSubscribing to log events...")
	sources := []string{"all"} // Subscribe to all log sources
	result, err := conn.Log().Subscribe(sources)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	fmt.Printf("Subscription result: %v\n", result)

	// Start listening for log events
	fmt.Println("\nListening for log events (press Ctrl+C to stop)...")
	fmt.Println("You can trigger events by performing actions on the IRC server (connections, commands, errors, etc.)")
	fmt.Println("========================================")

	// Event loop to receive streaming log events
	eventCount := 0
	for {
		event, err := conn.EventLoop()
		if err != nil {
			log.Printf("Error in event loop: %v", err)
			continue
		}

		// If event is nil, it's a timeout (no event received)
		if event == nil {
			// Print a dot to show we're still listening
			fmt.Print(".")
			continue
		}

		// Process the received event
		eventCount++
		fmt.Printf("\n[Event #%d] Received at %s:\n", eventCount, time.Now().Format("15:04:05"))

		// Try to parse the event as a map
		if eventMap, ok := event.(map[string]interface{}); ok {
			// Pretty print the event
			for key, value := range eventMap {
				fmt.Printf("  %s: %v\n", key, value)
			}
		} else {
			// If it's not a map, just print it as-is
			fmt.Printf("  Raw event: %v\n", event)
		}
		fmt.Println("----------------------------------------")
	}
}
