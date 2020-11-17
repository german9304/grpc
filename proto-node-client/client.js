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

const META = {
        java: {
            path: '../java-grpc/src/main/proto/helloworld.proto',
            port: 50051,
            service: 'HelloWorld',
        },
        go: {
            path: '../go-grpc/api/helloworld.proto',
            port: 8080,
            service: 'GreetService',
        }
}

async function main() {
    const arg = process.argv[2];
    const packageDefinition = await getProtoFile(META[arg].path);
    const { api } = grpc.loadPackageDefinition(packageDefinition);
    const client = new api[META[arg].service](`localhost:${META[arg].port}`, grpc.credentials.createInsecure());
    client.Hello({ message: 'hello world'}, (err, response) => {
        console.log(response);
    });
}

main();
