package deployer

import (
	"log/slog"
	"sync"
)

// Service заглушка для представления внешнего сервиса deployer.
// При необходимости легко заменяется на любой сервис
type Service struct {
	pods map[string]struct{}
	mu   sync.RWMutex
}

func NewService() *Service {
	return &Service{
		pods: make(map[string]struct{}),
	}
}

// CreatePod метод для создания пода.
// Пишет в логе информацию о созданном поде для наглядности
func (s *Service) CreatePod(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pods[name] = struct{}{}

	slog.Info("deployer", slog.String("pod", name), slog.String("action", "create"))
	return nil
}

// DeletePod метод для удаления пода.
// Пишет в логе информацию об удалении пода для наглядности
func (s *Service) DeletePod(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.pods, name)

	slog.Info("deployer", slog.String("pod", name), slog.String("action", "delete"))
	return nil
}

// GetPodList метод получения всех подов, запущенных в момент вызова
func (s *Service) GetPodList() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	podList := make([]string, 0, len(s.pods))
	for k := range s.pods {
		podList = append(podList, k)
	}
	return podList, nil
}
