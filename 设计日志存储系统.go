package pos

type LogSystem struct {
	ids []int
	ts  []string
}

func Constructor1() LogSystem {
	return LogSystem{
		ids: []int{},
		ts:  []string{},
	}
}
func (this *LogSystem) Put(id int, timestap string) {
	this.ids = append(this.ids, id)
	this.ts = append(this.ts, timestap)
}

var m map[string]int = map[string]int{
	"Year": 4, "Month": 7, "Day": 10, "Hour": 13, "Minute": 16, "Second": 19,
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	out := []int{}
	index := m[gra]
	for i, id := range this.ids {
		t := this.ts[i][:index]
		if t >= s[:index] && t <= e[:index] {
			out = append(out, id)
		}
	}
	return out
}
