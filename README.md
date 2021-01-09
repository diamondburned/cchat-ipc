# cchat-ipc

IPC for cchat backends to allow language-agnostic implementations of services.

## Regenerating

```sh
# Requires: bash, go, flatc
go generate
```

## Notes

- `flatc` Rust will segfault because of bitflags

```
#0  0x00000000005ab64f in flatbuffers::rust::RustGenerator::EscapeKeyword(std::__cxx11::basic_string<char, std::char_traits<char>, std::allocator<char> > const&) const ()
#1  0x00000000005abcea in flatbuffers::rust::RustGenerator::GetEnumValue[abi:cxx11](flatbuffers::EnumDef const&, flatbuffers::EnumVal const&) const ()
#2  0x00000000005b471e in flatbuffers::rust::RustGenerator::ForAllTableFields(flatbuffers::StructDef const&, std::function<void (flatbuffers::FieldDef const&)>, bool)::{lambda(flatbuffers::FieldDef const&)#1}::operator()(flatbuffers::FieldDef const&) const ()
```
