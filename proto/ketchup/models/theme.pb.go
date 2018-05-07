// Code generated by protoc-gen-go.
// source: theme.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ketchup_packages "github.com/ketchuphq/ketchup/proto/ketchup/packages"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A Theme can be entirely contained within a theme proto,
// however more usually `data` will be omitted for
// ThemeTemplate and ThemeAsset.
type Theme struct {
	Uuid             *string                   `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name             *string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description      *string                   `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Package          *ketchup_packages.Package `protobuf:"bytes,4,opt,name=package" json:"package,omitempty"`
	Templates        map[string]*ThemeTemplate `protobuf:"bytes,10,rep,name=templates" json:"templates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Assets           map[string]*ThemeAsset    `protobuf:"bytes,11,rep,name=assets" json:"assets,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Placeholders     []*Data                   `protobuf:"bytes,12,rep,name=placeholders" json:"placeholders,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Theme) Reset()                    { *m = Theme{} }
func (m *Theme) String() string            { return proto.CompactTextString(m) }
func (*Theme) ProtoMessage()               {}
func (*Theme) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *Theme) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Theme) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Theme) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Theme) GetPackage() *ketchup_packages.Package {
	if m != nil {
		return m.Package
	}
	return nil
}

func (m *Theme) GetTemplates() map[string]*ThemeTemplate {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *Theme) GetAssets() map[string]*ThemeAsset {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *Theme) GetPlaceholders() []*Data {
	if m != nil {
		return m.Placeholders
	}
	return nil
}

type ThemeTemplate struct {
	Uuid         *string             `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name         *string             `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Theme        *string             `protobuf:"bytes,3,opt,name=theme" json:"theme,omitempty"`
	Engine       *string             `protobuf:"bytes,4,opt,name=engine,def=html" json:"engine,omitempty"`
	HideContent  *bool               `protobuf:"varint,5,opt,name=hideContent,def=0" json:"hideContent,omitempty"`
	Description  *string             `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Placeholders []*ThemePlaceholder `protobuf:"bytes,10,rep,name=placeholders" json:"placeholders,omitempty"`
	// maybe an option for extensions?
	Data             *string `protobuf:"bytes,11,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ThemeTemplate) Reset()                    { *m = ThemeTemplate{} }
func (m *ThemeTemplate) String() string            { return proto.CompactTextString(m) }
func (*ThemeTemplate) ProtoMessage()               {}
func (*ThemeTemplate) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

const Default_ThemeTemplate_Engine string = "html"
const Default_ThemeTemplate_HideContent bool = false

func (m *ThemeTemplate) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ThemeTemplate) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ThemeTemplate) GetTheme() string {
	if m != nil && m.Theme != nil {
		return *m.Theme
	}
	return ""
}

func (m *ThemeTemplate) GetEngine() string {
	if m != nil && m.Engine != nil {
		return *m.Engine
	}
	return Default_ThemeTemplate_Engine
}

func (m *ThemeTemplate) GetHideContent() bool {
	if m != nil && m.HideContent != nil {
		return *m.HideContent
	}
	return Default_ThemeTemplate_HideContent
}

func (m *ThemeTemplate) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ThemeTemplate) GetPlaceholders() []*ThemePlaceholder {
	if m != nil {
		return m.Placeholders
	}
	return nil
}

func (m *ThemeTemplate) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

