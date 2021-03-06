package message

import (
	"bytes"
	"time"
)

type NullMessage struct {
}

type NullTransaction struct {
	NullMessage
}

var nullMessage = &NullMessage{}

var nullTransaction = &NullTransaction{}

func (m *NullMessage) Complete() {
}

func (m *NullMessage) GetType() string {
	return ""
}

func (m *NullMessage) GetName() string {
	return ""
}

func (m *NullMessage) SetName(name string) {
}

func (m *NullMessage) GetStatus() string {
	return ""
}

func (m *NullMessage) GetData() *bytes.Buffer {
	return nil
}

func (m *NullMessage) GetTime() time.Time {
	return time.Now()
}

func (m *NullMessage) SetTime(time time.Time) {
}

func (m *NullMessage) SetTimestamp(timestampMs int64) {
}

func (m *NullMessage) AddData(k string, v ...string) {
}

func (m *NullMessage) SetData(v string) {
}

func (m *NullMessage) SetStatus(status string) {
}

func (m *NullMessage) SetMessageId(id string) {
}

func (m *NullMessage) GetMessageId() string {
	return ""
}

func (t *NullTransaction) GetTopTrans() bool {
	return true
}

func (t *NullTransaction) SetTopTrans(isTopTrans bool) {
}

func (t *NullTransaction) GetChildren() []Messager {
	return []Messager{}
}

func (t *NullTransaction) GetDuration() time.Duration {
	return 0
}

func (t *NullTransaction) SetDuration(duration time.Duration) {
}

func (t *NullTransaction) SetDurationStart(duration time.Time) {
}

func (t *NullTransaction) NewEvent(mtype, name string) Messager {
	return nullMessage
}

func (t *NullTransaction) LogEvent(mtype, name string, args ...string) {
}

func (t *NullTransaction) AddChild(m Messager) {
}
