// Extracts relevant table of contents data out of your HTML markup.
// Takes all the heading and outputs a JSON of the relevant data (level,
// content and slug).
//
// NOTE: It doesn't handle headings that contains other tags correctly.

package html_toc

import (
	"log"
	"regexp"
	"strings"
)

// CreateTOC creates a table of contents out of the content.
func CreateTOC(s string) string {
	// The following regex will match every heading and its contents.
	// It doesn't handle headings that contains other elements.
	reg := regexp.MustCompile(`<h([1-6])?.*>(.*)</h[1-6]>`)

	matches := reg.FindAllStringSubmatch(s, -1)

	// Create a JSON out of the found heading with the following schema:
	// {
	//   {
	//     "lvl": 1,
	//     "content": "Text",
	//     "slug": "text",
	//   },
	//   {
	//     "lvl": 2,
	//     "content": "Text",
	//     "slug": "text",
	//   },
	// }
	var toc = "{"
	for _, v := range matches {
		toc += "{\"" + "lvl\":" + v[1] + ","
		toc += "\"" + "content\":\"" + v[2] + "\","
		toc += "\"" + "slug\":\"" + CreateSlug(v[2]) + "\",},"
	}
	toc += "}"

	return toc
}

// createSlug creates slug out of a title.
func CreateSlug(s string) string {
	// Remove all characters except for word characters, digits and white
	// space.
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	tmp := reg.ReplaceAllString(s, "")

	// Make all characters lowercase and replace white space with a hyphen
	tmp = strings.Replace(strings.ToLower(tmp), " ", "-", -1)
	return tmp
}

/*

Go HTML TOC extracts relevant table of contents data out of your HTML
markup.
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
