package chars

import "testing"

func TestNewID(t *testing.T) {
	id := NewBsonID()
	t.Log(id,len(id))

	id = NewUUID(true)
	t.Log(id,len(id))

	id = NewUUID(false)
	t.Log(id,len(id))

}
