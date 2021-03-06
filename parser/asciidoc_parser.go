package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/bytesparadise/libasciidoc/types"
)

// *****************************************************************************************
// This file is generated after its sibling `asciidoc-grammar.peg` file. DO NOT MODIFY !
// *****************************************************************************************

var g = &grammar{
	rules: []*rule{
		{
			name: "Document",
			pos:  position{line: 19, col: 1, offset: 501},
			expr: &actionExpr{
				pos: position{line: 19, col: 13, offset: 513},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 19, col: 13, offset: 513},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 19, col: 13, offset: 513},
							label: "frontMatter",
							expr: &zeroOrOneExpr{
								pos: position{line: 19, col: 26, offset: 526},
								expr: &ruleRefExpr{
									pos:  position{line: 19, col: 26, offset: 526},
									name: "FrontMatter",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 19, col: 40, offset: 540},
							label: "documentHeader",
							expr: &zeroOrOneExpr{
								pos: position{line: 19, col: 56, offset: 556},
								expr: &ruleRefExpr{
									pos:  position{line: 19, col: 56, offset: 556},
									name: "DocumentHeader",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 19, col: 73, offset: 573},
							label: "blocks",
							expr: &ruleRefExpr{
								pos:  position{line: 19, col: 81, offset: 581},
								name: "DocumentBlocks",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 19, col: 97, offset: 597},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlocks",
			pos:  position{line: 23, col: 1, offset: 685},
			expr: &choiceExpr{
				pos: position{line: 23, col: 19, offset: 703},
				alternatives: []interface{}{
					&labeledExpr{
						pos:   position{line: 23, col: 19, offset: 703},
						label: "content",
						expr: &seqExpr{
							pos: position{line: 23, col: 28, offset: 712},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 23, col: 28, offset: 712},
									name: "Preamble",
								},
								&oneOrMoreExpr{
									pos: position{line: 23, col: 37, offset: 721},
									expr: &ruleRefExpr{
										pos:  position{line: 23, col: 37, offset: 721},
										name: "Section",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 23, col: 49, offset: 733},
						run: (*parser).callonDocumentBlocks7,
						expr: &labeledExpr{
							pos:   position{line: 23, col: 49, offset: 733},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 23, col: 58, offset: 742},
								expr: &ruleRefExpr{
									pos:  position{line: 23, col: 58, offset: 742},
									name: "BlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BlockElement",
			pos:  position{line: 27, col: 1, offset: 786},
			expr: &choiceExpr{
				pos: position{line: 27, col: 17, offset: 802},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 27, col: 17, offset: 802},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 48, offset: 833},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 73, offset: 858},
						name: "TableOfContentsMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 96, offset: 881},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 109, offset: 894},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 116, offset: 901},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 131, offset: 916},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 148, offset: 933},
						name: "Admonition",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 161, offset: 946},
						name: "Paragraph",
					},
					&seqExpr{
						pos: position{line: 27, col: 174, offset: 959},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 27, col: 174, offset: 959},
								name: "ElementAttribute",
							},
							&ruleRefExpr{
								pos:  position{line: 27, col: 191, offset: 976},
								name: "EOL",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 198, offset: 983},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 29, col: 1, offset: 1038},
			expr: &actionExpr{
				pos: position{line: 29, col: 13, offset: 1050},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 29, col: 13, offset: 1050},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 29, col: 23, offset: 1060},
						expr: &ruleRefExpr{
							pos:  position{line: 29, col: 23, offset: 1060},
							name: "BlockElement",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 36, col: 1, offset: 1243},
			expr: &ruleRefExpr{
				pos:  position{line: 36, col: 16, offset: 1258},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 38, col: 1, offset: 1276},
			expr: &actionExpr{
				pos: position{line: 38, col: 16, offset: 1291},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 38, col: 16, offset: 1291},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 38, col: 16, offset: 1291},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 38, col: 37, offset: 1312},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 38, col: 46, offset: 1321},
								name: "YamlFrontMatterContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 38, col: 70, offset: 1345},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 42, col: 1, offset: 1425},
			expr: &seqExpr{
				pos: position{line: 42, col: 26, offset: 1450},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 42, col: 26, offset: 1450},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 42, col: 32, offset: 1456},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "YamlFrontMatterContent",
			pos:  position{line: 44, col: 1, offset: 1461},
			expr: &actionExpr{
				pos: position{line: 44, col: 27, offset: 1487},
				run: (*parser).callonYamlFrontMatterContent1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 44, col: 27, offset: 1487},
					expr: &seqExpr{
						pos: position{line: 44, col: 28, offset: 1488},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 44, col: 28, offset: 1488},
								expr: &ruleRefExpr{
									pos:  position{line: 44, col: 29, offset: 1489},
									name: "YamlFrontMatterToken",
								},
							},
							&anyMatcher{
								line: 44, col: 50, offset: 1510,
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 52, col: 1, offset: 1734},
			expr: &actionExpr{
				pos: position{line: 52, col: 19, offset: 1752},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 52, col: 19, offset: 1752},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 52, col: 19, offset: 1752},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 27, offset: 1760},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 42, offset: 1775},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 52, col: 51, offset: 1784},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 51, offset: 1784},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 69, offset: 1802},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 52, col: 79, offset: 1812},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 79, offset: 1812},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 98, offset: 1831},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 52, col: 115, offset: 1848},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 115, offset: 1848},
									name: "DocumentAttributeDeclaration",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentTitle",
			pos:  position{line: 56, col: 1, offset: 1979},
			expr: &actionExpr{
				pos: position{line: 56, col: 18, offset: 1996},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 56, col: 18, offset: 1996},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 56, col: 18, offset: 1996},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 56, col: 29, offset: 2007},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 30, offset: 2008},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 49, offset: 2027},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 56, col: 56, offset: 2034},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 56, col: 61, offset: 2039},
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 61, offset: 2039},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 65, offset: 2043},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 74, offset: 2052},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 56, col: 89, offset: 2067},
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 89, offset: 2067},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 93, offset: 2071},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 56, col: 96, offset: 2074},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 97, offset: 2075},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 56, col: 115, offset: 2093},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 60, col: 1, offset: 2208},
			expr: &choiceExpr{
				pos: position{line: 60, col: 20, offset: 2227},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 60, col: 20, offset: 2227},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 48, offset: 2255},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 62, col: 1, offset: 2285},
			expr: &actionExpr{
				pos: position{line: 62, col: 30, offset: 2314},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 62, col: 30, offset: 2314},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 62, col: 30, offset: 2314},
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 30, offset: 2314},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 62, col: 34, offset: 2318},
							expr: &litMatcher{
								pos:        position{line: 62, col: 35, offset: 2319},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 39, offset: 2323},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 48, offset: 2332},
								expr: &ruleRefExpr{
									pos:  position{line: 62, col: 48, offset: 2332},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 62, col: 65, offset: 2349},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 66, col: 1, offset: 2419},
			expr: &actionExpr{
				pos: position{line: 66, col: 33, offset: 2451},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 66, col: 33, offset: 2451},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 66, col: 33, offset: 2451},
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 33, offset: 2451},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 66, col: 37, offset: 2455},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 66, col: 48, offset: 2466},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 56, offset: 2474},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 70, col: 1, offset: 2565},
			expr: &actionExpr{
				pos: position{line: 70, col: 19, offset: 2583},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 19, offset: 2583},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 19, offset: 2583},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 19, offset: 2583},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 2587},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 34, offset: 2598},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 58, offset: 2622},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 68, offset: 2632},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 69, offset: 2633},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 94, offset: 2658},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 104, offset: 2668},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 105, offset: 2669},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 130, offset: 2694},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 136, offset: 2700},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 137, offset: 2701},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 159, offset: 2723},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 159, offset: 2723},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 70, col: 163, offset: 2727},
							expr: &litMatcher{
								pos:        position{line: 70, col: 163, offset: 2727},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 168, offset: 2732},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 168, offset: 2732},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 75, col: 1, offset: 2897},
			expr: &seqExpr{
				pos: position{line: 75, col: 27, offset: 2923},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 75, col: 27, offset: 2923},
						expr: &litMatcher{
							pos:        position{line: 75, col: 28, offset: 2924},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 75, col: 32, offset: 2928},
						expr: &litMatcher{
							pos:        position{line: 75, col: 33, offset: 2929},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 37, offset: 2933},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 75, col: 48, offset: 2944},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 48, offset: 2944},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 77, col: 1, offset: 2949},
			expr: &seqExpr{
				pos: position{line: 77, col: 24, offset: 2972},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 77, col: 24, offset: 2972},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 77, col: 28, offset: 2976},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 77, col: 34, offset: 2982},
							expr: &seqExpr{
								pos: position{line: 77, col: 35, offset: 2983},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 77, col: 35, offset: 2983},
										expr: &litMatcher{
											pos:        position{line: 77, col: 36, offset: 2984},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 77, col: 40, offset: 2988},
										expr: &ruleRefExpr{
											pos:  position{line: 77, col: 41, offset: 2989},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 77, col: 45, offset: 2993,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 77, col: 49, offset: 2997},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 81, col: 1, offset: 3133},
			expr: &actionExpr{
				pos: position{line: 81, col: 21, offset: 3153},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 81, col: 21, offset: 3153},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 21, offset: 3153},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 21, offset: 3153},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 25, offset: 3157},
							expr: &litMatcher{
								pos:        position{line: 81, col: 26, offset: 3158},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 3162},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 40, offset: 3172},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 41, offset: 3173},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 66, offset: 3198},
							expr: &litMatcher{
								pos:        position{line: 81, col: 66, offset: 3198},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 71, offset: 3203},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 79, offset: 3211},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 80, offset: 3212},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 103, offset: 3235},
							expr: &litMatcher{
								pos:        position{line: 81, col: 103, offset: 3235},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 108, offset: 3240},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 118, offset: 3250},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 119, offset: 3251},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 144, offset: 3276},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 86, col: 1, offset: 3449},
			expr: &choiceExpr{
				pos: position{line: 86, col: 27, offset: 3475},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 86, col: 27, offset: 3475},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 86, col: 27, offset: 3475},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 32, offset: 3480},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 39, offset: 3487},
								expr: &seqExpr{
									pos: position{line: 86, col: 40, offset: 3488},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 40, offset: 3488},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 41, offset: 3489},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 45, offset: 3493},
											expr: &litMatcher{
												pos:        position{line: 86, col: 46, offset: 3494},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 50, offset: 3498},
											expr: &litMatcher{
												pos:        position{line: 86, col: 51, offset: 3499},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 55, offset: 3503,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 86, col: 61, offset: 3509},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 86, col: 61, offset: 3509},
								expr: &litMatcher{
									pos:        position{line: 86, col: 61, offset: 3509},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 67, offset: 3515},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 74, offset: 3522},
								expr: &seqExpr{
									pos: position{line: 86, col: 75, offset: 3523},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 75, offset: 3523},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 76, offset: 3524},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 80, offset: 3528},
											expr: &litMatcher{
												pos:        position{line: 86, col: 81, offset: 3529},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 85, offset: 3533},
											expr: &litMatcher{
												pos:        position{line: 86, col: 86, offset: 3534},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 90, offset: 3538,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 94, offset: 3542},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 94, offset: 3542},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 86, col: 98, offset: 3546},
								expr: &litMatcher{
									pos:        position{line: 86, col: 99, offset: 3547},
									val:        ",",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionDate",
			pos:  position{line: 87, col: 1, offset: 3551},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 25, offset: 3575},
				expr: &seqExpr{
					pos: position{line: 87, col: 26, offset: 3576},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 26, offset: 3576},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 27, offset: 3577},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 87, col: 31, offset: 3581},
							expr: &litMatcher{
								pos:        position{line: 87, col: 32, offset: 3582},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 87, col: 36, offset: 3586,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 88, col: 1, offset: 3591},
			expr: &zeroOrMoreExpr{
				pos: position{line: 88, col: 27, offset: 3617},
				expr: &seqExpr{
					pos: position{line: 88, col: 28, offset: 3618},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 88, col: 28, offset: 3618},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 29, offset: 3619},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 88, col: 33, offset: 3623,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 93, col: 1, offset: 3743},
			expr: &choiceExpr{
				pos: position{line: 93, col: 33, offset: 3775},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 33, offset: 3775},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 76, offset: 3818},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 95, col: 1, offset: 3865},
			expr: &actionExpr{
				pos: position{line: 95, col: 45, offset: 3909},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 95, col: 45, offset: 3909},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 45, offset: 3909},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 95, col: 49, offset: 3913},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 55, offset: 3919},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 70, offset: 3934},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 95, col: 74, offset: 3938},
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 74, offset: 3938},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 78, offset: 3942},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 99, col: 1, offset: 4027},
			expr: &actionExpr{
				pos: position{line: 99, col: 49, offset: 4075},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 99, col: 49, offset: 4075},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 99, col: 49, offset: 4075},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 99, col: 53, offset: 4079},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 59, offset: 4085},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 99, col: 74, offset: 4100},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 99, col: 78, offset: 4104},
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 78, offset: 4104},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 99, col: 82, offset: 4108},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 99, col: 88, offset: 4114},
								expr: &seqExpr{
									pos: position{line: 99, col: 89, offset: 4115},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 99, col: 89, offset: 4115},
											expr: &ruleRefExpr{
												pos:  position{line: 99, col: 90, offset: 4116},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 99, col: 98, offset: 4124,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 99, col: 102, offset: 4128},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 103, col: 1, offset: 4231},
			expr: &choiceExpr{
				pos: position{line: 103, col: 27, offset: 4257},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 27, offset: 4257},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 78, offset: 4308},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 105, col: 1, offset: 4354},
			expr: &actionExpr{
				pos: position{line: 105, col: 53, offset: 4406},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 105, col: 53, offset: 4406},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 53, offset: 4406},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 105, col: 58, offset: 4411},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 64, offset: 4417},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 105, col: 79, offset: 4432},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 105, col: 83, offset: 4436},
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 83, offset: 4436},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 87, offset: 4440},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 109, col: 1, offset: 4514},
			expr: &actionExpr{
				pos: position{line: 109, col: 49, offset: 4562},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 109, col: 49, offset: 4562},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 109, col: 49, offset: 4562},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 109, col: 53, offset: 4566},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 59, offset: 4572},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 109, col: 74, offset: 4587},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 109, col: 79, offset: 4592},
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 79, offset: 4592},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 83, offset: 4596},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 113, col: 1, offset: 4670},
			expr: &actionExpr{
				pos: position{line: 113, col: 34, offset: 4703},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 113, col: 34, offset: 4703},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 113, col: 34, offset: 4703},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 113, col: 38, offset: 4707},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 113, col: 44, offset: 4713},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 113, col: 59, offset: 4728},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 120, col: 1, offset: 4982},
			expr: &seqExpr{
				pos: position{line: 120, col: 18, offset: 4999},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 120, col: 19, offset: 5000},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 120, col: 19, offset: 5000},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 27, offset: 5008},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 35, offset: 5016},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 120, col: 43, offset: 5024},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 120, col: 48, offset: 5029},
						expr: &choiceExpr{
							pos: position{line: 120, col: 49, offset: 5030},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 120, col: 49, offset: 5030},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 57, offset: 5038},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 65, offset: 5046},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 120, col: 73, offset: 5054},
									val:        "-",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "TableOfContentsMacro",
			pos:  position{line: 125, col: 1, offset: 5174},
			expr: &seqExpr{
				pos: position{line: 125, col: 25, offset: 5198},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 125, col: 25, offset: 5198},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 35, offset: 5208},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 130, col: 1, offset: 5321},
			expr: &choiceExpr{
				pos: position{line: 130, col: 12, offset: 5332},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 130, col: 12, offset: 5332},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 23, offset: 5343},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 34, offset: 5354},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 45, offset: 5365},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 56, offset: 5376},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 133, col: 1, offset: 5387},
			expr: &actionExpr{
				pos: position{line: 133, col: 13, offset: 5399},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 133, col: 13, offset: 5399},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 133, col: 13, offset: 5399},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 133, col: 21, offset: 5407},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 133, col: 36, offset: 5422},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 133, col: 46, offset: 5432},
								expr: &ruleRefExpr{
									pos:  position{line: 133, col: 46, offset: 5432},
									name: "Section1Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section1Block",
			pos:  position{line: 137, col: 1, offset: 5539},
			expr: &actionExpr{
				pos: position{line: 137, col: 18, offset: 5556},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 137, col: 18, offset: 5556},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 137, col: 18, offset: 5556},
							expr: &ruleRefExpr{
								pos:  position{line: 137, col: 19, offset: 5557},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 137, col: 28, offset: 5566},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 137, col: 37, offset: 5575},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 137, col: 37, offset: 5575},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 48, offset: 5586},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 59, offset: 5597},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 70, offset: 5608},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 81, offset: 5619},
										name: "BlockElement",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2",
			pos:  position{line: 141, col: 1, offset: 5681},
			expr: &actionExpr{
				pos: position{line: 141, col: 13, offset: 5693},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 141, col: 13, offset: 5693},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 141, col: 13, offset: 5693},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 141, col: 21, offset: 5701},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 141, col: 36, offset: 5716},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 141, col: 46, offset: 5726},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 46, offset: 5726},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 141, col: 62, offset: 5742},
							expr: &zeroOrMoreExpr{
								pos: position{line: 141, col: 63, offset: 5743},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 64, offset: 5744},
									name: "Section2",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Block",
			pos:  position{line: 145, col: 1, offset: 5846},
			expr: &actionExpr{
				pos: position{line: 145, col: 18, offset: 5863},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 145, col: 18, offset: 5863},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 145, col: 18, offset: 5863},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 19, offset: 5864},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 145, col: 28, offset: 5873},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 29, offset: 5874},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 38, offset: 5883},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 145, col: 47, offset: 5892},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 145, col: 47, offset: 5892},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 58, offset: 5903},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 69, offset: 5914},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 80, offset: 5925},
										name: "BlockElement",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3",
			pos:  position{line: 149, col: 1, offset: 5987},
			expr: &actionExpr{
				pos: position{line: 149, col: 13, offset: 5999},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 149, col: 13, offset: 5999},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 149, col: 13, offset: 5999},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 149, col: 21, offset: 6007},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 149, col: 36, offset: 6022},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 149, col: 46, offset: 6032},
								expr: &ruleRefExpr{
									pos:  position{line: 149, col: 46, offset: 6032},
									name: "Section3Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Block",
			pos:  position{line: 153, col: 1, offset: 6139},
			expr: &actionExpr{
				pos: position{line: 153, col: 18, offset: 6156},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 153, col: 18, offset: 6156},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 153, col: 18, offset: 6156},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 19, offset: 6157},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 153, col: 28, offset: 6166},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 29, offset: 6167},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 153, col: 38, offset: 6176},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 39, offset: 6177},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 153, col: 48, offset: 6186},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 153, col: 57, offset: 6195},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 153, col: 57, offset: 6195},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 153, col: 68, offset: 6206},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 153, col: 79, offset: 6217},
										name: "BlockElement",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4",
			pos:  position{line: 157, col: 1, offset: 6279},
			expr: &actionExpr{
				pos: position{line: 157, col: 13, offset: 6291},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 157, col: 13, offset: 6291},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 157, col: 13, offset: 6291},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 157, col: 21, offset: 6299},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 157, col: 36, offset: 6314},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 157, col: 46, offset: 6324},
								expr: &ruleRefExpr{
									pos:  position{line: 157, col: 46, offset: 6324},
									name: "Section4Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Block",
			pos:  position{line: 161, col: 1, offset: 6431},
			expr: &actionExpr{
				pos: position{line: 161, col: 18, offset: 6448},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 161, col: 18, offset: 6448},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 161, col: 18, offset: 6448},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 19, offset: 6449},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 28, offset: 6458},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 29, offset: 6459},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 38, offset: 6468},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 39, offset: 6469},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 48, offset: 6478},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 49, offset: 6479},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 58, offset: 6488},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 161, col: 67, offset: 6497},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 161, col: 67, offset: 6497},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 161, col: 78, offset: 6508},
										name: "BlockElement",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5",
			pos:  position{line: 165, col: 1, offset: 6570},
			expr: &actionExpr{
				pos: position{line: 165, col: 13, offset: 6582},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 165, col: 13, offset: 6582},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 165, col: 13, offset: 6582},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 21, offset: 6590},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 165, col: 36, offset: 6605},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 165, col: 46, offset: 6615},
								expr: &ruleRefExpr{
									pos:  position{line: 165, col: 46, offset: 6615},
									name: "Section5Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Block",
			pos:  position{line: 169, col: 1, offset: 6722},
			expr: &actionExpr{
				pos: position{line: 169, col: 18, offset: 6739},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 169, col: 18, offset: 6739},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 169, col: 18, offset: 6739},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 19, offset: 6740},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 28, offset: 6749},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 29, offset: 6750},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 38, offset: 6759},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 39, offset: 6760},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 48, offset: 6769},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 49, offset: 6770},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 58, offset: 6779},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 59, offset: 6780},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 169, col: 68, offset: 6789},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 77, offset: 6798},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 177, col: 1, offset: 6971},
			expr: &choiceExpr{
				pos: position{line: 177, col: 17, offset: 6987},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 177, col: 17, offset: 6987},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 33, offset: 7003},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 49, offset: 7019},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 65, offset: 7035},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 81, offset: 7051},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 179, col: 1, offset: 7066},
			expr: &actionExpr{
				pos: position{line: 179, col: 18, offset: 7083},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 179, col: 18, offset: 7083},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 179, col: 18, offset: 7083},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 179, col: 29, offset: 7094},
								expr: &ruleRefExpr{
									pos:  position{line: 179, col: 30, offset: 7095},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 49, offset: 7114},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 179, col: 56, offset: 7121},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 179, col: 62, offset: 7127},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 62, offset: 7127},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 66, offset: 7131},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 75, offset: 7140},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 179, col: 90, offset: 7155},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 90, offset: 7155},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 94, offset: 7159},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 179, col: 97, offset: 7162},
								expr: &ruleRefExpr{
									pos:  position{line: 179, col: 98, offset: 7163},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 179, col: 116, offset: 7181},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 116, offset: 7181},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 179, col: 120, offset: 7185},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 179, col: 125, offset: 7190},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 179, col: 125, offset: 7190},
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 125, offset: 7190},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 138, offset: 7203},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Title",
			pos:  position{line: 183, col: 1, offset: 7318},
			expr: &actionExpr{
				pos: position{line: 183, col: 18, offset: 7335},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 183, col: 18, offset: 7335},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 18, offset: 7335},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 183, col: 29, offset: 7346},
								expr: &ruleRefExpr{
									pos:  position{line: 183, col: 30, offset: 7347},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 49, offset: 7366},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 183, col: 56, offset: 7373},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 183, col: 63, offset: 7380},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 63, offset: 7380},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 67, offset: 7384},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 76, offset: 7393},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 183, col: 91, offset: 7408},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 91, offset: 7408},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 95, offset: 7412},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 183, col: 98, offset: 7415},
								expr: &ruleRefExpr{
									pos:  position{line: 183, col: 99, offset: 7416},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 183, col: 117, offset: 7434},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 117, offset: 7434},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 121, offset: 7438},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 183, col: 126, offset: 7443},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 183, col: 126, offset: 7443},
									expr: &ruleRefExpr{
										pos:  position{line: 183, col: 126, offset: 7443},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 183, col: 139, offset: 7456},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Title",
			pos:  position{line: 187, col: 1, offset: 7570},
			expr: &actionExpr{
				pos: position{line: 187, col: 18, offset: 7587},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 187, col: 18, offset: 7587},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 187, col: 18, offset: 7587},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 187, col: 29, offset: 7598},
								expr: &ruleRefExpr{
									pos:  position{line: 187, col: 30, offset: 7599},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 49, offset: 7618},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 187, col: 56, offset: 7625},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 187, col: 64, offset: 7633},
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 64, offset: 7633},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 68, offset: 7637},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 77, offset: 7646},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 187, col: 92, offset: 7661},
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 92, offset: 7661},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 96, offset: 7665},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 187, col: 99, offset: 7668},
								expr: &ruleRefExpr{
									pos:  position{line: 187, col: 100, offset: 7669},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 187, col: 118, offset: 7687},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 187, col: 123, offset: 7692},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 187, col: 123, offset: 7692},
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 123, offset: 7692},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 136, offset: 7705},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Title",
			pos:  position{line: 191, col: 1, offset: 7819},
			expr: &actionExpr{
				pos: position{line: 191, col: 18, offset: 7836},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 191, col: 18, offset: 7836},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 191, col: 18, offset: 7836},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 191, col: 29, offset: 7847},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 30, offset: 7848},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 49, offset: 7867},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 191, col: 56, offset: 7874},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 191, col: 65, offset: 7883},
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 65, offset: 7883},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 69, offset: 7887},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 78, offset: 7896},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 191, col: 93, offset: 7911},
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 93, offset: 7911},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 97, offset: 7915},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 191, col: 100, offset: 7918},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 101, offset: 7919},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 191, col: 119, offset: 7937},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 191, col: 124, offset: 7942},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 191, col: 124, offset: 7942},
									expr: &ruleRefExpr{
										pos:  position{line: 191, col: 124, offset: 7942},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 137, offset: 7955},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Title",
			pos:  position{line: 195, col: 1, offset: 8069},
			expr: &actionExpr{
				pos: position{line: 195, col: 18, offset: 8086},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 195, col: 18, offset: 8086},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 195, col: 18, offset: 8086},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 195, col: 29, offset: 8097},
								expr: &ruleRefExpr{
									pos:  position{line: 195, col: 30, offset: 8098},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 49, offset: 8117},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 195, col: 56, offset: 8124},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 195, col: 66, offset: 8134},
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 66, offset: 8134},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 70, offset: 8138},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 79, offset: 8147},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 195, col: 94, offset: 8162},
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 94, offset: 8162},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 98, offset: 8166},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 195, col: 101, offset: 8169},
								expr: &ruleRefExpr{
									pos:  position{line: 195, col: 102, offset: 8170},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 195, col: 120, offset: 8188},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 195, col: 125, offset: 8193},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 195, col: 125, offset: 8193},
									expr: &ruleRefExpr{
										pos:  position{line: 195, col: 125, offset: 8193},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 195, col: 138, offset: 8206},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "List",
			pos:  position{line: 202, col: 1, offset: 8421},
			expr: &actionExpr{
				pos: position{line: 202, col: 9, offset: 8429},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 202, col: 9, offset: 8429},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 202, col: 9, offset: 8429},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 202, col: 20, offset: 8440},
								expr: &ruleRefExpr{
									pos:  position{line: 202, col: 21, offset: 8441},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 204, col: 5, offset: 8530},
							label: "elements",
							expr: &ruleRefExpr{
								pos:  position{line: 204, col: 14, offset: 8539},
								name: "ListItems",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItems",
			pos:  position{line: 208, col: 1, offset: 8633},
			expr: &oneOrMoreExpr{
				pos: position{line: 208, col: 14, offset: 8646},
				expr: &choiceExpr{
					pos: position{line: 208, col: 15, offset: 8647},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 208, col: 15, offset: 8647},
							name: "OrderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 208, col: 33, offset: 8665},
							name: "UnorderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 208, col: 53, offset: 8685},
							name: "LabeledListItem",
						},
					},
				},
			},
		},
		{
			name: "ListAttribute",
			pos:  position{line: 210, col: 1, offset: 8704},
			expr: &actionExpr{
				pos: position{line: 210, col: 18, offset: 8721},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 210, col: 18, offset: 8721},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 210, col: 18, offset: 8721},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 210, col: 29, offset: 8732},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 210, col: 29, offset: 8732},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 210, col: 48, offset: 8751},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 210, col: 56, offset: 8759},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 214, col: 1, offset: 8798},
			expr: &actionExpr{
				pos: position{line: 214, col: 11, offset: 8808},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 214, col: 11, offset: 8808},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 214, col: 11, offset: 8808},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 214, col: 16, offset: 8813},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 20, offset: 8817},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 214, col: 24, offset: 8821},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 218, col: 1, offset: 8887},
			expr: &actionExpr{
				pos: position{line: 218, col: 21, offset: 8907},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 218, col: 21, offset: 8907},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 222, col: 1, offset: 8990},
			expr: &actionExpr{
				pos: position{line: 222, col: 19, offset: 9008},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 222, col: 19, offset: 9008},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 222, col: 25, offset: 9014},
						expr: &seqExpr{
							pos: position{line: 223, col: 5, offset: 9020},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 223, col: 5, offset: 9020},
									expr: &ruleRefExpr{
										pos:  position{line: 223, col: 7, offset: 9022},
										name: "OrderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 224, col: 5, offset: 9050},
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 7, offset: 9052},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 225, col: 5, offset: 9082},
									expr: &seqExpr{
										pos: position{line: 225, col: 7, offset: 9084},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 225, col: 7, offset: 9084},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 225, col: 27, offset: 9104},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&notExpr{
									pos: position{line: 226, col: 5, offset: 9135},
									expr: &ruleRefExpr{
										pos:  position{line: 226, col: 7, offset: 9137},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 227, col: 5, offset: 9164},
									expr: &ruleRefExpr{
										pos:  position{line: 227, col: 7, offset: 9166},
										name: "ElementAttribute",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 228, col: 5, offset: 9188},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 228, col: 37, offset: 9220},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContinuation",
			pos:  position{line: 232, col: 1, offset: 9289},
			expr: &actionExpr{
				pos: position{line: 232, col: 25, offset: 9313},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 232, col: 25, offset: 9313},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 232, col: 25, offset: 9313},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 232, col: 29, offset: 9317},
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 29, offset: 9317},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 232, col: 33, offset: 9321},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 236, col: 1, offset: 9373},
			expr: &actionExpr{
				pos: position{line: 236, col: 26, offset: 9398},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 236, col: 26, offset: 9398},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 236, col: 26, offset: 9398},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 236, col: 47, offset: 9419},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 55, offset: 9427},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItem",
			pos:  position{line: 243, col: 1, offset: 9583},
			expr: &actionExpr{
				pos: position{line: 243, col: 20, offset: 9602},
				run: (*parser).callonOrderedListItem1,
				expr: &seqExpr{
					pos: position{line: 243, col: 20, offset: 9602},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 243, col: 20, offset: 9602},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 243, col: 31, offset: 9613},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 32, offset: 9614},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 243, col: 51, offset: 9633},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 59, offset: 9641},
								name: "OrderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 243, col: 82, offset: 9664},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 91, offset: 9673},
								name: "OrderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 243, col: 115, offset: 9697},
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 115, offset: 9697},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItemPrefix",
			pos:  position{line: 247, col: 1, offset: 9845},
			expr: &choiceExpr{
				pos: position{line: 249, col: 1, offset: 9909},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 249, col: 1, offset: 9909},
						run: (*parser).callonOrderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 249, col: 1, offset: 9909},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 249, col: 1, offset: 9909},
									expr: &ruleRefExpr{
										pos:  position{line: 249, col: 1, offset: 9909},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 249, col: 5, offset: 9913},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 249, col: 12, offset: 9920},
										val:        ".",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 249, col: 17, offset: 9925},
									expr: &ruleRefExpr{
										pos:  position{line: 249, col: 17, offset: 9925},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 251, col: 5, offset: 10018},
						run: (*parser).callonOrderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 251, col: 5, offset: 10018},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 251, col: 5, offset: 10018},
									expr: &ruleRefExpr{
										pos:  position{line: 251, col: 5, offset: 10018},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 251, col: 9, offset: 10022},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 251, col: 16, offset: 10029},
										val:        "..",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 251, col: 22, offset: 10035},
									expr: &ruleRefExpr{
										pos:  position{line: 251, col: 22, offset: 10035},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 253, col: 5, offset: 10133},
						run: (*parser).callonOrderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 253, col: 5, offset: 10133},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 253, col: 5, offset: 10133},
									expr: &ruleRefExpr{
										pos:  position{line: 253, col: 5, offset: 10133},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 253, col: 9, offset: 10137},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 253, col: 16, offset: 10144},
										val:        "...",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 253, col: 23, offset: 10151},
									expr: &ruleRefExpr{
										pos:  position{line: 253, col: 23, offset: 10151},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 255, col: 5, offset: 10250},
						run: (*parser).callonOrderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 255, col: 5, offset: 10250},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 255, col: 5, offset: 10250},
									expr: &ruleRefExpr{
										pos:  position{line: 255, col: 5, offset: 10250},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 255, col: 9, offset: 10254},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 255, col: 16, offset: 10261},
										val:        "....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 255, col: 24, offset: 10269},
									expr: &ruleRefExpr{
										pos:  position{line: 255, col: 24, offset: 10269},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 257, col: 5, offset: 10369},
						run: (*parser).callonOrderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 257, col: 5, offset: 10369},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 257, col: 5, offset: 10369},
									expr: &ruleRefExpr{
										pos:  position{line: 257, col: 5, offset: 10369},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 257, col: 9, offset: 10373},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 257, col: 16, offset: 10380},
										val:        ".....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 257, col: 25, offset: 10389},
									expr: &ruleRefExpr{
										pos:  position{line: 257, col: 25, offset: 10389},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 260, col: 5, offset: 10512},
						run: (*parser).callonOrderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 260, col: 5, offset: 10512},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 260, col: 5, offset: 10512},
									expr: &ruleRefExpr{
										pos:  position{line: 260, col: 5, offset: 10512},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 260, col: 9, offset: 10516},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 260, col: 16, offset: 10523},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 260, col: 16, offset: 10523},
												expr: &seqExpr{
													pos: position{line: 260, col: 17, offset: 10524},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 260, col: 17, offset: 10524},
															expr: &litMatcher{
																pos:        position{line: 260, col: 18, offset: 10525},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 260, col: 22, offset: 10529},
															expr: &ruleRefExpr{
																pos:  position{line: 260, col: 23, offset: 10530},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 260, col: 26, offset: 10533},
															expr: &ruleRefExpr{
																pos:  position{line: 260, col: 27, offset: 10534},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 260, col: 35, offset: 10542},
															val:        "[0-9]",
															ranges:     []rune{'0', '9'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 260, col: 43, offset: 10550},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 260, col: 48, offset: 10555},
									expr: &ruleRefExpr{
										pos:  position{line: 260, col: 48, offset: 10555},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 262, col: 5, offset: 10650},
						run: (*parser).callonOrderedListItemPrefix60,
						expr: &seqExpr{
							pos: position{line: 262, col: 5, offset: 10650},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 262, col: 5, offset: 10650},
									expr: &ruleRefExpr{
										pos:  position{line: 262, col: 5, offset: 10650},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 262, col: 9, offset: 10654},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 262, col: 16, offset: 10661},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 262, col: 16, offset: 10661},
												expr: &seqExpr{
													pos: position{line: 262, col: 17, offset: 10662},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 262, col: 17, offset: 10662},
															expr: &litMatcher{
																pos:        position{line: 262, col: 18, offset: 10663},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 262, col: 22, offset: 10667},
															expr: &ruleRefExpr{
																pos:  position{line: 262, col: 23, offset: 10668},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 262, col: 26, offset: 10671},
															expr: &ruleRefExpr{
																pos:  position{line: 262, col: 27, offset: 10672},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 262, col: 35, offset: 10680},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 262, col: 43, offset: 10688},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 262, col: 48, offset: 10693},
									expr: &ruleRefExpr{
										pos:  position{line: 262, col: 48, offset: 10693},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 264, col: 5, offset: 10791},
						run: (*parser).callonOrderedListItemPrefix78,
						expr: &seqExpr{
							pos: position{line: 264, col: 5, offset: 10791},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 264, col: 5, offset: 10791},
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 5, offset: 10791},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 264, col: 9, offset: 10795},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 264, col: 16, offset: 10802},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 264, col: 16, offset: 10802},
												expr: &seqExpr{
													pos: position{line: 264, col: 17, offset: 10803},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 264, col: 17, offset: 10803},
															expr: &litMatcher{
																pos:        position{line: 264, col: 18, offset: 10804},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 264, col: 22, offset: 10808},
															expr: &ruleRefExpr{
																pos:  position{line: 264, col: 23, offset: 10809},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 264, col: 26, offset: 10812},
															expr: &ruleRefExpr{
																pos:  position{line: 264, col: 27, offset: 10813},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 264, col: 35, offset: 10821},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 264, col: 43, offset: 10829},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 264, col: 48, offset: 10834},
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 48, offset: 10834},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 266, col: 5, offset: 10932},
						run: (*parser).callonOrderedListItemPrefix96,
						expr: &seqExpr{
							pos: position{line: 266, col: 5, offset: 10932},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 266, col: 5, offset: 10932},
									expr: &ruleRefExpr{
										pos:  position{line: 266, col: 5, offset: 10932},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 266, col: 9, offset: 10936},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 266, col: 16, offset: 10943},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 266, col: 16, offset: 10943},
												expr: &seqExpr{
													pos: position{line: 266, col: 17, offset: 10944},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 266, col: 17, offset: 10944},
															expr: &litMatcher{
																pos:        position{line: 266, col: 18, offset: 10945},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 266, col: 22, offset: 10949},
															expr: &ruleRefExpr{
																pos:  position{line: 266, col: 23, offset: 10950},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 266, col: 26, offset: 10953},
															expr: &ruleRefExpr{
																pos:  position{line: 266, col: 27, offset: 10954},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 266, col: 35, offset: 10962},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 266, col: 43, offset: 10970},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 266, col: 48, offset: 10975},
									expr: &ruleRefExpr{
										pos:  position{line: 266, col: 48, offset: 10975},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 268, col: 5, offset: 11073},
						run: (*parser).callonOrderedListItemPrefix114,
						expr: &seqExpr{
							pos: position{line: 268, col: 5, offset: 11073},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 268, col: 5, offset: 11073},
									expr: &ruleRefExpr{
										pos:  position{line: 268, col: 5, offset: 11073},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 268, col: 9, offset: 11077},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 268, col: 16, offset: 11084},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 268, col: 16, offset: 11084},
												expr: &seqExpr{
													pos: position{line: 268, col: 17, offset: 11085},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 268, col: 17, offset: 11085},
															expr: &litMatcher{
																pos:        position{line: 268, col: 18, offset: 11086},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 268, col: 22, offset: 11090},
															expr: &ruleRefExpr{
																pos:  position{line: 268, col: 23, offset: 11091},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 268, col: 26, offset: 11094},
															expr: &ruleRefExpr{
																pos:  position{line: 268, col: 27, offset: 11095},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 268, col: 35, offset: 11103},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 268, col: 43, offset: 11111},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 268, col: 48, offset: 11116},
									expr: &ruleRefExpr{
										pos:  position{line: 268, col: 48, offset: 11116},
										name: "WS",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItemContent",
			pos:  position{line: 291, col: 1, offset: 11900},
			expr: &actionExpr{
				pos: position{line: 291, col: 27, offset: 11926},
				run: (*parser).callonOrderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 291, col: 27, offset: 11926},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 291, col: 37, offset: 11936},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 291, col: 37, offset: 11936},
								expr: &ruleRefExpr{
									pos:  position{line: 291, col: 37, offset: 11936},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 291, col: 52, offset: 11951},
								expr: &ruleRefExpr{
									pos:  position{line: 291, col: 52, offset: 11951},
									name: "ContinuedBlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 298, col: 1, offset: 12277},
			expr: &actionExpr{
				pos: position{line: 298, col: 22, offset: 12298},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 298, col: 22, offset: 12298},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 298, col: 22, offset: 12298},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 30, offset: 12306},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 298, col: 55, offset: 12331},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 64, offset: 12340},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 298, col: 90, offset: 12366},
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 90, offset: 12366},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 302, col: 1, offset: 12490},
			expr: &choiceExpr{
				pos: position{line: 302, col: 28, offset: 12517},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 302, col: 28, offset: 12517},
						run: (*parser).callonUnorderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 302, col: 28, offset: 12517},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 302, col: 28, offset: 12517},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 28, offset: 12517},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 302, col: 32, offset: 12521},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 302, col: 39, offset: 12528},
										val:        "*****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 302, col: 48, offset: 12537},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 48, offset: 12537},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 304, col: 5, offset: 12682},
						run: (*parser).callonUnorderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 304, col: 5, offset: 12682},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 304, col: 5, offset: 12682},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 5, offset: 12682},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 304, col: 9, offset: 12686},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 304, col: 16, offset: 12693},
										val:        "****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 304, col: 24, offset: 12701},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 24, offset: 12701},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 306, col: 5, offset: 12846},
						run: (*parser).callonUnorderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 306, col: 5, offset: 12846},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 306, col: 5, offset: 12846},
									expr: &ruleRefExpr{
										pos:  position{line: 306, col: 5, offset: 12846},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 306, col: 9, offset: 12850},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 306, col: 16, offset: 12857},
										val:        "***",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 306, col: 23, offset: 12864},
									expr: &ruleRefExpr{
										pos:  position{line: 306, col: 23, offset: 12864},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 308, col: 5, offset: 13010},
						run: (*parser).callonUnorderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 308, col: 5, offset: 13010},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 308, col: 5, offset: 13010},
									expr: &ruleRefExpr{
										pos:  position{line: 308, col: 5, offset: 13010},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 308, col: 9, offset: 13014},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 308, col: 16, offset: 13021},
										val:        "**",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 308, col: 22, offset: 13027},
									expr: &ruleRefExpr{
										pos:  position{line: 308, col: 22, offset: 13027},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 310, col: 5, offset: 13171},
						run: (*parser).callonUnorderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 310, col: 5, offset: 13171},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 310, col: 5, offset: 13171},
									expr: &ruleRefExpr{
										pos:  position{line: 310, col: 5, offset: 13171},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 310, col: 9, offset: 13175},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 310, col: 16, offset: 13182},
										val:        "*",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 310, col: 21, offset: 13187},
									expr: &ruleRefExpr{
										pos:  position{line: 310, col: 21, offset: 13187},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 312, col: 5, offset: 13330},
						run: (*parser).callonUnorderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 312, col: 5, offset: 13330},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 312, col: 5, offset: 13330},
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 5, offset: 13330},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 312, col: 9, offset: 13334},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 312, col: 16, offset: 13341},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 312, col: 21, offset: 13346},
									expr: &ruleRefExpr{
										pos:  position{line: 312, col: 21, offset: 13346},
										name: "WS",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 316, col: 1, offset: 13482},
			expr: &actionExpr{
				pos: position{line: 316, col: 29, offset: 13510},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 316, col: 29, offset: 13510},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 316, col: 39, offset: 13520},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 316, col: 39, offset: 13520},
								expr: &ruleRefExpr{
									pos:  position{line: 316, col: 39, offset: 13520},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 316, col: 54, offset: 13535},
								expr: &ruleRefExpr{
									pos:  position{line: 316, col: 54, offset: 13535},
									name: "ContinuedBlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItem",
			pos:  position{line: 323, col: 1, offset: 13859},
			expr: &choiceExpr{
				pos: position{line: 323, col: 20, offset: 13878},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 323, col: 20, offset: 13878},
						run: (*parser).callonLabeledListItem2,
						expr: &seqExpr{
							pos: position{line: 323, col: 20, offset: 13878},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 323, col: 20, offset: 13878},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 323, col: 26, offset: 13884},
										name: "LabeledListItemTerm",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 323, col: 47, offset: 13905},
									name: "LabeledListItemSeparator",
								},
								&labeledExpr{
									pos:   position{line: 323, col: 72, offset: 13930},
									label: "description",
									expr: &ruleRefExpr{
										pos:  position{line: 323, col: 85, offset: 13943},
										name: "LabeledListItemDescription",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 325, col: 6, offset: 14070},
						run: (*parser).callonLabeledListItem9,
						expr: &seqExpr{
							pos: position{line: 325, col: 6, offset: 14070},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 325, col: 6, offset: 14070},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 325, col: 12, offset: 14076},
										name: "LabeledListItemTerm",
									},
								},
								&litMatcher{
									pos:        position{line: 325, col: 33, offset: 14097},
									val:        "::",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 325, col: 38, offset: 14102},
									expr: &ruleRefExpr{
										pos:  position{line: 325, col: 38, offset: 14102},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 325, col: 42, offset: 14106},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemTerm",
			pos:  position{line: 329, col: 1, offset: 14243},
			expr: &actionExpr{
				pos: position{line: 329, col: 24, offset: 14266},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 329, col: 24, offset: 14266},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 329, col: 29, offset: 14271},
						expr: &seqExpr{
							pos: position{line: 329, col: 30, offset: 14272},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 329, col: 30, offset: 14272},
									expr: &ruleRefExpr{
										pos:  position{line: 329, col: 31, offset: 14273},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 329, col: 39, offset: 14281},
									expr: &litMatcher{
										pos:        position{line: 329, col: 40, offset: 14282},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 329, col: 45, offset: 14287,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 334, col: 1, offset: 14378},
			expr: &seqExpr{
				pos: position{line: 334, col: 30, offset: 14407},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 334, col: 30, offset: 14407},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 334, col: 35, offset: 14412},
						expr: &choiceExpr{
							pos: position{line: 334, col: 36, offset: 14413},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 334, col: 36, offset: 14413},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 334, col: 41, offset: 14418},
									name: "NEWLINE",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemDescription",
			pos:  position{line: 336, col: 1, offset: 14429},
			expr: &actionExpr{
				pos: position{line: 336, col: 31, offset: 14459},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 336, col: 31, offset: 14459},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 336, col: 40, offset: 14468},
						expr: &choiceExpr{
							pos: position{line: 336, col: 41, offset: 14469},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 336, col: 41, offset: 14469},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 336, col: 57, offset: 14485},
									name: "ContinuedBlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 345, col: 1, offset: 14942},
			expr: &actionExpr{
				pos: position{line: 345, col: 14, offset: 14955},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 345, col: 14, offset: 14955},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 345, col: 14, offset: 14955},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 345, col: 25, offset: 14966},
								expr: &ruleRefExpr{
									pos:  position{line: 345, col: 26, offset: 14967},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 345, col: 45, offset: 14986},
							expr: &seqExpr{
								pos: position{line: 345, col: 47, offset: 14988},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 345, col: 47, offset: 14988},
										expr: &litMatcher{
											pos:        position{line: 345, col: 47, offset: 14988},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 345, col: 52, offset: 14993},
										expr: &ruleRefExpr{
											pos:  position{line: 345, col: 52, offset: 14993},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 345, col: 57, offset: 14998},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 345, col: 63, offset: 15004},
								expr: &seqExpr{
									pos: position{line: 345, col: 64, offset: 15005},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 345, col: 64, offset: 15005},
											name: "InlineContentWithTrailingSpaces",
										},
										&ruleRefExpr{
											pos:  position{line: 345, col: 96, offset: 15037},
											name: "EOL",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineContentWithTrailingSpaces",
			pos:  position{line: 351, col: 1, offset: 15327},
			expr: &actionExpr{
				pos: position{line: 351, col: 36, offset: 15362},
				run: (*parser).callonInlineContentWithTrailingSpaces1,
				expr: &seqExpr{
					pos: position{line: 351, col: 36, offset: 15362},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 351, col: 36, offset: 15362},
							expr: &ruleRefExpr{
								pos:  position{line: 351, col: 37, offset: 15363},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 351, col: 52, offset: 15378},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 351, col: 61, offset: 15387},
								expr: &seqExpr{
									pos: position{line: 351, col: 62, offset: 15388},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 351, col: 62, offset: 15388},
											expr: &ruleRefExpr{
												pos:  position{line: 351, col: 62, offset: 15388},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 351, col: 66, offset: 15392},
											expr: &ruleRefExpr{
												pos:  position{line: 351, col: 67, offset: 15393},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 351, col: 83, offset: 15409},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 351, col: 97, offset: 15423},
											expr: &ruleRefExpr{
												pos:  position{line: 351, col: 97, offset: 15423},
												name: "WS",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineContent",
			pos:  position{line: 355, col: 1, offset: 15556},
			expr: &actionExpr{
				pos: position{line: 355, col: 18, offset: 15573},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 355, col: 18, offset: 15573},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 355, col: 18, offset: 15573},
							expr: &ruleRefExpr{
								pos:  position{line: 355, col: 19, offset: 15574},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 355, col: 34, offset: 15589},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 355, col: 43, offset: 15598},
								expr: &seqExpr{
									pos: position{line: 355, col: 44, offset: 15599},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 355, col: 44, offset: 15599},
											expr: &ruleRefExpr{
												pos:  position{line: 355, col: 44, offset: 15599},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 355, col: 48, offset: 15603},
											expr: &ruleRefExpr{
												pos:  position{line: 355, col: 49, offset: 15604},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 355, col: 65, offset: 15620},
											name: "InlineElement",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 359, col: 1, offset: 15742},
			expr: &choiceExpr{
				pos: position{line: 359, col: 18, offset: 15759},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 359, col: 18, offset: 15759},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 359, col: 35, offset: 15776},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 359, col: 49, offset: 15790},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 359, col: 63, offset: 15804},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 359, col: 76, offset: 15817},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 359, col: 83, offset: 15824},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 359, col: 115, offset: 15856},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "Admonition",
			pos:  position{line: 365, col: 1, offset: 15976},
			expr: &ruleRefExpr{
				pos:  position{line: 365, col: 15, offset: 15990},
				name: "AdmonitionParagraph",
			},
		},
		{
			name: "AdmonitionParagraph",
			pos:  position{line: 369, col: 1, offset: 16161},
			expr: &choiceExpr{
				pos: position{line: 369, col: 24, offset: 16184},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 369, col: 24, offset: 16184},
						run: (*parser).callonAdmonitionParagraph2,
						expr: &seqExpr{
							pos: position{line: 369, col: 24, offset: 16184},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 369, col: 24, offset: 16184},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 369, col: 35, offset: 16195},
										expr: &ruleRefExpr{
											pos:  position{line: 369, col: 36, offset: 16196},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 369, col: 55, offset: 16215},
									expr: &seqExpr{
										pos: position{line: 369, col: 57, offset: 16217},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 369, col: 57, offset: 16217},
												expr: &litMatcher{
													pos:        position{line: 369, col: 57, offset: 16217},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 369, col: 62, offset: 16222},
												expr: &ruleRefExpr{
													pos:  position{line: 369, col: 62, offset: 16222},
													name: "WS",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 369, col: 67, offset: 16227},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 369, col: 70, offset: 16230},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 369, col: 86, offset: 16246},
									val:        ": ",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 369, col: 91, offset: 16251},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 369, col: 100, offset: 16260},
										name: "AdmonitionParagraphContent",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 371, col: 5, offset: 16416},
						run: (*parser).callonAdmonitionParagraph18,
						expr: &seqExpr{
							pos: position{line: 371, col: 5, offset: 16416},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 371, col: 5, offset: 16416},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 371, col: 16, offset: 16427},
										expr: &ruleRefExpr{
											pos:  position{line: 371, col: 17, offset: 16428},
											name: "ElementAttribute",
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 371, col: 36, offset: 16447},
									label: "k",
									expr: &ruleRefExpr{
										pos:  position{line: 371, col: 39, offset: 16450},
										name: "AdmonitionMarker",
									},
								},
								&labeledExpr{
									pos:   position{line: 371, col: 57, offset: 16468},
									label: "otherAttributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 371, col: 73, offset: 16484},
										expr: &ruleRefExpr{
											pos:  position{line: 371, col: 74, offset: 16485},
											name: "ElementAttribute",
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 371, col: 93, offset: 16504},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 371, col: 102, offset: 16513},
										name: "AdmonitionParagraphContent",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionParagraphContent",
			pos:  position{line: 375, col: 1, offset: 16708},
			expr: &actionExpr{
				pos: position{line: 375, col: 31, offset: 16738},
				run: (*parser).callonAdmonitionParagraphContent1,
				expr: &labeledExpr{
					pos:   position{line: 375, col: 31, offset: 16738},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 375, col: 37, offset: 16744},
						expr: &seqExpr{
							pos: position{line: 375, col: 38, offset: 16745},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 375, col: 38, offset: 16745},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 375, col: 70, offset: 16777},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionMarker",
			pos:  position{line: 380, col: 1, offset: 16938},
			expr: &actionExpr{
				pos: position{line: 380, col: 21, offset: 16958},
				run: (*parser).callonAdmonitionMarker1,
				expr: &seqExpr{
					pos: position{line: 380, col: 21, offset: 16958},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 380, col: 21, offset: 16958},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 380, col: 25, offset: 16962},
							label: "k",
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 28, offset: 16965},
								name: "AdmonitionKind",
							},
						},
						&litMatcher{
							pos:        position{line: 380, col: 44, offset: 16981},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 380, col: 48, offset: 16985},
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 48, offset: 16985},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 380, col: 52, offset: 16989},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "AdmonitionKind",
			pos:  position{line: 384, col: 1, offset: 17020},
			expr: &choiceExpr{
				pos: position{line: 384, col: 19, offset: 17038},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 384, col: 19, offset: 17038},
						run: (*parser).callonAdmonitionKind2,
						expr: &litMatcher{
							pos:        position{line: 384, col: 19, offset: 17038},
							val:        "TIP",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 386, col: 5, offset: 17076},
						run: (*parser).callonAdmonitionKind4,
						expr: &litMatcher{
							pos:        position{line: 386, col: 5, offset: 17076},
							val:        "NOTE",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 388, col: 5, offset: 17116},
						run: (*parser).callonAdmonitionKind6,
						expr: &litMatcher{
							pos:        position{line: 388, col: 5, offset: 17116},
							val:        "IMPORTANT",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 390, col: 5, offset: 17166},
						run: (*parser).callonAdmonitionKind8,
						expr: &litMatcher{
							pos:        position{line: 390, col: 5, offset: 17166},
							val:        "WARNING",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 392, col: 5, offset: 17212},
						run: (*parser).callonAdmonitionKind10,
						expr: &litMatcher{
							pos:        position{line: 392, col: 5, offset: 17212},
							val:        "CAUTION",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 399, col: 1, offset: 17496},
			expr: &choiceExpr{
				pos: position{line: 399, col: 15, offset: 17510},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 399, col: 15, offset: 17510},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 399, col: 26, offset: 17521},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 399, col: 39, offset: 17534},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 400, col: 13, offset: 17562},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 400, col: 31, offset: 17580},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 400, col: 51, offset: 17600},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 402, col: 1, offset: 17622},
			expr: &choiceExpr{
				pos: position{line: 402, col: 13, offset: 17634},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 402, col: 13, offset: 17634},
						run: (*parser).callonBoldText2,
						expr: &seqExpr{
							pos: position{line: 402, col: 13, offset: 17634},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 402, col: 13, offset: 17634},
									expr: &litMatcher{
										pos:        position{line: 402, col: 14, offset: 17635},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 402, col: 19, offset: 17640},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 402, col: 24, offset: 17645},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 402, col: 33, offset: 17654},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 402, col: 52, offset: 17673},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 404, col: 5, offset: 17798},
						run: (*parser).callonBoldText10,
						expr: &seqExpr{
							pos: position{line: 404, col: 5, offset: 17798},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 404, col: 5, offset: 17798},
									expr: &litMatcher{
										pos:        position{line: 404, col: 6, offset: 17799},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 404, col: 11, offset: 17804},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 404, col: 16, offset: 17809},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 404, col: 25, offset: 17818},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 404, col: 44, offset: 17837},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 407, col: 5, offset: 18002},
						run: (*parser).callonBoldText18,
						expr: &seqExpr{
							pos: position{line: 407, col: 5, offset: 18002},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 407, col: 5, offset: 18002},
									expr: &litMatcher{
										pos:        position{line: 407, col: 6, offset: 18003},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 407, col: 10, offset: 18007},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 407, col: 14, offset: 18011},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 407, col: 23, offset: 18020},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 407, col: 42, offset: 18039},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 411, col: 1, offset: 18139},
			expr: &choiceExpr{
				pos: position{line: 411, col: 20, offset: 18158},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 411, col: 20, offset: 18158},
						run: (*parser).callonEscapedBoldText2,
						expr: &seqExpr{
							pos: position{line: 411, col: 20, offset: 18158},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 411, col: 20, offset: 18158},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 411, col: 33, offset: 18171},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 411, col: 33, offset: 18171},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 411, col: 38, offset: 18176},
												expr: &litMatcher{
													pos:        position{line: 411, col: 38, offset: 18176},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 411, col: 44, offset: 18182},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 411, col: 49, offset: 18187},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 411, col: 58, offset: 18196},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 411, col: 77, offset: 18215},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 413, col: 5, offset: 18370},
						run: (*parser).callonEscapedBoldText13,
						expr: &seqExpr{
							pos: position{line: 413, col: 5, offset: 18370},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 413, col: 5, offset: 18370},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 413, col: 18, offset: 18383},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 413, col: 18, offset: 18383},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 413, col: 22, offset: 18387},
												expr: &litMatcher{
													pos:        position{line: 413, col: 22, offset: 18387},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 413, col: 28, offset: 18393},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 413, col: 33, offset: 18398},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 413, col: 42, offset: 18407},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 413, col: 61, offset: 18426},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 416, col: 5, offset: 18620},
						run: (*parser).callonEscapedBoldText24,
						expr: &seqExpr{
							pos: position{line: 416, col: 5, offset: 18620},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 416, col: 5, offset: 18620},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 416, col: 18, offset: 18633},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 416, col: 18, offset: 18633},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 416, col: 22, offset: 18637},
												expr: &litMatcher{
													pos:        position{line: 416, col: 22, offset: 18637},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 416, col: 28, offset: 18643},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 416, col: 32, offset: 18647},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 416, col: 41, offset: 18656},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 416, col: 60, offset: 18675},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 420, col: 1, offset: 18827},
			expr: &choiceExpr{
				pos: position{line: 420, col: 15, offset: 18841},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 420, col: 15, offset: 18841},
						run: (*parser).callonItalicText2,
						expr: &seqExpr{
							pos: position{line: 420, col: 15, offset: 18841},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 420, col: 15, offset: 18841},
									expr: &litMatcher{
										pos:        position{line: 420, col: 16, offset: 18842},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 420, col: 21, offset: 18847},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 420, col: 26, offset: 18852},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 420, col: 35, offset: 18861},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 420, col: 54, offset: 18880},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 422, col: 5, offset: 18961},
						run: (*parser).callonItalicText10,
						expr: &seqExpr{
							pos: position{line: 422, col: 5, offset: 18961},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 422, col: 5, offset: 18961},
									expr: &litMatcher{
										pos:        position{line: 422, col: 6, offset: 18962},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 422, col: 11, offset: 18967},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 422, col: 16, offset: 18972},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 422, col: 25, offset: 18981},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 422, col: 44, offset: 19000},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 425, col: 5, offset: 19167},
						run: (*parser).callonItalicText18,
						expr: &seqExpr{
							pos: position{line: 425, col: 5, offset: 19167},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 425, col: 5, offset: 19167},
									expr: &litMatcher{
										pos:        position{line: 425, col: 6, offset: 19168},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 425, col: 10, offset: 19172},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 425, col: 14, offset: 19176},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 425, col: 23, offset: 19185},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 425, col: 42, offset: 19204},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 429, col: 1, offset: 19283},
			expr: &choiceExpr{
				pos: position{line: 429, col: 22, offset: 19304},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 429, col: 22, offset: 19304},
						run: (*parser).callonEscapedItalicText2,
						expr: &seqExpr{
							pos: position{line: 429, col: 22, offset: 19304},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 429, col: 22, offset: 19304},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 429, col: 35, offset: 19317},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 429, col: 35, offset: 19317},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 429, col: 40, offset: 19322},
												expr: &litMatcher{
													pos:        position{line: 429, col: 40, offset: 19322},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 429, col: 46, offset: 19328},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 429, col: 51, offset: 19333},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 429, col: 60, offset: 19342},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 429, col: 79, offset: 19361},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 431, col: 5, offset: 19516},
						run: (*parser).callonEscapedItalicText13,
						expr: &seqExpr{
							pos: position{line: 431, col: 5, offset: 19516},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 431, col: 5, offset: 19516},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 431, col: 18, offset: 19529},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 431, col: 18, offset: 19529},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 431, col: 22, offset: 19533},
												expr: &litMatcher{
													pos:        position{line: 431, col: 22, offset: 19533},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 431, col: 28, offset: 19539},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 431, col: 33, offset: 19544},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 431, col: 42, offset: 19553},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 431, col: 61, offset: 19572},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 434, col: 5, offset: 19766},
						run: (*parser).callonEscapedItalicText24,
						expr: &seqExpr{
							pos: position{line: 434, col: 5, offset: 19766},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 434, col: 5, offset: 19766},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 434, col: 18, offset: 19779},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 434, col: 18, offset: 19779},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 434, col: 22, offset: 19783},
												expr: &litMatcher{
													pos:        position{line: 434, col: 22, offset: 19783},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 434, col: 28, offset: 19789},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 434, col: 32, offset: 19793},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 434, col: 41, offset: 19802},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 434, col: 60, offset: 19821},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 438, col: 1, offset: 19973},
			expr: &choiceExpr{
				pos: position{line: 438, col: 18, offset: 19990},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 438, col: 18, offset: 19990},
						run: (*parser).callonMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 438, col: 18, offset: 19990},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 438, col: 18, offset: 19990},
									expr: &litMatcher{
										pos:        position{line: 438, col: 19, offset: 19991},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 438, col: 24, offset: 19996},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 438, col: 29, offset: 20001},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 438, col: 38, offset: 20010},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 438, col: 57, offset: 20029},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 440, col: 5, offset: 20159},
						run: (*parser).callonMonospaceText10,
						expr: &seqExpr{
							pos: position{line: 440, col: 5, offset: 20159},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 440, col: 5, offset: 20159},
									expr: &litMatcher{
										pos:        position{line: 440, col: 6, offset: 20160},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 440, col: 11, offset: 20165},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 440, col: 16, offset: 20170},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 440, col: 25, offset: 20179},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 440, col: 44, offset: 20198},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 443, col: 5, offset: 20368},
						run: (*parser).callonMonospaceText18,
						expr: &seqExpr{
							pos: position{line: 443, col: 5, offset: 20368},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 443, col: 5, offset: 20368},
									expr: &litMatcher{
										pos:        position{line: 443, col: 6, offset: 20369},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 443, col: 10, offset: 20373},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 443, col: 14, offset: 20377},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 443, col: 23, offset: 20386},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 443, col: 42, offset: 20405},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 447, col: 1, offset: 20532},
			expr: &choiceExpr{
				pos: position{line: 447, col: 25, offset: 20556},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 447, col: 25, offset: 20556},
						run: (*parser).callonEscapedMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 447, col: 25, offset: 20556},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 447, col: 25, offset: 20556},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 447, col: 38, offset: 20569},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 447, col: 38, offset: 20569},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 447, col: 43, offset: 20574},
												expr: &litMatcher{
													pos:        position{line: 447, col: 43, offset: 20574},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 447, col: 49, offset: 20580},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 447, col: 54, offset: 20585},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 447, col: 63, offset: 20594},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 447, col: 82, offset: 20613},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 449, col: 5, offset: 20768},
						run: (*parser).callonEscapedMonospaceText13,
						expr: &seqExpr{
							pos: position{line: 449, col: 5, offset: 20768},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 449, col: 5, offset: 20768},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 449, col: 18, offset: 20781},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 449, col: 18, offset: 20781},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 449, col: 22, offset: 20785},
												expr: &litMatcher{
													pos:        position{line: 449, col: 22, offset: 20785},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 449, col: 28, offset: 20791},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 449, col: 33, offset: 20796},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 449, col: 42, offset: 20805},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 449, col: 61, offset: 20824},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 452, col: 5, offset: 21018},
						run: (*parser).callonEscapedMonospaceText24,
						expr: &seqExpr{
							pos: position{line: 452, col: 5, offset: 21018},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 452, col: 5, offset: 21018},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 452, col: 18, offset: 21031},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 452, col: 18, offset: 21031},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 452, col: 22, offset: 21035},
												expr: &litMatcher{
													pos:        position{line: 452, col: 22, offset: 21035},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 452, col: 28, offset: 21041},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 452, col: 32, offset: 21045},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 452, col: 41, offset: 21054},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 452, col: 60, offset: 21073},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 456, col: 1, offset: 21225},
			expr: &seqExpr{
				pos: position{line: 456, col: 22, offset: 21246},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 456, col: 22, offset: 21246},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 456, col: 47, offset: 21271},
						expr: &seqExpr{
							pos: position{line: 456, col: 48, offset: 21272},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 456, col: 48, offset: 21272},
									expr: &ruleRefExpr{
										pos:  position{line: 456, col: 48, offset: 21272},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 456, col: 52, offset: 21276},
									name: "QuotedTextContentElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContentElement",
			pos:  position{line: 458, col: 1, offset: 21304},
			expr: &choiceExpr{
				pos: position{line: 458, col: 29, offset: 21332},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 458, col: 29, offset: 21332},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 458, col: 42, offset: 21345},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 458, col: 65, offset: 21368},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 460, col: 1, offset: 21503},
			expr: &oneOrMoreExpr{
				pos: position{line: 460, col: 25, offset: 21527},
				expr: &seqExpr{
					pos: position{line: 460, col: 26, offset: 21528},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 460, col: 26, offset: 21528},
							expr: &ruleRefExpr{
								pos:  position{line: 460, col: 27, offset: 21529},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 460, col: 35, offset: 21537},
							expr: &ruleRefExpr{
								pos:  position{line: 460, col: 36, offset: 21538},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 460, col: 39, offset: 21541},
							expr: &litMatcher{
								pos:        position{line: 460, col: 40, offset: 21542},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 460, col: 44, offset: 21546},
							expr: &litMatcher{
								pos:        position{line: 460, col: 45, offset: 21547},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 460, col: 49, offset: 21551},
							expr: &litMatcher{
								pos:        position{line: 460, col: 50, offset: 21552},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 460, col: 54, offset: 21556,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 462, col: 1, offset: 21599},
			expr: &actionExpr{
				pos: position{line: 462, col: 35, offset: 21633},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 462, col: 35, offset: 21633},
					expr: &seqExpr{
						pos: position{line: 462, col: 36, offset: 21634},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 462, col: 36, offset: 21634},
								expr: &ruleRefExpr{
									pos:  position{line: 462, col: 37, offset: 21635},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 462, col: 45, offset: 21643},
								expr: &ruleRefExpr{
									pos:  position{line: 462, col: 46, offset: 21644},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 462, col: 50, offset: 21648,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 467, col: 1, offset: 21893},
			expr: &choiceExpr{
				pos: position{line: 467, col: 31, offset: 21923},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 467, col: 31, offset: 21923},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 467, col: 37, offset: 21929},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 467, col: 43, offset: 21935},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 472, col: 1, offset: 22047},
			expr: &choiceExpr{
				pos: position{line: 472, col: 16, offset: 22062},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 472, col: 16, offset: 22062},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 472, col: 40, offset: 22086},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 472, col: 64, offset: 22110},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 474, col: 1, offset: 22128},
			expr: &actionExpr{
				pos: position{line: 474, col: 26, offset: 22153},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 474, col: 26, offset: 22153},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 474, col: 26, offset: 22153},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 474, col: 30, offset: 22157},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 474, col: 38, offset: 22165},
								expr: &seqExpr{
									pos: position{line: 474, col: 39, offset: 22166},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 474, col: 39, offset: 22166},
											expr: &ruleRefExpr{
												pos:  position{line: 474, col: 40, offset: 22167},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 474, col: 48, offset: 22175},
											expr: &litMatcher{
												pos:        position{line: 474, col: 49, offset: 22176},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 474, col: 53, offset: 22180,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 474, col: 57, offset: 22184},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 478, col: 1, offset: 22279},
			expr: &actionExpr{
				pos: position{line: 478, col: 26, offset: 22304},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 478, col: 26, offset: 22304},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 478, col: 26, offset: 22304},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 478, col: 32, offset: 22310},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 478, col: 40, offset: 22318},
								expr: &seqExpr{
									pos: position{line: 478, col: 41, offset: 22319},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 478, col: 41, offset: 22319},
											expr: &litMatcher{
												pos:        position{line: 478, col: 42, offset: 22320},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 478, col: 48, offset: 22326,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 478, col: 52, offset: 22330},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 482, col: 1, offset: 22427},
			expr: &choiceExpr{
				pos: position{line: 482, col: 21, offset: 22447},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 482, col: 21, offset: 22447},
						run: (*parser).callonPassthroughMacro2,
						expr: &seqExpr{
							pos: position{line: 482, col: 21, offset: 22447},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 482, col: 21, offset: 22447},
									val:        "pass:[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 482, col: 30, offset: 22456},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 482, col: 38, offset: 22464},
										expr: &ruleRefExpr{
											pos:  position{line: 482, col: 39, offset: 22465},
											name: "PassthroughMacroCharacter",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 482, col: 67, offset: 22493},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 484, col: 5, offset: 22584},
						run: (*parser).callonPassthroughMacro9,
						expr: &seqExpr{
							pos: position{line: 484, col: 5, offset: 22584},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 484, col: 5, offset: 22584},
									val:        "pass:q[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 484, col: 15, offset: 22594},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 484, col: 23, offset: 22602},
										expr: &choiceExpr{
											pos: position{line: 484, col: 24, offset: 22603},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 484, col: 24, offset: 22603},
													name: "QuotedText",
												},
												&ruleRefExpr{
													pos:  position{line: 484, col: 37, offset: 22616},
													name: "PassthroughMacroCharacter",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 484, col: 65, offset: 22644},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 488, col: 1, offset: 22734},
			expr: &seqExpr{
				pos: position{line: 488, col: 31, offset: 22764},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 488, col: 31, offset: 22764},
						expr: &litMatcher{
							pos:        position{line: 488, col: 32, offset: 22765},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 488, col: 36, offset: 22769,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 493, col: 1, offset: 22885},
			expr: &actionExpr{
				pos: position{line: 493, col: 19, offset: 22903},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 493, col: 19, offset: 22903},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 493, col: 19, offset: 22903},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 493, col: 24, offset: 22908},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 493, col: 28, offset: 22912},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 493, col: 32, offset: 22916},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 500, col: 1, offset: 23075},
			expr: &choiceExpr{
				pos: position{line: 500, col: 9, offset: 23083},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 500, col: 9, offset: 23083},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 500, col: 24, offset: 23098},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 502, col: 1, offset: 23113},
			expr: &actionExpr{
				pos: position{line: 502, col: 17, offset: 23129},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 502, col: 17, offset: 23129},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 502, col: 17, offset: 23129},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 502, col: 22, offset: 23134},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 502, col: 22, offset: 23134},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 502, col: 33, offset: 23145},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 502, col: 38, offset: 23150},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 502, col: 43, offset: 23155},
								expr: &seqExpr{
									pos: position{line: 502, col: 44, offset: 23156},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 502, col: 44, offset: 23156},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 502, col: 48, offset: 23160},
											expr: &ruleRefExpr{
												pos:  position{line: 502, col: 49, offset: 23161},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 502, col: 60, offset: 23172},
											val:        "]",
											ignoreCase: false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "RelativeLink",
			pos:  position{line: 509, col: 1, offset: 23333},
			expr: &actionExpr{
				pos: position{line: 509, col: 17, offset: 23349},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 509, col: 17, offset: 23349},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 509, col: 17, offset: 23349},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 509, col: 25, offset: 23357},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 509, col: 30, offset: 23362},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 509, col: 30, offset: 23362},
										expr: &ruleRefExpr{
											pos:  position{line: 509, col: 30, offset: 23362},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 509, col: 42, offset: 23374},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 509, col: 47, offset: 23379},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 509, col: 53, offset: 23385},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 509, col: 53, offset: 23385},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 509, col: 57, offset: 23389},
										expr: &ruleRefExpr{
											pos:  position{line: 509, col: 58, offset: 23390},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 509, col: 69, offset: 23401},
										val:        "]",
										ignoreCase: false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BlockImage",
			pos:  position{line: 519, col: 1, offset: 23663},
			expr: &actionExpr{
				pos: position{line: 519, col: 15, offset: 23677},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 519, col: 15, offset: 23677},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 519, col: 15, offset: 23677},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 519, col: 26, offset: 23688},
								expr: &ruleRefExpr{
									pos:  position{line: 519, col: 27, offset: 23689},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 519, col: 46, offset: 23708},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 519, col: 52, offset: 23714},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 519, col: 69, offset: 23731},
							expr: &ruleRefExpr{
								pos:  position{line: 519, col: 69, offset: 23731},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 519, col: 73, offset: 23735},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 524, col: 1, offset: 23894},
			expr: &actionExpr{
				pos: position{line: 524, col: 20, offset: 23913},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 524, col: 20, offset: 23913},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 524, col: 20, offset: 23913},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 524, col: 30, offset: 23923},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 524, col: 36, offset: 23929},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 524, col: 41, offset: 23934},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 524, col: 45, offset: 23938},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 524, col: 57, offset: 23950},
								expr: &ruleRefExpr{
									pos:  position{line: 524, col: 57, offset: 23950},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 524, col: 68, offset: 23961},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 528, col: 1, offset: 24028},
			expr: &actionExpr{
				pos: position{line: 528, col: 16, offset: 24043},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 528, col: 16, offset: 24043},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 528, col: 22, offset: 24049},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 533, col: 1, offset: 24194},
			expr: &actionExpr{
				pos: position{line: 533, col: 21, offset: 24214},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 533, col: 21, offset: 24214},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 533, col: 21, offset: 24214},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 533, col: 30, offset: 24223},
							expr: &litMatcher{
								pos:        position{line: 533, col: 31, offset: 24224},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 533, col: 35, offset: 24228},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 41, offset: 24234},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 533, col: 46, offset: 24239},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 533, col: 50, offset: 24243},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 533, col: 62, offset: 24255},
								expr: &ruleRefExpr{
									pos:  position{line: 533, col: 62, offset: 24255},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 533, col: 73, offset: 24266},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 540, col: 1, offset: 24596},
			expr: &choiceExpr{
				pos: position{line: 540, col: 19, offset: 24614},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 540, col: 19, offset: 24614},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 540, col: 33, offset: 24628},
						name: "ListingBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 540, col: 48, offset: 24643},
						name: "ExampleBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 542, col: 1, offset: 24657},
			expr: &choiceExpr{
				pos: position{line: 542, col: 19, offset: 24675},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 542, col: 19, offset: 24675},
						name: "LiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 542, col: 43, offset: 24699},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 542, col: 66, offset: 24722},
						name: "ListingBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 542, col: 90, offset: 24746},
						name: "ExampleBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 544, col: 1, offset: 24769},
			expr: &litMatcher{
				pos:        position{line: 544, col: 25, offset: 24793},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 546, col: 1, offset: 24800},
			expr: &actionExpr{
				pos: position{line: 546, col: 16, offset: 24815},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 546, col: 16, offset: 24815},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 546, col: 16, offset: 24815},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 546, col: 37, offset: 24836},
							expr: &ruleRefExpr{
								pos:  position{line: 546, col: 37, offset: 24836},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 546, col: 41, offset: 24840},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 546, col: 49, offset: 24848},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 546, col: 57, offset: 24856},
								expr: &seqExpr{
									pos: position{line: 546, col: 58, offset: 24857},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 546, col: 58, offset: 24857},
											expr: &ruleRefExpr{
												pos:  position{line: 546, col: 59, offset: 24858},
												name: "FencedBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 546, col: 80, offset: 24879,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 546, col: 84, offset: 24883},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 546, col: 105, offset: 24904},
							expr: &ruleRefExpr{
								pos:  position{line: 546, col: 105, offset: 24904},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 546, col: 109, offset: 24908},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 550, col: 1, offset: 25001},
			expr: &litMatcher{
				pos:        position{line: 550, col: 26, offset: 25026},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 552, col: 1, offset: 25034},
			expr: &actionExpr{
				pos: position{line: 552, col: 17, offset: 25050},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 552, col: 17, offset: 25050},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 552, col: 17, offset: 25050},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 552, col: 39, offset: 25072},
							expr: &ruleRefExpr{
								pos:  position{line: 552, col: 39, offset: 25072},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 552, col: 43, offset: 25076},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 552, col: 51, offset: 25084},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 552, col: 59, offset: 25092},
								expr: &seqExpr{
									pos: position{line: 552, col: 60, offset: 25093},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 552, col: 60, offset: 25093},
											expr: &ruleRefExpr{
												pos:  position{line: 552, col: 61, offset: 25094},
												name: "ListingBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 552, col: 83, offset: 25116,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 552, col: 87, offset: 25120},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 552, col: 109, offset: 25142},
							expr: &ruleRefExpr{
								pos:  position{line: 552, col: 109, offset: 25142},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 552, col: 113, offset: 25146},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ExampleBlockDelimiter",
			pos:  position{line: 556, col: 1, offset: 25240},
			expr: &litMatcher{
				pos:        position{line: 556, col: 26, offset: 25265},
				val:        "====",
				ignoreCase: false,
			},
		},
		{
			name: "ExampleBlock",
			pos:  position{line: 558, col: 1, offset: 25273},
			expr: &actionExpr{
				pos: position{line: 558, col: 17, offset: 25289},
				run: (*parser).callonExampleBlock1,
				expr: &seqExpr{
					pos: position{line: 558, col: 17, offset: 25289},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 558, col: 17, offset: 25289},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 558, col: 28, offset: 25300},
								expr: &ruleRefExpr{
									pos:  position{line: 558, col: 29, offset: 25301},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 558, col: 48, offset: 25320},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 558, col: 70, offset: 25342},
							expr: &ruleRefExpr{
								pos:  position{line: 558, col: 70, offset: 25342},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 558, col: 74, offset: 25346},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 558, col: 82, offset: 25354},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 558, col: 90, offset: 25362},
								expr: &choiceExpr{
									pos: position{line: 558, col: 91, offset: 25363},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 558, col: 91, offset: 25363},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 558, col: 98, offset: 25370},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 558, col: 110, offset: 25382},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 558, col: 123, offset: 25395},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 558, col: 145, offset: 25417},
							expr: &ruleRefExpr{
								pos:  position{line: 558, col: 145, offset: 25417},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 558, col: 149, offset: 25421},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 565, col: 1, offset: 25805},
			expr: &choiceExpr{
				pos: position{line: 565, col: 17, offset: 25821},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 565, col: 17, offset: 25821},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 565, col: 39, offset: 25843},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 565, col: 76, offset: 25880},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 568, col: 1, offset: 25975},
			expr: &actionExpr{
				pos: position{line: 568, col: 24, offset: 25998},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 568, col: 24, offset: 25998},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 568, col: 24, offset: 25998},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 568, col: 32, offset: 26006},
								expr: &ruleRefExpr{
									pos:  position{line: 568, col: 32, offset: 26006},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 568, col: 37, offset: 26011},
							expr: &ruleRefExpr{
								pos:  position{line: 568, col: 38, offset: 26012},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 568, col: 46, offset: 26020},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 568, col: 55, offset: 26029},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 568, col: 76, offset: 26050},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 573, col: 1, offset: 26231},
			expr: &actionExpr{
				pos: position{line: 573, col: 24, offset: 26254},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 573, col: 24, offset: 26254},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 573, col: 32, offset: 26262},
						expr: &seqExpr{
							pos: position{line: 573, col: 33, offset: 26263},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 573, col: 33, offset: 26263},
									expr: &seqExpr{
										pos: position{line: 573, col: 35, offset: 26265},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 573, col: 35, offset: 26265},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 573, col: 43, offset: 26273},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 573, col: 54, offset: 26284,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 578, col: 1, offset: 26369},
			expr: &choiceExpr{
				pos: position{line: 578, col: 22, offset: 26390},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 578, col: 22, offset: 26390},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 578, col: 22, offset: 26390},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 578, col: 30, offset: 26398},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 578, col: 42, offset: 26410},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 578, col: 52, offset: 26420},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 581, col: 1, offset: 26480},
			expr: &actionExpr{
				pos: position{line: 581, col: 39, offset: 26518},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 581, col: 39, offset: 26518},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 581, col: 39, offset: 26518},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 581, col: 61, offset: 26540},
							expr: &ruleRefExpr{
								pos:  position{line: 581, col: 61, offset: 26540},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 581, col: 65, offset: 26544},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 581, col: 73, offset: 26552},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 581, col: 81, offset: 26560},
								expr: &seqExpr{
									pos: position{line: 581, col: 82, offset: 26561},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 581, col: 82, offset: 26561},
											expr: &ruleRefExpr{
												pos:  position{line: 581, col: 83, offset: 26562},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 581, col: 105, offset: 26584,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 581, col: 109, offset: 26588},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 581, col: 131, offset: 26610},
							expr: &ruleRefExpr{
								pos:  position{line: 581, col: 131, offset: 26610},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 581, col: 135, offset: 26614},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 585, col: 1, offset: 26698},
			expr: &litMatcher{
				pos:        position{line: 585, col: 26, offset: 26723},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 588, col: 1, offset: 26785},
			expr: &actionExpr{
				pos: position{line: 588, col: 34, offset: 26818},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 588, col: 34, offset: 26818},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 588, col: 34, offset: 26818},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 588, col: 46, offset: 26830},
							expr: &ruleRefExpr{
								pos:  position{line: 588, col: 46, offset: 26830},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 588, col: 50, offset: 26834},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 588, col: 58, offset: 26842},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 588, col: 67, offset: 26851},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 588, col: 88, offset: 26872},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 595, col: 1, offset: 27084},
			expr: &actionExpr{
				pos: position{line: 595, col: 21, offset: 27104},
				run: (*parser).callonElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 595, col: 21, offset: 27104},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 595, col: 21, offset: 27104},
							expr: &ruleRefExpr{
								pos:  position{line: 595, col: 22, offset: 27105},
								name: "AdmonitionMarker",
							},
						},
						&labeledExpr{
							pos:   position{line: 595, col: 39, offset: 27122},
							label: "attr",
							expr: &choiceExpr{
								pos: position{line: 595, col: 45, offset: 27128},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 595, col: 45, offset: 27128},
										name: "ElementID",
									},
									&ruleRefExpr{
										pos:  position{line: 595, col: 57, offset: 27140},
										name: "ElementTitle",
									},
									&ruleRefExpr{
										pos:  position{line: 595, col: 72, offset: 27155},
										name: "AttributeGroup",
									},
									&ruleRefExpr{
										pos:  position{line: 595, col: 89, offset: 27172},
										name: "InvalidElementAttribute",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 595, col: 114, offset: 27197},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 599, col: 1, offset: 27288},
			expr: &choiceExpr{
				pos: position{line: 599, col: 14, offset: 27301},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 599, col: 14, offset: 27301},
						run: (*parser).callonElementID2,
						expr: &labeledExpr{
							pos:   position{line: 599, col: 14, offset: 27301},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 599, col: 18, offset: 27305},
								name: "InlineElementID",
							},
						},
					},
					&actionExpr{
						pos: position{line: 601, col: 5, offset: 27347},
						run: (*parser).callonElementID5,
						expr: &seqExpr{
							pos: position{line: 601, col: 5, offset: 27347},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 601, col: 5, offset: 27347},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 601, col: 10, offset: 27352},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 601, col: 14, offset: 27356},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 601, col: 18, offset: 27360},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 601, col: 22, offset: 27364},
									expr: &ruleRefExpr{
										pos:  position{line: 601, col: 22, offset: 27364},
										name: "WS",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElementID",
			pos:  position{line: 605, col: 1, offset: 27416},
			expr: &actionExpr{
				pos: position{line: 605, col: 20, offset: 27435},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 605, col: 20, offset: 27435},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 605, col: 20, offset: 27435},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 605, col: 25, offset: 27440},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 605, col: 29, offset: 27444},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 605, col: 33, offset: 27448},
							val:        "]]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 605, col: 38, offset: 27453},
							expr: &ruleRefExpr{
								pos:  position{line: 605, col: 38, offset: 27453},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 611, col: 1, offset: 27647},
			expr: &actionExpr{
				pos: position{line: 611, col: 17, offset: 27663},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 611, col: 17, offset: 27663},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 611, col: 17, offset: 27663},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 611, col: 21, offset: 27667},
							expr: &litMatcher{
								pos:        position{line: 611, col: 22, offset: 27668},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 611, col: 26, offset: 27672},
							expr: &ruleRefExpr{
								pos:  position{line: 611, col: 27, offset: 27673},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 611, col: 30, offset: 27676},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 611, col: 36, offset: 27682},
								expr: &seqExpr{
									pos: position{line: 611, col: 37, offset: 27683},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 611, col: 37, offset: 27683},
											expr: &ruleRefExpr{
												pos:  position{line: 611, col: 38, offset: 27684},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 611, col: 46, offset: 27692,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 611, col: 50, offset: 27696},
							expr: &ruleRefExpr{
								pos:  position{line: 611, col: 50, offset: 27696},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeGroup",
			pos:  position{line: 616, col: 1, offset: 27824},
			expr: &actionExpr{
				pos: position{line: 616, col: 19, offset: 27842},
				run: (*parser).callonAttributeGroup1,
				expr: &seqExpr{
					pos: position{line: 616, col: 19, offset: 27842},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 616, col: 19, offset: 27842},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 616, col: 23, offset: 27846},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 616, col: 34, offset: 27857},
								expr: &ruleRefExpr{
									pos:  position{line: 616, col: 35, offset: 27858},
									name: "GenericAttribute",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 616, col: 54, offset: 27877},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 616, col: 58, offset: 27881},
							expr: &ruleRefExpr{
								pos:  position{line: 616, col: 58, offset: 27881},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "GenericAttribute",
			pos:  position{line: 620, col: 1, offset: 27953},
			expr: &choiceExpr{
				pos: position{line: 620, col: 21, offset: 27973},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 620, col: 21, offset: 27973},
						run: (*parser).callonGenericAttribute2,
						expr: &seqExpr{
							pos: position{line: 620, col: 21, offset: 27973},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 620, col: 21, offset: 27973},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 620, col: 26, offset: 27978},
										name: "AttributeKey",
									},
								},
								&litMatcher{
									pos:        position{line: 620, col: 40, offset: 27992},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 620, col: 44, offset: 27996},
									label: "value",
									expr: &ruleRefExpr{
										pos:  position{line: 620, col: 51, offset: 28003},
										name: "AttributeValue",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 620, col: 67, offset: 28019},
									expr: &seqExpr{
										pos: position{line: 620, col: 68, offset: 28020},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 620, col: 68, offset: 28020},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 620, col: 72, offset: 28024},
												expr: &ruleRefExpr{
													pos:  position{line: 620, col: 72, offset: 28024},
													name: "WS",
												},
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 622, col: 5, offset: 28133},
						run: (*parser).callonGenericAttribute14,
						expr: &seqExpr{
							pos: position{line: 622, col: 5, offset: 28133},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 622, col: 5, offset: 28133},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 622, col: 10, offset: 28138},
										name: "AttributeKey",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 622, col: 24, offset: 28152},
									expr: &seqExpr{
										pos: position{line: 622, col: 25, offset: 28153},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 622, col: 25, offset: 28153},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 622, col: 29, offset: 28157},
												expr: &ruleRefExpr{
													pos:  position{line: 622, col: 29, offset: 28157},
													name: "WS",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeKey",
			pos:  position{line: 626, col: 1, offset: 28251},
			expr: &actionExpr{
				pos: position{line: 626, col: 17, offset: 28267},
				run: (*parser).callonAttributeKey1,
				expr: &seqExpr{
					pos: position{line: 626, col: 17, offset: 28267},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 626, col: 17, offset: 28267},
							label: "key",
							expr: &oneOrMoreExpr{
								pos: position{line: 626, col: 22, offset: 28272},
								expr: &seqExpr{
									pos: position{line: 626, col: 23, offset: 28273},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 626, col: 23, offset: 28273},
											expr: &ruleRefExpr{
												pos:  position{line: 626, col: 24, offset: 28274},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 626, col: 27, offset: 28277},
											expr: &litMatcher{
												pos:        position{line: 626, col: 28, offset: 28278},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 626, col: 32, offset: 28282},
											expr: &litMatcher{
												pos:        position{line: 626, col: 33, offset: 28283},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 626, col: 37, offset: 28287},
											expr: &litMatcher{
												pos:        position{line: 626, col: 38, offset: 28288},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 626, col: 42, offset: 28292,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 626, col: 46, offset: 28296},
							expr: &ruleRefExpr{
								pos:  position{line: 626, col: 46, offset: 28296},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeValue",
			pos:  position{line: 631, col: 1, offset: 28378},
			expr: &actionExpr{
				pos: position{line: 631, col: 19, offset: 28396},
				run: (*parser).callonAttributeValue1,
				expr: &seqExpr{
					pos: position{line: 631, col: 19, offset: 28396},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 631, col: 19, offset: 28396},
							expr: &ruleRefExpr{
								pos:  position{line: 631, col: 19, offset: 28396},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 631, col: 23, offset: 28400},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 631, col: 29, offset: 28406},
								expr: &seqExpr{
									pos: position{line: 631, col: 30, offset: 28407},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 631, col: 30, offset: 28407},
											expr: &ruleRefExpr{
												pos:  position{line: 631, col: 31, offset: 28408},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 631, col: 34, offset: 28411},
											expr: &litMatcher{
												pos:        position{line: 631, col: 35, offset: 28412},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 631, col: 39, offset: 28416},
											expr: &litMatcher{
												pos:        position{line: 631, col: 40, offset: 28417},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 631, col: 44, offset: 28421,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 631, col: 48, offset: 28425},
							expr: &ruleRefExpr{
								pos:  position{line: 631, col: 48, offset: 28425},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 636, col: 1, offset: 28512},
			expr: &actionExpr{
				pos: position{line: 636, col: 28, offset: 28539},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 636, col: 28, offset: 28539},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 636, col: 28, offset: 28539},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 636, col: 32, offset: 28543},
							expr: &ruleRefExpr{
								pos:  position{line: 636, col: 32, offset: 28543},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 636, col: 36, offset: 28547},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 636, col: 44, offset: 28555},
								expr: &seqExpr{
									pos: position{line: 636, col: 45, offset: 28556},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 636, col: 45, offset: 28556},
											expr: &litMatcher{
												pos:        position{line: 636, col: 46, offset: 28557},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 636, col: 50, offset: 28561,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 636, col: 54, offset: 28565},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 636, col: 58, offset: 28569},
							expr: &ruleRefExpr{
								pos:  position{line: 636, col: 58, offset: 28569},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 643, col: 1, offset: 28735},
			expr: &actionExpr{
				pos: position{line: 643, col: 14, offset: 28748},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 643, col: 14, offset: 28748},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 643, col: 14, offset: 28748},
							expr: &ruleRefExpr{
								pos:  position{line: 643, col: 15, offset: 28749},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 643, col: 19, offset: 28753},
							expr: &ruleRefExpr{
								pos:  position{line: 643, col: 19, offset: 28753},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 643, col: 23, offset: 28757},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 650, col: 1, offset: 28904},
			expr: &actionExpr{
				pos: position{line: 650, col: 15, offset: 28918},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 650, col: 15, offset: 28918},
					expr: &seqExpr{
						pos: position{line: 650, col: 16, offset: 28919},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 650, col: 16, offset: 28919},
								expr: &ruleRefExpr{
									pos:  position{line: 650, col: 17, offset: 28920},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 650, col: 25, offset: 28928},
								expr: &ruleRefExpr{
									pos:  position{line: 650, col: 26, offset: 28929},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 650, col: 29, offset: 28932,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 654, col: 1, offset: 28972},
			expr: &actionExpr{
				pos: position{line: 654, col: 8, offset: 28979},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 654, col: 8, offset: 28979},
					expr: &seqExpr{
						pos: position{line: 654, col: 9, offset: 28980},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 654, col: 9, offset: 28980},
								expr: &ruleRefExpr{
									pos:  position{line: 654, col: 10, offset: 28981},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 654, col: 18, offset: 28989},
								expr: &ruleRefExpr{
									pos:  position{line: 654, col: 19, offset: 28990},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 654, col: 22, offset: 28993},
								expr: &litMatcher{
									pos:        position{line: 654, col: 23, offset: 28994},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 654, col: 27, offset: 28998},
								expr: &litMatcher{
									pos:        position{line: 654, col: 28, offset: 28999},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 654, col: 32, offset: 29003,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 658, col: 1, offset: 29043},
			expr: &actionExpr{
				pos: position{line: 658, col: 7, offset: 29049},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 658, col: 7, offset: 29049},
					expr: &seqExpr{
						pos: position{line: 658, col: 8, offset: 29050},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 658, col: 8, offset: 29050},
								expr: &ruleRefExpr{
									pos:  position{line: 658, col: 9, offset: 29051},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 658, col: 17, offset: 29059},
								expr: &ruleRefExpr{
									pos:  position{line: 658, col: 18, offset: 29060},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 658, col: 21, offset: 29063},
								expr: &litMatcher{
									pos:        position{line: 658, col: 22, offset: 29064},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 658, col: 26, offset: 29068},
								expr: &litMatcher{
									pos:        position{line: 658, col: 27, offset: 29069},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 658, col: 31, offset: 29073},
								expr: &litMatcher{
									pos:        position{line: 658, col: 32, offset: 29074},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 658, col: 37, offset: 29079},
								expr: &litMatcher{
									pos:        position{line: 658, col: 38, offset: 29080},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 658, col: 42, offset: 29084,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 662, col: 1, offset: 29124},
			expr: &actionExpr{
				pos: position{line: 662, col: 13, offset: 29136},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 662, col: 13, offset: 29136},
					expr: &seqExpr{
						pos: position{line: 662, col: 14, offset: 29137},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 662, col: 14, offset: 29137},
								expr: &ruleRefExpr{
									pos:  position{line: 662, col: 15, offset: 29138},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 662, col: 23, offset: 29146},
								expr: &litMatcher{
									pos:        position{line: 662, col: 24, offset: 29147},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 662, col: 28, offset: 29151},
								expr: &litMatcher{
									pos:        position{line: 662, col: 29, offset: 29152},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 662, col: 33, offset: 29156,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 666, col: 1, offset: 29196},
			expr: &choiceExpr{
				pos: position{line: 666, col: 15, offset: 29210},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 666, col: 15, offset: 29210},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 666, col: 27, offset: 29222},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 666, col: 40, offset: 29235},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 666, col: 51, offset: 29246},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 666, col: 62, offset: 29257},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 668, col: 1, offset: 29268},
			expr: &charClassMatcher{
				pos:        position{line: 668, col: 10, offset: 29277},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 670, col: 1, offset: 29284},
			expr: &choiceExpr{
				pos: position{line: 670, col: 12, offset: 29295},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 670, col: 12, offset: 29295},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 670, col: 21, offset: 29304},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 670, col: 28, offset: 29311},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 672, col: 1, offset: 29317},
			expr: &choiceExpr{
				pos: position{line: 672, col: 7, offset: 29323},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 672, col: 7, offset: 29323},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 672, col: 13, offset: 29329},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 672, col: 13, offset: 29329},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 676, col: 1, offset: 29374},
			expr: &notExpr{
				pos: position{line: 676, col: 8, offset: 29381},
				expr: &anyMatcher{
					line: 676, col: 9, offset: 29382,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 678, col: 1, offset: 29385},
			expr: &choiceExpr{
				pos: position{line: 678, col: 8, offset: 29392},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 678, col: 8, offset: 29392},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 678, col: 18, offset: 29402},
						name: "EOF",
					},
				},
			},
		},
	},
}

