package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ConvertObj2Vertices(inputFilePath, outputFilePath string) {
	objFile, err := os.Open(inputFilePath)
	checkError(err)
	defer objFile.Close()

	objFileScanner := bufio.NewScanner(objFile)

	objFileScanner.Split(bufio.ScanLines)

	var allVertices = make([]Vertex, 0)

	for objFileScanner.Scan() {
		line := objFileScanner.Text()
		if strings.HasPrefix(line, "v") {
			coordinates := strings.Split(strings.TrimSpace(strings.TrimLeft(line, "v")), " ")
			if len(coordinates) < 3 {
				log.Fatalf("invalid line: %s\n", line)
			}
			x, err := strconv.ParseFloat(coordinates[0], 64)
			checkError(err)
			y, err := strconv.ParseFloat(coordinates[1], 64)
			checkError(err)
			z, err := strconv.ParseFloat(coordinates[2], 64)
			checkError(err)
			allVertices = append(allVertices, Vertex{x, y, z})
		}
	}

	allVertices = NormalizeSlice(allVertices)

	createDirectoryPath(outputFilePath)

	outputFile, err := os.Create(outputFilePath)
	checkError(err)
	defer outputFile.Close()

	for _, vertex := range allVertices {
		line := fmt.Sprintf("%f, %f, %f,\n", vertex.X, vertex.Y, vertex.Z)
		_, err = outputFile.WriteString(line)
		checkError(err)
	}

}

func main() {

	inputFilePath := flag.String("input", "", "the path to the .obj input file")
	outputFilePath := flag.String("output", "out/vertices.obj.txt", "the path to the vertex list output file")

	flag.Parse()

	if *inputFilePath == "" {
		log.Fatalln("an input file path must be provided")
	}

	if getExtension(*inputFilePath) != ".obj" {
		log.Fatalln("the input file must be a .obj file")
	}

	ConvertObj2Vertices(*inputFilePath, *outputFilePath)
}
