using Grpc.Core;
using Grpc.Health.V1;
using Proto;

namespace DriverSDK.Plugin
{
	public class Serve
	{
        private readonly GrpcDriver _grpcDriver;
        private readonly GrpcHealth _grpcHealth;

        public Serve(IDriver driver)
        {
            _grpcDriver = new GrpcDriver(driver);
            _grpcHealth = new GrpcHealth();
        }

        public void Start()
        {
            try
            {
                Server server = new()
                {
                    Services = { Health.BindService(_grpcHealth), Driver.BindService(_grpcDriver) },
                    Ports = { new ServerPort("localhost", 0, ServerCredentials.Insecure) }
                };
                // 启动服务器
                server.Start();

                // go-plugin 要求必须调用系统打印函数输出下面格式的数据
                int port = server.Ports.FirstOrDefault()?.BoundPort ?? 0;
                Console.WriteLine($"1|1|tcp|127.0.0.1:{port}|grpc");

                // 等待服务器终止
                server.ShutdownTask.Wait();
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
        }
    }
}

