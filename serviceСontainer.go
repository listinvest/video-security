package main

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/prohevg/video-security/controllers"
	"github.com/prohevg/video-security/deviceonvif"
	"github.com/prohevg/video-security/infrastructures"
	"github.com/prohevg/video-security/interfaces"
	"github.com/prohevg/video-security/logwriter"
	"github.com/prohevg/video-security/repositories"
	"github.com/prohevg/video-security/services"
	"github.com/prohevg/video-security/tasks/taskdispatcher"
	"github.com/prohevg/video-security/videostreamer"

	"github.com/syndtr/goleveldb/leveldb"
)

//IServiceContainer container
type IServiceContainer interface {
	InjectDeviceRepository() interfaces.IDeviceRepository
	InjectDeviceAuthRepository() interfaces.IDeviceAuthRepository

	InjectDeviceService() interfaces.IDeviceService
	InjectDeviceAuthService() interfaces.IDeviceAuthService
	InjectSearchService() interfaces.ISearchService

	InjectDispatcher() taskdispatcher.ITaskDispatcher
	InjectDeviceOnvif() deviceonvif.IDeviceOnvif

	InjectSearchController() controllers.SearchController
	InjectVideoController() controllers.VideoController
	InjectDeviceController() controllers.DeviceController
	InjectDeviceAuthController() controllers.DeviceAuthController
}

//kernel
type kernel struct {
	Db     *leveldb.DB
	Logger *logwriter.Logger
}

//Inject InjectDispatcher
func (k *kernel) InjectDispatcher() taskdispatcher.ITaskDispatcher {
	return &taskdispatcher.TaskDispatcher{}
}

//Inject InjectDeviceOnvif
func (k *kernel) InjectDeviceOnvif() deviceonvif.IDeviceOnvif {
	return &deviceonvif.DeviceOnvif{}
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

//Inject DeviceAuthRepository
func (k *kernel) InjectDeviceAuthRepository() interfaces.IDeviceAuthRepository {
	return &repositories.DeviceAuthRepository{
		k.CreateBaseRepository("DeviceAuth_"),
	}
}

//Inject DeviceService
func (k *kernel) InjectDeviceService() interfaces.IDeviceService {
	return &services.DeviceService{
		k.Logger,
		k.InjectDeviceRepository(),
		k.InjectDeviceAuthRepository(),
		k.InjectDeviceOnvif(),
	}
}

//Inject DeviceAuthService
func (k *kernel) InjectDeviceAuthService() interfaces.IDeviceAuthService {
	return &services.DeviceAuthService{
		k.Logger,
		k.InjectDeviceAuthRepository(),
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

//Inject DeviceAuthController
func (k *kernel) InjectDeviceAuthController() controllers.DeviceAuthController {
	return controllers.DeviceAuthController{
		k.Logger,
		k.InjectDeviceAuthService(),
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
