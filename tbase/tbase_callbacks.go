package tbase

/*
#include <Vset.h>
extern void tbEnumerateProxy(void*, void*);

int tb_enumerateCgo(Vset *v, void *user_data) {
  tbEnumerateProxy(v, user_data);

  return 0;
}
*/
import "C"
