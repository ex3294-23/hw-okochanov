package unpkr

import "testing"

func TestUnpack001(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = ps.Unpack()

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack002(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = "abcd"
	want = "abcd"
	got = ps.Unpack()

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack003(t *testing.T) {
        var ps PackedString
        var got, want string

        ps = "3abc"
        want = ""
        got = ps.Unpack()

        if got != want {
                t.Errorf("ps.Unpack() == %q, want %q", got, want)
        }
}

func TestUnpack004(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = "45"
	want = ""
	got = ps.Unpack()

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack005(t *testing.T) {
        var ps PackedString
        var got, want string

        ps = "aaa10b"
        want = ""
        got = ps.Unpack()

        if got != want {
                t.Errorf("ps.Unpack() == %q, want %q", got, want)
        }
}

func TestUnpack006(t *testing.T) {
        var ps PackedString
        var got, want string

        ps = "aaa0b"
        want = "aab"
        got = ps.Unpack()

        if got != want {
                t.Errorf("ps.Unpack() == %q, want %q", got, want)
        }
}

func TestUnpack007(t *testing.T) {
        var ps PackedString
        var got, want string

        ps = " "
        want = " "
        got = ps.Unpack()

        if got != want {
                t.Errorf("ps.Unpack() == %q, want %q", got, want)
        }
}

func TestUnpack008(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = `d\n5abc`
	want = `d\n\n\n\n\nabc`
	got = ps.Unpack()

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack009(t *testing.T) {
	var ps PackedString
	var got, want string

	ps = `qwe\4\5`
	want = `qwe45`
	got = ps.Unpack()

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack010(t *testing.T) {
        var ps PackedString
        var got, want string

        ps = `qwe\45`
        want = `qwe44444`
        got = ps.Unpack()

        if got != want {
                t.Errorf("ps.Unpack() == %q, want %q", got, want)
        }
}

func TestUnpack011(t *testing.T) {
        var ps PackedString
        var got, want string

        ps = `qwe\\5`
        want = `qwe\\\\\`
        got = ps.Unpack()

        if got != want {
                t.Errorf("ps.Unpack() == %q, want %q", got, want)
        }
}

func TestUnpack012(t *testing.T) {
        var ps PackedString
        var got, want string

        ps = `qw\ne`
	want = " "
        got = ps.Unpack()

        if got != want {
                t.Errorf("ps.Unpack() == %q, want %q", got, want)
        }
}

