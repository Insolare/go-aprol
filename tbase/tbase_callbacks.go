package tbase

/*
#include <stdlib.h>
#include <Vset.h>
extern void tbEnumerateProxy(void*, void*);
extern void tbReferProxy(void*, void*, void *);

int tb_enumerateCgo(Vset *v, void *user_data) {
  tbEnumerateProxy(v, user_data);

  return 0;
}

int tb_referCgo(Vset *s, char *field, void *closure) {
  tbReferProxy(s, field, closure);

  free(field);

  return 0;
}
*/
import "C"
