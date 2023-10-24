package proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.56.0)",
    comments = "Source: driver.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class ReportGrpc {

  private ReportGrpc() {}

  public static final String SERVICE_NAME = "proto.Report";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getPostMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Post",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getPostMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getPostMethod;
    if ((getPostMethod = ReportGrpc.getPostMethod) == null) {
      synchronized (ReportGrpc.class) {
        if ((getPostMethod = ReportGrpc.getPostMethod) == null) {
          ReportGrpc.getPostMethod = getPostMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Post"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new ReportMethodDescriptorSupplier("Post"))
              .build();
        }
      }
    }
    return getPostMethod;
  }

  private static volatile io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getStateMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "State",
      requestType = proto.DriverOuterClass.RequestArgs.class,
      responseType = proto.DriverOuterClass.ResponseResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs,
      proto.DriverOuterClass.ResponseResult> getStateMethod() {
    io.grpc.MethodDescriptor<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult> getStateMethod;
    if ((getStateMethod = ReportGrpc.getStateMethod) == null) {
      synchronized (ReportGrpc.class) {
        if ((getStateMethod = ReportGrpc.getStateMethod) == null) {
          ReportGrpc.getStateMethod = getStateMethod =
              io.grpc.MethodDescriptor.<proto.DriverOuterClass.RequestArgs, proto.DriverOuterClass.ResponseResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "State"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.RequestArgs.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.DriverOuterClass.ResponseResult.getDefaultInstance()))
              .setSchemaDescriptor(new ReportMethodDescriptorSupplier("State"))
              .build();
        }
      }
    }
    return getStateMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ReportStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ReportStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ReportStub>() {
        @java.lang.Override
        public ReportStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ReportStub(channel, callOptions);
        }
      };
    return ReportStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ReportBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ReportBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ReportBlockingStub>() {
        @java.lang.Override
        public ReportBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ReportBlockingStub(channel, callOptions);
        }
      };
    return ReportBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ReportFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ReportFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ReportFutureStub>() {
        @java.lang.Override
        public ReportFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ReportFutureStub(channel, callOptions);
        }
      };
    return ReportFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * 驱动（client） --&gt; 宿主（server）
     * </pre>
     */
    default void post(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPostMethod(), responseObserver);
    }

    /**
     */
    default void state(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStateMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Report.
   */
  public static abstract class ReportImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return ReportGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Report.
   */
  public static final class ReportStub
      extends io.grpc.stub.AbstractAsyncStub<ReportStub> {
    private ReportStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReportStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ReportStub(channel, callOptions);
    }

    /**
     * <pre>
     * 驱动（client） --&gt; 宿主（server）
     * </pre>
     */
    public void post(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPostMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void state(proto.DriverOuterClass.RequestArgs request,
        io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStateMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Report.
   */
  public static final class ReportBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<ReportBlockingStub> {
    private ReportBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReportBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ReportBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 驱动（client） --&gt; 宿主（server）
     * </pre>
     */
    public proto.DriverOuterClass.ResponseResult post(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPostMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.DriverOuterClass.ResponseResult state(proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStateMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Report.
   */
  public static final class ReportFutureStub
      extends io.grpc.stub.AbstractFutureStub<ReportFutureStub> {
    private ReportFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ReportFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ReportFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 驱动（client） --&gt; 宿主（server）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> post(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPostMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.DriverOuterClass.ResponseResult> state(
        proto.DriverOuterClass.RequestArgs request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStateMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_POST = 0;
  private static final int METHODID_STATE = 1;

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
        case METHODID_POST:
          serviceImpl.post((proto.DriverOuterClass.RequestArgs) request,
              (io.grpc.stub.StreamObserver<proto.DriverOuterClass.ResponseResult>) responseObserver);
          break;
        case METHODID_STATE:
          serviceImpl.state((proto.DriverOuterClass.RequestArgs) request,
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
          getPostMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_POST)))
        .addMethod(
          getStateMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              proto.DriverOuterClass.RequestArgs,
              proto.DriverOuterClass.ResponseResult>(
                service, METHODID_STATE)))
        .build();
  }

  private static abstract class ReportBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ReportBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return proto.DriverOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Report");
    }
  }

  private static final class ReportFileDescriptorSupplier
      extends ReportBaseDescriptorSupplier {
    ReportFileDescriptorSupplier() {}
  }

  private static final class ReportMethodDescriptorSupplier
      extends ReportBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ReportMethodDescriptorSupplier(String methodName) {
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
      synchronized (ReportGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ReportFileDescriptorSupplier())
              .addMethod(getPostMethod())
              .addMethod(getStateMethod())
              .build();
        }
      }
    }
    return result;
  }
}
