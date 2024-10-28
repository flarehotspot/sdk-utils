package plugins

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	sse "core/internal/utils/sse"
)

type BootProgData struct {
	Logs []string `json:"logs"`
	Done bool     `json:"done"`
}

type BootProgress struct {
	mu      sync.RWMutex
	logs    []string
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

func (bp *BootProgress) AppendLog(s string) {
	go func() {
		bp.mu.Lock()
		defer bp.mu.Unlock()
		// add timestamp to log
		s = fmt.Sprintf("%s %s", time.Now().Format("2006-01-02 15:04:05"), s)
		bp.logs = append(bp.logs, s)
		bp.emit()
	}()
}

func (bp *BootProgress) Logs() []string {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return bp.logs
}

func (bp *BootProgress) IsDone() bool {
	return bp.done.Load()
}

func (bp *BootProgress) Done(err error) {
	log.Println("Setting boot progress to done...")
	bp.done.Store(true)

	bp.mu.Lock()
	if err != nil {
		bp.logs = append(bp.logs, err.Error())
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
	data := BootProgData{bp.logs, bp.done.Load()}
	// log.Println(data)

	for _, s := range bp.sockets {
		err := s.Emit("boot:progress", data)
		if err != nil {
			log.Println("Boot Progress socket error:", err)
		}
	}
}
