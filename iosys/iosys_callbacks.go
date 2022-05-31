package iosys

/*
#include <PccIosys.h>

extern void iosysChangedProxy(void*, void*);
extern void iosysChangeRequestProxy(void*, void*);
extern void iosysIdlerProxy(void*, void*);

void iosys_changedCgo(IosVar *v, void *user_data) {
	iosysChangedProxy(v, user_data);
}

void iosys_changerequestCgo(IosVar *v, void *user_data, IosVar *request) {
	iosysChangeRequestProxy(v, user_data);
}

void iosys_idlerCgo(IosVar *v, void *user_data, int idle) {
	iosysIdlerProxy(v, user_data);
}
*/
import "C"
