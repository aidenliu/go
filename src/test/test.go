package main

import (
	"fmt"
	"net/url"
	_ "strconv"
	_ "strings"
)

func main() {
	rawurl := "http://115.com?A=1&B=3&C=4&=5"
	escape_url := url.QueryEscape(rawurl)
	fmt.Println(escape_url)
	unescape_url, err := url.QueryUnescape(escape_url)
	if err == nil {
		fmt.Println(unescape_url)
	} else {
		fmt.Println(err)
	}

	url_info, err := url.ParseRequestURI(rawurl)
	if err == nil {
		fmt.Println(url_info.Host)
		fmt.Println(url_info.Scheme)
		fmt.Println(url_info.User)
		fmt.Println(url_info.RawQuery)
		fmt.Println(url_info.Path)
		fmt.Println(url_info.Fragment)
		fmt.Println(url_info.Opaque)

		is_abs := url_info.IsAbs()
		if is_abs {
			fmt.Println(rawurl + " is a abs url")
		}

		values := url_info.Query()
		fmt.Println(values)

		request_url := url_info.RequestURI()
		fmt.Println("RequestURI:", request_url)

		Parse_values, _ := url.ParseQuery("A=1&B=3&C=4&C=5")
		Parse_values.Add("F", "6")
		fmt.Println("Parse_values:", Parse_values)

		Parse_values.Del("A")
		fmt.Println("Parse_values:", Parse_values)

		encode_values := Parse_values.Encode()
		fmt.Println("encode_values:", encode_values)

		Parse_values.Set("B", "9")
		fmt.Println("Parse_values:", Parse_values)

		c_value := Parse_values.Get("C")
		fmt.Println("c_value:", c_value)
		fmt.Println("c_value_map", Parse_values["C"])

	} else {
		fmt.Print("err:")
		fmt.Println(err)
	}
}
