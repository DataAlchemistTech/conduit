// Copyright © 2022 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/conduitio/conduit/pkg/foundation/assert"
	"github.com/conduitio/conduit/pkg/foundation/metrics"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"google.golang.org/protobuf/proto"
)

func TestTimer(t *testing.T) {
	testCases := []struct {
		name       string
		observe    func(m metrics.Timer)
		wantCount  uint64
		wantSum    float64
		wantBucket map[float64]uint64
	}{{
		name:    "empty timer",
		observe: func(m metrics.Timer) {},
	}, {
		name:       "update once",
		observe:    func(m metrics.Timer) { m.Update(time.Millisecond * 2500) },
		wantCount:  1,
		wantSum:    2.5,
		wantBucket: map[float64]uint64{2.5: 1, 5: 1, 10: 1},
	}, {
		name: "update twice",
		observe: func(m metrics.Timer) {
			m.Update(time.Millisecond * 2500)
			m.Update(time.Millisecond * 1000)
		},
		wantCount:  2,
		wantSum:    3.5,
		wantBucket: map[float64]uint64{1: 1, 2.5: 2, 5: 2, 10: 2},
	}, {
		name: "update 10 times",
		observe: func(m metrics.Timer) {
			for i := 1; i <= 10; i++ {
				m.Update(time.Millisecond * 25 * time.Duration(i))
			}
		},
		wantCount:  10,
		wantSum:    1.375,
		wantBucket: map[float64]uint64{0.025: 1, 0.05: 2, 0.1: 4, 0.25: 10, 0.5: 10, 1: 10, 2.5: 10, 5: 10, 10: 10},
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reg := NewRegistry(nil)

			m := reg.NewTimer("my_timer", "test timer")
			tc.observe(m)

			mt := dto.MetricType_HISTOGRAM
			want := []*dto.MetricFamily{{
				Name: proto.String("my_timer"),
				Help: proto.String("test timer"),
				Type: &mt,
				Metric: []*dto.Metric{{
					Label: nil, // NB: inconsistency in prometheus lib - other metric types have an empty slice here
					Histogram: &dto.Histogram{
						SampleCount: proto.Uint64(tc.wantCount),
						SampleSum:   proto.Float64(tc.wantSum),
						Bucket:      buildHistogramBuckets(tc.wantBucket),
					},
				}},
			}}

			promRegistry := prometheus.NewRegistry()
			err := promRegistry.Register(reg)
			assert.Ok(t, err)

			got, err := promRegistry.Gather()
			assert.Ok(t, err)
			assert.Equal(t, want, got)
		})
	}
}

func TestTimer_Since(t *testing.T) {
	reg := NewRegistry(nil)

	m := reg.NewTimer("my_timer", "test timer")
	m.UpdateSince(time.Now().Add(-time.Millisecond * 2501))

	mt := dto.MetricType_HISTOGRAM
	want := []*dto.MetricFamily{{
		Name: proto.String("my_timer"),
		Help: proto.String("test timer"),
		Type: &mt,
		Metric: []*dto.Metric{{
			Label: nil, // NB: inconsistency in prometheus lib - other metric types have an empty slice here
			Histogram: &dto.Histogram{
				SampleCount: proto.Uint64(1),
				SampleSum:   proto.Float64(2.501),
				Bucket:      buildHistogramBuckets(map[float64]uint64{5: 1, 10: 1}),
			},
		}},
	}}

	promRegistry := prometheus.NewRegistry()
	err := promRegistry.Register(reg)
	assert.Ok(t, err)

	got, err := promRegistry.Gather()
	assert.Ok(t, err)

	// first manually check the difference, we can't know the actual sum
	diff := *got[0].Metric[0].Histogram.SampleSum - *want[0].Metric[0].Histogram.SampleSum
	if diff > 0.01 { // we leave room for 10 milliseconds
		t.Errorf("sample sum diff was greater than expected: %v", diff)
	}

	// add difference to our estimate and make sure everything else matches
	*want[0].Metric[0].Histogram.SampleSum += diff
	assert.Equal(t, want, got)
}

