package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	fmt.Printf("FileCode: %d\n", i.FileCode)
	fmt.Printf("FileLength: %d\n", i.FileLength)
	fmt.Printf("ShapeType: %v\n", i.ShapeType)
	fmt.Printf("BoundingBox: %v\n", i.boundingBoxInfo)
}

func (i *Record) printRecords() {
	fmt.Printf("RecordNumber: %d\n", i.RecordHeader.RecordNumber)
	fmt.Printf("Record ContentLength: %d\n", i.RecordHeader.ContentLength)
	fmt.Printf("Record Numpoints: %d\n", i.RecordBody.NumPoints)
	fmt.Printf("Points: %v\n", i.RecordBody.Points)
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

func ReadRecords(r *bytes.Reader, fileLength uint32) (Record, error) {
	var record Record
	var points []Point = make([]Point, 0)
	binary.Read(r, binary.BigEndian, &record.RecordHeader.RecordNumber)
	binary.Read(r, binary.BigEndian, &record.RecordHeader.ContentLength)
	binary.Read(r, binary.LittleEndian, &record.RecordBody.ShapeType)
	binary.Read(r, binary.LittleEndian, &record.RecordBody.BoundingBox)
	binary.Read(r, binary.LittleEndian, &record.RecordBody.NumPoints)
	for i := 0; i < int(record.RecordBody.NumPoints); i++ {
		var point Point
		binary.Read(r, binary.LittleEndian, &point.X)
		binary.Read(r, binary.LittleEndian, &point.Y)
		points = append(points, point)
	}
	record.RecordBody.Points = points
	fmt.Printf("Reader index: %d\n", r.Len())
	fmt.Printf("FileLength: %d\n", fileLength)
	return record, nil
}

func main() {

	// This icon is from OpenEmu app, you can get it inside the example repository
	data, err := ioutil.ReadFile("A1_NODE.shp")

	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)
	header, err := ReadFileHeader(reader)
	record, err := ReadRecords(reader, header.FileLength)
	if err != nil {
		panic(err)
	}
	/* header.printHeader() // Dump the information */
	record.printRecords()

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
