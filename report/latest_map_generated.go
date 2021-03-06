// Generated file, do not edit.
// To regenerate, run ../extras/generate_latest_map ./latest_map_generated.go string NodeControlData

package report

import (
	"fmt"
	"time"

	"github.com/ugorji/go/codec"
	"github.com/weaveworks/ps"
)

type stringLatestEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Value     string    `json:"value"`
	dummySelfer
}

// String returns the StringLatestEntry's string representation.
func (e *stringLatestEntry) String() string {
	return fmt.Sprintf("%v (%s)", e.Value, e.Timestamp.String())
}

// Equal returns true if the supplied StringLatestEntry is equal to this one.
func (e *stringLatestEntry) Equal(e2 *stringLatestEntry) bool {
	return e.Timestamp.Equal(e2.Timestamp) && e.Value == e2.Value
}

// StringLatestMap holds latest string instances.
type StringLatestMap struct{ ps.Map }

var emptyStringLatestMap = StringLatestMap{ps.NewMap()}

// MakeStringLatestMap makes an empty StringLatestMap.
func MakeStringLatestMap() StringLatestMap {
	return emptyStringLatestMap
}

// Size returns the number of elements.
func (m StringLatestMap) Size() int {
	if m.Map == nil {
		return 0
	}
	return m.Map.Size()
}

// Merge produces a fresh StringLatestMap containing the keys from both inputs.
// When both inputs contain the same key, the newer value is used.
func (m StringLatestMap) Merge(other StringLatestMap) StringLatestMap {
	output := mergeMaps(m.Map, other.Map, func(a, b interface{}) bool {
		return a.(*stringLatestEntry).Timestamp.Before(b.(*stringLatestEntry).Timestamp)
	})
	return StringLatestMap{output}
}

// Lookup the value for the given key.
func (m StringLatestMap) Lookup(key string) (string, bool) {
	v, _, ok := m.LookupEntry(key)
	if !ok {
		var zero string
		return zero, false
	}
	return v, true
}

// LookupEntry returns the raw entry for the given key.
func (m StringLatestMap) LookupEntry(key string) (string, time.Time, bool) {
	if m.Map == nil {
		var zero string
		return zero, time.Time{}, false
	}
	value, ok := m.Map.Lookup(key)
	if !ok {
		var zero string
		return zero, time.Time{}, false
	}
	e := value.(*stringLatestEntry)
	return e.Value, e.Timestamp, true
}

// Set the value for the given key.
func (m StringLatestMap) Set(key string, timestamp time.Time, value string) StringLatestMap {
	if m.Map == nil {
		m.Map = ps.NewMap()
	}
	return StringLatestMap{m.Map.Set(key, &stringLatestEntry{Timestamp: timestamp, Value: value})}
}

// ForEach executes fn on each key value pair in the map.
func (m StringLatestMap) ForEach(fn func(k string, timestamp time.Time, v string)) {
	if m.Map != nil {
		m.Map.ForEach(func(key string, value interface{}) {
			fn(key, value.(*stringLatestEntry).Timestamp, value.(*stringLatestEntry).Value)
		})
	}
}

// String returns the StringLatestMap's string representation.
func (m StringLatestMap) String() string {
	return mapToString(m.Map)
}

// DeepEqual tests equality with other StringLatestMap.
func (m StringLatestMap) DeepEqual(n StringLatestMap) bool {
	return mapEqual(m.Map, n.Map, func(val, otherValue interface{}) bool {
		return val.(*stringLatestEntry).Equal(otherValue.(*stringLatestEntry))
	})
}

// CodecEncodeSelf implements codec.Selfer.
func (m *StringLatestMap) CodecEncodeSelf(encoder *codec.Encoder) {
	mapWrite(m.Map, encoder, func(encoder *codec.Encoder, val interface{}) {
		val.(*stringLatestEntry).CodecEncodeSelf(encoder)
	})
}

// CodecDecodeSelf implements codec.Selfer.
func (m *StringLatestMap) CodecDecodeSelf(decoder *codec.Decoder) {
	out := mapRead(decoder, func(isNil bool) interface{} {
		value := &stringLatestEntry{}
		if !isNil {
			value.CodecDecodeSelf(decoder)
		}
		return value
	})
	*m = StringLatestMap{out}
}

// MarshalJSON shouldn't be used, use CodecEncodeSelf instead.
func (StringLatestMap) MarshalJSON() ([]byte, error) {
	panic("MarshalJSON shouldn't be used, use CodecEncodeSelf instead")
}

// UnmarshalJSON shouldn't be used, use CodecDecodeSelf instead.
func (*StringLatestMap) UnmarshalJSON(b []byte) error {
	panic("UnmarshalJSON shouldn't be used, use CodecDecodeSelf instead")
}

type nodeControlDataLatestEntry struct {
	Timestamp time.Time       `json:"timestamp"`
	Value     NodeControlData `json:"value"`
	dummySelfer
}

