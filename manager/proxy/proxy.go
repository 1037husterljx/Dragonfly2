/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proxy

import (
	"fmt"
	"io"
	"net"
	"sync"

	logger "d7y.io/dragonfly/v2/internal/dflog"
	"d7y.io/dragonfly/v2/manager/config"
)

type Proxy interface {
	Serve() error
	Stop()
}

type redisProxy struct {
	from string
	to   string
	done chan struct{}
}

func New(cfg *config.RedisConfig) Proxy {
	return &redisProxy{
		from: fmt.Sprintf(":%d", cfg.Port),
		to:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		done: make(chan struct{}),
	}
}

func (p *redisProxy) Serve() error {
	listener, err := net.Listen("tcp", p.from)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		select {
		case <-p.done:
			return nil
		default:
			conn, err := listener.Accept()
			if err != nil {
				logger.Errorf("error accepting conn %v", err)
			} else {
				go p.handleConn(conn)
			}
		}
	}
}

func (p *redisProxy) Stop() {
	if p.done == nil {
		return
	}
	close(p.done)
	p.done = nil
}

func (p *redisProxy) handleConn(conn net.Conn) {
	logger.Debugf("handling conn from %v", conn.RemoteAddr())
	defer logger.Debugf("done handling %v", conn)
	rConn, err := net.Dial("tcp", p.to)
	if err != nil {
		logger.Debugf("failed dialing remote host %v", err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go p.copy(rConn, conn, wg)
	go p.copy(conn, rConn, wg)
	wg.Wait()
}

func (p *redisProxy) copy(from, to net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer from.Close()
	defer to.Close()
	select {
	case <-p.done:
		return
	default:
		if _, err := io.Copy(to, from); err != nil {
			logger.Debugf("copy failed%v", err)
			return
		}
	}
}
