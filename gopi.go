package main

import (
  "fmt"
  "math"
  "math/rand"
  "unicode"

  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  "github.com/faiface/pixel/imdraw"
  "github.com/faiface/pixel/text"
  "golang.org/x/image/font/basicfont"
)

const MAX_GENERATIONS = 1337
const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768
const SQUARE_SIZE = 400
const CIRCLE_RADIUS = SQUARE_SIZE / 2
const MARKER_SIZE float64 = 2

type Vector struct {
  X float64
  Y float64
}

func main() {
  pixelgl.Run(run)
}

func run() {
  cfg := pixelgl.WindowConfig {
    Bounds: pixel.R(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT),
    VSync: true,
  }

  window, err := pixelgl.NewWindow(cfg)

  if err != nil {
    panic(err)
  }

  atlas := text.NewAtlas(
    basicfont.Face7x13,
    text.ASCII,
    text.RangeTable(unicode.Latin),
  )

  drawRenderer := imdraw.New(nil)
  textRenderer := text.New(pixel.V(20, 40), atlas)
  generations := 0
  vectors := []Vector {}

  midPoint := Vector {
    X: WINDOW_WIDTH / 2,
    Y: WINDOW_HEIGHT / 2,
  }

  renderBoundingBox(midPoint, drawRenderer)
  renderInnerCircle(midPoint, drawRenderer)

  for !window.Closed() {
    qualifiedPoints := 0
    window.Clear(pixel.RGB(0, 0, 0))
    textRenderer.Clear()

    if generations < MAX_GENERATIONS {
      generations += 1
      vectors = append(vectors, makeVector())
    }

    for _, vector := range vectors {
      if (inCircle(vector)) {
        qualifiedPoints += 1
        drawMarker(vector, true, midPoint, drawRenderer)
      } else {
        drawMarker(vector, false, midPoint, drawRenderer)
      }
    }

    pi := (float64(qualifiedPoints) / float64(generations)) * 4
    renderOutput(pi, generations, textRenderer)
    textRenderer.Draw(window, pixel.IM.Scaled(textRenderer.Orig, 2))
    drawRenderer.Draw(window)
    window.Update()
  }
}

func inCircle(vector Vector) bool {
  circleMidPoint := float64(SQUARE_SIZE / 2)
  distance := math.Sqrt(math.Pow(vector.X - circleMidPoint, 2) + math.Pow(vector.Y - circleMidPoint, 2))

  return distance <= CIRCLE_RADIUS
}

func renderBoundingBox(midPoint Vector, drawRenderer *imdraw.IMDraw) {
  boundingBoxMidPoint := float64(SQUARE_SIZE / 2)
  drawRenderer.Color = pixel.RGB(0, 0, 0.5)
  drawRenderer.Push(pixel.V(midPoint.X - boundingBoxMidPoint, midPoint.Y - boundingBoxMidPoint))
  drawRenderer.Push(pixel.V(midPoint.X + boundingBoxMidPoint, midPoint.Y - boundingBoxMidPoint))
  drawRenderer.Push(pixel.V(midPoint.X + boundingBoxMidPoint, midPoint.Y + boundingBoxMidPoint))
  drawRenderer.Push(pixel.V(midPoint.X - boundingBoxMidPoint, midPoint.Y + boundingBoxMidPoint))
  drawRenderer.Polygon(2)
}

func renderInnerCircle(midPoint Vector, imd *imdraw.IMDraw) {
  imd.Color = pixel.RGB(1, 1, 1)
  imd.Push(pixel.V(midPoint.X, midPoint.Y))
  imd.Circle(CIRCLE_RADIUS, 1)
}

func renderOutput(pi float64, generations int, textRenderer *text.Text) {
  fmt.Fprintln(textRenderer, "Pi:", pi)
  fmt.Fprintln(textRenderer, "Gen:", generations, "/", MAX_GENERATIONS)
}

func drawMarker(vector Vector, inCircle bool, midPoint Vector, imd *imdraw.IMDraw) {
  boundingBoxMidPoint := float64(SQUARE_SIZE / 2)

  if (inCircle) {
    imd.Color = pixel.RGB(0, 1, 0)
  } else {
    imd.Color = pixel.RGB(1, 0, 0)
  }

  vec1TopLeft := pixel.V(
    midPoint.X - boundingBoxMidPoint + vector.X - MARKER_SIZE,
    midPoint.Y - boundingBoxMidPoint + vector.Y - MARKER_SIZE,
  )

  vec1BottomRight := pixel.V(
    midPoint.X - boundingBoxMidPoint + vector.X + MARKER_SIZE,
    midPoint.Y - boundingBoxMidPoint + vector.Y + MARKER_SIZE,
  )

  vec2TopRight := pixel.V(
    midPoint.X - boundingBoxMidPoint + vector.X - MARKER_SIZE,
    midPoint.Y - boundingBoxMidPoint + vector.Y + MARKER_SIZE,
  )

  vec2BottomLeft := pixel.V(
    midPoint.X - boundingBoxMidPoint + vector.X + MARKER_SIZE,
    midPoint.Y - boundingBoxMidPoint + vector.Y - MARKER_SIZE,
  )

  imd.Push(vec1TopLeft)
  imd.Push(vec1BottomRight)
  imd.Line(1)
  imd.Push(vec2TopRight)
  imd.Push(vec2BottomLeft)
  imd.Line(1)
}

func makeVector() Vector {
  x := SQUARE_SIZE * rand.Float64()
  y := SQUARE_SIZE * rand.Float64()

  return Vector {
    X: x,
    Y: y,
  }
}
