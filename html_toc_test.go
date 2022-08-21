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

	rawHTML, err := ioutil.ReadFile("./raw.html")
	if err != nil {
		log.Fatal(err)
	}

	exJSON, err := ioutil.ReadFile("./result.json")
	if err != nil {
		log.Fatal(err)
	}

	exHTML, err := ioutil.ReadFile("./result.html")
	if err != nil {
		log.Fatal(err)
	}

	resJSON, resHTML := ht.CreateTOC(string(rawHTML))

	is.Equal(string(exJSON), resJSON)
	is.Equal(string(exHTML), resHTML)
}

func TestCreateSamePageSlug(t *testing.T) {
	is := is.New(t)

	resA := ht.CreateSamePageSlug("Looking to become a front-end web developer?")
	resB := ht.CreateSamePageSlug("Introduction to HTML")

	exA := "looking_to_become_a_frontend_web_developer"
	exB := "introduction_to_html"

	is.Equal(exA, resA)
	is.Equal(exB, resB)
}

func TestInsertAnchorTag(t *testing.T) {
	is := is.New(t)

	res := ht.InsertAnchorTag(`<h2>Teknis Pelaksanaan</h2>`)

	ex := `<h2 id="teknis_pelaksanaan"><a href="#teknis_pelaksanaan" title="Permalink to Teknis Pelaksanaan">Teknis Pelaksanaan</a></h2>`

	is.Equal(ex, res)
}

func TestRemoveAnchorTag(t *testing.T) {
	is := is.New(t)

	res := ht.RemoveAnchorTag(`<h2 id="teknis_pelaksanaan"><a href="#teknis_pelaksanaan" title="Permalink to Teknis Pelaksanaan">Teknis Pelaksanaan</a></h2>`)

	ex := `<h2>Teknis Pelaksanaan</h2>`

	is.Equal(ex, res)
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
