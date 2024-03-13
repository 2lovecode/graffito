package mode0

import "github.com/2lovecode/graffito/internal/app/experiment/mode0/handler"

type OutData struct {
	CommonData interface{}
	Data       []handler.DataItem
}

type DataSource interface {
	Output() OutData
}

type DataOptions struct {
	Source DataSource
}

type DataAggregationLayer struct {
	ds DataSource
}

type DataOption func(d *DataOptions)

func Source(s DataSource) DataOption {
	return func(d *DataOptions) {
		d.Source = s
	}
}

func NewDataAggregationLayer(opts ...DataOption) DataAggregationLayer {
	opt := DataOptions{
		Source: nil,
	}

	for _, o := range opts {
		o(&opt)
	}

	dataAggLayer := DataAggregationLayer{ds: opt.Source}

	return dataAggLayer
}

func (al DataAggregationLayer) Output() OutData {
	return al.ds.Output()
}
