package main

import (
	"fmt"

	"go.examples.tasks/cmd/interface_examples/coding/task10_hard/decorator"
	"go.examples.tasks/cmd/interface_examples/coding/task10_hard/storage"
)

func NewStorage(kind string, config interface{}) (storage.Storage, error) {
	switch kind {
	case "memory":
		return storage.NewMapStorage(), nil
	case "file":
		path, ok := config.(string)
		if !ok {
			return nil, fmt.Errorf("config for file must be string path")
		}
		return storage.NewFileStorage(path)
	case "logging":
		inner, ok := config.(storage.Storage)
		if !ok {
			return nil, fmt.Errorf("config for logging must be Storage")
		}
		return decorator.NewLoggingStorage(inner), nil
	case "metrics":
		inner, ok := config.(storage.Storage)
		if !ok {
			return nil, fmt.Errorf("config for metrics must be Storage")
		}
		return decorator.NewMetricsStorage(inner), nil
	default:
		return nil, fmt.Errorf("unknown storage kind: %s", kind)
	}
}
