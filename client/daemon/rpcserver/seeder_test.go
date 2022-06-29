/*
 *     Copyright 2022 The Dragonfly Authors
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

package rpcserver

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phayes/freeport"
	testifyassert "github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"d7y.io/dragonfly/v2/client/clientutil"
	"d7y.io/dragonfly/v2/client/config"
	"d7y.io/dragonfly/v2/client/daemon/peer"
	"d7y.io/dragonfly/v2/client/daemon/storage"
	"d7y.io/dragonfly/v2/client/daemon/storage/mocks"
	"d7y.io/dragonfly/v2/pkg/dfnet"
	"d7y.io/dragonfly/v2/pkg/rpc/base"
	"d7y.io/dragonfly/v2/pkg/rpc/base/common"
	"d7y.io/dragonfly/v2/pkg/rpc/cdnsystem"
	cdnclient "d7y.io/dragonfly/v2/pkg/rpc/cdnsystem/client"
	dfdaemonserver "d7y.io/dragonfly/v2/pkg/rpc/dfdaemon/server"
	"d7y.io/dragonfly/v2/pkg/rpc/scheduler"
)

func Test_ObtainSeeds(t *testing.T) {
	assert := testifyassert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		pieceSize = uint32(1024)
	)

	type pieceRange struct {
		start int
		end   int
	}
	var tests = []struct {
		name            string
		existTaskID     string       // test for non-exists task
		existPieces     []pieceRange // already exist pieces in storage
		followingPieces []pieceRange // following pieces in running task subscribe channel
		limit           uint32
		totalPieces     uint32
		success         bool
		verify          func(t *testing.T, assert *testifyassert.Assertions)
	}{
		{
			name: "already exists in storage",
			existPieces: []pieceRange{
				{
					start: 0,
					end:   10,
				},
			},
			totalPieces: 11,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
		{
			name: "already exists in storage with extra get piece request",
			existPieces: []pieceRange{
				{
					start: 0,
					end:   10,
				},
			},
			totalPieces: 11,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
		{
			name: "already exists in storage - large",
			existPieces: []pieceRange{
				{
					start: 0,
					end:   1000,
				},
			},
			totalPieces: 1001,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
		{
			name: "already exists in storage - large with extra get piece request",
			existPieces: []pieceRange{
				{
					start: 0,
					end:   1000,
				},
			},
			totalPieces: 1001,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
		{
			name: "partial exists in storage",
			existPieces: []pieceRange{
				{
					start: 0,
					end:   10,
				},
			},
			followingPieces: []pieceRange{
				{
					start: 11,
					end:   20,
				},
			},
			totalPieces: 21,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
		{
			name: "partial exists in storage - large",
			existPieces: []pieceRange{
				{
					start: 0,
					end:   1000,
				},
			},
			followingPieces: []pieceRange{
				{
					start: 1001,
					end:   2000,
				},
			},
			totalPieces: 2001,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
		{
			name: "not exists in storage",
			followingPieces: []pieceRange{
				{
					start: 0,
					end:   20,
				},
			},
			totalPieces: 21,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
		{
			name: "not exists in storage - large",
			followingPieces: []pieceRange{
				{
					start: 0,
					end:   2000,
				},
			},
			totalPieces: 2001,
			success:     true,
			verify: func(t *testing.T, assert *testifyassert.Assertions) {
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, delay := range []bool{false, true} {
				delay := delay
				mockStorageManger := mocks.NewMockManager(ctrl)

				if tc.limit == 0 {
					tc.limit = 1024
				}

				var (
					totalPieces []*base.PieceInfo
					lock        sync.Mutex
				)

				var addedPieces = make(map[uint32]*base.PieceInfo)
				for _, p := range tc.existPieces {
					if p.end == 0 {
						p.end = p.start
					}
					for i := p.start; i <= p.end; i++ {
						if _, ok := addedPieces[uint32(i)]; ok {
							continue
						}
						piece := &base.PieceInfo{
							PieceNum:    int32(i),
							RangeStart:  uint64(i) * uint64(pieceSize),
							RangeSize:   pieceSize,
							PieceOffset: uint64(i) * uint64(pieceSize),
							PieceStyle:  base.PieceStyle_PLAIN,
						}
						totalPieces = append(totalPieces, piece)
						addedPieces[uint32(i)] = piece
					}
				}

				mockStorageManger.EXPECT().GetPieces(gomock.Any(),
					gomock.Any()).AnyTimes().DoAndReturn(
					func(ctx context.Context, req *base.PieceTaskRequest) (*base.PiecePacket, error) {
						var pieces []*base.PieceInfo
						lock.Lock()
						for i := req.StartNum; i < tc.totalPieces; i++ {
							if piece, ok := addedPieces[i]; ok {
								if piece.PieceNum >= int32(req.StartNum) && len(pieces) < int(req.Limit) {
									pieces = append(pieces, piece)
								}
							}
						}
						lock.Unlock()
						return &base.PiecePacket{
							TaskId:        req.TaskId,
							DstPid:        req.DstPid,
							DstAddr:       "",
							PieceInfos:    pieces,
							TotalPiece:    int32(tc.totalPieces),
							ContentLength: int64(tc.totalPieces) * int64(pieceSize),
							PieceMd5Sign:  "",
						}, nil
					})
				mockStorageManger.EXPECT().GetExtendAttribute(gomock.Any(),
					gomock.Any()).AnyTimes().DoAndReturn(
					func(ctx context.Context, req *storage.PeerTaskMetadata) (*base.ExtendAttribute, error) {
						return &base.ExtendAttribute{
							Header: map[string]string{
								"Test": "test",
							},
						}, nil
					})
				mockTaskManager := peer.NewMockTaskManager(ctrl)
				mockTaskManager.EXPECT().StartSeedTask(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx context.Context, req *peer.SeedTaskRequest) (*peer.SeedTaskResponse, bool, error) {
						ch := make(chan *peer.PieceInfo)
						success := make(chan struct{})
						fail := make(chan struct{})

						go func(followingPieces []pieceRange) {
							for i, p := range followingPieces {
								if p.end == 0 {
									p.end = p.start
								}
								for j := p.start; j <= p.end; j++ {
									lock.Lock()
									if _, ok := addedPieces[uint32(j)]; ok {
										continue
									}
									piece := &base.PieceInfo{
										PieceNum:    int32(j),
										RangeStart:  uint64(j) * uint64(pieceSize),
										RangeSize:   pieceSize,
										PieceOffset: uint64(j) * uint64(pieceSize),
										PieceStyle:  base.PieceStyle_PLAIN,
									}
									totalPieces = append(totalPieces, piece)
									addedPieces[uint32(j)] = piece
									lock.Unlock()

									var finished bool
									if i == len(followingPieces)-1 && j == p.end {
										finished = true
									}
									if !delay {
										ch <- &peer.PieceInfo{
											Num:      int32(j),
											Finished: finished,
										}
									}
								}
							}
							close(success)
						}(tc.followingPieces)

						tracer := otel.Tracer("test")
						ctx, span := tracer.Start(ctx, config.SpanSeedTask, trace.WithSpanKind(trace.SpanKindClient))
						return &peer.SeedTaskResponse{
							SubscribeResponse: peer.SubscribeResponse{
								Storage:          mockStorageManger,
								PieceInfoChannel: ch,
								Success:          success,
								Fail:             fail,
							},
							Context: ctx,
							Span:    span,
							TaskID:  "fake-task-id",
						}, false, nil
					})

				s := &server{
					KeepAlive:       clientutil.NewKeepAlive("test"),
					peerHost:        &scheduler.PeerHost{},
					storageManager:  mockStorageManger,
					peerTaskManager: mockTaskManager,
				}
				sd := &seeder{server: s}

				_, client := setupSeederServerAndClient(t, s, sd, assert, s.ServePeer)

				pps, err := client.ObtainSeeds(
					context.Background(),
					&cdnsystem.SeedRequest{
						TaskId:  "fake-task-id",
						Url:     "http://localhost/path/to/file",
						UrlMeta: nil,
					})
				assert.Nil(err, "client obtain seeds grpc call should be ok")

				var (
					total  = make(map[int32]bool)
					maxNum int32
				)

				for {
					p, err := pps.Recv()
					if err == io.EOF {
						break
					}
					if p.PieceInfo.PieceNum == common.BeginOfPiece {
						continue
					}
					total[p.PieceInfo.PieceNum] = true
					if p.PieceInfo.PieceNum >= maxNum {
						maxNum = p.PieceInfo.PieceNum
					}
					if tc.success {
						assert.Nil(err, "receive seed info should be ok")
					}
				}
				if tc.success {
					assert.Equal(int(maxNum+1), len(total))
				}
				s.peerServer.GracefulStop()
			}

		})
	}
}

func setupSeederServerAndClient(t *testing.T, srv *server, sd *seeder, assert *testifyassert.Assertions, serveFunc func(listener net.Listener) error) (int, cdnclient.CdnClient) {
	srv.peerServer = dfdaemonserver.New(srv)
	cdnsystem.RegisterSeederServer(srv.peerServer, sd)

	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	assert.Nil(err, "get free port should be ok")
	go func() {
		if err := serveFunc(ln); err != nil {
			t.Error(err)
		}
	}()

	client := cdnclient.GetClientByAddr([]dfnet.NetAddr{
		{
			Type: dfnet.TCP,
			Addr: fmt.Sprintf(":%d", port),
		},
	})
	return port, client
}
