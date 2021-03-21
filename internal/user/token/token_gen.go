package token

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Token) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "type":
			{
				var zb0002 int8
				zb0002, err = dc.ReadInt8()
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = Types(zb0002)
			}
		case "selector":
			z.Selector, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Selector")
				return
			}
		case "tokenHash":
			z.TokenHash, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "TokenHash")
				return
			}
		case "created":
			z.Created, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Created")
				return
			}
		case "expiration":
			z.Expiration, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Expiration")
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
func (z *Token) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "type"
	err = en.Append(0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt8(int8(z.Type))
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "selector"
	err = en.Append(0xa8, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Selector)
	if err != nil {
		err = msgp.WrapError(err, "Selector")
		return
	}
	// write "tokenHash"
	err = en.Append(0xa9, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteString(z.TokenHash)
	if err != nil {
		err = msgp.WrapError(err, "TokenHash")
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
	// write "expiration"
	err = en.Append(0xaa, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Expiration)
	if err != nil {
		err = msgp.WrapError(err, "Expiration")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Token) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "id"
	o = append(o, 0x87, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.ID)
	// string "userId"
	o = append(o, 0xa6, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.UserID)
	// string "type"
	o = append(o, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendInt8(o, int8(z.Type))
	// string "selector"
	o = append(o, 0xa8, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72)
	o = msgp.AppendString(o, z.Selector)
	// string "tokenHash"
	o = append(o, 0xa9, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x61, 0x73, 0x68)
	o = msgp.AppendString(o, z.TokenHash)
	// string "created"
	o = append(o, 0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Created)
	// string "expiration"
	o = append(o, 0xaa, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendInt64(o, z.Expiration)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Token) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "type":
			{
				var zb0002 int8
				zb0002, bts, err = msgp.ReadInt8Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = Types(zb0002)
			}
		case "selector":
			z.Selector, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Selector")
				return
			}
		case "tokenHash":
			z.TokenHash, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TokenHash")
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Created")
				return
			}
		case "expiration":
			z.Expiration, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Expiration")
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
func (z *Token) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 7 + msgp.Int64Size + 5 + msgp.Int8Size + 9 + msgp.StringPrefixSize + len(z.Selector) + 10 + msgp.StringPrefixSize + len(z.TokenHash) + 8 + msgp.Int64Size + 11 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Types) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int8
		zb0001, err = dc.ReadInt8()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Types(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Types) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt8(int8(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Types) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt8(o, int8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Types) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int8
		zb0001, bts, err = msgp.ReadInt8Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Types(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Types) Msgsize() (s int) {
	s = msgp.Int8Size
	return
}
