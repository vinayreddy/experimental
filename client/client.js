const { AddRequest, AddResponse } = require("../bazel-bin/proto/calculator_pb.js")
const { CalculatorClient } = require("../bazel-bin/proto/calculator_grpc_web_pb.js")
var client = new CalculatorClient('http://localhost:8080');

var request = new AddRequest()

request.setNum1(2)
request.setNum2(3)

client.add(request, {}, (err, response) => {
    console.log("Result of Add : ",response.getResult())
})

export function GRPCAdd(a, b) {
  var request = new AddRequest()
  request.setNum1(a)
  request.setNum2(b)
  client.add(request, {}, (err, response) => {
    console.log("Result of gRPC add: ", response.getResult())
  })
}

GRPCAdd(10, 15);
