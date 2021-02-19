package asciitransport

import (
	"fmt"

	"github.com/btwiuse/pretty"

	"k0s.io/pkg/asciitransport/cast"
)

type ResizeEvent cast.Header
type OutputEvent Event
type PingEvent Event
type InputEvent Event
type Event cast.Event

func (e *Event) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&e.Time, &e.Type, &e.Data}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, w := len(tmp), wantLen; g != w {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, w)
	}
	return nil
}

func (e *ResizeEvent) String() string { return pretty.JsonString(e) }
func (e *InputEvent) String() string {
	return pretty.JsonString([]interface{}{&e.Time, &e.Type, &e.Data})
}
func (e *OutputEvent) String() string {
	return pretty.JsonString([]interface{}{&e.Time, &e.Type, &e.Data})
}
func (e *PingEvent) String() string {
	return pretty.JsonString([]interface{}{&e.Time, &e.Type, &e.Data})
}
