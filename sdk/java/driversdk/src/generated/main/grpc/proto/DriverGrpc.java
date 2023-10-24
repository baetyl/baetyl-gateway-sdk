package proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.56.0)",
    comments = "Source: driver.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class DriverGrpc {

  private DriverGrpc() {}

  public static final String SERVICE_NAME = "proto.Driver";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getGetDriverInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDriverInfo",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getGetDriverInfoMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getGetDriverInfoMethod;
    if ((getGetDriverInfoMethod = DriverGrpc.getGetDriverInfoMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getGetDriverInfoMethod = DriverGrpc.getGetDriverInfoMethod) == null) {
          DriverGrpc.getGetDriverInfoMethod = getGetDriverInfoMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDriverInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("GetDriverInfo"))
              .build();
        }
      }
    }
    return getGetDriverInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getSetConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetConfig",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getSetConfigMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getSetConfigMethod;
    if ((getSetConfigMethod = DriverGrpc.getSetConfigMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getSetConfigMethod = DriverGrpc.getSetConfigMethod) == null) {
          DriverGrpc.getSetConfigMethod = getSetConfigMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("SetConfig"))
              .build();
        }
      }
    }
    return getSetConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getSetupMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Setup",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getSetupMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getSetupMethod;
    if ((getSetupMethod = DriverGrpc.getSetupMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getSetupMethod = DriverGrpc.getSetupMethod) == null) {
          DriverGrpc.getSetupMethod = getSetupMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Setup"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("Setup"))
              .build();
        }
      }
    }
    return getSetupMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getStartMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Start",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getStartMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getStartMethod;
    if ((getStartMethod = DriverGrpc.getStartMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getStartMethod = DriverGrpc.getStartMethod) == null) {
          DriverGrpc.getStartMethod = getStartMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Start"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("Start"))
              .build();
        }
      }
    }
    return getStartMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getRestartMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Restart",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getRestartMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getRestartMethod;
    if ((getRestartMethod = DriverGrpc.getRestartMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getRestartMethod = DriverGrpc.getRestartMethod) == null) {
          DriverGrpc.getRestartMethod = getRestartMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Restart"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("Restart"))
              .build();
        }
      }
    }
    return getRestartMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getStopMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Stop",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getStopMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getStopMethod;
    if ((getStopMethod = DriverGrpc.getStopMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getStopMethod = DriverGrpc.getStopMethod) == null) {
          DriverGrpc.getStopMethod = getStopMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Stop"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("Stop"))
              .build();
        }
      }
    }
    return getStopMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getGetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Get",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getGetMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getGetMethod;
    if ((getGetMethod = DriverGrpc.getGetMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getGetMethod = DriverGrpc.getGetMethod) == null) {
          DriverGrpc.getGetMethod = getGetMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Get"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("Get"))
              .build();
        }
      }
    }
    return getGetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getSetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Set",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getSetMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getSetMethod;
    if ((getSetMethod = DriverGrpc.getSetMethod) == null) {
      synchronized (DriverGrpc.class) {
        if ((getSetMethod = DriverGrpc.getSetMethod) == null) {
          DriverGrpc.getSetMethod = getSetMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Set"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new DriverMethodDescriptorSupplier("Set"))
              .build();
        }
      }
    }
    return getSetMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DriverStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DriverStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DriverStub>() {
        @java.lang.Override
        public DriverStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DriverStub(channel, callOptions);
        }
      };
    return DriverStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DriverBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DriverBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DriverBlockingStub>() {
        @java.lang.Override
        public DriverBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DriverBlockingStub(channel, callOptions);
        }
      };
    return DriverBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DriverFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DriverFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DriverFutureStub>() {
        @java.lang.Override
        public DriverFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DriverFutureStub(channel, callOptions);
        }
      };
    return DriverFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * 宿主（client） --&gt; 驱动（server）
     * </pre>
     */
    default void getDriverInfo(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDriverInfoMethod(), responseObserver);
    }

    /**
     */
    default void setConfig(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetConfigMethod(), responseObserver);
    }

    /**
     */
    default void setup(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetupMethod(), responseObserver);
    }

    /**
     */
    default void start(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStartMethod(), responseObserver);
    }

    /**
     */
    default void restart(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRestartMethod(), responseObserver);
    }

    /**
     */
    default void stop(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStopMethod(), responseObserver);
    }

    /**
     */
    default void get(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetMethod(), responseObserver);
    }

    /**
     */
    default void set(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Driver.
   */
  public static abstract class DriverImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return DriverGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Driver.
   */
  public static final class DriverStub
      extends io.grpc.stub.AbstractAsyncStub<DriverStub> {
    private DriverStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DriverStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DriverStub(channel, callOptions);
    }

    /**
     * <pre>
     * 宿主（client） --&gt; 驱动（server）
     * </pre>
     */
    public void getDriverInfo(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDriverInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setConfig(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setup(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetupMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void start(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStartMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void restart(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRestartMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void stop(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStopMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void get(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void set(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Driver.
   */
  public static final class DriverBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<DriverBlockingStub> {
    private DriverBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DriverBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DriverBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 宿主（client） --&gt; 驱动（server）
     * </pre>
     */
    public proto.DriverOuterClass.ResponseResult getDriverInfo(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDriverInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult setConfig(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult setup(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetupMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult start(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStartMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult restart(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRestartMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult stop(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStopMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult get(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult set(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Driver.
   */
  public static final class DriverFutureStub
      extends io.grpc.stub.AbstractFutureStub<DriverFutureStub> {
    private DriverFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DriverFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DriverFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 宿主（client） --&gt; 驱动（server）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> getDriverInfo(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDriverInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> setConfig(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> setup(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetupMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> start(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStartMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> restart(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRestartMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> stop(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStopMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> get(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> set(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_DRIVER_INFO = 0;
  private static final int METHODID_SET_CONFIG = 1;
  private static final int METHODID_SETUP = 2;
  private static final int METHODID_START = 3;
  private static final int METHODID_RESTART = 4;
  private static final int METHODID_STOP = 5;
  private static final int METHODID_GET = 6;
  private static final int METHODID_SET = 7;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_DRIVER_INFO:
          serviceImpl.getDriverInfo((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_SET_CONFIG:
          serviceImpl.setConfig((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_SETUP:
          serviceImpl.setup((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_START:
          serviceImpl.start((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_RESTART:
          serviceImpl.restart((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_STOP:
          serviceImpl.stop((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_GET:
          serviceImpl.get((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_SET:
          serviceImpl.set((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getGetDriverInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_GET_DRIVER_INFO)))
        .addMethod(
          getSetConfigMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_SET_CONFIG)))
        .addMethod(
          getSetupMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_SETUP)))
        .addMethod(
          getStartMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_START)))
        .addMethod(
          getRestartMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_RESTART)))
        .addMethod(
          getStopMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_STOP)))
        .addMethod(
          getGetMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_GET)))
        .addMethod(
          getSetMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_SET)))
        .build();
  }

  private static abstract class DriverBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DriverBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return proto.DriverOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Driver");
    }
  }

  private static final class DriverFileDescriptorSupplier
      extends DriverBaseDescriptorSupplier {
    DriverFileDescriptorSupplier() {}
  }

  private static final class DriverMethodDescriptorSupplier
      extends DriverBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DriverMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (DriverGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DriverFileDescriptorSupplier())
              .addMethod(getGetDriverInfoMethod())
              .addMethod(getSetConfigMethod())
              .addMethod(getSetupMethod())
              .addMethod(getStartMethod())
              .addMethod(getRestartMethod())
              .addMethod(getStopMethod())
              .addMethod(getGetMethod())
              .addMethod(getSetMethod())
              .build();
        }
      }
    }
    return result;
  }
}
