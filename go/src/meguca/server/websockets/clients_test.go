package websockets

import (
	"testing"

	. "meguca/test"
)

func newClientMap() *ClientMap {
	return &ClientMap{
		clients: make(map[*Client]SyncID),
	}
}

func TestMapAddHasRemove(t *testing.T) {
	t.Parallel()

	m := newClientMap()
	sv := newWSServer(t)
	defer sv.Close()
	id := SyncID{
		OP:    1,
		Board: "a",
	}

	// Add client
	cl, _ := sv.NewClient()
	m.add(cl, id)
	assertSyncID(t, m, cl, id)
	// Remove client
	m.remove(cl)
	if synced, _ := m.GetSync(cl); synced {
		t.Error("client still synced")
	}
}

func assertSyncID(t *testing.T, m *ClientMap, cl *Client, id SyncID) {
	synced, sync := m.GetSync(cl)
	if !synced {
		t.Error("client not synced")
	}
	if sync != id {
		LogUnexpected(t, id, sync)
	}
}

func TestMapChangeSync(t *testing.T) {
	t.Parallel()

	oldSync := SyncID{
		OP:    1,
		Board: "a",
	}
	newSync := SyncID{
		OP:    2,
		Board: "g",
	}
	m := newClientMap()
	sv := newWSServer(t)
	defer sv.Close()

	cl, _ := sv.NewClient()
	m.add(cl, oldSync)
	assertSyncID(t, m, cl, oldSync)

	m.changeSync(cl, newSync)
	assertSyncID(t, m, cl, newSync)
}
