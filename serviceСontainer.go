package main

import (
	"os"
	"path/filepath"
	"sync"

	"videoSecurity/controllers"
	"videoSecurity/infrastructures"
	"videoSecurity/interfaces"
	"videoSecurity/logwriter"
	"videoSecurity/repositories"
	"videoSecurity/services"
	"videoSecurity/tasks/taskdispatcher"
	"videoSecurity/videostreamer"

	"github.com/syndtr/goleveldb/leveldb"
)

//IServiceContainer container
type IServiceContainer interface {
	InjectDeviceRepository() interfaces.IDeviceRepository

	InjectDeviceService() interfaces.IDeviceService
	InjectSearchService() interfaces.ISearchService

	InjectDispatcher() taskdispatcher.ITaskDispatcher

	InjectSearchController() controllers.SearchController
	InjectVideoController() controllers.VideoController
	InjectDeviceController() controllers.DeviceController	
}

//kernel
type kernel struct {
	Db     *leveldb.DB
	Logger *logwriter.Logger
}

//Inject InjectSearchService
func (k *kernel) InjectDispatcher() taskdispatcher.ITaskDispatcher {
	return &taskdispatcher.TaskDispatcher{}
}

//Inject CreateDbHandler
func (k *kernel) CreateDbHandler() interfaces.IDbHandler {
	return &infrastructures.GoLevelDbHandler{
		Db:     k.Db,
		Logger: k.Logger,
	}
}

//Inject CreateBaseRepository
func (k *kernel) CreateBaseRepository(keyPrefix string) interfaces.IBaseRepository {
	return &repositories.BaseRepository{
		k.CreateDbHandler(),
		keyPrefix,
		k.Logger,
	}
}

//Inject DeviceRepository
func (k *kernel) InjectDeviceRepository() interfaces.IDeviceRepository {
	return &repositories.DeviceRepository{
		k.CreateBaseRepository("Device_"),
	}
}

//Inject DeviceService
func (k *kernel) InjectDeviceService() interfaces.IDeviceService {
	return &services.DeviceService{
		k.Logger,
		k.InjectDeviceRepository(),
	}
}

//Inject InjectSearchService
func (k *kernel) InjectSearchService() interfaces.ISearchService {
	return &services.SearchService{
		k.Logger,
		k.InjectDispatcher(),
	}
}

//Inject SearchController
func (k *kernel) InjectSearchController() controllers.SearchController {
	return controllers.SearchController{
		k.Logger,
		k.InjectSearchService(),
	}
}

//Inject VideoController
func (k *kernel) InjectVideoController() controllers.VideoController {
	return controllers.VideoController{
		k.Logger,
		&videostreamer.VideoStreamer{
			Mutex:      &sync.RWMutex{},
			Dispatcher: &videostreamer.StreamerDispatcher{},
		},
	}
}


//Inject DeviceController
func (k *kernel) InjectDeviceController() controllers.DeviceController {
	return controllers.DeviceController{
		k.Logger,
		k.InjectDeviceService(),
	}
}


var (
	k             *kernel
	containerOnce sync.Once
)

//ServiceContainer create container
func ServiceContainer(logger *logwriter.Logger, db *leveldb.DB) IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{
				Logger: logger,
				Db:     db,
			}
		})
	}
	return k
}

//getCurrentPath get current path of currently running exe
func getCurrentPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
