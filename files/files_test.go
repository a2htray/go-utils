package files

import (
	"testing"
)

func TestExtension(t *testing.T) {
	if Extension("/aaa") != "" {
		t.Fatal("need to be empty")
	}

	if Extension("\\file\\a.txt") != "txt" {
		t.Fatal("need to be txt")
	}
}

func TestExtensionWithDot(t *testing.T) {
	if ExtensionWithDot("/aaa") != "" {
		t.Fatal("need to be empty")
	}

	if ExtensionWithDot("\\file\\a.txt") != ".txt" {
		t.Fatal("need to be .txt")
	}
}
