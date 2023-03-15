package list

import (
	"strings"
)

func (w *List) addCache(username string) {
	w.cMutex.Lock()
	defer w.cMutex.Unlock()
	w.List = append(w.List, username)
}
func (w *List) add(username string) error {
	w.addCache(username)
	return w.settings.Gophig.SetConf(w)
}

func (w *List) removeCache(username string) {
	w.cMutex.Lock()
	defer w.cMutex.Unlock()
	for i, u := range w.List {
		if strings.ToLower(u) == strings.ToLower(username) {
			w.List = append(w.List[:i], w.List[i+1:]...)
		}
	}
}
func (w *List) remove(username string) error {
	w.removeCache(username)
	return w.settings.Gophig.SetConf(w)
}

func (w *List) listedCache(username string) bool {
	w.cMutex.Lock()
	defer w.cMutex.Unlock()
	for _, u := range w.List {
		if strings.ToLower(u) == strings.ToLower(username) {
			return true
		}
	}
	return false
}

func (w *List) listed(username string) bool {
	err := w.settings.Gophig.GetConf(w)
	return err == nil && w.listedCache(username)
}

func (w *List) closeCache() {
	w = nil
}
func (w *List) close() error {
	err := w.settings.Gophig.SetConf(w)
	w.closeCache()
	return err
}
