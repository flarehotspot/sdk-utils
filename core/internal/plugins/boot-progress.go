package plugins

import (
	"log"
	"sync"
	"sync/atomic"

	sse "core/internal/utils/sse"
)

type BootProgData struct {
	Status string `json:"status"`
	Done   bool   `json:"done"`
}

type BootProgress struct {
	mu      sync.RWMutex
	status  string
	done    atomic.Bool
	sockets []*sse.SseSocket
	DONE_C  chan error // should only be used once
}

func NewBootProgress() *BootProgress {
	return &BootProgress{
		sockets: []*sse.SseSocket{},
		DONE_C:  make(chan error),
	}
}

func (bp *BootProgress) SetStatus(s string) {
	go func() {
		bp.mu.Lock()
		defer bp.mu.Unlock()
		bp.status = s
		bp.emit()
	}()
}

func (bp *BootProgress) Status() string {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return bp.status
}

func (bp *BootProgress) IsDone() bool {
	return bp.done.Load()
}

func (bp *BootProgress) SetDone(err error) {
	log.Println("Setting boot progress to done...")
	bp.done.Store(true)

	bp.mu.Lock()
	if err != nil {
		bp.status = err.Error()
	}
	bp.emit()
	bp.mu.Unlock()

	bp.DONE_C <- err
}

func (bp *BootProgress) AddSocket(s *sse.SseSocket) {
	bp.mu.Lock()
	defer bp.mu.Unlock()
	bp.sockets = append(bp.sockets, s)

	log.Println("Socket added to boot progress: ", s.Id())

	go func() {
		<-s.Done()
		// remove socket if disconnected
		bp.mu.Lock()
		defer bp.mu.Unlock()

		sockets := []*sse.SseSocket{}
		for _, ss := range sockets {
			if s.Id() != ss.Id() {
				sockets = append(sockets, s)
			}
		}

		bp.sockets = sockets

		log.Println("Socket removed from boot progress: ", s.Id())
	}()
}

func (bp *BootProgress) emit() {
	data := BootProgData{bp.status, bp.done.Load()}
	log.Println(data)

	for _, s := range bp.sockets {
		err := s.Emit("boot:progress", data)
		if err != nil {
			log.Println("Boot Progress socket error:", err)
		}
	}
}
