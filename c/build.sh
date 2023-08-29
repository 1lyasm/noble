#!/usr/bin/bash
clang -O2 -Wall -Werror -std=c11 \
    -fsanitize=address -Wno-declaration-after-statement \
    src/main/mrsequential.c src/mr/mr.c src/apps/wc.c -o mrsequential
