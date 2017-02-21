package runeio

// RuneReadUnreader is the underlying interface Reader will use.
type RuneReadUnreader interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
	String() string
}

type Reader struct {
	rr    RuneReadUnreader
	Runes []rune
}

func NewRuneio(r RuneReadUnreader) *Reader {
	return &Reader{r, []rune{}}
}

func (r *Reader) Discard(n uint) (discarded uint, err error) {
	runes, err := r.ReadRunes(n)
	return uint(len(runes)), err
}

func (r *Reader) ReadRunes(n uint) (runes []rune, err error) {
	if err = r.readFromReader(n); err != nil {
		n = uint(len(r.Runes))
	}

	runes = r.Runes[0:n]
	r.Runes = r.Runes[n:]

	return runes, err
}

func (r *Reader) PeekRunes(n uint) (runes []rune, err error) {
	if err := r.readFromReader(n); err != nil {
		return r.Runes, err
	}

	return r.Runes[0:n], nil
}

func (r *Reader) String() string {
	return string(r.Runes) + r.rr.String()
}

func (r *Reader) Reset(bufReader RuneReadUnreader) {
	r.rr = bufReader
}

func (r *Reader) readFromReader(n uint) error {
	l := int(n) - len(r.Runes)

	// check if we've already read enough runes
	if l <= 0 {
		return nil
	}

	// if not, read the remaining amount of runes
	for i := 0; i < l; i++ {
		ru, _, err := r.rr.ReadRune()
		if err != nil {
			return err
		}
		r.Runes = append(r.Runes, ru)
	}

	return nil
}
