// Extracts relevant table of contents data out of your HTML markup and add
// permalinks to your markup's headings.
// Takes your HTML markup, outputs a string in JSON format of the relevant
// data (level, content and slug) and inserts permalinks to the markup's
// headings.

package html_toc

import (
	"regexp"
	"strings"
)

// CreateTOC creates a table of contents out of the content.
func CreateTOC(s string) (string, string) {
	// Workaround for single line HTML markup.
	// Appends newline on every closing tag.
	reg := regexp.MustCompile("(</[^ ][^<]*>)")
	s = reg.ReplaceAllString(s, "$1\n")

	// Deletes newline after every heading closing tag.
	reg = regexp.MustCompile("\n(</h[1-6][^<]*>)")
	s = reg.ReplaceAllString(s, "$1")

	s = InsertAnchorTag(s)

	// The following regex will match every heading, the anchor tags inside
	// it and the actual content.
	reg = regexp.MustCompile(`<h([1-6])?.*><a?.*>(.*)</a></h[1-6]>`)

	matches := reg.FindAllStringSubmatch(s, -1)

	// Create a JSON out of the found heading with the following schema:
	// {
	//   {
	//     "level": 1,
	//     "content": "Text",
	//     "slug": "text",
	//   },
	//   {
	//     "level": 2,
	//     "content": "Text",
	//     "slug": "text",
	//   },
	// }
	var toc = "["
	for i, v := range matches {
		toc += "{\"" + "level\":" + v[1] + ","
		toc += "\"" + "content\":\"" + v[2] + "\","
		if len(matches) == i+1 {
			toc += "\"" + "slug\":\"" + CreateSamePageSlug(v[2]) + "\"}"
		}
		if len(matches) > i+1 {
			toc += "\"" + "slug\":\"" + CreateSamePageSlug(v[2]) + "\"},"
		}
	}
	toc += "]"

	// Remove all newlines.
	reg = regexp.MustCompile("\n")
	s = reg.ReplaceAllString(s, "")

	return toc, s
}

// CreateSamePageSlug creates slug out of a title.
func CreateSamePageSlug(s string) string {
	// Remove all characters except for word characters, digits and white
	// space.
	reg := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	tmp := reg.ReplaceAllString(s, "")

	// Make all characters lowercase and replace white space with a hyphen
	tmp = strings.Replace(strings.ToLower(tmp), " ", "_", -1)
	return tmp
}

// InsertAnchorTag inserts an anchor tag inside a heading.
func InsertAnchorTag(s string) string {
	// Get all headings and capture the opening tag, content and closing
	// tag.
	reg := regexp.MustCompile(`(<h[1-6]?.*>)(.*)(</h[1-6]>)`)
	matches := reg.FindAllStringSubmatch(s, -1)

	// Insert the anchor tag between the captured parts from all of the
	// headings.
	for _, v := range matches {
		slug := CreateSamePageSlug(v[2])

		reg := regexp.MustCompile(v[1] + v[2] + v[3])

		s = reg.ReplaceAllString(s, v[1][:3]+" id=\""+slug+"\""+v[1][3:]+"<a href=\"#"+slug+"\" title=\"Permalink to "+v[2]+"\">"+v[2]+"</a>"+v[3])
	}

	return s
}

/*

Go HTML TOC extracts relevant table of contents data out of your HTML
markup and add permalinks to your markup's headings.
Copyright (C) 2022  Aranggi J. Toar

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; only version 2 of the License.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

*/
