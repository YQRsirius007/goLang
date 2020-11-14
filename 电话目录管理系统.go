package pos

type PhoneDirectory struct {
	Phones []int
	Used   []bool
}

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor3(maxNumbers int) PhoneDirectory {
	phone := make([]int, maxNumbers)
	for i := range phone {
		phone[i] = i
	}
	used := make([]bool, maxNumbers)
	return PhoneDirectory{
		Phones: phone,
		Used:   used,
	}
}

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
	for i, v := range this.Used {
		if !v {
			this.Used[i] = true
			return this.Phones[i]
		}
	}
	return -1
}

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
	for i, v := range this.Phones {
		if v == number {
			if this.Used[i] {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int) {
	for i, v := range this.Phones {
		if v == number {
			this.Used[i] = false
		}
	}
}
