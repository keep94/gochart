package gochart

import (
  "fmt"
  "io"
  "math/big"
  "os"
  "strconv"
  "strings"
)

const (
  kIdxOutOfRange = "idx out of range"
)

// Ints is a sequence of integer X values.
// Note that Ints implements the Values interface.
type Ints struct {
  start int64
  inc int64
  count int
}

// NewInts returns a sequence of count integers starting at start and
// incrementing by inc.
func NewInts(start, inc int64, count int) *Ints {
  return &Ints{start: start, inc: inc, count: count}
}

// Apply applies f to each of these X values and returns the resulting
// Y values.
func (i *Ints) Apply(f func(int64) int64) Values {
  result := make(valueSlice, i.count)
  for j := 0; j < i.count; j++ {
    result[j] = f(i.value(j))
  }
  return result
}

// ApplyBigInt applies f to each of these X values and returns the resulting
// Y values. f must return a new *big.Int each time.
func (i *Ints) ApplyBigInt(f func(int64) *big.Int) Values {
  result := make(valueSlice, i.count)
  for j := 0; j < i.count; j++ {
    result[j] = f(i.value(j))
  }
  return result
}

// ApplyBigIntCh uses ch to return the resulting Y values.
// If the X value is 1, the corresponding Y value will be the first value
// off ch. If the X value is 2, the corresponding Y value will be the second
// value off of ch etc. If the X value is less than 1, the corresponding Y
// value will be 0. If the X value is greater than the number of values in Ch,
// the corresponding Y value is also 0.
func (i *Ints) ApplyBigIntCh(ch <-chan *big.Int) Values {
  result := make(valueSlice, i.count)
  ok := true
  var val *big.Int
  var valIndex int64
  for j := 0; j < i.count; j++ {
    idx := i.value(j)
    for ok && valIndex < idx {
      val, ok = <-ch
      valIndex++
    }
    if val == nil {
      result[j] = big.NewInt(0)
    } else {
      result[j] = val
    }
  }
  return result
}

func (i *Ints) Value(idx int) interface{} {
  if idx<0 || idx>=i.count {
    panic(kIdxOutOfRange)
  }
  return i.value(idx)
}

func (i *Ints) Len() int {
  return i.count
}

func (i *Ints) value(idx int) int64 {
  return i.start + int64(idx)*i.inc
}

// Floats is a sequence of floating point X values.
// Note that Floats implements the Values interface.
type Floats struct {
  start float64
  inc float64
  count int
}

// NewFloats returns a sequence of count floats starting at start and
// incrementing by inc.
func NewFloats(start, inc float64, count int) *Floats {
  return &Floats{start: start, inc: inc, count: count}
}

// Apply applies fn to each of these X values and returns the resulting
// Y values.
func (f *Floats) Apply(fn func(float64) float64) Values {
  result := make(valueSlice, f.count)
  for i := 0; i < f.count; i++ {
    result[i] = fn(f.value(i))
  }
  return result
}

// ApplyInv applies the inverse of fn to each of these X values and returns the
// resulting Y values. The Y values that ApplyInv produces will be between
// start and end. fn must be monotone increasing or decreasing between
// start and end. end > start or ApplyInv panics.
func (f *Floats) ApplyInv(fn func(float64) float64, start, end float64) Values {
  if end <= start {
    panic("end must be greater than start")
  }
  result := make(valueSlice, f.count)
  for i := 0; i < f.count; i++ {
    result[i] = inverse(f.value(i), fn, start, end)
  }
  return result
}

func (f *Floats) Value(idx int) interface{} {
  if idx<0 || idx>=f.count {
    panic(kIdxOutOfRange)
  }
  return f.value(idx)
}

func (f *Floats) Len() int {
  return f.count
}

func (f *Floats) value(idx int) float64 {
  return f.start + float64(idx)*f.inc
}

// Interface Values represents a sequence values for either the X or Y column
// of a chart.
type Values interface {

  // Returns the 0-based idx value in this sequence of values
  Value(idx int) interface{}

  // Returns the number of values
  Len() int
}

// Option represents an option for creating a chart.
type Option option

// XFormat sets the format string for formatting X values, default is "%v"
func XFormat(fmtStr string) Option {
  return func(s *settingsType) {
    s.xFormat = fmtStr
  }
}

// YFormat sets the format string for formatting Y values, default is "%v"
func YFormat(fmtStr string) Option {
  return func(s *settingsType) {
    s.yFormat = fmtStr
  }
}

// RowCount sets the number of rows in the chart. The default number of rows
// is the minimum number of rows needed to show all the values given the
// number of columns. If neither rowCount or colCount are set, rowCount
// defaults to the number of values and colCount defaults to 1.
func RowCount(count int) Option {
  return func(s *settingsType) {
    s.rowCount = count
  }
}

// ColCount sets the number of columns in the chart. The default number of
// columns is the minimum number of columns needed to show all the values
// given the number of rows. If neither rowCount or colCount are set, colCount
// defaults to 1 and rowCount defaults to the number of values.
func ColCount(count int) Option {
  return func(s *settingsType) {
    s.colCount = count
  }
}

