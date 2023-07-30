#if defined _rakstar_included
	#endinput
#endif

#define _rakstar_included

// Should return 1 if this works out all fine.
native GoInt32: rakstar_CallEvent(const function[32], const format[] = "", {Float,_}:...);