func (c *current) onDocument1(frontMatter, documentHeader, blocks interface{}) (interface{}, error) {
	return types.NewDocument(frontMatter, documentHeader, blocks.([]interface{}))
}

func (p *parser) callonDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument1(stack["frontMatter"], stack["documentHeader"], stack["blocks"])
}

func (c *current) onDocumentBlocks7(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonDocumentBlocks7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentBlocks7(stack["content"])
}

func (c *current) onPreamble1(elements interface{}) (interface{}, error) {
	return types.NewPreamble(elements.([]interface{}))
}

func (p *parser) callonPreamble1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPreamble1(stack["elements"])
}

func (c *current) onFrontMatter1(content interface{}) (interface{}, error) {
	return types.NewYamlFrontMatter(content.(string))
}

func (p *parser) callonFrontMatter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrontMatter1(stack["content"])
}

func (c *current) onYamlFrontMatterContent1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonYamlFrontMatterContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onYamlFrontMatterContent1()
}

func (c *current) onDocumentHeader1(header, authors, revision, otherAttributes interface{}) (interface{}, error) {

	return types.NewDocumentHeader(header, authors, revision, otherAttributes.([]interface{}))
}

func (p *parser) callonDocumentHeader1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentHeader1(stack["header"], stack["authors"], stack["revision"], stack["otherAttributes"])
}

