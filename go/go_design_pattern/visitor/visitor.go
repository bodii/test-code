package visitor

import (
	"fmt"
	"path"
)

// IResourceFile resource file interface
type IResourceFile interface {
	Accept(Visitor) error
}

// Visitor visitor interface
type Visitor interface {
	Visit(IResourceFile) error
}

// PPTFile ppt file struct
type PPTFile struct {
	path string
}

// Accept PPTFile accept visiter function
func (p *PPTFile) Accept(visit Visitor) error {
	return visit.Visit(p)
}

// PDFFile pdf file struct
type PDFFile struct {
	path string
}

// Accept PDFFile accept visiter function
func (p *PDFFile) Accept(visit Visitor) error {
	return visit.Visit(p)
}

// NewResourceFile create an new ResourceFile
func NewResourceFile(filepath string) (IResourceFile, error) {
	switch path.Ext(filepath) {
	case ".ppt":
		return &PPTFile{path: filepath}, nil
	case ".pdf":
		return &PDFFile{path: filepath}, nil
	default:
		return nil, fmt.Errorf("not found file type: %s", filepath)
	}
}

// Compressor struct
type Compressor struct{}

// Visit Compressor visit function
func (c *Compressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PPTFile:
		return c.PPTFileVisit(f)
	case *PDFFile:
		return c.PDFFileVisit(f)
	default:
		return fmt.Errorf("not found resource type: %#v", r)
	}
}

// PPTFileVisit ppt file visit function
func (c *Compressor) PPTFileVisit(p *PPTFile) error {
	fmt.Println("this is ppt file")
	return nil
}

// PDFFileVisit pdf file visit function
func (c *Compressor) PDFFileVisit(p *PDFFile) error {
	fmt.Println("this is pdf file")
	return nil
}
