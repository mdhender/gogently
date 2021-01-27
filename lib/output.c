#define PRIVATE static

#include <stdio.h>

/*--------------------------------------------------------------------*/

#define OutBufSize 6000
#define FlushPos 5000

/*--------------------------------------------------------------------*/

PRIVATE long OutFileOpen = 0;
PRIVATE char OutBuf[OutBufSize];
PRIVATE char *OutBufPtr;
PRIVATE FILE *OutFile;

/*--------------------------------------------------------------------*/

int OpenOutput (char *Name) {
    CloseOutput ();
    OutFile = fopen(Name, "w");

    if (OutFile == NULL) {
        printf("cannot open %s\n", Name); exit(1);
    }
    OutBufPtr = &OutBuf[0];
    OutFileOpen = 1;
    return 1;
}

/*--------------------------------------------------------------------*/
void CloseOutput () {
    if (OutFileOpen) {
        fwrite(OutBuf, 1, OutBufPtr - &OutBuf[0], OutFile);
        fclose(OutFile);
    }
}

/*--------------------------------------------------------------------*/
void Put(char *Str) {
    while(*Str) {
        *OutBufPtr++ = *Str++;
    }
}

/*--------------------------------------------------------------------*/
void PutI (long N) {
    if (N < 0) {
        *OutBufPtr++ = '-';
        N = -N;
    }
    if (N < 10) {
        *OutBufPtr++ = N + '0';
    } else if (N < 100) {
        *OutBufPtr++ = N/10 + '0';
        *OutBufPtr++ = N%10 + '0';
    } else {
        char buf[50];
        register i;
        register n = N;

        i = 0;

        do {
            buf[++i] = (n % 10) + '0';
            n = n / 10;
        } while (n);

        while (i) {
            *OutBufPtr++ = buf[i--];
        }
    }
}

/*--------------------------------------------------------------------*/
void Nl () {
    if (OutBufPtr > &OutBuf[FlushPos]) {
        fwrite(OutBuf, 1, OutBufPtr - &OutBuf[0], OutFile);
        OutBufPtr = &OutBuf[0];
    }

    #ifdef EMIT_CR
        *OutBufPtr++ = '\r';
    #endif
    *OutBufPtr++ = '\n';
}

/*--------------------------------------------------------------------*/