type ThemePlaceholder struct {
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*ThemePlaceholder_Short
	//	*ThemePlaceholder_Text
	//	*ThemePlaceholder_Multiple
	Type             IsThemePlaceholder_Type `protobuf_oneof:"type"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ThemePlaceholder) Reset()                    { *m = ThemePlaceholder{} }
func (m *ThemePlaceholder) String() string            { return proto.CompactTextString(m) }
func (*ThemePlaceholder) ProtoMessage()               {}
func (*ThemePlaceholder) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type IsThemePlaceholder_Type interface {
	IsThemePlaceholder_Type()
}

type ThemePlaceholder_Short struct {
	Short *ContentString `protobuf:"bytes,11,opt,name=short,oneof"`
}
type ThemePlaceholder_Text struct {
	Text *ContentText `protobuf:"bytes,12,opt,name=text,oneof"`
}
type ThemePlaceholder_Multiple struct {
	Multiple *ContentMultiple `protobuf:"bytes,13,opt,name=multiple,oneof"`
}

func (*ThemePlaceholder_Short) IsThemePlaceholder_Type()    {}
func (*ThemePlaceholder_Text) IsThemePlaceholder_Type()     {}
func (*ThemePlaceholder_Multiple) IsThemePlaceholder_Type() {}

func (m *ThemePlaceholder) GetType() IsThemePlaceholder_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ThemePlaceholder) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ThemePlaceholder) GetShort() *ContentString {
	if x, ok := m.GetType().(*ThemePlaceholder_Short); ok {
		return x.Short
	}
	return nil
}

func (m *ThemePlaceholder) GetText() *ContentText {
	if x, ok := m.GetType().(*ThemePlaceholder_Text); ok {
		return x.Text
	}
	return nil
}

func (m *ThemePlaceholder) GetMultiple() *ContentMultiple {
	if x, ok := m.GetType().(*ThemePlaceholder_Multiple); ok {
		return x.Multiple
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ThemePlaceholder) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ThemePlaceholder_OneofMarshaler, _ThemePlaceholder_OneofUnmarshaler, _ThemePlaceholder_OneofSizer, []interface{}{
		(*ThemePlaceholder_Short)(nil),
		(*ThemePlaceholder_Text)(nil),
		(*ThemePlaceholder_Multiple)(nil),
	}
}

func _ThemePlaceholder_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ThemePlaceholder)
	// type
	switch x := m.Type.(type) {
	case *ThemePlaceholder_Short:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Short); err != nil {
			return err
		}
	case *ThemePlaceholder_Text:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *ThemePlaceholder_Multiple:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Multiple); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ThemePlaceholder.Type has unexpected type %T", x)
	}
	return nil
}

func _ThemePlaceholder_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ThemePlaceholder)
	switch tag {
	case 11: // type.short
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentString)
		err := b.DecodeMessage(msg)
		m.Type = &ThemePlaceholder_Short{msg}
		return true, err
	case 12: // type.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentText)
		err := b.DecodeMessage(msg)
		m.Type = &ThemePlaceholder_Text{msg}
		return true, err
	case 13: // type.multiple
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentMultiple)
		err := b.DecodeMessage(msg)
		m.Type = &ThemePlaceholder_Multiple{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ThemePlaceholder_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ThemePlaceholder)
	// type
	switch x := m.Type.(type) {
	case *ThemePlaceholder_Short:
		s := proto.Size(x.Short)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ThemePlaceholder_Text:
		s := proto.Size(x.Text)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ThemePlaceholder_Multiple:
		s := proto.Size(x.Multiple)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ThemeAsset struct {
	Uuid             *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Theme            *string `protobuf:"bytes,3,opt,name=theme" json:"theme,omitempty"`
	Data             *string `protobuf:"bytes,10,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ThemeAsset) Reset()                    { *m = ThemeAsset{} }
func (m *ThemeAsset) String() string            { return proto.CompactTextString(m) }
func (*ThemeAsset) ProtoMessage()               {}
func (*ThemeAsset) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *ThemeAsset) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ThemeAsset) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ThemeAsset) GetTheme() string {
	if m != nil && m.Theme != nil {
		return *m.Theme
	}
	return ""
}

func (m *ThemeAsset) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Theme)(nil), "ketchup.models.Theme")
	proto.RegisterType((*ThemeTemplate)(nil), "ketchup.models.ThemeTemplate")
	proto.RegisterType((*ThemePlaceholder)(nil), "ketchup.models.ThemePlaceholder")
	proto.RegisterType((*ThemeAsset)(nil), "ketchup.models.ThemeAsset")
}
func (m *Theme) SetUuid(v *string) {
	m.Uuid = v
}

func (m *Theme) SetName(v *string) {
	m.Name = v
}

func (m *Theme) SetDescription(v *string) {
	m.Description = v
}

func (m *Theme) SetPackage(v *ketchup_packages.Package) {
	m.Package = v
}

