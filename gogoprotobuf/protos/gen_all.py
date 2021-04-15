import os
import sys

project_path = f"{os.environ['GOPATH']}/src/proto"
gogoprotobuf_demo_path = f'{project_path}/gogoprotobuf'
vendor_path = f'{project_path}/vendor'
proto_file_path = f'{gogoprotobuf_demo_path}/protos'

file_names = ['enum_prefix', 'getters', 'face', 'nullable', 'customname', 'stringer', 'gostring']

index = int(sys.argv[1])

# generate gogoproto file
command = f'protoc -I {vendor_path} -I {proto_file_path} --gogo_out {gogoprotobuf_demo_path}/types/gogoprotobuf {index}_{file_names[index]}_gogo.proto'
print(command)
os.system(command)

# generate proto file
command = f'protoc -I {proto_file_path} --go_out {gogoprotobuf_demo_path}/types/protobuf {index}_{file_names[index]}.proto'
print(command)
os.system(command)
