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

package dfcache

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/pkg/errors"

	"d7y.io/dragonfly/v2/client/config"
	"d7y.io/dragonfly/v2/internal/dferrors"
	logger "d7y.io/dragonfly/v2/internal/dflog"
	"d7y.io/dragonfly/v2/pkg/basic"
	"d7y.io/dragonfly/v2/pkg/rpc/base"
	"d7y.io/dragonfly/v2/pkg/rpc/dfdaemon"
	daemonclient "d7y.io/dragonfly/v2/pkg/rpc/dfdaemon/client"
)

// Format that's used to cast the given cid to URI.
// "d7y" as scheme, and absolute path starts with "/"
const cidURIFormat = "d7y:/%s"

func newCid(cid string) string {
	return fmt.Sprintf(cidURIFormat, url.QueryEscape(cid))
}

// Stat checks if the given cache entry exists in local storage and/or in P2P network, and returns
// os.ErrNotExist if cache is not found.
func Stat(cfg *config.DfcacheConfig, client daemonclient.DaemonClient) error {
	var (
		ctx       = context.Background()
		cancel    context.CancelFunc
		statError error
	)

	if err := cfg.Validate(config.CmdStat); err != nil {
		return errors.Wrap(err, "validate stat option failed")
	}

	wLog := logger.With("Cid", cfg.Cid, "Tag", cfg.Tag)
	wLog.Info("init success and start to stat")

	if cfg.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, cfg.Timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	go func() {
		statError = statTask(ctx, client, cfg, wLog)
		cancel()
	}()

	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		return errors.Errorf("stat timeout(%s)", cfg.Timeout)
	}
	return statError
}

func statTask(ctx context.Context, client daemonclient.DaemonClient, cfg *config.DfcacheConfig, wLog *logger.SugaredLoggerOnWith) error {
	if client == nil {
		return errors.New("stat has no daemon client")
	}

	start := time.Now()
	statError := client.StatTask(ctx, newStatRequest(cfg))
	if statError == nil {
		wLog.Infof("task found in %.6f s", time.Since(start).Seconds())
		return nil
	}

	// Task not found, return os.ErrNotExist
	if dferrors.CheckError(statError, base.Code_PeerTaskNotFound) {
		return os.ErrNotExist
	}

	// Otherwise hit internal error
	wLog.Errorf("daemon stat file error: %s", statError)
	return statError
}

func newStatRequest(cfg *config.DfcacheConfig) *dfdaemon.StatTaskRequest {
	return &dfdaemon.StatTaskRequest{
		Url: newCid(cfg.Cid),
		UrlMeta: &base.UrlMeta{
			Tag: cfg.Tag,
		},
		LocalOnly: cfg.LocalOnly,
	}
}

// Import imports the given cache into P2P network.
func Import(cfg *config.DfcacheConfig, client daemonclient.DaemonClient) error {
	var (
		ctx         = context.Background()
		cancel      context.CancelFunc
		importError error
	)

	if err := cfg.Validate(config.CmdImport); err != nil {
		return errors.Wrap(err, "validate import option failed")
	}

	wLog := logger.With("Cid", cfg.Cid, "Tag", cfg.Tag, "file", cfg.Path)
	wLog.Info("init success and start to import")

	if cfg.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, cfg.Timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	go func() {
		importError = importTask(ctx, client, cfg, wLog)
		cancel()
	}()

	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		return errors.Errorf("import timeout(%s)", cfg.Timeout)
	}
	return importError
}

func importTask(ctx context.Context, client daemonclient.DaemonClient, cfg *config.DfcacheConfig, wLog *logger.SugaredLoggerOnWith) error {
	if client == nil {
		return errors.New("import has no daemon client")
	}

	start := time.Now()
	importError := client.ImportTask(ctx, newImportRequest(cfg))
	if importError != nil {
		wLog.Errorf("daemon import file error: %s", importError)
		return importError
	}

	wLog.Infof("task imported successfully in %.6f s", time.Since(start).Seconds())
	return nil
}

