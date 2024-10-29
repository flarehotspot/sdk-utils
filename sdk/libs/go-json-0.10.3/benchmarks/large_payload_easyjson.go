// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package benchmark

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD519278DecodeBenchmark(in *jlexer.Lexer, out *LargePayloadEasyJson) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make(DSUsers, 0, 8)
					} else {
						out.Users = DSUsers{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DSUser
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DSUser)
						}
						easyjsonD519278DecodeBenchmark1(in, v1)
					}
					out.Users = append(out.Users, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				if out.Topics == nil {
					out.Topics = new(DSTopicsList)
				}
				easyjsonD519278DecodeBenchmark2(in, out.Topics)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD519278EncodeBenchmark(out *jwriter.Writer, in LargePayloadEasyJson) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Users\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Users {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjsonD519278EncodeBenchmark1(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Topics\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Topics == nil {
			out.RawString("null")
		} else {
			easyjsonD519278EncodeBenchmark2(out, *in.Topics)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LargePayloadEasyJson) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD519278EncodeBenchmark(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LargePayloadEasyJson) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD519278EncodeBenchmark(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LargePayloadEasyJson) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD519278DecodeBenchmark(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LargePayloadEasyJson) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD519278DecodeBenchmark(l, v)
}
func easyjsonD519278DecodeBenchmark2(in *jlexer.Lexer, out *DSTopicsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				in.Delim('[')
				if out.Topics == nil {
					if !in.IsDelim(']') {
						out.Topics = make(DSTopics, 0, 8)
					} else {
						out.Topics = DSTopics{}
					}
				} else {
					out.Topics = (out.Topics)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *DSTopic
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(DSTopic)
						}
						easyjsonD519278DecodeBenchmark3(in, v4)
					}
					out.Topics = append(out.Topics, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "MoreTopicsUrl":
			out.MoreTopicsUrl = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD519278EncodeBenchmark2(out *jwriter.Writer, in DSTopicsList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Topics\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Topics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Topics {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjsonD519278EncodeBenchmark3(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"MoreTopicsUrl\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MoreTopicsUrl))
	}
	out.RawByte('}')
}
func easyjsonD519278DecodeBenchmark3(in *jlexer.Lexer, out *DSTopic) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = int(in.Int())
		case "Slug":
			out.Slug = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD519278EncodeBenchmark3(out *jwriter.Writer, in DSTopic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"Slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	out.RawByte('}')
}
func easyjsonD519278DecodeBenchmark1(in *jlexer.Lexer, out *DSUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Username":
			out.Username = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD519278EncodeBenchmark1(out *jwriter.Writer, in DSUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	out.RawByte('}')
}
