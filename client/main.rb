require './proto/hellogrpc_services_pb'

END_POINT = "localhost:50051"

stub = Hellogrpc::Greeter::Stub.new(END_POINT, :this_channel_is_insecure)

p stub.say_hello(Hellogrpc::HelloRequest.new(name: "world"))