func newImportRequest(cfg *config.DfcacheConfig) *dfdaemon.ImportTaskRequest {
	return &dfdaemon.ImportTaskRequest{
		Type: base.TaskType_DfCache,
		Url:  newCid(cfg.Cid),
		Path: cfg.Path,
		UrlMeta: &base.UrlMeta{
			Tag: cfg.Tag,
		},
	}
}

// Export exports or downloads the given cache from P2P network, and return os.ErrNotExist if cache
// doesn't exist.
func Export(cfg *config.DfcacheConfig, client daemonclient.DaemonClient) error {
	var (
		ctx         = context.Background()
		cancel      context.CancelFunc
		exportError error
	)

	if err := cfg.Validate(config.CmdExport); err != nil {
		return errors.Wrap(err, "validate export option failed")
	}

	wLog := logger.With("Cid", cfg.Cid, "Tag", cfg.Tag, "output", cfg.Output)
	wLog.Info("init success and start to export")

	if cfg.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, cfg.Timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	go func() {
		exportError = exportTask(ctx, client, cfg, wLog)
		cancel()
	}()

	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		return errors.Errorf("export timeout(%s)", cfg.Timeout)
	}
	return exportError
}

func exportTask(ctx context.Context, client daemonclient.DaemonClient, cfg *config.DfcacheConfig, wLog *logger.SugaredLoggerOnWith) error {
	if client == nil {
		return errors.New("export has no daemon client")
	}

	start := time.Now()
	exportError := client.ExportTask(ctx, newExportRequest(cfg))
	if exportError == nil {
		wLog.Infof("task exported successfully in %.6f s", time.Since(start).Seconds())
		return nil
	}

	// Task not found, return os.ErrNotExist
	if dferrors.CheckError(exportError, base.Code_PeerTaskNotFound) {
		return os.ErrNotExist
	}

	// Otherwise hit internal error
	wLog.Errorf("daemon export file error: %s", exportError)
	return exportError
}

func newExportRequest(cfg *config.DfcacheConfig) *dfdaemon.ExportTaskRequest {
	return &dfdaemon.ExportTaskRequest{
		Url:     newCid(cfg.Cid),
		Output:  cfg.Output,
		Timeout: uint64(cfg.Timeout),
		Limit:   float64(cfg.RateLimit),
		UrlMeta: &base.UrlMeta{
			Tag: cfg.Tag,
		},
		Uid:       int64(basic.UserID),
		Gid:       int64(basic.UserGroup),
		LocalOnly: cfg.LocalOnly,
	}
}

func Delete(cfg *config.DfcacheConfig, client daemonclient.DaemonClient) error {
	var (
		ctx         = context.Background()
		cancel      context.CancelFunc
		deleteError error
	)

	if err := cfg.Validate(config.CmdDelete); err != nil {
		return errors.Wrap(err, "validate delete option failed")
	}

	wLog := logger.With("Cid", cfg.Cid, "Tag", cfg.Tag)
	wLog.Info("init success and start to delete")

	if cfg.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, cfg.Timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	go func() {
		deleteError = deleteTask(ctx, client, cfg, wLog)
		cancel()
	}()

	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		return errors.Errorf("delete timeout(%s)", cfg.Timeout)
	}
	return deleteError
}

func deleteTask(ctx context.Context, client daemonclient.DaemonClient, cfg *config.DfcacheConfig, wLog *logger.SugaredLoggerOnWith) error {
	if client == nil {
		return errors.New("delete has no daemon client")
	}

	start := time.Now()
	deleteError := client.DeleteTask(ctx, newDeleteRequest(cfg))
	if deleteError != nil {
		wLog.Errorf("daemon delete file error: %s", deleteError)
		return deleteError
	}

	wLog.Infof("task deleted successfully in %.6f s", time.Since(start).Seconds())
	return nil
}

func newDeleteRequest(cfg *config.DfcacheConfig) *dfdaemon.DeleteTaskRequest {
	return &dfdaemon.DeleteTaskRequest{
		Url: newCid(cfg.Cid),
		UrlMeta: &base.UrlMeta{
			Tag: cfg.Tag,
		},
	}
}
