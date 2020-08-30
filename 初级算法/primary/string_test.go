package primary

import "testing"

func TestReverse(t *testing.T) {
	s := []byte("12345678")
	reverseString(s)
	t.Logf("%s", s)
}

func TestReverseInt(t *testing.T) {
	t.Logf("%d", reverse(-12344))
	t.Logf("%d", reverse(12344))
	t.Logf("%d", reverse(92008859454349999))

}
