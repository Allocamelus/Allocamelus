package media

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Media) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "fileType":
			err = z.FileType.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "FileType")
				return
			}
		case "meta":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					err = msgp.WrapError(err, "Meta")
					return
				}
				switch msgp.UnsafeString(field) {
				case "alt":
					z.Meta.Alt, err = dc.ReadString()
					if err != nil {
						err = msgp.WrapError(err, "Meta", "Alt")
						return
					}
				case "width":
					z.Meta.Width, err = dc.ReadInt64()
					if err != nil {
						err = msgp.WrapError(err, "Meta", "Width")
						return
					}
				case "height":
					z.Meta.Height, err = dc.ReadInt64()
					if err != nil {
						err = msgp.WrapError(err, "Meta", "Height")
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						err = msgp.WrapError(err, "Meta")
						return
					}
				}
			}
		case "url":
			z.Url, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Url")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Media) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "fileType"
	err = en.Append(0x83, 0xa9, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = z.FileType.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "FileType")
		return
	}
	// write "meta"
	err = en.Append(0xa4, 0x6d, 0x65, 0x74, 0x61)
	if err != nil {
		return
	}
	// map header, size 3
	// write "alt"
	err = en.Append(0x83, 0xa3, 0x61, 0x6c, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Meta.Alt)
	if err != nil {
		err = msgp.WrapError(err, "Meta", "Alt")
		return
	}
	// write "width"
	err = en.Append(0xa5, 0x77, 0x69, 0x64, 0x74, 0x68)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Meta.Width)
	if err != nil {
		err = msgp.WrapError(err, "Meta", "Width")
		return
	}
	// write "height"
	err = en.Append(0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Meta.Height)
	if err != nil {
		err = msgp.WrapError(err, "Meta", "Height")
		return
	}
	// write "url"
	err = en.Append(0xa3, 0x75, 0x72, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Url)
	if err != nil {
		err = msgp.WrapError(err, "Url")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Media) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "fileType"
	o = append(o, 0x83, 0xa9, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65)
	o, err = z.FileType.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "FileType")
		return
	}
	// string "meta"
	o = append(o, 0xa4, 0x6d, 0x65, 0x74, 0x61)
	// map header, size 3
	// string "alt"
	o = append(o, 0x83, 0xa3, 0x61, 0x6c, 0x74)
	o = msgp.AppendString(o, z.Meta.Alt)
	// string "width"
	o = append(o, 0xa5, 0x77, 0x69, 0x64, 0x74, 0x68)
	o = msgp.AppendInt64(o, z.Meta.Width)
	// string "height"
	o = append(o, 0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt64(o, z.Meta.Height)
	// string "url"
	o = append(o, 0xa3, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.Url)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Media) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "fileType":
			bts, err = z.FileType.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "FileType")
				return
			}
		case "meta":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					err = msgp.WrapError(err, "Meta")
					return
				}
				switch msgp.UnsafeString(field) {
				case "alt":
					z.Meta.Alt, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Meta", "Alt")
						return
					}
				case "width":
					z.Meta.Width, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Meta", "Width")
						return
					}
				case "height":
					z.Meta.Height, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Meta", "Height")
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						err = msgp.WrapError(err, "Meta")
						return
					}
				}
			}
		case "url":
			z.Url, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Url")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Media) Msgsize() (s int) {
	s = 1 + 10 + z.FileType.Msgsize() + 5 + 1 + 4 + msgp.StringPrefixSize + len(z.Meta.Alt) + 6 + msgp.Int64Size + 7 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Url)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Meta) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "alt":
			z.Alt, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Alt")
				return
			}
		case "width":
			z.Width, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Width")
				return
			}
		case "height":
			z.Height, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Meta) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "alt"
	err = en.Append(0x83, 0xa3, 0x61, 0x6c, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Alt)
	if err != nil {
		err = msgp.WrapError(err, "Alt")
		return
	}
	// write "width"
	err = en.Append(0xa5, 0x77, 0x69, 0x64, 0x74, 0x68)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Width)
	if err != nil {
		err = msgp.WrapError(err, "Width")
		return
	}
	// write "height"
	err = en.Append(0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Height)
	if err != nil {
		err = msgp.WrapError(err, "Height")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Meta) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "alt"
	o = append(o, 0x83, 0xa3, 0x61, 0x6c, 0x74)
	o = msgp.AppendString(o, z.Alt)
	// string "width"
	o = append(o, 0xa5, 0x77, 0x69, 0x64, 0x74, 0x68)
	o = msgp.AppendInt64(o, z.Width)
	// string "height"
	o = append(o, 0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt64(o, z.Height)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Meta) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "alt":
			z.Alt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Alt")
				return
			}
		case "width":
			z.Width, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Width")
				return
			}
		case "height":
			z.Height, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Meta) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Alt) + 6 + msgp.Int64Size + 7 + msgp.Int64Size
	return
}
