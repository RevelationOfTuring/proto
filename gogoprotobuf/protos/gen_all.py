import os
import sys

project_path = f"{os.environ['GOPATH']}/src/proto"
gogoprotobuf_demo_path = f'{project_path}/gogoprotobuf'
vendor_path = f'{project_path}/vendor'
proto_file_path = f'{gogoprotobuf_demo_path}/protos'

proto_file = ''
gogoproto_file = ''
if sys.argv[1] == '1':
    proto_file = '1_enum_prefix.proto'
    gogoproto_file = '1_enum_prefix_gogo.proto'
elif sys.argv[1] == 2:
    pass


# generate gogoproto file
command = f'protoc -I {vendor_path} -I {proto_file_path} --gogo_out {gogoprotobuf_demo_path}/types/gogoprotobuf {gogoproto_file}'
os.system(command)
print(command)

# generate proto file
command = f'protoc -I {proto_file_path} --go_out {gogoprotobuf_demo_path}/types/protobuf {proto_file}'
os.system(command)
print(command)
