package server

func (gb *Goroutine) Submit(task func()) {
	pool.Submit(task)
}
