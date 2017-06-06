package addressline

import "testing"

func TestSimpleAddress(t *testing.T) {
	tables := []struct {
		val string
		res string
	}{
		{"Winterallee 3", "Winterallee,3"},
		{"Musterstrasse 45", "Musterstrasse,45"},
		{"Blaufeldweg 123B", "Blaufeldweg,123B"},
		{"Klosterstraße 62", "Klosterstraße,62"},
	}

	for _, table := range tables {
		result := Parse(table.val)
		if result != table.res {
			t.Errorf("Incorrect Output, got: %d, expected: %d", result, table.res)
		}

	}
}

func TestMoreComplicatedAddress(t *testing.T) {
	tables := []struct {
		val string
		res string
	}{
		{"Am Bächle 23", "Am Bächle,23"},
		{"Auf der Vogelwiese 23 b", "Auf der Vogelwiese,23 b"},
	}

	for _, table := range tables {
		result := Parse(table.val)
		if result != table.res {
			t.Errorf("Incorrect Output, got: %d, expected: %d", result, table.res)
		}

	}
}

func TestComplexAddress(t *testing.T) {
	tables := []struct {
		val string
		res string
	}{
		{"Calle Aduana, 29", "Calle Aduana,29"},
	}

	for _, table := range tables {
		result := Parse(table.val)
		if result != table.res {
			t.Errorf("Incorrect Output, got: %v, expected: %v", result, table.res)
		}

	}
}
