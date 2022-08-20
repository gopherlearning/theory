package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

const workers = 10

type Record struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type InputWorker struct {
	r  io.Reader
	ch chan []string
}

func (w *InputWorker) Do(ctx context.Context) error {
	headerRead := false // флаг сигнализирует о считывании заголовка csv-файла (первая строка)
	r := csv.NewReader(w.r)

	for {
		recordRaw, err := r.Read() // ["ivan", "ivanov"]
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		if !headerRead {
			headerRead = true
			continue
		}

		select {
		case <-ctx.Done():
			return nil
		case w.ch <- recordRaw:
		}
	}
}

type OutputWorker struct {
	ch chan []string
	w  io.Writer
	mu *sync.Mutex
}

func (w *OutputWorker) Do() error {
	for recordRaw := range w.ch {
		if len(recordRaw) < 2 {
			return fmt.Errorf("wrong amount of csv segments: %d", len(recordRaw))
		}
		err := func() error {
			w.mu.Lock()
			defer w.mu.Unlock()

			record := Record{FirstName: recordRaw[0], LastName: recordRaw[1]}
			err := json.NewEncoder(w.w).Encode(&record)
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	inputFilename := flag.String("f", "input.csv", "input file")      // задаём входной csv-файл
	outputFilename := flag.String("o", "output.jsonl", "output file") // задаём выходной JSON-файл
	flag.Parse()

	inputFile, err := os.Open(*inputFilename)
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile(*outputFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("unable to open output file: %v", err)
	}
	defer outputFile.Close()

	ctx, cancel := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(ctx) // используем errgroup
	recordCh := make(chan []string)
	mu := &sync.Mutex{} // нужен для потокобезопасной записи в файл

	for i := 0; i < workers; i++ {
		w := &OutputWorker{ch: recordCh, w: outputFile, mu: mu}
		g.Go(w.Do)
	}

	w := &InputWorker{ch: recordCh, r: inputFile}
	err = w.Do(ctx)
	if err != nil {
		log.Println(err)
		cancel()
	}
	close(recordCh)

	err = g.Wait()
	if err != nil {
		log.Println(err)
	}
}