func (c *current) onDocumentTitle1(attributes, level, content, id interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonDocumentTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentTitle1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onDocumentAuthorsInlineForm1(authors interface{}) (interface{}, error) {
	return types.NewDocumentAuthors(authors.([]interface{}))
}

func (p *parser) callonDocumentAuthorsInlineForm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthorsInlineForm1(stack["authors"])
}

func (c *current) onDocumentAuthorsAttributeForm1(author interface{}) (interface{}, error) {
	return []types.DocumentAuthor{author.(types.DocumentAuthor)}, nil
}

func (p *parser) callonDocumentAuthorsAttributeForm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthorsAttributeForm1(stack["author"])
}

func (c *current) onDocumentAuthor1(namePart1, namePart2, namePart3, email interface{}) (interface{}, error) {
	return types.NewDocumentAuthor(namePart1, namePart2, namePart3, email)
}

func (p *parser) callonDocumentAuthor1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthor1(stack["namePart1"], stack["namePart2"], stack["namePart3"], stack["email"])
}

func (c *current) onDocumentRevision1(revnumber, revdate, revremark interface{}) (interface{}, error) {
	return types.NewDocumentRevision(revnumber, revdate, revremark)
}

func (p *parser) callonDocumentRevision1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentRevision1(stack["revnumber"], stack["revdate"], stack["revremark"])
}

