package com.grpc.app;

import java.io.IOException;
import com.grpc.api.HelloWorldGrpc;
import com.grpc.api.Helloworld.HelloRequest;
import com.grpc.api.Helloworld.HelloResponse;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;
import java.util.concurrent.TimeUnit;

class HelloWorldImpl extends HelloWorldGrpc.HelloWorldImplBase {

    @Override
    public void hello(HelloRequest request, StreamObserver<HelloResponse> responseObserver) {
        String message = request.getMessage();
        HelloResponse.Builder builder = HelloResponse.newBuilder();
    
        HelloResponse response = builder.setMessage(message)
        .build();

        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}

class HelloWorldService {

    public Server start() throws IOException {
        Server server = ServerBuilder
        .forPort(50051)
        .addService(new HelloWorldImpl())
        .build();

        System.out.println("Starting server...");
        server.start();
        System.out.println("Server started!");

        try {
            if (server != null) {
                server.awaitTermination();
            }
        }catch(InterruptedException e) {

        }
        return server;
    }
}


public class App {

    public static void main(String[] args) {
        try {
            HelloWorldService s = new HelloWorldService();
            s.start();
        } catch(IOException e) {

        }
            
    }
}

