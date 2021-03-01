package session

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Session) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "loggedIn":
			z.LoggedIn, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "LoggedIn")
				return
			}
		case "userId":
			z.UserID, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "UserID")
				return
			}
		case "perms":
			err = z.Perms.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Perms")
				return
			}
		case "privateKey":
			z.PrivateKey, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "PrivateKey")
				return
			}
		case "loginToken":
			z.LoginToken, err = dc.ReadBytes(z.LoginToken)
			if err != nil {
				err = msgp.WrapError(err, "LoginToken")
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
func (z *Session) EncodeMsg(en *msgp.Writer) (err error) {
	// omitempty: check for empty values
	zb0001Len := uint32(5)
	var zb0001Mask uint8 /* 5 bits */
	if z.PrivateKey == "" {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "loggedIn"
	err = en.Append(0xa8, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteBool(z.LoggedIn)
	if err != nil {
		err = msgp.WrapError(err, "LoggedIn")
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
	// write "perms"
	err = en.Append(0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	if err != nil {
		return
	}
	err = z.Perms.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Perms")
		return
	}
	if (zb0001Mask & 0x8) == 0 { // if not empty
		// write "privateKey"
		err = en.Append(0xaa, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79)
		if err != nil {
			return
		}
		err = en.WriteString(z.PrivateKey)
		if err != nil {
			err = msgp.WrapError(err, "PrivateKey")
			return
		}
	}
	// write "loginToken"
	err = en.Append(0xaa, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.LoginToken)
	if err != nil {
		err = msgp.WrapError(err, "LoginToken")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Session) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(5)
	var zb0001Mask uint8 /* 5 bits */
	if z.PrivateKey == "" {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	// string "loggedIn"
	o = append(o, 0xa8, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e)
	o = msgp.AppendBool(o, z.LoggedIn)
	// string "userId"
	o = append(o, 0xa6, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.UserID)
	// string "perms"
	o = append(o, 0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	o, err = z.Perms.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Perms")
		return
	}
	if (zb0001Mask & 0x8) == 0 { // if not empty
		// string "privateKey"
		o = append(o, 0xaa, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79)
		o = msgp.AppendString(o, z.PrivateKey)
	}
	// string "loginToken"
	o = append(o, 0xaa, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendBytes(o, z.LoginToken)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Session) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "loggedIn":
			z.LoggedIn, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LoggedIn")
				return
			}
		case "userId":
			z.UserID, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UserID")
				return
			}
		case "perms":
			bts, err = z.Perms.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Perms")
				return
			}
		case "privateKey":
			z.PrivateKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PrivateKey")
				return
			}
		case "loginToken":
			z.LoginToken, bts, err = msgp.ReadBytesBytes(bts, z.LoginToken)
			if err != nil {
				err = msgp.WrapError(err, "LoginToken")
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
func (z *Session) Msgsize() (s int) {
	s = 1 + 9 + msgp.BoolSize + 7 + msgp.Int64Size + 6 + z.Perms.Msgsize() + 11 + msgp.StringPrefixSize + len(z.PrivateKey) + 11 + msgp.BytesPrefixSize + len(z.LoginToken)
	return
}