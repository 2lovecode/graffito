package trait

type Comparable interface {
	Eq(target Comparable) bool
	Less(target Comparable) bool
	More(target Comparable) bool
}

type ComparableInt8 int8
type ComparableInt int

func (source ComparableInt) Eq(target Comparable) bool {
	if t, ok := target.(ComparableInt); ok {
		return source == t
	}
	return false
}

func (source ComparableInt) Less(target Comparable) bool {
	if t, ok := target.(ComparableInt); ok {
		return source < t
	}
	return false
}

func (source ComparableInt) More(target Comparable) bool {
	if t, ok := target.(ComparableInt); ok {
		return source > t
	}
	return false
}
