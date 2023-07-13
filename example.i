%module example
%{
#include <stdio.h>
#include "example.h"
%}

extern void print_string(char* string);