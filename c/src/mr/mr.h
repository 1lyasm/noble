#ifndef MR_H
#define MR_H

typedef struct {
    char *k;
    char *v;
} Kv;

Kv *allocKv(int kLen, int vLen);
Kv *copyKv(Kv *kv);
void printKv(Kv *kv);
void freeKv(Kv *kv);

typedef struct {
    Kv **data;
    int len;
    int size;
} Kva;

Kva *allocKva(void);
void expandKva(Kva *kva);
void addKv(Kva *kva, Kv *kv);
void printKva(Kva *kva);
void freeKva(Kva *kva);

void map(char *fName, char *str, Kva *kva);
char *reduce(char *key, char **vals, int valsLen);

#endif
