/*
 * Copyright (c) 2025 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package regexp

import (
	"github.com/dlclark/regexp2"
)

// Expr is the representation of a uncompiled regular expression.
type Expr string

// RegexOptions impact the runtime and parsing behavior
// for each specific regex.  They are setable in code as well
// as in the regex pattern itself.
type RegexOptions = regexp2.RegexOptions

const (
	None                    = regexp2.None
	IgnoreCase              = regexp2.IgnoreCase              // "i"
	Multiline               = regexp2.Multiline               // "m"
	ExplicitCapture         = regexp2.ExplicitCapture         // "n"
	Compiled                = regexp2.Compiled                // "c"
	Singleline              = regexp2.Singleline              // "s"
	IgnorePatternWhitespace = regexp2.IgnorePatternWhitespace // "x"
	RightToLeft             = regexp2.RightToLeft             // "r"
	Debug                   = regexp2.Debug                   // "d"
	ECMAScript              = regexp2.ECMAScript              // "e"
	RE2                     = regexp2.RE2                     // RE2 compatibility mode
	Unicode                 = regexp2.Unicode                 // "u"
)

// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines.
type Regexp = regexp2.Regexp

// Match is a single regex result match that contains groups and repeated captures
//
//		-Groups
//	   -Capture
type Match = regexp2.Match

// MatchEvaluator is a function that takes a match and returns a replacement string to be used
type MatchEvaluator = regexp2.MatchEvaluator

// Group is an explicit or implit (group 0) matched group within the pattern
type Group = regexp2.Group

// Capture is a single capture of text within the larger original string
type Capture = regexp2.Capture

// New returns a new uncompiled regular expression.
func New(expr string) Expr {
	return Expr(expr)
}

// Compile parses the regular expression and returns, if successful,
// a Regexp object that can be used to match against text.
func (e Expr) Compile(opts ...RegexOptions) (*Regexp, error) {
	opt := None
	if len(opts) > 0 {
		opt = opts[0]
	}
	return regexp2.Compile(string(e), opt)
}
