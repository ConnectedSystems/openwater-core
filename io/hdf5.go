package io

import (
	"bytes"
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/flowmatters/openwater-core/conv"
	"github.com/flowmatters/openwater-core/data"
	"gonum.org/v1/hdf5"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "ArrayType=float64,float32,int32,uint32,int64,uint64"

type H5RefArrayType struct {
	Filename string
	Dataset  string
}

func (h H5RefArrayType) Load() (data.NDArrayType, error) {
	f, err := hdf5.OpenFile(h.Filename, hdf5.F_ACC_RDONLY)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ds, err := f.OpenDataset(h.Dataset)
	if err != nil {
		return nil, err
	}
	defer ds.Close()

	space := ds.Space()
	dims, _, err := space.SimpleExtentDims()
	if err != nil {
		return nil, err
	}

	shape := conv.UintsToInts(dims)
	result := data.NewArrayArrayType(shape)
	impl := result.Unroll()
	ds.Read(&impl)
	return result, nil
}

func (h H5RefArrayType) Write(data data.NDArrayType) error {
	f, err := hdf5.OpenFile(h.Filename, hdf5.F_ACC_RDWR)
	if err != nil {
		if _, err := os.Stat(h.Filename); os.IsNotExist(err) {
			f, err = hdf5.CreateFile(h.Filename, hdf5.F_ACC_TRUNC)
			if err != nil {
				return err
			}
		}
	}
	defer f.Close()

	ds, err := f.OpenDataset(h.Dataset)
	if err == nil {
		// Ensure dataspace is the write size...
		// OR. just complain for now...

	} else {
		dtype, err := hdf5.NewDatatypeFromValue(0.0)
		if err != nil {
			return err
		}

		dims := conv.IntsToUints(data.Shape())
		space, err := hdf5.CreateSimpleDataspace(dims, nil)
		if err != nil {
			return err
		}
		defer space.Close()

		ds, err = f.CreateDataset(h.Dataset, dtype, space)
		if err != nil {
			return err
		}
	}
	defer ds.Close()

	arrAsSlice := data.Unroll()
	err = ds.Write(&arrAsSlice)
	if err != nil {
		return err
	}

	return nil
}

func (h H5RefArrayType) LoadText() ([]string, error) {
	f, err := hdf5.OpenFile(h.Filename, hdf5.F_ACC_RDONLY)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ds, err := f.OpenDataset(h.Dataset)
	if err != nil {
		return nil, err
	}
	defer ds.Close()

	dt, err := ds.Datatype()
	defer dt.Close()

	if dt.GoType() != reflect.TypeOf("a string") {
		return nil, errors.New("Not a string type")
	}

	space := ds.Space()
	dims, _, err := space.SimpleExtentDims()
	if err != nil {
		return nil, err
	}

	if len(dims) > 1 {
		return nil, errors.New("Can only read 1D data as text")
	}
	maxLen := int(dt.Size())
	nStrings := int(dims[0])
	characters := make([]byte, nStrings*maxLen)
	ds.Read(&characters)

	result := make([]string, dims[0])
	for i := 0; i < nStrings; i++ {
		theBytes := characters[(i * maxLen):((i + 1) * maxLen)]
		end := bytes.Index(theBytes, []byte{0})
		if end < 0 {
			end = maxLen
		}
		theBytes = theBytes[0:end]
		result[i] = string(theBytes)
	}
	return result, nil
}

func (h H5RefArrayType) GetDatasets() ([]string, error) {
	f, err := hdf5.OpenFile(h.Filename, hdf5.F_ACC_RDONLY)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	g, err := f.OpenGroup(h.Dataset)
	if err != nil {
		return nil, err
	}
	defer g.Close()

	n, err := g.NumObjects()
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for i := 0; i < int(n); i++ {
		name, err := g.ObjectNameByIndex(uint(i))
		if err != nil {
			return nil, err
		}
		ds, err := g.OpenDataset(name)
		if err == nil {
			ds.Close()
			result = append(result, name)
		}
	}
	return result, nil
}

func ParseH5RefArrayType(path string) H5RefArrayType {
	components := strings.Split(path, ":")
	return H5RefArrayType{components[0], components[1]}
}
