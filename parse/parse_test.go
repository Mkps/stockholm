package parse

import (
	"testing"
	"crypto/sha256"
	"bytes"
)


func TestParserNormal(t *testing.T) {
		testing_key := "mysecretkeyisverylongandbig"
    args := []string{"program_name", testing_key}
    opt, err := Parser(args)

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

		hash := sha256.Sum256([]byte(testing_key))
		test_hash := hash[:]
		if !bytes.Equal(opt.Key, test_hash) {
			t.Errorf("expected hash [%s] got [%s]", test_hash, opt.Key)
    }

    if opt.Reverse {
        t.Error("expected reverse=false")
    }
}

func TestParserReverse(t *testing.T) {
		testing_key := "mysecretkeyisverylongandbig"
    args := []string{"program_name", "-r", testing_key}
    opt, err := Parser(args)

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if !opt.Reverse {
        t.Error("expected reverse=true")
    }

		hash := sha256.Sum256([]byte(testing_key))
		test_hash := hash[:]
		if !bytes.Equal(opt.Key, test_hash) {
			t.Errorf("expected hash [%s] got [%s]", test_hash, opt.Key)
    }
}

func TestParserSilent(t *testing.T) {
		testing_key := "mysecretkeyisverylongandbig"
    args := []string{"program_name", testing_key, "-s"}
    opt, err := Parser(args)

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if !opt.Silent {
        t.Error("expected silent=true")
    }

		hash := sha256.Sum256([]byte(testing_key))
		test_hash := hash[:]
		if !bytes.Equal(opt.Key, test_hash) {
			t.Errorf("expected hash [%s] got [%s]", test_hash, opt.Key)
    }
}

func TestParserMissingKey(t *testing.T) {
    _, err := Parser([]string{})
    if err == nil {
        t.Error("expected an error for missing key")
    }
}

func TestParserKeyTooSmall(t *testing.T) {
		testing_key := "smallkey"
    args := []string{"program_name", testing_key}
    _, err := Parser(args)
    if err == nil {
        t.Error("expected an error for key too small")
    }
}
