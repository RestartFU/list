package list

import (
	"os"
	"sync"
)

type List struct {
	settings *Settings
	cMutex   sync.RWMutex

	List    []string
	Enabled bool
}

func New(settings *Settings) (*List, error) {
	whitelist := &List{settings: settings}

	if err := settings.Gophig.GetConf(whitelist); os.IsNotExist(err) {
		if err = settings.Gophig.SetConf(&List{List: make([]string, 0), Enabled: true}); err != nil {
			return whitelist, err
		}
		if err = settings.Gophig.GetConf(whitelist); err != nil {
			return whitelist, err
		}
	}
	return whitelist, nil
}

// Add adds a new username to the whitelist
func (w *List) Add(username string) error {
	if !w.Listed(username) {
		if w.settings.CacheOnly {
			w.addCache(username)
			return nil
		}
		return w.add(username)
	}
	return nil
}

// Remove removes the username provided from the whitelist
func (w *List) Remove(username string) error {
	if w.Listed(username) {
		if w.settings.CacheOnly {
			w.removeCache(username)
			return nil
		}
		return w.remove(username)
	}
	return nil
}

// Whitelisted returns a bool of if the player is whitelisted or not
func (w *List) Listed(username string) bool {
	if username == "" {
		return false
	}
	if w.settings.CacheOnly {
		return w.listedCache(username)
	}
	return w.listed(username)
}

func (w *List) Close() error {
	return w.close()
}
