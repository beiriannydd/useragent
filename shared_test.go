package useragent

import (
	"testing"
)

// TestRun tests the run function really basically
// a more thorough test would ensure split by line
// and test error response
func TestRun(t *testing.T) {
	result, err := run("bash", "-c", "echo hello world")
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	if len(result) != 1 {
		t.Errorf("expected 1 line, got %d", len(result))
		return
	}
	t.Logf("result: %+v", result)
}

// TestWindowsVersion tests the parsing of the version string for Windows
func TestWindowsVersion(t *testing.T) {
	tests := map[string]string{
		"Microsoft Windows [Version 10.0.43544.0]": "Windows NT 10.0",
		"95.2": "Windows 95.2",
	}
	for in := range tests {
		t.Run(in, func(t *testing.T) {
			if parseWindowsVersion(in) != tests[in] {
				t.Errorf("%s generated %s, not %s", in, parseWindowsVersion(in), tests[in])
			}
		})
	}
}

// TestString just outputs a string for visual confirmation
func TestString(t *testing.T) {
	t.Logf("%s\n", DefaultUserAgent.String())
}
