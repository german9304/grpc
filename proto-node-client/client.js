const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

async function getProtoFile(protoFile) {
    try {
        const proto = await protoLoader.load(protoFile);
        return proto;
    } catch(err) {
        return err
    }
}

async function main() {
    const packageDefinition = await getProtoFile('../java-grpc/src/main/proto/helloworld.proto');
    const { api } = grpc.loadPackageDefinition(packageDefinition);
    const client = new api.HelloWorld('localhost:50051', grpc.credentials.createInsecure());
    client.Hello({ message: 'hello world'}, (err, response) => {
        console.log(response);
    });
}

main()