package main

import (
	"fmt"
	"indexer/globals"
	"indexer/helpers"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const nChunks int = 10

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Ingrese un path válido!")
	}

	path := os.Args[1]

	start := time.Now()

	//SE ASEGURA QUE EL TEMPDIR SEA ÚNICO
	if _, err := os.Stat(globals.TEMPDIR); err == nil {
		if err := os.RemoveAll(globals.TEMPDIR); err != nil {
			log.Fatal(err)
		}
	}

	err := os.Mkdir(globals.TEMPDIR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	//SE LEVANTA EL PROCESO DE ZINCSEARCH EN LA TERMINAL
	cmd := exec.Command("zinc")
	cmd.Env = append(os.Environ(),
		"ZINC_FIRST_ADMIN_USER="+globals.ZINC_USER,
		"ZINC_FIRST_ADMIN_PASSWORD="+globals.ZINC_PWD,
	)

	err = cmd.Start()
	if err != nil {
		fmt.Println("ERROR: Cannot start Zincsearch process ")
	}

	fmt.Println("DEBUG: Starting Zincsearch process")
	time.Sleep(2 * time.Second)

	//SE CREA EL INDEX EN ZINCSEARCG
	err = helpers.CreateIndex(globals.ZINC_CRTIDX_ENDPOINT, "POST", globals.ZINC_IDXMAP)
	if err != nil {
		fmt.Printf("Zinc server is down!: %v \n", err)
		return
	}

	//GENERAR SLICE CON TODOS LOS PATHS DE EMAILS EN LA DB
	fileList := []string{}

	error := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if !d.IsDir() && d.Name() != ".DS_Store" && d.Name() != "DELETIONS.txt" {
			fileList = append(fileList, path)
		}
		return nil
	})

	if error != nil {
		fmt.Println(error)
	}

	//GENERAR SLICE DE CHUNKS PARA PROCESAR CON GORUTINA
	chunks := helpers.ChunkSlice(fileList, nChunks)
	fmt.Printf("---------- NO. OF CHUNKS: %v\n---------- NO. OF FILES:  %v\n", len(chunks), len(fileList))

	//SE GENERA DOCUMENTO JSON VALIDO Y SE CARGAR BULK POR CHUNKS USANDO CONCURRENCIA
	wg.Add(nChunks)

	for idx, chunk := range chunks {
		go helpers.UploadChunk(chunk, idx, &wg)
	}

	wg.Wait() //Se bloquea la ejecución hasta que se terminen todos los threads

	//SE BORRA EL DIRECTORIO TEMPORAL
	if err := os.RemoveAll(globals.TEMPDIR); err != nil {
		log.Fatal(err)
	}

	helpers.TimeTrack(start, "indexEmailDBProcess") //Aquí finaliza la medición de tiempo del proceso de indexación

	//SE DEJA VIVO EL PROCESO DE ZINCSEARCH
	fmt.Println("DEBUG: Zincsearch UI running in http://localhost:4080")
	err = cmd.Wait()
	if err != nil {
		fmt.Println("ERR: Cannot wait Zincsearch process")
	}

	fmt.Println("DEBUG: Zincsearch proccess finished")
}
