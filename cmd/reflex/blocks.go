/*
 * gogently - a universal toolkit for interpreters
 * Copyright (C) 2021  Michael D Henderson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

/* ( 1) %{ */
/* ( 2) YYSTYPE block */
/* ( 3) SETPOS block */
/* ( 4) LITBLOCK block */
/* ( 5) %} */
/* ( 6) LEXDEF block */
/* ( 7) %% */
/* ( 8) gen.lit */
/* ( 9) <token>.t for each <token> in gen.tkn */
/* (10) COMMENTS block */
/* (11) LAYOUT block */
/* (12) ILLEGAL block */
/* (13) %% */
/* (14) LEXFUNC block */
/* (15) YYWRAP block */

/* ( 1) %{ */
var leftpar []string = []string{
	"%{\n",
	"\n",
}

/* ( 2) YYSTYPE block */
var yystype []string = []string{
	"#include \"gen.h\"\n",
	"extern YYSTYPE yylval;\n",
	"\n",
}

/* ( 3) SETPOS block */
var setpos []string = []string{
	"extern long yypos;\n",
	"#define yysetpos() { yylval.attr[0] = yypos; yypos += yyleng; }\n",
	"\n",
}

/* ( 4) LITBLOCK block */
var litblock []string = []string{
	"\n",
}

/* ( 5) %} */
var rightpar []string = []string{
	"%}\n",
	"\n",
}

/* ( 6) LEXDEF block */
var lexdef []string = []string{
	"\n",
}

/* ( 7) %% */
var separator []string = []string{
	"%%\n",
	"\n",
}

/* ( 8) gen.lit */
/* ( 9) <token>.t for each <token> in gen.tkn */
/* (10) COMMENTS block */
var comments []string = []string{
	"\n",
}

/* (11) LAYOUT block */
var layout []string = []string{
	"\\  { yypos += 1; }\n",
	"\\t { yypos += 1; }\n",
	"\\r { yypos += 1; }\n",
	"\\n { yyPosToNextLine(); }\n",
	"\n",
}

/* (12) ILLEGAL block */
var illegal []string = []string{
	". { yysetpos(); yyerror(\"illegal token\"); }\n",
	"\n",
}

/* (14) LEXFUNC block */
var lexfunc []string = []string{
	"\n",
}

/* (15) YYWRAP block */
var yywrap []string = []string{
	"#ifndef yywrap\n",
	"yywrap() { return 1; }\n",
	"#endif\n",
	"\n",
}
