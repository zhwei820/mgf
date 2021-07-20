// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gview_test

// func Test_I18n(t *testing.T) {
// 	gtest.C(t, func(t *gtest.T) {
// 		content := `{{.name}} says "{#hello}{#world}!"`
// 		expect1 := `john says "你好世界!"`
// 		expect2 := `john says "こんにちは世界!"`
// 		expect3 := `john says "{#hello}{#world}!"`

// 		g.I18n().SetPath(gdebug.TestDataPath("i18n"))

// 		g.I18n().SetLanguage("zh-CN")
// 		result1, err := g.View().ParseContent(context.TODO(), content, g.Map{
// 			"name": "john",
// 		})
// 		t.Assert(err, nil)
// 		t.Assert(result1, expect1)

// 		g.I18n().SetLanguage("ja")
// 		result2, err := g.View().ParseContent(context.TODO(), content, g.Map{
// 			"name": "john",
// 		})
// 		t.Assert(err, nil)
// 		t.Assert(result2, expect2)

// 		g.I18n().SetLanguage("none")
// 		result3, err := g.View().ParseContent(context.TODO(), content, g.Map{
// 			"name": "john",
// 		})
// 		t.Assert(err, nil)
// 		t.Assert(result3, expect3)
// 	})
// 	gtest.C(t, func(t *gtest.T) {
// 		content := `{{.name}} says "{#hello}{#world}!"`
// 		expect1 := `john says "你好世界!"`
// 		expect2 := `john says "こんにちは世界!"`
// 		expect3 := `john says "{#hello}{#world}!"`

// 		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")

// 		result1, err := g.View().ParseContent(context.TODO(), content, g.Map{
// 			"name":         "john",
// 			"I18nLanguage": "zh-CN",
// 		})
// 		t.Assert(err, nil)
// 		t.Assert(result1, expect1)

// 		result2, err := g.View().ParseContent(context.TODO(), content, g.Map{
// 			"name":         "john",
// 			"I18nLanguage": "ja",
// 		})
// 		t.Assert(err, nil)
// 		t.Assert(result2, expect2)

// 		result3, err := g.View().ParseContent(context.TODO(), content, g.Map{
// 			"name":         "john",
// 			"I18nLanguage": "none",
// 		})
// 		t.Assert(err, nil)
// 		t.Assert(result3, expect3)
// 	})
// }
