package sdksse

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type SseSocket struct {
	mu      sync.RWMutex
	id      string
	res     http.ResponseWriter
	req     *http.Request
	flusher http.Flusher
	msgId   int32
	msgCh   chan SseData
}

type SseData struct {
	MsgType string
	Data    []byte
}

func (s *SseSocket) Id() string {
	return s.id
}

func (s *SseSocket) Emit(t string, jsonData interface{}) (err error) {
	bytes, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	s.msgCh <- SseData{t, bytes}
	return nil
}

func (s *SseSocket) Done() <-chan struct{} {
	return s.req.Context().Done()
}

func (s *SseSocket) Flush() {
	s.flusher.Flush()
}

func (s *SseSocket) Listen() {
	for {
		select {
		case d := <-s.msgCh:
			data := string(d.Data)
			payload := fmt.Sprintf("id: %d\nevent: %s\ndata: %s\n\n", s.msgId, d.MsgType, data)
			log.Println("Socket data:", payload)
			fmt.Fprint(s.res, payload)
			s.Flush()
			s.msgId += 1
		case <-s.Done():
			return
		}
	}
}

func RespondError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, err)
}

func NewSocket(w http.ResponseWriter, r *http.Request) (s *SseSocket, err error) {
	f, ok := w.(http.Flusher)
	if !ok {
		log.Println("Streaming not supported in path: ", r.URL.Path)
		err = errors.New("streaming not supported")
		return nil, err
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	f.Flush()

	id := strconv.Itoa(int(rand.Int()))

	return &SseSocket{
		id:      id,
		res:     w,
		req:     r,
		msgCh:   make(chan SseData),
		flusher: f,
	}, nil
}