func (c *current) onDocumentAttributeDeclarationWithNameOnly1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeDeclaration(name.([]interface{}), nil)
}

func (p *parser) callonDocumentAttributeDeclarationWithNameOnly1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeDeclarationWithNameOnly1(stack["name"])
}

func (c *current) onDocumentAttributeDeclarationWithNameAndValue1(name, value interface{}) (interface{}, error) {
	return types.NewDocumentAttributeDeclaration(name.([]interface{}), value.([]interface{}))
}

func (p *parser) callonDocumentAttributeDeclarationWithNameAndValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeDeclarationWithNameAndValue1(stack["name"], stack["value"])
}

func (c *current) onDocumentAttributeResetWithSectionTitleBangSymbol1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeReset(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeResetWithSectionTitleBangSymbol1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeResetWithSectionTitleBangSymbol1(stack["name"])
}

func (c *current) onDocumentAttributeResetWithTrailingBangSymbol1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeReset(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeResetWithTrailingBangSymbol1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeResetWithTrailingBangSymbol1(stack["name"])
}

func (c *current) onDocumentAttributeSubstitution1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeSubstitution(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeSubstitution1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeSubstitution1(stack["name"])
}

func (c *current) onSection11(header, elements interface{}) (interface{}, error) {
	return types.NewSection(1, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection11(stack["header"], stack["elements"])
}

func (c *current) onSection1Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection1Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Block1(stack["content"])
}

