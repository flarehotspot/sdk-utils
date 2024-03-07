/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sse

import (
	"log"
	"sync"
)

var (
	mu            = sync.RWMutex{}
	deviceSockets = map[string][]*SseSocket{}
)

func AddSocket(key string, socket *SseSocket) {
	mu.Lock()
	defer mu.Unlock()

	_, ok := deviceSockets[key]
	if !ok {
		deviceSockets[key] = []*SseSocket{}
	}

	deviceSockets[key] = append(deviceSockets[key], socket)
	log.Printf("Socket %s attached to \"%s\"...", socket.Id(), key)

	go func() {
		<-socket.Done()
		RemoveSocket(key, socket)
	}()
}

func RemoveSocket(key string, socket *SseSocket) {
	mu.Lock()
	defer mu.Unlock()

	_, ok := deviceSockets[key]
	if ok {
		sockets := deviceSockets[key]
		for i, s := range sockets {
			if socket.Id() == s.Id() {
				sockets[i] = sockets[len(sockets)-1]
				sockets = sockets[:len(sockets)-1]
				if len(sockets) == 0 {
					delete(deviceSockets, key)
				} else {
					deviceSockets[key] = sockets
				}
				log.Printf("Socket %s removed from device %s...", socket.Id(), key)
				break
			}
		}
	}
}

func Emit(key string, event string, d interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	sockets, ok := deviceSockets[key]
	if ok {
		for _, s := range sockets {
			s.Emit(event, d)
		}
	}
}

func Broadcast(event string, d interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	for _, sockets := range deviceSockets {
		for _, s := range sockets {
			s.Emit(event, d)
		}
	}
}
