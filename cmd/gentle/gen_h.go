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

type yyATTRIBUTES struct {
	attr [2]int
}

type YYSTYPE = yyATTRIBUTES

//extern YYSTYPE yylval;

// defines, errr, enums?
const (
	MODULE       = 257
	EXPORTTOKEN  = 258
	IMPORTTOKEN  = 259
	USE          = 260
	END          = 261
	VAR          = 262
	TYPETOKEN    = 263
	PROC         = 264
	COND         = 265
	NONTERMTOKEN = 266
	TOKENTOKEN   = 267
	CHOICE       = 268
	CLASSTOKEN   = 269
	SWEEP        = 270
	ROOT         = 271
	RULETOKEN    = 272
	TABLE        = 273
	KEY          = 274
	EQ           = 275
	TIMES        = 276
	DIV          = 277
	PLUS         = 278
	MINUS        = 279
	UNDERSCORE   = 280
	QUOTE        = 281
	DOLLAR       = 282
	BEGINDISJ    = 283
	ENDDISJ      = 284
	BEGINLOOP    = 285
	ENDLOOP      = 286
	DISJDELIM    = 287
	BEGINCOND    = 288
	ENDCOND      = 289
	RIGHTARROW   = 290
	LEFTARROW    = 291
	BECOMES      = 292
	COMESBE      = 293
	SMALLBECOMES = 294
	SMALLCOMESBE = 295
	COLON        = 296
	LEFTPAREN    = 297
	RIGHTPAREN   = 298
	LEFTBRACKET  = 299
	RIGHTBRACKET = 300
	COMMA        = 301
	DOT          = 302
	AMPERSAND    = 303
	EOFTOKEN     = 304
	FILESEP      = 305
	INTEGERCONST = 306
	STRINGCONST  = 307
	SMALLID      = 308
	LARGEID      = 309
)