func (m *Theme) SetTemplates(v map[string]*ThemeTemplate) {
	m.Templates = v
}

func (m *Theme) SetAssets(v map[string]*ThemeAsset) {
	m.Assets = v
}

func (m *Theme) SetPlaceholders(v []*Data) {
	m.Placeholders = v
}

func (m *ThemeTemplate) SetUuid(v *string) {
	m.Uuid = v
}

func (m *ThemeTemplate) SetName(v *string) {
	m.Name = v
}

func (m *ThemeTemplate) SetTheme(v *string) {
	m.Theme = v
}

func (m *ThemeTemplate) SetEngine(v *string) {
	m.Engine = v
}

func (m *ThemeTemplate) SetHideContent(v *bool) {
	m.HideContent = v
}

func (m *ThemeTemplate) SetDescription(v *string) {
	m.Description = v
}

func (m *ThemeTemplate) SetPlaceholders(v []*ThemePlaceholder) {
	m.Placeholders = v
}

func (m *ThemeTemplate) SetData(v *string) {
	m.Data = v
}

func (m *ThemePlaceholder) SetKey(v *string) {
	m.Key = v
}

func (m *ThemeAsset) SetUuid(v *string) {
	m.Uuid = v
}

func (m *ThemeAsset) SetName(v *string) {
	m.Name = v
}

func (m *ThemeAsset) SetTheme(v *string) {
	m.Theme = v
}

func (m *ThemeAsset) SetData(v *string) {
	m.Data = v
}

