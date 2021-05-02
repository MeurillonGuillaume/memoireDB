package datastore

import (
	"encoding/gob"
	"fmt"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

type (
	persistedDatastore struct {
		memStore *memoryDatastore
		closing  *atomic.Bool
		wg       sync.WaitGroup
		cfg      *PersistenceConfig
	}

	PersistenceConfig struct {
		Directory string `required:"true" flagUsage:"Which directory should be used to store database snapshots"`
		Retention int    `default:"100" flagUsage:"After how many snapshots should a certain snapshot become outdated"`
	}
)

var _ Store = (*persistedDatastore)(nil)

func newPersistedDatastore() (s Store, err error) {
	var cfg PersistenceConfig
	configLoader := multiconfig.New()
	if err = configLoader.Load(&cfg); err != nil {
		return
	}

	if err = configLoader.Validate(&cfg); err != nil {
		return
	}

	if cfg.Retention < 1 {
		err = fmt.Errorf("expected a positive retention value of at least 1")
		return
	}

	ds := persistedDatastore{
		cfg:      &cfg,
		closing:  atomic.NewBool(false),
		wg:       sync.WaitGroup{},
		memStore: newMemoryDatastore().(*memoryDatastore),
	}
	ds.restoreFromLatestSnapshot()

	s = &ds
	return
}

func (pd *persistedDatastore) LoadKeyValue(m model.RetrieveModel) (interface{}, error) {
	if pd.closing.Load() {
		return nil, ErrClosing
	}

	pd.wg.Add(1)
	go pd.takeSnapshot()
	defer pd.wg.Done()

	return pd.memStore.LoadKeyValue(m)
}

func (pd *persistedDatastore) StoreKeyValue(m model.InsertModel) (interface{}, error) {
	if pd.closing.Load() {
		return nil, ErrClosing
	}

	pd.wg.Add(1)
	go pd.takeSnapshot()
	defer pd.wg.Done()

	return pd.memStore.StoreKeyValue(m)
}

func (pd *persistedDatastore) ListKeys(m model.ListKeysModel) ([]string, error) {
	if pd.closing.Load() {
		return nil, ErrClosing
	}

	pd.wg.Add(1)
	defer pd.wg.Done()
	return pd.memStore.ListKeys(m)
}

func (pd *persistedDatastore) Close() error {
	if !pd.closing.Load() {
		pd.closing.Store(true)
	}

	pd.wg.Wait()
	return nil
}

func (pd *persistedDatastore) restoreFromLatestSnapshot() {}

func (pd *persistedDatastore) takeSnapshot() {
	pd.wg.Add(1)
	defer pd.wg.Done()

	pd.memStore.mux.Lock()
	snapshot := pd.memStore.store
	pd.memStore.mux.Unlock()

	snapPath := path.Join(
		pd.cfg.Directory,
		"snap_"+strconv.FormatInt(time.Now().UnixNano(), 10)+_snapExt,
	)

	snapfile, err := os.Create(snapPath)
	if err != nil {
		logrus.WithError(err).Error("Could not open snapshot file, no snapshot taken")
		return
	}

	defer func() {
		if err := snapfile.Close(); err != nil {
			logrus.WithError(err).Errorf("Could not properly close snapshot file %s, it might be corrupted", snapfile.Name())
		}
	}()

	encoder := gob.NewEncoder(snapfile)
	if err := encoder.Encode(&snapshot); err != nil {
		logrus.WithError(err).Error("Could not encode state to snapshot")
	} else {
		logrus.Infof("Successfully created snapshot %s", snapfile.Name())
	}
}
