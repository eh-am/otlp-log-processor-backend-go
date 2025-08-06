package keycounter_test

import (
	"bytes"
	"fmt"
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

	fmt.Println("buf string", buf.String())
	assert.Equal(t, `"myval" - 2
`, buf.String())
	buf.Reset()

	kc.Add("myval")
	kc.Add("myval")
	kc.Add("myval")
	kc.Add("myval2")
	kc.Flush()
	assert.Equal(t, `"myval" - 3
"myval2" - 1
`, buf.String())
	buf.Reset()
}