// Chart represents a chart of X and Y values.
type Chart struct {
  header string
  xyFormat string
  rowCount int
  colCount int
  xyValues xyValuesType
}

// NewChart creates a new chart. xs are the X values and ys are the Y values.
// The number of X and Y values must be the same or else NewChart panics.
// options are the options for creating the chart.
func NewChart(xs, ys Values, options ...Option) *Chart {
  if xs.Len() != ys.Len() {
    panic("xs and ys must have same length")
  }
  settings := &settingsType{xFormat: "%v", yFormat: "%v"}
  settings.applyOptions(options)
  settings.computeDimensions(xs.Len())
  xyValues := createXYValues(xs, ys, settings.xFormat, settings.yFormat)
  xwidth, ywidth := xyValues.widths()
  return &Chart{
      header: createHeader(xwidth, ywidth, settings.colCount),
      xyFormat: createXYFormat(xwidth, ywidth),
      rowCount: settings.rowCount,
      colCount: settings.colCount,
      xyValues: xyValues}
}

func createHeader(xwidth, ywidth, colCount int) string {
  piece := "+" + strings.Repeat("-", xwidth) + "+" + strings.Repeat("-", ywidth)
  return fmt.Sprintf("%s+", strings.Repeat(piece, colCount))
}

func createXYFormat(xwidth, ywidth int) string {
  return "|%" + strconv.Itoa(xwidth) + "s|%" + strconv.Itoa(ywidth) + "s"
}

// WriteTo writes the chart to writer w. If w is nil, WriteTo writes the
// chart to stdout. WriteTo returns the number of bytes written and any
// error encountered.
func (c *Chart) WriteTo(w io.Writer) (n int, err error) {
  if w == nil {
    w = os.Stdout
  }
  var nn int
  nn, err = fmt.Fprintln(w, c.header)
  n += nn
  if err != nil {
    return
  }
  for i := 0; i < c.rowCount; i++ {
    for j := 0; j < c.colCount; j++ {
      xyValue := c.xy(i, j)
      nn, err = fmt.Fprintf(w, c.xyFormat, xyValue.x, xyValue.y)
      n += nn
      if err != nil {
        return
      }
    }
    nn, err = fmt.Fprintln(w, "|")
    n += nn
    if err != nil {
      return
    }
  }
  nn, err = fmt.Fprintln(w, c.header)
  n += nn
  return
}

// RowCount returns the number of rows in this chart.
func (c *Chart) RowCount() int {
  return c.rowCount
}

// ColCount returns the number of columns in this chart.
func (c *Chart) ColCount() int {
  return c.colCount
}

func (c *Chart) xy(row, col int) xyValueType {
  idx := row + c.rowCount*col
  var result xyValueType
  if idx < len(c.xyValues) {
    result = c.xyValues[idx]
  }
  return result
}

type xyValueType struct {
  x string
  y string
}

type xyValuesType []xyValueType

func createXYValues(
    xs, ys Values, xformat, yformat string) xyValuesType {
  result := make(xyValuesType, xs.Len())
  for i := 0; i < xs.Len(); i++ {
    result[i].x = fmt.Sprintf(xformat, xs.Value(i))
    result[i].y = fmt.Sprintf(yformat, ys.Value(i))
  }
  return result
}

func (xy xyValuesType) widths() (xwidth int, ywidth int) {
  for i := 0; i < len(xy); i++ {
    if len(xy[i].x) > xwidth {
      xwidth = len(xy[i].x)
    }
    if len(xy[i].y) > ywidth {
      ywidth = len(xy[i].y)
    }
  }
  return
}

type option func(s *settingsType)

type settingsType struct {
  xFormat string
  yFormat string
  rowCount int
  colCount int
}

func (s *settingsType) applyOptions(options []Option) {
  for _, option := range options {
    option(s)
  }
}

func (s *settingsType) computeDimensions(count int) {
  if s.rowCount <= 0 && s.colCount <= 0 {
    s.rowCount = count
    s.colCount = 1
    return
  }
  if s.rowCount <= 0 {
    colCount := s.colCount
    rowCount := (count + colCount - 1) / colCount
    s.rowCount = rowCount
    return
  }
  if s.colCount <= 0 {
    rowCount := s.rowCount
    colCount := (count + rowCount - 1) / rowCount
    s.colCount = colCount
    return
  }
}

type valueSlice []interface{}

func (v valueSlice) Value(idx int) interface{} {
  if idx < 0 || idx >= len(v) {
    panic(kIdxOutOfRange)
  }
  return v[idx]
}

func (v valueSlice) Len() int {
  return len(v)
}

func inverse(
    x float64,
    f func(float64) float64,
    lower float64,
    upper float64) float64 {
  var g func(float64) float64
  if f(lower) > f(upper) {
    g = func(val float64) float64 {
      return x - f(val)
    }
  } else {
    g = func(val float64) float64 {
      return f(val) - x
    }
  }
  step := (upper - lower) / 2.0
  result := (upper + lower) / 2.0
  for i := 0; i < 53; i++ {
    step /= 2.0
    if g(result) > 0.0 {
      result -= step
    } else {
      result += step
    }
  }
  return result
}