package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
)

type FileHeader struct {
	FileCode   uint32
	Unused     [20]byte
	FileLength uint32
	Version    uint32
	ShapeType  uint32
	// just read to get to the end of the header
	boundingBoxInfo [8]float64
}

type RecordHeader struct {
	RecordNumber uint32
	/* contentLength measured in 16Bit words */
	ContentLength uint32
}

type RecordBody struct {
	/* Xmin, Ymin, Xmax, Ymax */
	ShapeType   uint32
	BoundingBox [4]float64
	NumPoints   uint32
	Points      []Point
}

type Point struct {
	/* both in little */
	X float64
	Y float64
}

type Record struct {
	RecordHeader
	RecordBody
}

/*
type AppleIcon struct {
	Header
	Icons []IconData // All Icons
} */

// Here is a utility function for dump the information about the file
func (i *FileHeader) printHeader() {
	fmt.Print("\n")
	fmt.Printf("FileHeader:\n")
	fmt.Printf("FileCode: %d\n", i.FileCode)
	fmt.Printf("FileLength: %d\n", i.FileLength)
	fmt.Printf("ShapeType: %v\n", i.ShapeType)
	fmt.Printf("BoundingBox: %v\n", i.boundingBoxInfo)
}

func printRecords(records []Record) {
	for _, record := range records {
		fmt.Printf("RecordNumber: %d\n", record.RecordHeader.RecordNumber)
		fmt.Printf("Record ContentLength: %d\n", record.RecordHeader.ContentLength)
		fmt.Printf("Record Numpoints: %d\n", record.RecordBody.NumPoints)
		fmt.Printf("Points: %v\n", record.RecordBody.Points)
	}

}

// ReadFileHeader uses the reader to read bytes into de AppleIcon structure
func ReadFileHeader(r *bytes.Reader) (*FileHeader, error) {
	var header FileHeader
	binary.Read(r, binary.BigEndian, &header.FileCode)
	binary.Read(r, binary.BigEndian, &header.Unused)
	binary.Read(r, binary.BigEndian, &header.FileLength)
	binary.Read(r, binary.LittleEndian, &header.Version)
	binary.Read(r, binary.LittleEndian, &header.ShapeType)
	binary.Read(r, binary.LittleEndian, &header.boundingBoxInfo)
	return &header, nil
}

func ReadRecords(r *bytes.Reader) {
	var record_number uint32
	var content_length uint32
	var shapeType uint32
	var box [4]float64
	var NumPoints uint32
	var x__coord float64
	var y__coord float64
	for {
		err := binary.Read(r, binary.BigEndian, &record_number)
		binary.Read(r, binary.BigEndian, &content_length)

		binary.Read(r, binary.LittleEndian, &shapeType)
		binary.Read(r, binary.LittleEndian, &box)
		binary.Read(r, binary.LittleEndian, &NumPoints)
		binary.Read(r, binary.LittleEndian, &x__coord)
		binary.Read(r, binary.LittleEndian, &y__coord)
		var count int = 24
		var unused []byte = make([]byte, count)
		binary.Read(r, binary.LittleEndian, &unused)

		fmt.Println(record_number)
		fmt.Println(NumPoints)
		fmt.Println(x__coord)
		fmt.Println(y__coord)

		fmt.Print("\n")
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else {
				fmt.Print(err)
			}
		}
		/* 		count := 8 + 8 + (8 * record.NumPoints) + 8 + 8 + (8 * record.NumPoints) */
		//var unused []byte = make([]byte, count)
		/* fmt.Printf("NumPoints: %d\n", record.RecordBody.NumPoints)
		for i := 0; i < int(record.RecordBody.NumPoints); i++ {
			var point Point
			binary.Read(r, binary.LittleEndian, &point.X)
			binary.Read(r, binary.LittleEndian, &point.Y)
			points = append(points, point)
		}
		binary.Read(r, binary.LittleEndian, &record.RecordBody.Points)
		binary.Read(r, binary.LittleEndian, &unused) */
	}
}

func main() {

	// This icon is from OpenEmu app, you can get it inside the example repository
	data, err := ioutil.ReadFile("A1_NODE.shp")

	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)
	header, err := ReadFileHeader(reader)
	ReadRecords(reader)
	if err != nil {
		panic(err)
	}
	header.printHeader() // Dump the information

	/* 	printRecords(records) */

	/* TODO to read a complete file
	1. read file header and get the shapetype
	2. depending on the shapetype:
		2.1 skip the record header (no usable information)
		2.2 jump to count of entries
		2.3 loop through every entry and calculate on which adress in loop we can read the point
		2.4 add the point to current shape record
		2.5 add the record back to shape_array
		2.6 return shape_array
		2.7 convert to flattent json
	*/
}
