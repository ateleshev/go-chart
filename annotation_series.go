package chart

import (
	"math"
	"time"

	"github.com/wcharczuk/go-chart"
)

// CreateContinuousSeriesLastValueLabel returns a (1) value annotation series.
func CreateContinuousSeriesLastValueLabel(name string, xvalues, yvalues []float64, valueFormatter ValueFormatter) AnnotationSeries {
	return AnnotationSeries{
		Name: name,
		Style: Style{
			Show:        true,
			StrokeColor: chart.GetDefaultSeriesStrokeColor(0),
		},
		Annotations: []chart.Annotation{
			Annotation{
				X:     xvalues[len(xvalues)-1],
				Y:     yvalues[len(yvalues)-1],
				Label: valueFormatter(yvalues[len(yvalues)-1]),
			},
		},
	}
}

// CreateTimeSeriesLastValueLabel returns a (1) value annotation series.
func CreateTimeSeriesLastValueLabel(name string, xvalues []time.Time, yvalues []float64, valueFormatter ValueFormatter) AnnotationSeries {
	return AnnotationSeries{
		Name: name,
		Style: Style{
			Show:        true,
			StrokeColor: chart.GetDefaultSeriesStrokeColor(0),
		},
		Annotations: []chart.Annotation{
			Annotation{
				X:     xvalues[len(xvalues)-1],
				Y:     yvalues[len(yvalues)-1],
				Label: valueFormatter(yvalues[len(yvalues)-1]),
			},
		},
	}
}

// Annotation is a label on the chart.
type Annotation struct {
	X, Y  float64
	Label string
}

// AnnotationSeries is a series of labels on the chart.
type AnnotationSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	Annotations []Annotation
}

// GetName returns the name of the time series.
func (as AnnotationSeries) GetName() string {
	return as.Name
}

// GetStyle returns the line style.
func (as AnnotationSeries) GetStyle() Style {
	return as.Style
}

// GetYAxis returns which YAxis the series draws on.
func (as AnnotationSeries) GetYAxis() YAxisType {
	return as.YAxis
}

// Measure returns a bounds box of the series.
func (as AnnotationSeries) Measure(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) Box {
	box := Box{
		Top:    math.MaxInt32,
		Left:   math.MaxInt32,
		Right:  0,
		Bottom: 0,
	}
	if as.Style.Show {
		style := as.Style.WithDefaultsFrom(Style{
			Font:        defaults.Font,
			FillColor:   DefaultAnnotationFillColor,
			FontSize:    DefaultAnnotationFontSize,
			StrokeColor: defaults.StrokeColor,
			StrokeWidth: defaults.StrokeWidth,
			Padding:     DefaultAnnotationPadding,
		})
		for _, a := range as.Annotations {
			lx := canvasBox.Right - xrange.Translate(a.X)
			ly := yrange.Translate(a.Y) + canvasBox.Top
			ab := MeasureAnnotation(r, canvasBox, xrange, yrange, style, lx, ly, a.Label)
			if ab.Top < box.Top {
				box.Top = ab.Top
			}
			if ab.Left < box.Left {
				box.Left = ab.Left
			}
			if ab.Right > box.Right {
				box.Right = ab.Right
			}
			if ab.Bottom > box.Bottom {
				box.Bottom = ab.Bottom
			}
		}
	}
	return box
}

// Render draws the series.
func (as AnnotationSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	if as.Style.Show {
		style := as.Style.WithDefaultsFrom(Style{
			Font:        defaults.Font,
			FillColor:   DefaultAnnotationFillColor,
			FontSize:    DefaultAnnotationFontSize,
			StrokeColor: defaults.StrokeColor,
			StrokeWidth: defaults.StrokeWidth,
			Padding:     DefaultAnnotationPadding,
		})
		for _, a := range as.Annotations {
			lx := canvasBox.Right - xrange.Translate(a.X)
			ly := yrange.Translate(a.Y) + canvasBox.Top
			DrawAnnotation(r, canvasBox, xrange, yrange, style, lx, ly, a.Label)
		}
	}
}
