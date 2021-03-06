package gltf

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/go-test/deep"
)

func readFile(path string) []uint8 {
	r, _ := ioutil.ReadFile(path)
	return r
}

func TestOpen(t *testing.T) {
	deep.FloatPrecision = 5
	type args struct {
		name     string
		embedded string
	}
	tests := []struct {
		args    args
		want    *Document
		wantErr bool
	}{
		{args{"openError", ""}, nil, true},
		{args{"testdata/Cube/glTF/Cube.gltf", ""}, &Document{
			Accessors: []Accessor{
				{BufferView: Index(0), ByteOffset: 0, ComponentType: UnsignedShort, Count: 36, Max: []float64{35}, Min: []float64{0}, Type: Scalar},
				{BufferView: Index(1), ByteOffset: 0, ComponentType: Float, Count: 36, Max: []float64{1, 1, 1}, Min: []float64{-1, -1, -1}, Type: Vec3},
				{BufferView: Index(2), ByteOffset: 0, ComponentType: Float, Count: 36, Max: []float64{1, 1, 1}, Min: []float64{-1, -1, -1}, Type: Vec3},
				{BufferView: Index(3), ByteOffset: 0, ComponentType: Float, Count: 36, Max: []float64{1, 0, 0, 1}, Min: []float64{0, 0, -1, -1}, Type: Vec4},
				{BufferView: Index(4), ByteOffset: 0, ComponentType: Float, Count: 36, Max: []float64{1, 1}, Min: []float64{-1, -1}, Type: Vec2}},
			Asset: Asset{Generator: "VKTS glTF 2.0 exporter", Version: "2.0"},
			BufferViews: []BufferView{
				{Buffer: 0, ByteLength: 72, ByteOffset: 0, Target: ElementArrayBuffer},
				{Buffer: 0, ByteLength: 432, ByteOffset: 72, Target: ArrayBuffer},
				{Buffer: 0, ByteLength: 432, ByteOffset: 504, Target: ArrayBuffer},
				{Buffer: 0, ByteLength: 576, ByteOffset: 936, Target: ArrayBuffer},
				{Buffer: 0, ByteLength: 288, ByteOffset: 1512, Target: ArrayBuffer},
			},
			Buffers:   []Buffer{{ByteLength: 1800, URI: "Cube.bin", Data: readFile("testdata/Cube/glTF/Cube.bin")}},
			Images:    []Image{{URI: "Cube_BaseColor.png"}, {URI: "Cube_MetallicRoughness.png"}},
			Materials: []Material{{Name: "Cube", AlphaMode: Opaque, AlphaCutoff: Float64(0.5), PBRMetallicRoughness: &PBRMetallicRoughness{BaseColorFactor: NewRGBA(), MetallicFactor: Float64(1), RoughnessFactor: Float64(1), BaseColorTexture: &TextureInfo{Index: 0}, MetallicRoughnessTexture: &TextureInfo{Index: 1}}}},
			Meshes:    []Mesh{{Name: "Cube", Primitives: []Primitive{{Indices: Index(0), Material: Index(0), Mode: Triangles, Attributes: map[string]uint32{"NORMAL": 2, "POSITION": 1, "TANGENT": 3, "TEXCOORD_0": 4}}}}},
			Nodes:     []Node{{Mesh: Index(0), Name: "Cube", Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{0, 0, 0, 1}, Scale: [3]float64{1, 1, 1}}},
			Samplers:  []Sampler{{WrapS: Repeat, WrapT: Repeat}},
			Scene:     Index(0),
			Scenes:    []Scene{{Nodes: []uint32{0}}},
			Textures: []Texture{
				{Sampler: Index(0), Source: Index(0)}, {Sampler: Index(0), Source: Index(1)},
			},
		}, false},
		{args{"testdata/Cameras/glTF/Cameras.gltf", "testdata/Cameras/glTF-Embedded/Cameras.gltf"}, &Document{
			Accessors: []Accessor{
				{BufferView: Index(0), ByteOffset: 0, ComponentType: UnsignedShort, Count: 6, Max: []float64{3}, Min: []float64{0}, Type: Scalar},
				{BufferView: Index(1), ByteOffset: 0, ComponentType: Float, Count: 4, Max: []float64{1, 1, 0}, Min: []float64{0, 0, 0}, Type: Vec3},
			},
			Asset: Asset{Version: "2.0"},
			BufferViews: []BufferView{
				{Buffer: 0, ByteLength: 12, ByteOffset: 0, Target: ElementArrayBuffer},
				{Buffer: 0, ByteLength: 48, ByteOffset: 12, Target: ArrayBuffer},
			},
			Buffers: []Buffer{{ByteLength: 60, URI: "simpleSquare.bin", Data: readFile("testdata/Cameras/glTF/simpleSquare.bin")}},
			Cameras: []Camera{
				{Perspective: &Perspective{AspectRatio: Float64(1.0), Yfov: 0.7, Zfar: Float64(100), Znear: 0.01}},
				{Orthographic: &Orthographic{Xmag: 1.0, Ymag: 1.0, Zfar: 100, Znear: 0.01}},
			},
			Meshes: []Mesh{{Primitives: []Primitive{{Indices: Index(0), Mode: Triangles, Attributes: map[string]uint32{"POSITION": 1}}}}},
			Nodes: []Node{
				{Mesh: Index(0), Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{-0.3, 0, 0, 0.9}, Scale: [3]float64{1, 1, 1}},
				{Camera: Index(0), Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{0, 0, 0, 1}, Scale: [3]float64{1, 1, 1}, Translation: [3]float64{0.5, 0.5, 3.0}},
				{Camera: Index(1), Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{0, 0, 0, 1}, Scale: [3]float64{1, 1, 1}, Translation: [3]float64{0.5, 0.5, 3.0}},
			},
			Scene:  nil,
			Scenes: []Scene{{Nodes: []uint32{0, 1, 2}}},
		}, false},
		{args{"testdata/BoxVertexColors/glTF-Binary/BoxVertexColors.glb", ""}, &Document{
			Accessors: []Accessor{
				{BufferView: Index(0), ByteOffset: 0, ComponentType: UnsignedShort, Count: 36, Type: Scalar},
				{BufferView: Index(1), ByteOffset: 0, ComponentType: Float, Count: 24, Max: []float64{0.5, 0.5, 0.5}, Min: []float64{-0.5, -0.5, -0.5}, Type: Vec3},
				{BufferView: Index(2), ByteOffset: 0, ComponentType: Float, Count: 24, Type: Vec3},
				{BufferView: Index(3), ByteOffset: 0, ComponentType: Float, Count: 24, Type: Vec4},
				{BufferView: Index(4), ByteOffset: 0, ComponentType: Float, Count: 24, Type: Vec2},
			},
			Asset: Asset{Version: "2.0", Generator: "FBX2glTF"},
			BufferViews: []BufferView{
				{Buffer: 0, ByteLength: 72, ByteOffset: 0, Target: ElementArrayBuffer},
				{Buffer: 0, ByteLength: 288, ByteOffset: 72, Target: ArrayBuffer},
				{Buffer: 0, ByteLength: 288, ByteOffset: 360, Target: ArrayBuffer},
				{Buffer: 0, ByteLength: 384, ByteOffset: 648, Target: ArrayBuffer},
				{Buffer: 0, ByteLength: 192, ByteOffset: 1032, Target: ArrayBuffer},
			},
			Buffers:   []Buffer{{ByteLength: 1224, Data: readFile("testdata/BoxVertexColors/glTF-Binary/BoxVertexColors.glb")[1628+20+8:]}},
			Materials: []Material{{Name: "Default", AlphaMode: Opaque, AlphaCutoff: Float64(0.5), PBRMetallicRoughness: &PBRMetallicRoughness{BaseColorFactor: &RGBA{R: 0.8, G: 0.8, B: 0.8, A: 1}, MetallicFactor: Float64(0.1), RoughnessFactor: Float64(0.99)}}},
			Meshes:    []Mesh{{Name: "Cube", Primitives: []Primitive{{Indices: Index(0), Material: Index(0), Mode: Triangles, Attributes: map[string]uint32{"POSITION": 1, "COLOR_0": 3, "NORMAL": 2, "TEXCOORD_0": 4}}}}},
			Nodes: []Node{
				{Name: "RootNode", Children: []uint32{1, 2, 3}, Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{0, 0, 0, 1}, Scale: [3]float64{1, 1, 1}},
				{Name: "Mesh", Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{0, 0, 0, 1}, Scale: [3]float64{1, 1, 1}},
				{Name: "Cube", Mesh: Index(0), Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{0, 0, 0, 1}, Scale: [3]float64{1, 1, 1}},
				{Name: "Texture Group", Matrix: [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, Rotation: [4]float64{0, 0, 0, 1}, Scale: [3]float64{1, 1, 1}},
			},
			Samplers: []Sampler{{WrapS: Repeat, WrapT: Repeat}},
			Scene:    Index(0),
			Scenes:   []Scene{{Name: "Root Scene", Nodes: []uint32{0}}},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.args.name, func(t *testing.T) {
			got, err := Open(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("Open() = %v", diff)
				return
			}
			if tt.args.embedded != "" {
				got, err = Open(tt.args.embedded)
				for i, b := range got.Buffers {
					if b.IsEmbeddedResource() {
						tt.want.Buffers[i].EmbeddedResource()
					}
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if diff := deep.Equal(got, tt.want); diff != nil {
					t.Errorf("Open() = %v", diff)
					return
				}
			}
		})
	}
}

func readCallback(name string) (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewBufferString("a")), nil
}

func TestDecoder_decodeBuffer(t *testing.T) {
	type args struct {
		buffer *Buffer
	}
	tests := []struct {
		name    string
		d       *Decoder
		args    args
		wantErr bool
	}{
		{"byteLength_0", &Decoder{quotas: ReadQuotas{MaxMemoryAllocation: 2}}, args{&Buffer{ByteLength: 0, URI: "a.bin"}}, true},
		{"noURI", &Decoder{quotas: ReadQuotas{MaxMemoryAllocation: 2}}, args{&Buffer{ByteLength: 1, URI: ""}}, true},
		{"invalidURI", &Decoder{quotas: ReadQuotas{MaxMemoryAllocation: 2}}, args{&Buffer{ByteLength: 1, URI: "../a.bin"}}, true},
		{"maxQuota", &Decoder{quotas: ReadQuotas{MaxMemoryAllocation: 2}}, args{&Buffer{ByteLength: 3, URI: "a.bin"}}, true},
		{"cbErr", NewDecoder(nil, func(name string) (io.ReadCloser, error) { return nil, errors.New("") }), args{&Buffer{ByteLength: 3, URI: "a.bin"}}, true},
		{"base", NewDecoder(nil, readCallback), args{&Buffer{ByteLength: 3, URI: "a.bin"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.decodeBuffer(tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeBuffer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_decodeBinaryBuffer(t *testing.T) {
	type args struct {
		buffer *Buffer
	}
	tests := []struct {
		name    string
		d       *Decoder
		args    args
		wantErr bool
	}{
		{"invalidBuffer", new(Decoder), args{&Buffer{ByteLength: 0, URI: "a.bin"}}, true},
		{"readErr", NewDecoder(bytes.NewBufferString(""), nil), args{&Buffer{ByteLength: 1, URI: "a.bin"}}, true},
		{"invalidHeader", NewDecoder(bytes.NewBufferString("aaaaaaaa"), nil), args{&Buffer{ByteLength: 1, URI: "a.bin"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.decodeBinaryBuffer(tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeBinaryBuffer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_Decode(t *testing.T) {
	type args struct {
		doc *Document
	}
	tests := []struct {
		name    string
		d       *Decoder
		args    args
		wantErr bool
	}{
		{"baseJSON", NewDecoder(bytes.NewBufferString("{\"buffers\": [{\"byteLength\": 1, \"URI\": \"a.bin\"}]}"), readCallback), args{new(Document)}, false},
		{"onlyGLBHeader", NewDecoder(bytes.NewBuffer([]byte{0x67, 0x6c, 0x54, 0x46, 0x02, 0x00, 0x00, 0x00, 0x40, 0x0b, 0x00, 0x00, 0x5c, 0x06, 0x00, 0x00, 0x4a, 0x53, 0x4f, 0x4e}), readCallback), args{new(Document)}, true},
		{"glbMaxMemory", NewDecoder(bytes.NewBuffer([]byte{0x67, 0x6c, 0x54, 0x46, 0x02, 0x00, 0x00, 0x00, 0x40, 0x0b, 0x00, 0x00, 0x5c, 0x06, 0x00, 0x00, 0x4a, 0x53, 0x4f, 0x4e}), readCallback).SetQuotas(ReadQuotas{MaxMemoryAllocation: 0}), args{new(Document)}, true},
		{"glbNoJSONChunk", NewDecoder(bytes.NewBuffer([]byte{0x67, 0x6c, 0x54, 0x46, 0x02, 0x00, 0x00, 0x00, 0x40, 0x0b, 0x00, 0x00, 0x5c, 0x06, 0x00, 0x00, 0x4a, 0x52, 0x4f, 0x4e}), readCallback), args{new(Document)}, true},
		{"empty", NewDecoder(bytes.NewBufferString(""), nil), args{new(Document)}, true},
		{"invalidJSON", NewDecoder(bytes.NewBufferString("{asset: {}}"), nil), args{new(Document)}, true},
		{"invalidBuffer", NewDecoder(bytes.NewBufferString("{\"buffers\": [{\"byteLength\": 0}]}"), nil), args{new(Document)}, true},
		{"maxBuffers", NewDecoder(bytes.NewBufferString("{\"buffers\": [{\"byteLength\": 0}]}"), nil).SetQuotas(ReadQuotas{MaxBufferCount: 0}), args{new(Document)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Decode(tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
