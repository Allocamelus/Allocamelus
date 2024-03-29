package comment

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	_ "embed"

	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Comment) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "postId":
			z.PostID, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "PostID")
				return
			}
		case "parentId":
			z.ParentID, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ParentID")
				return
			}
		case "created":
			z.Created, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Created")
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
		case "replies":
			z.Replies, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Replies")
				return
			}
		case "depth":
			z.Depth, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Depth")
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
func (z *Comment) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "id"
	err = en.Append(0x89, 0xa2, 0x69, 0x64)
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
	// write "postId"
	err = en.Append(0xa6, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.PostID)
	if err != nil {
		err = msgp.WrapError(err, "PostID")
		return
	}
	// write "parentId"
	err = en.Append(0xa8, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ParentID)
	if err != nil {
		err = msgp.WrapError(err, "ParentID")
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
	// write "replies"
	err = en.Append(0xa7, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Replies)
	if err != nil {
		err = msgp.WrapError(err, "Replies")
		return
	}
	// write "depth"
	err = en.Append(0xa5, 0x64, 0x65, 0x70, 0x74, 0x68)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Depth)
	if err != nil {
		err = msgp.WrapError(err, "Depth")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Comment) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "id"
	o = append(o, 0x89, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.ID)
	// string "userId"
	o = append(o, 0xa6, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.UserID)
	// string "postId"
	o = append(o, 0xa6, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.PostID)
	// string "parentId"
	o = append(o, 0xa8, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.ParentID)
	// string "created"
	o = append(o, 0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Created)
	// string "updated"
	o = append(o, 0xa7, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Updated)
	// string "content"
	o = append(o, 0xa7, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Content)
	// string "replies"
	o = append(o, 0xa7, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73)
	o = msgp.AppendInt64(o, z.Replies)
	// string "depth"
	o = append(o, 0xa5, 0x64, 0x65, 0x70, 0x74, 0x68)
	o = msgp.AppendInt64(o, z.Depth)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Comment) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "postId":
			z.PostID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PostID")
				return
			}
		case "parentId":
			z.ParentID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ParentID")
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Created")
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
		case "replies":
			z.Replies, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Replies")
				return
			}
		case "depth":
			z.Depth, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Depth")
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
func (z *Comment) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 7 + msgp.Int64Size + 7 + msgp.Int64Size + 9 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.Int64Size + 8 + msgp.StringPrefixSize + len(z.Content) + 8 + msgp.Int64Size + 6 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *List) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "ListComments":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "ListComments")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					err = msgp.WrapError(err, "ListComments")
					return
				}
				switch msgp.UnsafeString(field) {
				default:
					err = dc.Skip()
					if err != nil {
						err = msgp.WrapError(err, "ListComments")
						return
					}
				}
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
func (z *List) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "ListComments"
	err = en.Append(0x81, 0xac, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *List) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "ListComments"
	o = append(o, 0x81, 0xac, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73)
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *List) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ListComments":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ListComments")
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					err = msgp.WrapError(err, "ListComments")
					return
				}
				switch msgp.UnsafeString(field) {
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						err = msgp.WrapError(err, "ListComments")
						return
					}
				}
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
func (z *List) Msgsize() (s int) {
	s = 1 + 13 + 1
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ListComments) DecodeMsg(dc *msgp.Reader) (err error) {
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
func (z ListComments) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ListComments) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ListComments) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
func (z ListComments) Msgsize() (s int) {
	s = 1
	return
}