func (c *current) onSection21(header, elements interface{}) (interface{}, error) {
	return types.NewSection(2, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection21(stack["header"], stack["elements"])
}

func (c *current) onSection2Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection2Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Block1(stack["content"])
}

func (c *current) onSection31(header, elements interface{}) (interface{}, error) {
	return types.NewSection(3, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection31() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection31(stack["header"], stack["elements"])
}

func (c *current) onSection3Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection3Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Block1(stack["content"])
}

func (c *current) onSection41(header, elements interface{}) (interface{}, error) {
	return types.NewSection(4, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection41() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection41(stack["header"], stack["elements"])
}

func (c *current) onSection4Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection4Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Block1(stack["content"])
}

func (c *current) onSection51(header, elements interface{}) (interface{}, error) {
	return types.NewSection(5, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection51() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection51(stack["header"], stack["elements"])
}

func (c *current) onSection5Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection5Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Block1(stack["content"])
}

func (c *current) onSection1Title1(attributes, level, content, id interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection1Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection2Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection2Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection3Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection3Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection4Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection4Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection5Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection5Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onList1(attributes, elements interface{}) (interface{}, error) {
	return types.NewList(elements.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonList1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onList1(stack["attributes"], stack["elements"])
}

func (c *current) onListAttribute1(attribute interface{}) (interface{}, error) {
	return attribute, nil
}

func (p *parser) callonListAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListAttribute1(stack["attribute"])
}

func (c *current) onListID1(id interface{}) (interface{}, error) {
	return map[string]interface{}{"ID": id.(string)}, nil
}

func (p *parser) callonListID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListID1(stack["id"])
}

