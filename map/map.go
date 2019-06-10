package maps

type entry struct {
	key     int
	value   int
	isValid bool
}

func NewEntry(key, value int) *entry {
	return &entry{key: key, value: value, isValid: true}
}

type MyMap struct {
	datas    []*entry
	capacity int
}

func ToPow2(i int) int {
	if i < 1<<3 {
		return 1 << 3
	} else if i&(i-1) == 0 {
		return i
	} else {
		var j uint
		// 求最左边的1是第几位
		for ; i > 0; i >>= 1 {
			j++
		}
		return 1 << j
	}
}

func NewMap(lens int) *MyMap {
	s := make([]*entry, ToPow2(lens))
	return &MyMap{datas: s}
}

func (this *MyMap) put0(key, value int) bool {
	for index := key & (len(this.datas) - 1); ; index ++ {
		entry := this.datas[index]

		if entry == nil {
			entry = NewEntry(key, value)
			this.datas[index] = entry
			return true
		} else {
			if key == entry.key {
				entry.value = value
				entry.isValid = true
				return true
			} else {
				if !entry.isValid {
					entry.value = value
					entry.isValid = true
					return true
				} else {
					return false
				}
			}
		}
	}
	return false
}

func (this *MyMap) Put(key, value int) {
	isSuccess := this.put0(key, value)
	if isSuccess {
		this.capacity ++
	}

	if this.capacity == len(this.datas) {
		this.rehash()
	}
}

func (this *MyMap) rehash() {
	entries := this.datas
	this.datas = make([]*entry, len(this.datas)*2)
	for _, v := range entries {
		this.put0(v.key, v.value)
	}
}
func (this *MyMap) Get(key int) int {
	originIndex := key & (len(this.datas) - 1)
	for index := originIndex; ; {
		entry := this.datas[index]
		if entry == nil {
			return -1
		}

		if entry.isValid && entry.key == key {
			return entry.value
		}

		index ++

		if index == originIndex {
			return -1
		}
	}
}

func (this *MyMap) Remove(key int) {
	for index := key & (len(this.datas) - 1); ; index ++ {
		entry := this.datas[index]

		if entry == nil {
			return
		} else {
			if entry.isValid {
				if entry.key == key {
					entry.isValid = false
				}
			}
		}
	}
}

func (this *MyMap) Size() int {
	return this.capacity
}

func (this *MyMap) ContainKey(key int) bool {
	originIndex := key & (len(this.datas) - 1)
	for index := originIndex; ; {
		entry := this.datas[index]
		if entry == nil {
			return false
		}

		if entry.isValid && entry.key == key {
			return true
		}

		index ++

		if index == originIndex {
			return false
		}
	}
	return false
}
