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

// Here is a utility function for dump the information about the file
func (i *FileHeader) printHeader() {
	fmt.Print("\n")
	fmt.Printf("FileHeader:\n")
	fmt.Printf("FileCode: %d\n", i.FileCode)
	fmt.Printf("FileLength: %d\n", i.FileLength)
	fmt.Printf("ShapeType: %v\n", i.ShapeType)
	fmt.Printf("BoundingBox: %v\n", i.boundingBoxInfo)
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

type shape struct {
	shapeType uint32
	points    []Point
}

func ReadRecords(r *bytes.Reader) : {
	var record Record
	/* read new record if the file not has ended yet */
	for {
		err := binary.Read(r, binary.BigEndian, &record.RecordHeader.RecordNumber)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else {
				fmt.Print(err)
			}
		}
		binary.Read(r, binary.BigEndian, &record.RecordHeader.RecordNumber)

		binary.Read(r, binary.LittleEndian, &record.RecordBody.ShapeType)
		binary.Read(r, binary.LittleEndian, &record.RecordBody.BoundingBox)
		binary.Read(r, binary.LittleEndian, &record.RecordBody.NumPoints)
		record.RecordBody.Points = make([]Point, record.RecordBody.NumPoints)
		for i := 0; i < int(record.RecordBody.NumPoints); i++ {
			var new_point = Point{}
			binary.Read(r, binary.LittleEndian, &new_point.X)
			binary.Read(r, binary.LittleEndian, &new_point.Y)
			record.RecordBody.Points[i] = new_point
		}

		var count int = 24
		var unused []byte = make([]byte, count)
		binary.Read(r, binary.LittleEndian, &unused)
	}
}

func main() {
	data, err := ioutil.ReadFile("A1_NODE.shp")
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(data)
	/* always start with reading the fileHeader in the same way */
	header, err := ReadFileHeader(reader)
	if header.ShapeType == 1 {
		fmt.Println("Point")
	} else if header.ShapeType == 3 {
		fmt.Println("PolyLine")
	} else if header.ShapeType == 5 {
		fmt.Println("Polygon")
	} else if header.ShapeType == 8 {
		fmt.Println("MultiPoint")
	} else if header.ShapeType == 11 {
		fmt.Println("PointZ")
	} else if header.ShapeType == 13 {
		fmt.Println("PolyLineZ")
	} else if header.ShapeType == 15 {
		fmt.Println("PolygonZ")
	} else if header.ShapeType == 18 {
		fmt.Println("MultiPointZ")
	} else if header.ShapeType == 21 {
		fmt.Println("PointM")
	} else if header.ShapeType == 23 {
		fmt.Println("PolyLineM")
	} else if header.ShapeType == 25 {
		fmt.Println("PolygonM")
	} else if header.ShapeType == 28 {
		fmt.Println("MultiPointM")
	} else if header.ShapeType == 31 {
		fmt.Println("MultiPatch")
	}

	ReadRecords(reader)
	if err != nil {
		panic(err)
	}
}
