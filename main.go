package main

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/y-yu/kindle-clock-go/inject"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func convertSVGTest() bytes.Buffer {
	// Generate your SVG data here
	svgData := []byte(`<svg width="758" height="1024" xmlns="http://www.w3.org/2000/svg">
      <defs>
      <style>
        @import url('https://fonts.googleapis.com/css2?family=Roboto+Mono:wght@400&amp;family=Roboto+Slab&amp;display=swap');
      </style>
  </defs>
      <style>
          text { fill: black; }
          .imageColor { fill: black; }
        </style>
      <title>Kindle Clock</title>

      <g font-family="Roboto Slab">
        <g transform="scale(3) translate(15 0)" class="imageColor">
          <svg height="100" width="100" xmlns="http://www.w3.org/2000/svg">
      <path d="M33.469,15v11.938h3.125v-11.938h-3.125zm-13.813,3.406l-2.656,1.656,6.312,10.126,2.657-1.657-6.313-10.125zm30.938,0.313l-5.782,10.437,2.75,1.532,5.782-10.438-2.75-1.531zm-44.532,11.375l-1.5313,2.75,10.407,5.781,1.5-2.75-10.376-5.781zm28.968,0.468c-10.675,0-19.312,8.695-19.312,19.407,0,9.054,6.173,16.619,14.531,18.75-3.104,1.593-5.25,4.819-5.25,8.531,0,2.64,1.214,4.853,3.062,6.094,1.508,1.012,3.31,1.457,5.219,1.594v0.062h1.313,1.812,22.75,13.375,12.281c8.363,0,15.188-6.823,15.188-15.188,0-8.364-6.824-15.156-15.188-15.156-1.204,0-2.335,0.198-3.437,0.469-2.148-7.959-9.313-13.875-17.937-13.875-3.692,0-7.108,1.081-10,2.938-2.449-7.888-9.748-13.626-18.407-13.626zm29.219,0.376l-10.281,5.968,1.562,2.719,10.313-5.969-1.594-2.718zm-0.812,13.874c7.565,0,13.762,5.572,14.906,12.813l0.312,2.156,2.063-0.781c1.3-0.493,2.668-0.781,4.093-0.781,6.436,0,11.626,5.159,11.626,11.593,0,6.435-5.191,11.626-11.626,11.626h-12.281-13.375-24.562c-1.933,0-3.553-0.406-4.532-1.063-0.978-0.657-1.5-1.449-1.5-3.125,0-3.347,2.681-6.031,6.032-6.031,0.554,0,1.13,0.126,1.75,0.312l2.281,0.688v-2.407-0.187c0-4.067,3.246-7.313,7.313-7.313-0.059,0,0.097,0.036,0.5,0.063l2,0.125-0.094-2c-0.022-0.488-0.063-0.646-0.063-0.531,0-8.384,6.767-15.157,15.157-15.157zm-63.438,2.688v3.125h11.906v-3.125h-11.906zm13.562,11.156l-10.062,6.344,1.625,2.656,10.094-6.344-1.657-2.656zm8.438,10.282l-5.75,10.468,2.719,1.5,5.781-10.437-2.75-1.531z"/>
    </svg>
        </g>
        <text font-size="10px" y="270" x="490" text-anchor="end">
          2025-02-23T13:40:56+09:00[Asia/Tokyo]
        </text>

        <rect y="25" x="485" width="145" height="60" rx="5" style="stroke:gray;stroke-width:5;fill-opacity:0;"/>
        <text font-size="40px" font-family="Roboto Mono" y="70" x="500">
          13:44
        </text>
        
        <text font-size="35px" y="160" x="490" text-anchor="middle">Score:</text>
        <text font-size="90px" y="230" x="630" text-anchor="middle">
          82
        </text>

        <text y="360" x="20" font-size="30px">
          <tspan x="20" dy="2em">AWAIR</tspan>
          <tspan x="20" dy="4em">Nature</tspan>
          <tspan x="40" dy="1em">Remo</tspan>
          <tspan x="20" dy="4em">SwitchBot</tspan>
        </text>
        
        <text y="360" x="170" font-size="90px">
          <tspan x="200" dy="-1em" dx="0.5em" font-size="30px">Temperature</tspan>
          <tspan x="200" dy="1.2em">
            17.7
            <tspan font-size="40px" dx="-0.5em" text-anchor="end">°C</tspan>
          </tspan>
          <tspan x="200" dy="1.5em">
            19.2
            <tspan font-size="40px" dx="-0.5em" text-anchor="end">°C</tspan>
          </tspan>
          <tspan x="200" dy="1.5em">
            19.6
            <tspan font-size="40px" dx="-0.5em" text-anchor="end">°C</tspan>
          </tspan>
        </text>

        <text y="360" x="520" font-size="90px">
          <tspan x="490" dy="-1em" dx="1.2em" font-size="30px">Humidity</tspan>
          <tspan x="490" dy="1.2em">
            28.4
            <tspan font-size="40px" dx="-0.5em" text-anchor="end">%</tspan>
          </tspan>
          <tspan x="490" dy="1.5em">
            30.0
            <tspan font-size="40px" dx="-0.5em" text-anchor="end">%</tspan>
          </tspan>
          <tspan x="490" dy="1.5em">
            23.0
            <tspan font-size="40px" dx="-0.5em" text-anchor="end">%</tspan>
          </tspan>
        </text>
        
        <text font-size="35px" y="780" x="250" text-anchor="middle">Electric Energy:</text>
        <text font-size="90px" y="850" x="565" text-anchor="end">
          301
        </text>
        <text font-size="64px" y="850" x="570" text-anchor="start">W</text>

        <text font-size="35px" y="923" x="50" text-anchor="middle">CO<tspan baseline-shift="sub" font-size="25">2</tspan>:</text>
        <text font-size="64px" y="993" x="180" text-anchor="end">
          765
        </text>
        <text font-size="35px" y="983" x="220" text-anchor="middle">ppm</text>

        <text font-size="35px" y="923" x="300" text-anchor="middle">VOC:</text>
        <text font-size="64px" y="993" x="440" text-anchor="end">
          277
        </text>
        <text font-size="35px" y="983" x="475" text-anchor="middle">ppb</text>

        <text font-size="35px" y="923" x="575" text-anchor="middle">PM2.5:</text>
        <text font-size="64px" y="993" x="630" text-anchor="end">
          0.0
        </text>
        <text font-size="35px" y="983" x="690" text-anchor="middle">μg/m<tspan baseline-shift="super" font-size="25px">3</tspan></text>
      </g>
    </svg>`)

	// Create a pipe
	pr, pw := io.Pipe()

	var wg sync.WaitGroup
	wg.Add(1)
	var buf bytes.Buffer
	// Start GraphicsMagick in a separate goroutine
	go func() {
		defer wg.Done()

		cmd := exec.Command("inkscape", "--pipe", "--export-filename=-", "--export-type=png")
		cmd.Stdin = pr
		cmd.Stdout = &buf
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Write SVG data to the pipe
	_, err := pw.Write(svgData)
	if err != nil {
		fmt.Println(err)
	}

	// Close the pipe writer to signal end of data
	err = pw.Close()
	if err != nil {
		fmt.Println(err)
	}

	wg.Wait()
	return buf
}

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(error); ok {
					// スタックトレースを含むエラーを生成
					stackTrace := errors.WithStack(e)
					slog.Error("Recovered", "stacktrace", stackTrace)
				}
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

const port = 8080

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(errorMiddleware)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		repo := inject.Initialize(ctx)
		response, err := repo.GetRoomInfo(ctx)
		if err != nil {
			slog.Warn("Error!", "error", err)
		}
		slog.Info("GetRoomInfo", "response", response)

		buf := convertSVGTest()
		w.Header().Set("Content-Type", "image/png")
		w.Write(buf.Bytes())
		//w.Write([]byte("welcome"))
	})

	slog.Info("Server started!", "port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}
