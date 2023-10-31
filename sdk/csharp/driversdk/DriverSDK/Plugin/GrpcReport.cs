using Grpc.Core;
using Proto;

namespace DriverSDK.Plugin
{
	public class GrpcReport: IReport
	{
        private readonly Channel _channel;
        private readonly Report.ReportClient _client;

        public GrpcReport(string host, int port)
        {
            _channel = new Channel(host, port, ChannelCredentials.Insecure);
            _client = new Report.ReportClient(_channel);
        }

        public string Post(string data)
        {
            var request = new RequestArgs { Request = data };
            var response = _client.Post(request);
            return response.Data;
        }

        public string State(string data)
        {
            var request = new RequestArgs { Request = data };
            var response = _client.State(request);
            return response.Data;
        }

        public void Dispose()
        {
            _channel.ShutdownAsync().Wait();
        }

        ~GrpcReport()
        {
            Dispose();
        }
    }
}

