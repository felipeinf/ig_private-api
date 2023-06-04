package scrap

import (
	"log"
	"os"
	"time"

	"github.com/ahmdrz/goinsta"
)

const (
	sleepTime = 5 * time.Second
)

////////////////////////////////////////////////////////////////////////////////////////////
////// COMMON //////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////

func GetInstagram(private bool) (*goinsta.Instagram, error) {
	// Get bot credentials
	username := "lahaipway"
	password := "alohahaip"

	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %s\n", err)
		return nil, err
	}

	// Check for Instagram configuration file
	inst, err := goinsta.Import(dir + "/.bot_" + username)
	if err != nil {
		// Login to Instagram
		inst = goinsta.New(username, password)

		if err := inst.Login(); err != nil {
			log.Printf("Error logging in to Instagram: %s\n", err)
			return nil, err
		}

		// Sleep
		time.Sleep(sleepTime)
		inst.Export(dir + "/.bot_" + username)
	}

	return inst, nil
}
