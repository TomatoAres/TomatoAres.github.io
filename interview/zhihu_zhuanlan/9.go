package main

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch) /* close 之后就不能正常使用了，返回没用*/
		set.RUnlock()

	}()
	return ch
}
