%{

#include "gen.h"
extern YYSTYPE yylval;

extern long yypos;
#define yysetpos() { yylval.attr[0] = yypos; yypos += yyleng; }


%}


%%


\  { yypos += 1; }
\t { yypos += 1; }
\r { yypos += 1; }
\n { yyPosToNextLine(); }

. { yysetpos(); yyerror("illegal token"); }

%%


#ifndef yywrap
yywrap() { return 1; }
#endif

