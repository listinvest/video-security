package services

import (
	"errors"
	"strconv"
	"net/url"
	"github.com/prohevg/video-security/logwriter"
	"github.com/prohevg/video-security/models"
	"github.com/prohevg/video-security/tasks/base"
	tmodels "github.com/prohevg/video-security/tasks/models"
	"github.com/prohevg/video-security/tasks/autosearch"
	"github.com/prohevg/video-security/tasks/manualsearch"
	"github.com/prohevg/video-security/tasks/taskdispatcher"
)

//SearchService search device in network
type SearchService struct {
	Logger *logwriter.Logger
	ITaskDispatcher taskdispatcher.ITaskDispatcher
}

//Manual search by parameters in network
func (s *SearchService) Manual(ips string, ports string) ([]models.Device, error) {
	bizTask := manualsearch.DeviceTask {
		Ips:   ips,
		Ports: ports,
		Task:  base.BizTask{
			ID:   "1",
			Name: "manual search devices",
		},
	}

	s.Logger.Debug("Run task %s with params: id=%s, ips=%s, ports=%s", bizTask.Task.Name, bizTask.Task.ID, ips, ports)

	s.ITaskDispatcher.RunTask(&bizTask)

	s.Logger.Debug("Executed task %s with params: id=%s, devices=%d", bizTask.Task.Name, bizTask.Task.ID, len(bizTask.Result.Devices))

	result := s.getDevices(bizTask.Result)

	if bizTask.Result.Result.IsError {
		return result, errors.New(bizTask.Result.Result.Error)
	}

	return result, nil
}

//Auto search all in network
func (s *SearchService) Auto() ([]models.Device, error) {
	bizTask := autosearch.DeviceTask{
		Task: base.BizTask{
			ID:   "1",
			Name: "auto search devices",
		},
	}

	s.Logger.Debug("Run task %s with params: id=%s", bizTask.Task.Name, bizTask.Task.ID)

	s.ITaskDispatcher.RunTask(&bizTask)

	s.Logger.Debug("Executed task %s with params: id=%s, devices=%d", bizTask.Task.Name, bizTask.Task.ID, len(bizTask.Result.Devices))

	result := s.getDevices(bizTask.Result)

	if bizTask.Result.Result.IsError {
		return result, errors.New(bizTask.Result.Result.Error)
	}

	return result, nil
}

//getDevices device from search
func  (s *SearchService) getDevices(taskResult tmodels.DeviceTaskResult) []models.Device {
	res := make([]models.Device, len(taskResult.Devices))
	
	for i, dev := range taskResult.Devices {
		ip, port, err := s.getIPAndPort(dev.Xaddr)
		if err != nil {
			s.Logger.Error(err)
			continue
		}

		res[i] = models.Device {
			IP: ip,
			Port: port,
		}
	}

	return res
}

//getIPAndPort ip and port
func  (s *SearchService) getIPAndPort(xaddr string) (string, int, error) {
	u, err := url.Parse(xaddr)
	if err != nil {
		return "", 0, err
	}
  
	ip := u.Hostname()
				
	port, _ := strconv.Atoi(u.Port())
	if port == 0 {
		port = 80
	}

	return ip, port, nil
}
