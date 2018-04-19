package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		indent++
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				for j := 0; j < indent; j++ {
					buf.WriteByte(' ')
				}
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct:
		buf.WriteByte('(')
		indent++
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				for j := 0; j < indent; j++ {
					buf.WriteByte(' ')
				}
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			x := len(v.Type().Field(i).Name) + 2 // '(Key '
			if err := encode(buf, v.Field(i), indent+x); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map:
		buf.WriteByte('(')
		indent++
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte('\n')
				for j := 0; j < indent; j++ {
					buf.WriteByte(' ')
				}
			}
			buf.WriteByte('(')
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%v", v.Float())
	case reflect.Complex64, reflect.Complex128:
		z := v.Complex()
		fmt.Fprintf(buf, "#C(%v %v)", real(z), imag(z))
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}
	case reflect.Interface:
		buf.WriteByte('(')
		indent++

		elemstr := strconv.Quote(v.Elem().Type().String())
		buf.WriteString(elemstr)
		indent += len(elemstr)

		buf.WriteByte(' ')
		indent++
		if err := encode(buf, v.Elem(), indent); err != nil {
			return err
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
