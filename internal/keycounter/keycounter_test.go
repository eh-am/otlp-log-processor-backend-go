package keycounter_test

import (
	"bytes"
	"testing"

	"dash0.com/otlp-log-processor-backend/internal/keycounter"
	"github.com/alecthomas/assert/v2"
)

func TestCount(t *testing.T) {
	var buf bytes.Buffer
	kc := keycounter.NewKeyCounter("mykey", &buf)

	// Initially it shouldn't report anything
	kc.Flush()
	assert.Equal(t, 0, buf.Len())
	buf.Reset()

	kc.Add("myval")
	kc.Add("myval")
	kc.Flush()
	assert.Equal(t, `"myval" - 2\n`, buf.String())
	buf.Reset()

	kc.Add("myval")
	kc.Add("myval")
	kc.Add("myval")
	kc.Add("myval2")
	kc.Flush()
	assert.Equal(t, `"myval" - 3\n"myval2" - 1\n`, buf.String())
	buf.Reset()
}