func init() { proto.RegisterFile("theme.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x8a, 0x13, 0x4d,
	0x10, 0xcd, 0x24, 0x99, 0x7c, 0x9b, 0x9a, 0x24, 0x2c, 0xcd, 0x5e, 0xf4, 0x17, 0x15, 0xc7, 0x20,
	0x98, 0xab, 0x41, 0xb3, 0x08, 0x6b, 0xc0, 0x0b, 0xd7, 0x15, 0x72, 0x23, 0x2c, 0x6d, 0xbc, 0x12,
	0x84, 0x26, 0x53, 0x66, 0x86, 0xcc, 0x1f, 0xd3, 0x15, 0xd9, 0xbc, 0x85, 0x2f, 0xe7, 0x93, 0xf8,
	0x02, 0x32, 0xdd, 0x9d, 0x4d, 0x26, 0xce, 0x82, 0xe0, 0x5d, 0x75, 0xd5, 0x39, 0x67, 0xaa, 0xea,
	0xd4, 0x80, 0x47, 0x11, 0xa6, 0x18, 0x14, 0x65, 0x4e, 0x39, 0x1b, 0x6d, 0x90, 0x56, 0xd1, 0xb6,
	0x08, 0xd2, 0x3c, 0xc4, 0x44, 0x8d, 0xa1, 0x90, 0x6b, 0x5b, 0x1b, 0x0f, 0x57, 0x79, 0x46, 0x98,
	0x91, 0x7d, 0x8e, 0x0a, 0xb9, 0xda, 0xc8, 0x35, 0x2a, 0xfb, 0x86, 0x50, 0x92, 0x34, 0xf1, 0xe4,
	0x57, 0x07, 0xdc, 0x65, 0x25, 0xcb, 0x18, 0x74, 0xb7, 0xdb, 0x38, 0xe4, 0x8e, 0xef, 0x4c, 0xfb,
	0x42, 0xc7, 0x55, 0x2e, 0x93, 0x29, 0xf2, 0xb6, 0xc9, 0x55, 0x31, 0xf3, 0xc1, 0x0b, 0x51, 0xad,
	0xca, 0xb8, 0xa0, 0x38, 0xcf, 0x78, 0x47, 0x97, 0x8e, 0x53, 0xec, 0x12, 0xfe, 0xb3, 0x5f, 0xe4,
	0x5d, 0xdf, 0x99, 0x7a, 0xb3, 0xff, 0x83, 0x7d, 0xb3, 0xf7, 0x9d, 0xdc, 0x9a, 0x40, 0xec, 0x91,
	0xec, 0x1a, 0xfa, 0x84, 0x69, 0x91, 0x48, 0x42, 0xc5, 0xc1, 0xef, 0x4c, 0xbd, 0xd9, 0xf3, 0xa0,
	0x3e, 0x63, 0xa0, 0x1b, 0x0d, 0x96, 0x7b, 0xd8, 0x87, 0x8c, 0xca, 0x9d, 0x38, 0xd0, 0xd8, 0x1b,
	0xe8, 0x49, 0xa5, 0x90, 0x14, 0xf7, 0xb4, 0xc0, 0xb3, 0x66, 0x81, 0x77, 0x1a, 0x63, 0xd8, 0x96,
	0xc0, 0xae, 0x60, 0x50, 0x24, 0x72, 0x85, 0x51, 0x9e, 0x84, 0x58, 0x2a, 0x3e, 0xd0, 0x02, 0x17,
	0xa7, 0x02, 0x37, 0x92, 0xa4, 0xa8, 0x21, 0xc7, 0x5f, 0x60, 0x54, 0xef, 0x88, 0x9d, 0x43, 0x67,
	0x83, 0x3b, 0xbb, 0xc8, 0x2a, 0x64, 0x97, 0xe0, 0x7e, 0x97, 0xc9, 0xd6, 0x2c, 0xd2, 0x9b, 0x3d,
	0x69, 0xec, 0x6b, 0xaf, 0x22, 0x0c, 0x76, 0xde, 0xbe, 0x72, 0xc6, 0x9f, 0xc1, 0x3b, 0xea, 0xb6,
	0x41, 0xf9, 0x65, 0x5d, 0x79, 0xdc, 0xa8, 0xac, 0x25, 0x8e, 0x64, 0x27, 0x3f, 0xda, 0x30, 0xac,
	0x7d, 0xf3, 0xaf, 0xdd, 0xbf, 0x00, 0x57, 0x5f, 0xa1, 0xf5, 0xdd, 0x3c, 0xd8, 0x63, 0xe8, 0x61,
	0xb6, 0x8e, 0x33, 0x63, 0x78, 0x7f, 0xde, 0x8d, 0x28, 0x4d, 0x84, 0xcd, 0xb1, 0x17, 0xe0, 0x45,
	0x71, 0x88, 0xef, 0xcd, 0x51, 0x72, 0xd7, 0x77, 0xa6, 0x67, 0x73, 0xf7, 0x9b, 0x4c, 0x14, 0x8a,
	0xe3, 0xca, 0xe9, 0x69, 0xf5, 0xfe, 0x3c, 0xad, 0x9b, 0x13, 0x9b, 0xcc, 0xa1, 0xf8, 0x8d, 0x53,
	0xdf, 0x1e, 0x80, 0x75, 0xcb, 0xaa, 0xc1, 0xaa, 0x5f, 0x80, 0x7b, 0x66, 0xb0, 0x2a, 0x9e, 0xfc,
	0x74, 0xe0, 0xfc, 0x94, 0xd6, 0xb0, 0xef, 0xd7, 0xe0, 0xaa, 0x28, 0x2f, 0x49, 0x73, 0x1b, 0x9c,
	0xb4, 0xa3, 0x7c, 0xa2, 0x32, 0xce, 0xd6, 0x8b, 0x96, 0x30, 0x68, 0xf6, 0x0a, 0xba, 0x84, 0x77,
	0xc4, 0x07, 0x9a, 0xf5, 0xe8, 0x01, 0xd6, 0x12, 0xef, 0x68, 0xd1, 0x12, 0x1a, 0xca, 0xde, 0xc2,
	0x59, 0xba, 0x4d, 0x28, 0x2e, 0x12, 0xe4, 0x43, 0x4d, 0x7b, 0xfa, 0x00, 0xed, 0xa3, 0x85, 0x2d,
	0x5a, 0xe2, 0x9e, 0x72, 0xdd, 0x83, 0x2e, 0xed, 0x0a, 0x9c, 0x7c, 0x05, 0x38, 0xdc, 0xc0, 0x3f,
	0xda, 0xbc, 0xdf, 0x1b, 0x1c, 0xf6, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x12, 0xfa, 0x7e, 0x2a,
	0x95, 0x04, 0x00, 0x00,
}
