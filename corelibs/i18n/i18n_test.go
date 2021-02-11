// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package i18n

import (
	"testing"

	"github.com/mattermost/go-i18n/i18n/bundle"
	"github.com/mattermost/go-i18n/i18n/language"
	"github.com/mattermost/go-i18n/i18n/translation"
	"github.com/stretchr/testify/assert"
)

var htmlTestTranslationBundle *bundle.Bundle

func init() {
	htmlTestTranslationBundle = bundle.New()
	fooBold, _ := translation.NewTranslation(map[string]interface{}{
		"id":          "foo.bold",
		"translation": "<p>[[{{ .Foo }}]]</p>",
	})
	htmlTestTranslationBundle.AddTranslation(&language.Language{Tag: "en"}, fooBold)
}

func TestTranslateAsHtml(t *testing.T) {
	assert.EqualValues(t, "<p><strong>&lt;i&gt;foo&lt;/i&gt;</strong></p>", TranslateAsHtml(TranslateFunc(htmlTestTranslationBundle.MustTfunc("en")), "foo.bold", map[string]interface{}{
		"Foo": "<i>foo</i>",
	}))
}

func TestEscapeForHtml(t *testing.T) {
	stringForPointer := "<b>abc</b>"
	for name, tc := range map[string]struct {
		In       interface{}
		Expected interface{}
	}{
		"NoHTML": {
			In:       "abc",
			Expected: "abc",
		},
		"String": {
			In:       "<b>abc</b>",
			Expected: "&lt;b&gt;abc&lt;/b&gt;",
		},
		"StringPointer": {
			In:       &stringForPointer,
			Expected: "&lt;b&gt;abc&lt;/b&gt;",
		},
		"Map": {
			In: map[string]interface{}{
				"abc": "abc",
				"123": "<b>123</b>",
			},
			Expected: map[string]interface{}{
				"abc": "abc",
				"123": "&lt;b&gt;123&lt;/b&gt;",
			},
		},
		"Unsupported": {
			In:       struct{ string }{"<b>abc</b>"},
			Expected: "",
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, escapeForHtml(tc.In))
		})
	}
}
