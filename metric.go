package main

import (
	"errors"
	"github.com/rcrowley/go-metrics"
	"strconv"
)

type Metric struct {
	Kind  string
	Key   string
	Value int64

	Counter metrics.Counter
	Gauge   metrics.Gauge
	Meter   metrics.Meter
}

func (m *Metric) Track() error {
	switch m.Kind {
	case `counter`:
		m.Counter = metrics.GetOrRegisterCounter(m.Key, metrics.DefaultRegistry)
		m.Counter.Inc(m.Value)
	case `gauge`:
		m.Gauge = metrics.GetOrRegisterGauge(m.Key, metrics.DefaultRegistry)
		m.Gauge.Update(m.Value)
	case `meter`:
		m.Meter = metrics.GetOrRegisterMeter(m.Key, metrics.DefaultRegistry)
		m.Meter.Mark(m.Value)
	default:
		return errors.New("Invalid kind provided")
	}

	return nil
}

func NewMetric(kind string, key string, value string) (*Metric, error) {
	iValue, err := strconv.ParseInt(value, 0, 0)

	if err != nil {
		return nil, err
	}

	metric := &Metric{
		Kind:    kind,
		Key:     key,
		Value:   iValue,
		Counter: metrics.NewCounter(),
		Gauge:   metrics.NewGauge(),
		Meter:   metrics.NewMeter(),
	}

	return metric, nil
}
