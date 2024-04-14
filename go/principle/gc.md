# Garbage collector (GC)

## 原理
实现原理
- A Guide to the Go Garbage Collector https://go.dev/doc/gc-guide
- Getting to Go: The Journey of Go's Garbage Collector https://go.dev/blog/ismmkeynote

src/runtime 源码
- Garbage collector (GC). https://go.dev/src/runtime/mgc.go
- Garbage collector: stack objects and stack tracing https://go.dev/src/runtime/mgcstack.go
- Garbage collector: marking and scanning https://go.dev/src/runtime/mgcmark.go
- Garbage collector: sweeping https://go.dev/src/runtime/mgcsweep.go
- Garbage collector: write barriers. https://go.dev/src/runtime/mbarrier.go
- Garbage collector: type and heap bitmaps. https://go.dev/src/runtime/mbitmap_allocheaders.go
- Garbage collector: type and heap bitmaps. https://go.dev/src/runtime/mbitmap_noallocheaders.go
- Scavenging free pages. https://go.dev/src/runtime/mgcscavenge.go
- Garbage collector work pool abstraction. https://go.dev/src/runtime/mgcwork.go
- https://go.dev/src/runtime/mgclimit.go
- https://go.dev/src/runtime/mgcpacer.go

