package system

import (
	"goinaction/k8sapp/logger"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

type SignalType int

const (
	Shutdown SignalType = iota
	Reload
	Maintenance
)

func (s SignalType) String() string {
	switch s {
	case Shutdown:
		return "SHUTDOWN"
	case Reload:
		return "RELOAD"
	case Maintenance:
		return "MAINTENANCE"
	default:
		return strconv.Itoa(int(s))
	}
}

type Signals struct {
	mutex sync.RWMutex

	interrupt chan os.Signal
	quit      chan struct{}

	shutdown    []os.Signal
	reload      []os.Signal
	maintenance []os.Signal
}

func NewSignals() *Signals {
	signals := &Signals{
		interrupt: make(chan os.Signal, 1),
		quit:      make(chan struct{}, 1),

		shutdown:    []os.Signal{syscall.SIGINT, syscall.SIGTERM},
		reload:      []os.Signal{syscall.SIGHUP},
		maintenance: []os.Signal{syscall.SIGUSR1},
	}
	signal.Notify(signals.interrupt)
	return signals
}

func (s *Signals) Get(sigType SignalType) (signals []os.Signal) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	switch sigType {
	case Shutdown:
		signals = make([]os.Signal, len(s.shutdown))
		copy(signals, s.shutdown)
	case Reload:
		signals = make([]os.Signal, len(s.reload))
		copy(signals, s.reload)
	case Maintenance:
		signals = make([]os.Signal, len(s.maintenance))
		copy(signals, s.maintenance)
	}

	return
}

func (s *Signals) Add(sig os.Signal, sigType SignalType) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	switch sigType {
	case Shutdown:
		s.shutdown = append(s.shutdown, sig)
	case Reload:
		s.reload = append(s.reload, sig)
	case Maintenance:
		s.maintenance = append(s.maintenance, sig)
	}
}

func (s *Signals) Remove(sig os.Signal, sigType SignalType) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	switch sigType {
	case Shutdown:
		s.shutdown = removeSignal(sig, s.shutdown)
	case Reload:
		s.reload = removeSignal(sig, s.reload)
	case Maintenance:
		s.maintenance = removeSignal(sig, s.maintenance)
	}
}

func (s *Signals) Wait(logger logger.Logger, operator Operator) error {
	for {
		select {
		case <-s.quit:
			logger.Info("Gracefully closed")
			return nil
		case sig := <-s.interrupt:
			s.mutex.RLock()
			logger.Info("Got signal: %s", sig)
			switch {
			case IsSignalAvailable(sig, s.maintenance):
				s.mutex.RUnlock()
				logger.Info("Maintenance request")
				err := operator.Maintenance()
				if err != nil {
					logger.Error(err)
				}
			case IsSignalAvailable(sig, s.reload):
				s.mutex.RUnlock()
				logger.Info("Reloading configuration...")
				err := operator.Reload()
				if err != nil {
					logger.Error(err)
				}
			case IsSignalAvailable(sig, s.shutdown):
				s.mutex.RUnlock()
				logger.Info("Service was terminated by system signal")
				err := operator.Shutdown()
				if err != nil {
					logger.Error(err)
				}
				s.quit <- struct{}{}
			}
		}
	}
}

func IsSignalAvailable(signal os.Signal, list []os.Signal) bool {
	for _, s := range list {
		if s == signal {
			return true
		}
	}
	return false
}

func removeSignal(signal os.Signal, list []os.Signal) (signals []os.Signal) {
	for ind, sig := range list {
		if sig == signal {
			signals = append(list[:ind], list[ind+1:]...)
			return
		}
	}
	return list
}
