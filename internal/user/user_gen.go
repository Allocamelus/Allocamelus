package user

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *FollowStruct) DecodeMsg(dc *msgp.Reader) (err error) {
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
func (z FollowStruct) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z FollowStruct) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FollowStruct) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
func (z FollowStruct) Msgsize() (s int) {
	s = 1
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

// DecodeMsg implements msgp.Decodable
func (z *User) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "userName":
			z.UserName, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "UserName")
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "email":
			z.Email, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Email")
				return
			}
		case "bio":
			z.Bio, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Bio")
				return
			}
		case "followers":
			z.Followers, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Followers")
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
		case "permissions":
			err = z.Permissions.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Permissions")
				return
			}
		case "created":
			z.Created, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Created")
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
func (z *User) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "userName"
	err = en.Append(0xa8, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.UserName)
	if err != nil {
		err = msgp.WrapError(err, "UserName")
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "email"
	err = en.Append(0xa5, 0x65, 0x6d, 0x61, 0x69, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Email)
	if err != nil {
		err = msgp.WrapError(err, "Email")
		return
	}
	// write "bio"
	err = en.Append(0xa3, 0x62, 0x69, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bio)
	if err != nil {
		err = msgp.WrapError(err, "Bio")
		return
	}
	// write "followers"
	err = en.Append(0xa9, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Followers)
	if err != nil {
		err = msgp.WrapError(err, "Followers")
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
	// write "permissions"
	err = en.Append(0xab, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return
	}
	err = z.Permissions.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Permissions")
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *User) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "id"
	o = append(o, 0x89, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.ID)
	// string "userName"
	o = append(o, 0xa8, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.UserName)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "email"
	o = append(o, 0xa5, 0x65, 0x6d, 0x61, 0x69, 0x6c)
	o = msgp.AppendString(o, z.Email)
	// string "bio"
	o = append(o, 0xa3, 0x62, 0x69, 0x6f)
	o = msgp.AppendString(o, z.Bio)
	// string "followers"
	o = append(o, 0xa9, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73)
	o = msgp.AppendInt64(o, z.Followers)
	// string "type"
	o = append(o, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendInt8(o, int8(z.Type))
	// string "permissions"
	o = append(o, 0xab, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73)
	o, err = z.Permissions.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Permissions")
		return
	}
	// string "created"
	o = append(o, 0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Created)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *User) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "userName":
			z.UserName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UserName")
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "email":
			z.Email, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Email")
				return
			}
		case "bio":
			z.Bio, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bio")
				return
			}
		case "followers":
			z.Followers, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Followers")
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
		case "permissions":
			bts, err = z.Permissions.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Permissions")
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Created")
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
func (z *User) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 9 + msgp.StringPrefixSize + len(z.UserName) + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.StringPrefixSize + len(z.Email) + 4 + msgp.StringPrefixSize + len(z.Bio) + 10 + msgp.Int64Size + 5 + msgp.Int8Size + 12 + z.Permissions.Msgsize() + 8 + msgp.Int64Size
	return
}