func (c *current) onHorizontalLayout1() (interface{}, error) {
	return map[string]interface{}{"layout": "horizontal"}, nil
}

func (p *parser) callonHorizontalLayout1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHorizontalLayout1()
}

func (c *current) onListParagraph1(lines interface{}) (interface{}, error) {
	return types.NewListParagraph(lines.([]interface{}))
}

func (p *parser) callonListParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListParagraph1(stack["lines"])
}

func (c *current) onListItemContinuation1() (interface{}, error) {
	return types.NewListItemContinuation()
}

func (p *parser) callonListItemContinuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListItemContinuation1()
}

func (c *current) onContinuedBlockElement1(element interface{}) (interface{}, error) {
	return element, nil
}

func (p *parser) callonContinuedBlockElement1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onContinuedBlockElement1(stack["element"])
}

func (c *current) onOrderedListItem1(attributes, prefix, content interface{}) (interface{}, error) {
	return types.NewOrderedListItem(prefix.(types.OrderedListItemPrefix), content.([]types.DocElement), attributes.([]interface{}))
}

func (p *parser) callonOrderedListItem1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItem1(stack["attributes"], stack["prefix"], stack["content"])
}

func (c *current) onOrderedListItemPrefix2(style interface{}) (interface{}, error) {
	// numbering style: "."
	return types.NewOrderedListItemPrefix(types.Arabic, 1)
}