// String returns the StringLatestEntry's string representation.
func (e *nodeControlDataLatestEntry) String() string {
	return fmt.Sprintf("%v (%s)", e.Value, e.Timestamp.String())
}

// Equal returns true if the supplied StringLatestEntry is equal to this one.
func (e *nodeControlDataLatestEntry) Equal(e2 *nodeControlDataLatestEntry) bool {
	return e.Timestamp.Equal(e2.Timestamp) && e.Value == e2.Value
}

// NodeControlDataLatestMap holds latest NodeControlData instances.
type NodeControlDataLatestMap struct{ ps.Map }

var emptyNodeControlDataLatestMap = NodeControlDataLatestMap{ps.NewMap()}

// MakeNodeControlDataLatestMap makes an empty NodeControlDataLatestMap.
func MakeNodeControlDataLatestMap() NodeControlDataLatestMap {
	return emptyNodeControlDataLatestMap
}

// Size returns the number of elements.
func (m NodeControlDataLatestMap) Size() int {
	if m.Map == nil {
		return 0
	}
	return m.Map.Size()
}

// Merge produces a fresh NodeControlDataLatestMap containing the keys from both inputs.
// When both inputs contain the same key, the newer value is used.
func (m NodeControlDataLatestMap) Merge(other NodeControlDataLatestMap) NodeControlDataLatestMap {
	output := mergeMaps(m.Map, other.Map, func(a, b interface{}) bool {
		return a.(*nodeControlDataLatestEntry).Timestamp.Before(b.(*nodeControlDataLatestEntry).Timestamp)
	})
	return NodeControlDataLatestMap{output}
}

// Lookup the value for the given key.
func (m NodeControlDataLatestMap) Lookup(key string) (NodeControlData, bool) {
	v, _, ok := m.LookupEntry(key)
	if !ok {
		var zero NodeControlData
		return zero, false
	}
	return v, true
}

// LookupEntry returns the raw entry for the given key.
func (m NodeControlDataLatestMap) LookupEntry(key string) (NodeControlData, time.Time, bool) {
	if m.Map == nil {
		var zero NodeControlData
		return zero, time.Time{}, false
	}
	value, ok := m.Map.Lookup(key)
	if !ok {
		var zero NodeControlData
		return zero, time.Time{}, false
	}
	e := value.(*nodeControlDataLatestEntry)
	return e.Value, e.Timestamp, true
}

// Set the value for the given key.
func (m NodeControlDataLatestMap) Set(key string, timestamp time.Time, value NodeControlData) NodeControlDataLatestMap {
	if m.Map == nil {
		m.Map = ps.NewMap()
	}
	return NodeControlDataLatestMap{m.Map.Set(key, &nodeControlDataLatestEntry{Timestamp: timestamp, Value: value})}
}

// ForEach executes fn on each key value pair in the map.
func (m NodeControlDataLatestMap) ForEach(fn func(k string, timestamp time.Time, v NodeControlData)) {
	if m.Map != nil {
		m.Map.ForEach(func(key string, value interface{}) {
			fn(key, value.(*nodeControlDataLatestEntry).Timestamp, value.(*nodeControlDataLatestEntry).Value)
		})
	}
}

// String returns the NodeControlDataLatestMap's string representation.
func (m NodeControlDataLatestMap) String() string {
	return mapToString(m.Map)
}

// DeepEqual tests equality with other NodeControlDataLatestMap.
func (m NodeControlDataLatestMap) DeepEqual(n NodeControlDataLatestMap) bool {
	return mapEqual(m.Map, n.Map, func(val, otherValue interface{}) bool {
		return val.(*nodeControlDataLatestEntry).Equal(otherValue.(*nodeControlDataLatestEntry))
	})
}

// CodecEncodeSelf implements codec.Selfer.
func (m *NodeControlDataLatestMap) CodecEncodeSelf(encoder *codec.Encoder) {
	mapWrite(m.Map, encoder, func(encoder *codec.Encoder, val interface{}) {
		val.(*nodeControlDataLatestEntry).CodecEncodeSelf(encoder)
	})
}

// CodecDecodeSelf implements codec.Selfer.
func (m *NodeControlDataLatestMap) CodecDecodeSelf(decoder *codec.Decoder) {
	out := mapRead(decoder, func(isNil bool) interface{} {
		value := &nodeControlDataLatestEntry{}
		if !isNil {
			value.CodecDecodeSelf(decoder)
		}
		return value
	})
	*m = NodeControlDataLatestMap{out}
}

// MarshalJSON shouldn't be used, use CodecEncodeSelf instead.
func (NodeControlDataLatestMap) MarshalJSON() ([]byte, error) {
	panic("MarshalJSON shouldn't be used, use CodecEncodeSelf instead")
}

// UnmarshalJSON shouldn't be used, use CodecDecodeSelf instead.
func (*NodeControlDataLatestMap) UnmarshalJSON(b []byte) error {
	panic("UnmarshalJSON shouldn't be used, use CodecDecodeSelf instead")
}
