import os

project_path = f"{os.environ['GOPATH']}/src/proto"
performance_demo_path = f'{project_path}/performance'
vendor_path = f'{project_path}/vendor'
proto_file_path = f'{performance_demo_path}/protos'

# generate gogoproto file
command = f'protoc -I {vendor_path} -I {proto_file_path} --gogo_out {performance_demo_path}/types tokenlist_gogo.proto'
print(command)
os.system(command)

# generate proto file
command = f'protoc -I {proto_file_path} --go_out {performance_demo_path}/types tokenlist.proto'
print(command)
os.system(command)