func (p *parser) callonOrderedListItemPrefix2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix2(stack["style"])
}

func (c *current) onOrderedListItemPrefix10(style interface{}) (interface{}, error) {
	// numbering style: ".."
	return types.NewOrderedListItemPrefix(types.LowerAlpha, 2)
}

func (p *parser) callonOrderedListItemPrefix10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix10(stack["style"])
}

func (c *current) onOrderedListItemPrefix18(style interface{}) (interface{}, error) {
	// numbering style: "..."
	return types.NewOrderedListItemPrefix(types.LowerRoman, 3)
}

func (p *parser) callonOrderedListItemPrefix18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix18(stack["style"])
}

func (c *current) onOrderedListItemPrefix26(style interface{}) (interface{}, error) {
	// numbering style: "...."
	return types.NewOrderedListItemPrefix(types.UpperAlpha, 4)
}

func (p *parser) callonOrderedListItemPrefix26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix26(stack["style"])
}

func (c *current) onOrderedListItemPrefix34(style interface{}) (interface{}, error) {
	// numbering style: "....."
	return types.NewOrderedListItemPrefix(types.UpperRoman, 5)
	// explicit numbering
}

func (p *parser) callonOrderedListItemPrefix34() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix34(stack["style"])
}

func (c *current) onOrderedListItemPrefix42(style interface{}) (interface{}, error) {
	// numbering style: "1."
	return types.NewOrderedListItemPrefix(types.Arabic, 1)
}

func (p *parser) callonOrderedListItemPrefix42() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix42(stack["style"])
}

func (c *current) onOrderedListItemPrefix60(style interface{}) (interface{}, error) {
	// numbering style: "a."
	return types.NewOrderedListItemPrefix(types.LowerAlpha, 1)
}

func (p *parser) callonOrderedListItemPrefix60() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix60(stack["style"])
}

func (c *current) onOrderedListItemPrefix78(style interface{}) (interface{}, error) {
	// numbering style: "A."
	return types.NewOrderedListItemPrefix(types.UpperAlpha, 1)
}

func (p *parser) callonOrderedListItemPrefix78() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix78(stack["style"])
}

func (c *current) onOrderedListItemPrefix96(style interface{}) (interface{}, error) {
	// numbering style: "i)"
	return types.NewOrderedListItemPrefix(types.LowerRoman, 1)
}

func (p *parser) callonOrderedListItemPrefix96() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix96(stack["style"])
}

func (c *current) onOrderedListItemPrefix114(style interface{}) (interface{}, error) {
	// numbering style: "I)"
	return types.NewOrderedListItemPrefix(types.UpperRoman, 1)
}

func (p *parser) callonOrderedListItemPrefix114() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix114(stack["style"])
}

func (c *current) onOrderedListItemContent1(elements interface{}) (interface{}, error) {
	// Another list or a literal paragraph immediately following a list item will be implicitly included in the list item
	return types.NewListItemContent(elements.([]interface{}))
}

func (p *parser) callonOrderedListItemContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemContent1(stack["elements"])
}

func (c *current) onUnorderedListItem1(prefix, content interface{}) (interface{}, error) {
	return types.NewUnorderedListItem(prefix.(types.UnorderedListItemPrefix), content.([]types.DocElement))
}

func (p *parser) callonUnorderedListItem1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItem1(stack["prefix"], stack["content"])
}

func (c *current) onUnorderedListItemPrefix2(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.FiveAsterisks, 5)
}

func (p *parser) callonUnorderedListItemPrefix2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix2(stack["level"])
}

func (c *current) onUnorderedListItemPrefix10(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.FourAsterisks, 4)
}

func (p *parser) callonUnorderedListItemPrefix10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix10(stack["level"])
}

func (c *current) onUnorderedListItemPrefix18(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.ThreeAsterisks, 3)
}

func (p *parser) callonUnorderedListItemPrefix18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix18(stack["level"])
}

func (c *current) onUnorderedListItemPrefix26(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.TwoAsterisks, 2)
}

func (p *parser) callonUnorderedListItemPrefix26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix26(stack["level"])
}

func (c *current) onUnorderedListItemPrefix34(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.OneAsterisk, 1)
}

func (p *parser) callonUnorderedListItemPrefix34() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix34(stack["level"])
}

func (c *current) onUnorderedListItemPrefix42(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.Dash, 1)
}

func (p *parser) callonUnorderedListItemPrefix42() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix42(stack["level"])
}

func (c *current) onUnorderedListItemContent1(elements interface{}) (interface{}, error) {
	// Another list or a literal paragraph immediately following a list item will be implicitly included in the list item
	return types.NewListItemContent(elements.([]interface{}))
}

func (p *parser) callonUnorderedListItemContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemContent1(stack["elements"])
}

func (c *current) onLabeledListItem2(term, description interface{}) (interface{}, error) {
	return types.NewLabeledListItem(term.([]interface{}), description.([]types.DocElement))
}

func (p *parser) callonLabeledListItem2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItem2(stack["term"], stack["description"])
}

func (c *current) onLabeledListItem9(term interface{}) (interface{}, error) {
	// here, WS is optional since there is no description afterwards
	return types.NewLabeledListItem(term.([]interface{}), nil)
}

func (p *parser) callonLabeledListItem9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItem9(stack["term"])
}

func (c *current) onLabeledListItemTerm1(term interface{}) (interface{}, error) {
	return term, nil
}

func (p *parser) callonLabeledListItemTerm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemTerm1(stack["term"])
}

func (c *current) onLabeledListItemDescription1(elements interface{}) (interface{}, error) {
	// TODO: replace with (ListParagraph+ ContinuedBlockElement*) and use a single rule for all item contents ?
	return types.NewListItemContent(elements.([]interface{}))
}

func (p *parser) callonLabeledListItemDescription1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemDescription1(stack["elements"])
}

func (c *current) onParagraph1(attributes, lines interface{}) (interface{}, error) {
	return types.NewParagraph(lines.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph1(stack["attributes"], stack["lines"])
}

func (c *current) onInlineContentWithTrailingSpaces1(elements interface{}) (interface{}, error) {
	// includes heading and trailing spaces in the elements arg
	return types.NewInlineContent(elements.([]interface{}))
}

func (p *parser) callonInlineContentWithTrailingSpaces1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineContentWithTrailingSpaces1(stack["elements"])
}

func (c *current) onInlineContent1(elements interface{}) (interface{}, error) {
	// absorbs heading and trailing spaces
	return types.NewInlineContent(elements.([]interface{}))
}

func (p *parser) callonInlineContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineContent1(stack["elements"])
}

func (c *current) onAdmonitionParagraph2(attributes, t, content interface{}) (interface{}, error) {
	// paragraph style
	return types.NewAdmonitionParagraph(t.(types.AdmonitionKind), content, attributes.([]interface{}))
}

func (p *parser) callonAdmonitionParagraph2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionParagraph2(stack["attributes"], stack["t"], stack["content"])
}

func (c *current) onAdmonitionParagraph18(attributes, k, otherAttributes, content interface{}) (interface{}, error) {
	// block style
	return types.NewAdmonitionParagraph(k.(types.AdmonitionKind), content, append(attributes.([]interface{}), otherAttributes.([]interface{})...))
}

func (p *parser) callonAdmonitionParagraph18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionParagraph18(stack["attributes"], stack["k"], stack["otherAttributes"], stack["content"])
}

func (c *current) onAdmonitionParagraphContent1(lines interface{}) (interface{}, error) {
	return types.NewAdmonitionParagraphContent(lines.([]interface{}))
}

func (p *parser) callonAdmonitionParagraphContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionParagraphContent1(stack["lines"])
}

func (c *current) onAdmonitionMarker1(k interface{}) (interface{}, error) {
	return k, nil
}

func (p *parser) callonAdmonitionMarker1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionMarker1(stack["k"])
}

func (c *current) onAdmonitionKind2() (interface{}, error) {
	return types.Tip, nil
}

func (p *parser) callonAdmonitionKind2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind2()
}

func (c *current) onAdmonitionKind4() (interface{}, error) {
	return types.Note, nil
}

func (p *parser) callonAdmonitionKind4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind4()
}

func (c *current) onAdmonitionKind6() (interface{}, error) {
	return types.Important, nil
}

func (p *parser) callonAdmonitionKind6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind6()
}

func (c *current) onAdmonitionKind8() (interface{}, error) {
	return types.Warning, nil
}

func (p *parser) callonAdmonitionKind8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind8()
}

func (c *current) onAdmonitionKind10() (interface{}, error) {
	return types.Caution, nil
}

func (p *parser) callonAdmonitionKind10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind10()
}

func (c *current) onBoldText2(content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldText2(stack["content"])
}

func (c *current) onBoldText10(content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewQuotedText(types.Bold, result)
}

func (p *parser) callonBoldText10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldText10(stack["content"])
}

func (c *current) onBoldText18(content interface{}) (interface{}, error) {
	// single punctuation
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldText18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldText18(stack["content"])
}

func (c *current) onEscapedBoldText2(backslashes, content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "**", content.([]interface{}))
}

func (p *parser) callonEscapedBoldText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldText2(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldText13(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", result)
}

func (p *parser) callonEscapedBoldText13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldText13(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldText24(backslashes, content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", content.([]interface{}))
}

func (p *parser) callonEscapedBoldText24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldText24(stack["backslashes"], stack["content"])
}

func (c *current) onItalicText2(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicText2(stack["content"])
}

func (c *current) onItalicText10(content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewQuotedText(types.Italic, result)
}

func (p *parser) callonItalicText10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicText10(stack["content"])
}

func (c *current) onItalicText18(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicText18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicText18(stack["content"])
}

func (c *current) onEscapedItalicText2(backslashes, content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "__", content.([]interface{}))
}

func (p *parser) callonEscapedItalicText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicText2(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicText13(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", result)
}

func (p *parser) callonEscapedItalicText13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicText13(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicText24(backslashes, content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", content.([]interface{}))
}

func (p *parser) callonEscapedItalicText24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicText24(stack["backslashes"], stack["content"])
}

func (c *current) onMonospaceText2(content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceText2(stack["content"])
}

func (c *current) onMonospaceText10(content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewQuotedText(types.Monospace, result)
}

func (p *parser) callonMonospaceText10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceText10(stack["content"])
}

func (c *current) onMonospaceText18(content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceText18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceText18(stack["content"])
}

func (c *current) onEscapedMonospaceText2(backslashes, content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "``", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceText2(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceText13(backslashes, content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", result)
}

func (p *parser) callonEscapedMonospaceText13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceText13(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceText24(backslashes, content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceText24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceText24(stack["backslashes"], stack["content"])
}

func (c *current) onCharactersWithQuotePunctuation1() (interface{}, error) {
	// can have "*", "_" or "`" within, maybe because the user inserted another quote, or made an error (extra or missing space, for example)
	return c.text, nil
}

func (p *parser) callonCharactersWithQuotePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharactersWithQuotePunctuation1()
}

func (c *current) onSinglePlusPassthrough1(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.SinglePlusPassthrough, content.([]interface{}))
}

func (p *parser) callonSinglePlusPassthrough1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinglePlusPassthrough1(stack["content"])
}

func (c *current) onTriplePlusPassthrough1(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.TriplePlusPassthrough, content.([]interface{}))
}

func (p *parser) callonTriplePlusPassthrough1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTriplePlusPassthrough1(stack["content"])
}

func (c *current) onPassthroughMacro2(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.PassthroughMacro, content.([]interface{}))
}

func (p *parser) callonPassthroughMacro2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPassthroughMacro2(stack["content"])
}

func (c *current) onPassthroughMacro9(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.PassthroughMacro, content.([]interface{}))
}

func (p *parser) callonPassthroughMacro9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPassthroughMacro9(stack["content"])
}

func (c *current) onCrossReference1(id interface{}) (interface{}, error) {
	return types.NewCrossReference(id.(string))
}

func (p *parser) callonCrossReference1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCrossReference1(stack["id"])
}

func (c *current) onExternalLink1(url, text interface{}) (interface{}, error) {
	if text != nil {
		return types.NewLink(url.([]interface{}), text.([]interface{}))
	}
	return types.NewLink(url.([]interface{}), nil)
}

func (p *parser) callonExternalLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternalLink1(stack["url"], stack["text"])
}

func (c *current) onRelativeLink1(url, text interface{}) (interface{}, error) {
	if text != nil {
		return types.NewLink(url.([]interface{}), text.([]interface{}))
	}
	return types.NewLink(url.([]interface{}), nil)
}

func (p *parser) callonRelativeLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRelativeLink1(stack["url"], stack["text"])
}

func (c *current) onBlockImage1(attributes, image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewBlockImage(image.(types.ImageMacro), attributes.([]interface{}))
}

func (p *parser) callonBlockImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImage1(stack["attributes"], stack["image"])
}

