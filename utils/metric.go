package utils

import (
	"sort"

	f "github.com/open-falcon/falcon-plus/modules/api/app/model/falcon_portal"
)

type By func(m1, m2 *f.Metric) bool

func (by By) Sort(m []f.Metric) {
	ms := &MetricSorter{
		metrics: m,
		by:      by,
	}
	sort.Sort(ms)
}

type MetricSorter struct {
	metrics []f.Metric
	index   int
	by      func(p1, p2 *f.Metric) bool
}

func (s *MetricSorter) Len() int {
	return len(s.metrics)
}

func (s *MetricSorter) Swap(i, j int) {
	s.metrics[i], s.metrics[j] = s.metrics[j], s.metrics[i]
}

func (s *MetricSorter) Less(i, j int) bool {
	return s.by(&s.metrics[i], &s.metrics[j])
}
