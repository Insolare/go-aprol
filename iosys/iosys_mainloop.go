package iosys

/*
#cgo CFLAGS: -I /opt/aprol/include

#include <PccIosys.h>

void goaprol_ios_mainloop() {
	fd_set rMask, wMask;
	struct timeval timeout = {10, 0};
	int maxHandle = -1;

	FD_ZERO(&rMask);
	FD_ZERO(&wMask);

	Ios_get_fds(&rMask, &wMask, &timeout, &maxHandle);

	select(1 + maxHandle, &rMask, &wMask, NULL, &timeout);
	Ios_run(&rMask, &wMask);
}
*/
import "C"