func (c *current) onBlockImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(path.(string), attributes)
}

func (p *parser) callonBlockImageMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImageMacro1(stack["path"], stack["attributes"])
}

func (c *current) onInlineImage1(image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewInlineImage(image.(types.ImageMacro))
}

func (p *parser) callonInlineImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineImage1(stack["image"])
}

func (c *current) onInlineImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(path.(string), attributes)
}

func (p *parser) callonInlineImageMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineImageMacro1(stack["path"], stack["attributes"])
}

func (c *current) onFencedBlock1(content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.FencedBlock, content.([]interface{}), nil)
}

func (p *parser) callonFencedBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFencedBlock1(stack["content"])
}

func (c *current) onListingBlock1(content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.ListingBlock, content.([]interface{}), nil)
}

func (p *parser) callonListingBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListingBlock1(stack["content"])
}

func (c *current) onExampleBlock1(attributes, content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.ExampleBlock, content.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonExampleBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExampleBlock1(stack["attributes"], stack["content"])
}

func (c *current) onParagraphWithSpaces1(spaces, content interface{}) (interface{}, error) {
	return types.NewLiteralBlock(spaces.([]interface{}), content.([]interface{}))
}

func (p *parser) callonParagraphWithSpaces1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithSpaces1(stack["spaces"], stack["content"])
}

func (c *current) onLiteralBlockContent1(content interface{}) (interface{}, error) {

	return content, nil
}

func (p *parser) callonLiteralBlockContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLiteralBlockContent1(stack["content"])
}

func (c *current) onParagraphWithLiteralBlockDelimiter1(content interface{}) (interface{}, error) {
	return types.NewLiteralBlock([]interface{}{}, content.([]interface{}))
}

func (p *parser) callonParagraphWithLiteralBlockDelimiter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithLiteralBlockDelimiter1(stack["content"])
}

func (c *current) onParagraphWithLiteralAttribute1(content interface{}) (interface{}, error) {
	return types.NewLiteralBlock([]interface{}{}, content.([]interface{}))
}

func (p *parser) callonParagraphWithLiteralAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithLiteralAttribute1(stack["content"])
}

func (c *current) onElementAttribute1(attr interface{}) (interface{}, error) {
	return attr, nil // avoid returning something like `[]interface{}{attr, EOL}`
}

func (p *parser) callonElementAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementAttribute1(stack["attr"])
}

func (c *current) onElementID2(id interface{}) (interface{}, error) {
	return id, nil
}

func (p *parser) callonElementID2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID2(stack["id"])
}

func (c *current) onElementID5(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonElementID5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID5(stack["id"])
}

func (c *current) onInlineElementID1(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonInlineElementID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineElementID1(stack["id"])
}

func (c *current) onElementTitle1(title interface{}) (interface{}, error) {
	return types.NewElementTitle(title.([]interface{}))
}

func (p *parser) callonElementTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementTitle1(stack["title"])
}

func (c *current) onAttributeGroup1(attributes interface{}) (interface{}, error) {
	return types.NewAttributeGroup(attributes.([]interface{}))
}

func (p *parser) callonAttributeGroup1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttributeGroup1(stack["attributes"])
}

func (c *current) onGenericAttribute2(key, value interface{}) (interface{}, error) {
	// value is set
	return types.NewGenericAttribute(key.([]interface{}), value.([]interface{}))
}

func (p *parser) callonGenericAttribute2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGenericAttribute2(stack["key"], stack["value"])
}

func (c *current) onGenericAttribute14(key interface{}) (interface{}, error) {
	// value is not set
	return types.NewGenericAttribute(key.([]interface{}), nil)
}

func (p *parser) callonGenericAttribute14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGenericAttribute14(stack["key"])
}

func (c *current) onAttributeKey1(key interface{}) (interface{}, error) {
	// fmt.Printf("found attribute key: %v\n", key)
	return key, nil
}

func (p *parser) callonAttributeKey1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttributeKey1(stack["key"])
}

func (c *current) onAttributeValue1(value interface{}) (interface{}, error) {
	// fmt.Printf("found attribute value: %v\n", value)
	return value, nil
}

func (p *parser) callonAttributeValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttributeValue1(stack["value"])
}

func (c *current) onInvalidElementAttribute1(content interface{}) (interface{}, error) {
	return types.NewInvalidElementAttribute(c.text)
}

func (p *parser) callonInvalidElementAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInvalidElementAttribute1(stack["content"])
}

func (c *current) onBlankLine1() (interface{}, error) {
	return types.NewBlankLine()
}

func (p *parser) callonBlankLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlankLine1()
}

func (c *current) onCharacters1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonCharacters1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharacters1()
}

func (c *current) onURL1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonURL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onURL1()
}

func (c *current) onID1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onID1()
}

func (c *current) onURL_TEXT1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonURL_TEXT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onURL_TEXT1()
}

func (c *current) onWS3() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonWS3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWS3()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEntrypoint is returned when the specified entrypoint rule
	// does not exit.
	errInvalidEntrypoint = errors.New("invalid entrypoint")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errMaxExprCnt is used to signal that the maximum number of
	// expressions have been parsed.
	errMaxExprCnt = errors.New("max number of expresssions parsed")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// MaxExpressions creates an Option to stop parsing after the provided
// number of expressions have been parsed, if the value is 0 then the parser will
// parse for as many steps as needed (possibly an infinite number).
//
// The default for maxExprCnt is 0.
func MaxExpressions(maxExprCnt uint64) Option {
	return func(p *parser) Option {
		oldMaxExprCnt := p.maxExprCnt
		p.maxExprCnt = maxExprCnt
		return MaxExpressions(oldMaxExprCnt)
	}
}

// Entrypoint creates an Option to set the rule name to use as entrypoint.
// The rule name must have been specified in the -alternate-entrypoints
// if generating the parser with the -optimize-grammar flag, otherwise
// it may have been optimized out. Passing an empty string sets the
// entrypoint to the first rule in the grammar.
//
// The default is to start parsing at the first rule in the grammar.
func Entrypoint(ruleName string) Option {
	return func(p *parser) Option {
		oldEntrypoint := p.entrypoint
		p.entrypoint = ruleName
		if ruleName == "" {
			p.entrypoint = g.rules[0].name
		}
		return Entrypoint(oldEntrypoint)
	}
}

// Statistics adds a user provided Stats struct to the parser to allow
// the user to process the results after the parsing has finished.
// Also the key for the "no match" counter is set.
//
// Example usage:
//
//     input := "input"
//     stats := Stats{}
//     _, err := Parse("input-file", []byte(input), Statistics(&stats, "no match"))
//     if err != nil {
//         log.Panicln(err)
//     }
//     b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
//     if err != nil {
//         log.Panicln(err)
//     }
//     fmt.Println(string(b))
//
func Statistics(stats *Stats, choiceNoMatch string) Option {
	return func(p *parser) Option {
		oldStats := p.Stats
		p.Stats = stats
		oldChoiceNoMatch := p.choiceNoMatch
		p.choiceNoMatch = choiceNoMatch
		if p.Stats.ChoiceAltCnt == nil {
			p.Stats.ChoiceAltCnt = make(map[string]map[string]int)
		}
		return Statistics(oldStats, oldChoiceNoMatch)
	}
}

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// AllowInvalidUTF8 creates an Option to allow invalid UTF-8 bytes.
// Every invalid UTF-8 byte is treated as a utf8.RuneError (U+FFFD)
// by character class matchers and is matched by the any matcher.
// The returned matched value, c.text and c.offset are NOT affected.
//
// The default is false.
func AllowInvalidUTF8(b bool) Option {
	return func(p *parser) Option {
		old := p.allowInvalidUTF8
		p.allowInvalidUTF8 = b
		return AllowInvalidUTF8(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// GlobalStore creates an Option to set a key to a certain value in
// the globalStore.
func GlobalStore(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.globalStore[key]
		p.cur.globalStore[key] = value
		return GlobalStore(key, old)
	}
}

// InitState creates an Option to set a key to a certain value in
// the global "state" store.
func InitState(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.state[key]
		p.cur.state[key] = value
		return InitState(key, old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match

	// state is a store for arbitrary key,value pairs that the user wants to be
	// tied to the backtracking of the parser.
	// This is always rolled back if a parsing rule fails.
	state storeDict

	// globalStore is a general store for the user to store arbitrary key-value
	// pairs that they need to manage and that they do not want tied to the
	// backtracking of the parser. This is only modified by the user and never
	// rolled back by the parser. It is always up to the user to keep this in a
	// consistent state.
	globalStore storeDict
}

type storeDict map[string]interface{}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type recoveryExpr struct {
	pos          position
	expr         interface{}
	recoverExpr  interface{}
	failureLabel []string
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type throwExpr struct {
	pos   position
	label string
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type stateCodeExpr struct {
	pos position
	run func(*parser) error
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos             position
	val             string
	basicLatinChars [128]bool
	chars           []rune
	ranges          []rune
	classes         []*unicode.RangeTable
	ignoreCase      bool
	inverted        bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	stats := Stats{
		ChoiceAltCnt: make(map[string]map[string]int),
	}

	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
		cur: current{
			state:       make(storeDict),
			globalStore: make(storeDict),
		},
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make([]string, 0, 20),
		Stats:           &stats,
		// start rule is rule [0] unless an alternate entrypoint is specified
		entrypoint: g.rules[0].name,
		emptyState: make(storeDict),
	}
	p.setOptions(opts)

	if p.maxExprCnt == 0 {
		p.maxExprCnt = math.MaxUint64
	}

	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

const choiceNoMatch = -1

// Stats stores some statistics, gathered during parsing
type Stats struct {
	// ExprCnt counts the number of expressions processed during parsing
	// This value is compared to the maximum number of expressions allowed
	// (set by the MaxExpressions option).
	ExprCnt uint64

	// ChoiceAltCnt is used to count for each ordered choice expression,
	// which alternative is used how may times.
	// These numbers allow to optimize the order of the ordered choice expression
	// to increase the performance of the parser
	//
	// The outer key of ChoiceAltCnt is composed of the name of the rule as well
	// as the line and the column of the ordered choice.
	// The inner key of ChoiceAltCnt is the number (one-based) of the matching alternative.
	// For each alternative the number of matches are counted. If an ordered choice does not
	// match, a special counter is incremented. The name of this counter is set with
	// the parser option Statistics.
	// For an alternative to be included in ChoiceAltCnt, it has to match at least once.
	ChoiceAltCnt map[string]map[string]int
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	depth   int
	recover bool
	debug   bool

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// parse fail
	maxFailPos            position
	maxFailExpected       []string
	maxFailInvertExpected bool

	// max number of expressions to be parsed
	maxExprCnt uint64
	// entrypoint for the parser
	entrypoint string

	allowInvalidUTF8 bool

	*Stats

	choiceNoMatch string
	// recovery expression stack, keeps track of the currently available recovery expression, these are traversed in reverse
	recoveryStack []map[string]interface{}

	// emptyState contains an empty storeDict, which is used to optimize cloneState if global "state" store is not used.
	emptyState storeDict
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

// push a recovery expression with its labels to the recoveryStack
func (p *parser) pushRecovery(labels []string, expr interface{}) {
	if cap(p.recoveryStack) == len(p.recoveryStack) {
		// create new empty slot in the stack
		p.recoveryStack = append(p.recoveryStack, nil)
	} else {
		// slice to 1 more
		p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)+1]
	}

	m := make(map[string]interface{}, len(labels))
	for _, fl := range labels {
		m[fl] = expr
	}
	p.recoveryStack[len(p.recoveryStack)-1] = m
}

// pop a recovery expression from the recoveryStack
func (p *parser) popRecovery() {
	// GC that map
	p.recoveryStack[len(p.recoveryStack)-1] = nil

	p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = p.maxFailExpected[:0]
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected = append(p.maxFailExpected, want)
	}
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError && n == 1 { // see utf8.DecodeRune
		if !p.allowInvalidUTF8 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// Cloner is implemented by any value that has a Clone method, which returns a
// copy of the value. This is mainly used for types which are not passed by
// value (e.g map, slice, chan) or structs that contain such types.
//
// This is used in conjunction with the global state feature to create proper
// copies of the state to allow the parser to properly restore the state in
// the case of backtracking.
type Cloner interface {
	Clone() interface{}
}

// clone and return parser current state.
func (p *parser) cloneState() storeDict {
	if p.debug {
		defer p.out(p.in("cloneState"))
	}

	if len(p.cur.state) == 0 {
		if len(p.emptyState) > 0 {
			p.emptyState = make(storeDict)
		}
		return p.emptyState
	}

	state := make(storeDict, len(p.cur.state))
	for k, v := range p.cur.state {
		if c, ok := v.(Cloner); ok {
			state[k] = c.Clone()
		} else {
			state[k] = v
		}
	}
	return state
}

// restore parser current state to the state storeDict.
// every restoreState should applied only one time for every cloned state
func (p *parser) restoreState(state storeDict) {
	if p.debug {
		defer p.out(p.in("restoreState"))
	}
	p.cur.state = state
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	startRule, ok := p.rules[p.entrypoint]
	if !ok {
		p.addErr(errInvalidEntrypoint)
		return nil, p.errs.err()
	}

	p.read() // advance to first rune
	val, ok = p.parseRule(startRule)
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			maxFailExpectedMap := make(map[string]struct{}, len(p.maxFailExpected))
			for _, v := range p.maxFailExpected {
				maxFailExpectedMap[v] = struct{}{}
			}
			expected := make([]string, 0, len(maxFailExpectedMap))
			eof := false
			if _, ok := maxFailExpectedMap["!."]; ok {
				delete(maxFailExpectedMap, "!.")
				eof = true
			}
			for k := range maxFailExpectedMap {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}

		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.ExprCnt++
	if p.ExprCnt > p.maxExprCnt {
		panic(errMaxExprCnt)
	}

	var val interface{}
	var ok bool
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *recoveryExpr:
		val, ok = p.parseRecoveryExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *stateCodeExpr:
		val, ok = p.parseStateCodeExpr(expr)
	case *throwExpr:
		val, ok = p.parseThrowExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		state := p.cloneState()
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		p.restoreState(state)

		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	state := p.cloneState()

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn == utf8.RuneError && p.pt.w == 0 {
		// EOF - see utf8.DecodeRune
		p.failAt(false, p.pt.position, ".")
		return nil, false
	}
	start := p.pt
	p.read()
	p.failAt(true, start.position, ".")
	return p.sliceFrom(start), true
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt

	// can't match EOF
	if cur == utf8.RuneError && p.pt.w == 0 { // see utf8.DecodeRune
		p.failAt(false, start.position, chr.val)
		return nil, false
	}

	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
	return nil, false
}

func (p *parser) incChoiceAltCnt(ch *choiceExpr, altI int) {
	choiceIdent := fmt.Sprintf("%s %d:%d", p.rstack[len(p.rstack)-1].name, ch.pos.line, ch.pos.col)
	m := p.ChoiceAltCnt[choiceIdent]
	if m == nil {
		m = make(map[string]int)
		p.ChoiceAltCnt[choiceIdent] = m
	}
	// We increment altI by 1, so the keys do not start at 0
	alt := strconv.Itoa(altI + 1)
	if altI == choiceNoMatch {
		alt = p.choiceNoMatch
	}
	m[alt]++
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for altI, alt := range ch.alternatives {
		// dummy assignment to prevent compile error if optimized
		_ = altI

		state := p.cloneState()

		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			p.incChoiceAltCnt(ch, altI)
			return val, ok
		}
		p.restoreState(state)
	}
	p.incChoiceAltCnt(ch, choiceNoMatch)
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	state := p.cloneState()

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRecoveryExpr(recover *recoveryExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRecoveryExpr (" + strings.Join(recover.failureLabel, ",") + ")"))
	}

	p.pushRecovery(recover.failureLabel, recover.recoverExpr)
	val, ok := p.parseExpr(recover.expr)
	p.popRecovery()

	return val, ok
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	vals := make([]interface{}, 0, len(seq.exprs))

	pt := p.pt
	state := p.cloneState()
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restoreState(state)
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseStateCodeExpr(state *stateCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseStateCodeExpr"))
	}

	err := state.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, true
}

func (p *parser) parseThrowExpr(expr *throwExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseThrowExpr"))
	}

	for i := len(p.recoveryStack) - 1; i >= 0; i-- {
		if recoverExpr, ok := p.recoveryStack[i][expr.label]; ok {
			if val, ok := p.parseExpr(recoverExpr); ok {
				return val, ok
			}
		}
	}

	return nil, false
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}