func TestLabeledTimer(t *testing.T) {
	testCases := []struct {
		name        string
		observe     func(m metrics.LabeledTimer)
		wantCounts  map[string]uint64
		wantSums    map[string]float64
		wantBuckets map[string]map[float64]uint64
	}{{
		name:    "no observed timers",
		observe: func(m metrics.LabeledTimer) {},
	}, {
		name:       "only labels observed",
		observe:    func(m metrics.LabeledTimer) { m.WithValues("val1") },
		wantCounts: map[string]uint64{"val1": 0},
		wantSums:   map[string]float64{"val1": 0},
	}, {
		name: "one observed",
		observe: func(m metrics.LabeledTimer) {
			m1 := m.WithValues("val1")
			m1.Update(time.Millisecond * 2500)
		},
		wantCounts:  map[string]uint64{"val1": 1},
		wantSums:    map[string]float64{"val1": 2.5},
		wantBuckets: map[string]map[float64]uint64{"val1": {2.5: 1, 5: 1, 10: 1}},
	}, {
		name: "10 observed",
		observe: func(m metrics.LabeledTimer) {
			for i := 1; i <= 10; i++ {
				m1 := m.WithValues(fmt.Sprintf("val%d", i))
				m1.Update(time.Millisecond * 250 * time.Duration(i))
			}
		},
		wantCounts: map[string]uint64{
			"val1": 1, "val2": 1, "val3": 1, "val4": 1, "val5": 1, "val6": 1, "val7": 1, "val8": 1, "val9": 1, "val10": 1,
		},
		wantSums: map[string]float64{
			"val1": 0.25, "val2": 0.5, "val3": 0.75, "val4": 1, "val5": 1.25, "val6": 1.5, "val7": 1.75, "val8": 2, "val9": 2.25, "val10": 2.5,
		},
		wantBuckets: map[string]map[float64]uint64{
			"val1":  {0.25: 1, 0.5: 1, 1: 1, 2.5: 1, 5: 1, 10: 1},
			"val2":  {0.5: 1, 1: 1, 2.5: 1, 5: 1, 10: 1},
			"val3":  {1: 1, 2.5: 1, 5: 1, 10: 1},
			"val4":  {1: 1, 2.5: 1, 5: 1, 10: 1},
			"val5":  {2.5: 1, 5: 1, 10: 1},
			"val6":  {2.5: 1, 5: 1, 10: 1},
			"val7":  {2.5: 1, 5: 1, 10: 1},
			"val8":  {2.5: 1, 5: 1, 10: 1},
			"val9":  {2.5: 1, 5: 1, 10: 1},
			"val10": {2.5: 1, 5: 1, 10: 1},
		},
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reg := NewRegistry(nil)
			m := reg.NewLabeledTimer("my_labeled_timer", "test labeled timer", []string{"test_label"})
			tc.observe(m)

			mt := dto.MetricType_HISTOGRAM
			want := make([]*dto.MetricFamily, 0)
			if len(tc.wantCounts) > 0 {
				mf := &dto.MetricFamily{
					Name:   proto.String("my_labeled_timer"),
					Help:   proto.String("test labeled timer"),
					Type:   &mt,
					Metric: []*dto.Metric{},
				}

				// iterate through map in an ordered way
				keys := make([]string, 0)
				for k := range tc.wantCounts {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				for _, label := range keys {
					mf.Metric = append(mf.Metric, &dto.Metric{
						Label: []*dto.LabelPair{{
							Name:  proto.String("test_label"),
							Value: proto.String(label),
						}},
						Histogram: &dto.Histogram{
							SampleCount: proto.Uint64(tc.wantCounts[label]),
							SampleSum:   proto.Float64(tc.wantSums[label]),
							Bucket:      buildHistogramBuckets(tc.wantBuckets[label]),
						},
					})
				}
				want = append(want, mf)
			}

			promRegistry := prometheus.NewRegistry()
			err := promRegistry.Register(reg)
			assert.Ok(t, err)

			got, err := promRegistry.Gather()
			assert.Ok(t, err)
			assert.Equal(t, want, got)
		})
	}
}
