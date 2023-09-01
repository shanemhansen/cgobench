package cgobench

/*
#include <time.h>
int trivial_add(int a, int b) {
  return a+b;
}
*/
import "C"
func Call() {
	// I don't do much. Yet.
}
// wow this is easy
// import "C"
func CgoCall() {
	C.trivial_add(1,2)
}
