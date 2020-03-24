package experiment

//DFA AC 中文分词

type DFATree struct {
	First *CharEntry
	Current *CharEntry
}

func NewDFATree() *DFATree {
	char := &CharEntry{}
	return &DFATree{
		First:   char,
		Current: char,
	}
}

//func (mine *DFATree) Add(s []rune) {
//	charLen := len(s)
//	for i := 0; i < charLen; i++ {
//		char := &CharEntry{
//			Value:   s[i],
//			IsEnd:   i == (charLen - 1),
//			NextLen: 0,
//			Next:    nil,
//		}
//	}
//}
//
//func (mine *DFATree) Find(s rune) bool {
//
//}

type CharEntry struct {
	Value rune
	IsEnd bool
	NextLen int
	Next []*CharEntry
}

func ParticipleRun() {
	//dict := []string{"is", "wrong", "file", "filepath", "path"}
	//
	////有穷自动机算法，实现分词
	//
	//
	//str := "filepathiswrong"


}
