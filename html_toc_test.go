package html_toc_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/matryer/is"

	ht "git.sr.ht/~toar/go-html-toc"
)

func TestCreateTOC(t *testing.T) {
	is := is.New(t)

	// Taken from
	// https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML
	html, err := ioutil.ReadFile("./test.html")
	if err != nil {
		log.Fatal(err)
	}

	json, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Fatal(err)
	}

	is.Equal(string(json), ht.CreateTOC(string(html)))
}

func TestCreateSlug(t *testing.T) {
	is := is.New(t)

	testStringOne := ht.CreateSlug("Looking to become a front-end web developer?")
	testStringTwo := ht.CreateSlug("Introduction to HTML")

	is.Equal("looking-to-become-a-frontend-web-developer", testStringOne)
	is.Equal("introduction-to-html", testStringTwo)
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
