package post

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Post) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "id":
			z.ID, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "userId":
			z.UserID, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "UserID")
				return
			}
		case "created":
			z.Created, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Created")
				return
			}
		case "published":
			z.Published, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Published")
				return
			}
		case "updated":
			z.Updated, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Updated")
				return
			}
		case "content":
			z.Content, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
		case "media":
			z.Media, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Media")
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
func (z *Post) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "id"
	err = en.Append(0x87, 0xa2, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ID)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	// write "userId"
	err = en.Append(0xa6, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.UserID)
	if err != nil {
		err = msgp.WrapError(err, "UserID")
		return
	}
	// write "created"
	err = en.Append(0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Created)
	if err != nil {
		err = msgp.WrapError(err, "Created")
		return
	}
	// write "published"
	err = en.Append(0xa9, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Published)
	if err != nil {
		err = msgp.WrapError(err, "Published")
		return
	}
	// write "updated"
	err = en.Append(0xa7, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Updated)
	if err != nil {
		err = msgp.WrapError(err, "Updated")
		return
	}
	// write "content"
	err = en.Append(0xa7, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Content)
	if err != nil {
		err = msgp.WrapError(err, "Content")
		return
	}
	// write "media"
	err = en.Append(0xa5, 0x6d, 0x65, 0x64, 0x69, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Media)
	if err != nil {
		err = msgp.WrapError(err, "Media")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Post) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "id"
	o = append(o, 0x87, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.ID)
	// string "userId"
	o = append(o, 0xa6, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.UserID)
	// string "created"
	o = append(o, 0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Created)
	// string "published"
	o = append(o, 0xa9, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Published)
	// string "updated"
	o = append(o, 0xa7, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Updated)
	// string "content"
	o = append(o, 0xa7, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Content)
	// string "media"
	o = append(o, 0xa5, 0x6d, 0x65, 0x64, 0x69, 0x61)
	o = msgp.AppendBool(o, z.Media)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Post) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "id":
			z.ID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "userId":
			z.UserID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UserID")
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Created")
				return
			}
		case "published":
			z.Published, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Published")
				return
			}
		case "updated":
			z.Updated, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Updated")
				return
			}
		case "content":
			z.Content, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
		case "media":
			z.Media, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Media")
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
func (z *Post) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 7 + msgp.Int64Size + 8 + msgp.Int64Size + 10 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.Content) + 6 + msgp.BoolSize
	return
}
