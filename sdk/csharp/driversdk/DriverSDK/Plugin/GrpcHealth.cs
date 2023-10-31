using Grpc.Core;
using Grpc.Health.V1;

namespace DriverSDK.Plugin
{
    public class GrpcHealth : Health.HealthBase
    {
        private Dictionary<String, HealthCheckResponse.Types.ServingStatus> _statusMap;

        public GrpcHealth()
        {
            _statusMap = new Dictionary<string, HealthCheckResponse.Types.ServingStatus> { };
            _statusMap.Add("plugin",HealthCheckResponse.Types.ServingStatus.Serving);
        }

        public override Task<HealthCheckResponse> Check(HealthCheckRequest request, ServerCallContext context)
        {
            if (_statusMap.TryGetValue(request.Service, out var status))
            {
                return Task.FromResult(new HealthCheckResponse { Status = status });
            }
            return Task.FromResult(new HealthCheckResponse { Status = HealthCheckResponse.Types.ServingStatus.Serving });
        }
    }
}

