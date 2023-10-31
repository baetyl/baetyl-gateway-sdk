using Grpc.Core;
using Proto;

namespace DriverSDK.Plugin
{
	public class GrpcDriver: Driver.DriverBase
	{
        private IDriver _driver; 
        private IReport? _report; 

        public GrpcDriver(IDriver driver)
        {
            _driver = driver; 
        }

        public override Task<ResponseResult> GetDriverInfo(RequestArgs request, ServerCallContext context)
        {
            return Task.FromResult(new ResponseResult { Data = _driver.GetDriverInfo(request.Request) });
        }

        public override Task<ResponseResult> SetConfig(RequestArgs request, ServerCallContext context)
        {
            return Task.FromResult(new ResponseResult { Data = _driver.SetConfig(request.Request) });
        }

        public override Task<ResponseResult> Setup(RequestArgs request, ServerCallContext context)
        {
            _report = new GrpcReport("0.0.0.0", (int)request.Brokerid);
            return Task.FromResult(new ResponseResult { Data = _driver.Setup(request.Request, _report) });
        }

        public override Task<ResponseResult> Start(RequestArgs request, ServerCallContext context)
        {
            return Task.FromResult(new ResponseResult { Data = _driver.Start(request.Request) });
        }

        public override Task<ResponseResult> Restart(RequestArgs request, ServerCallContext context)
        {
            return Task.FromResult(new ResponseResult { Data = _driver.Restart(request.Request) });
        }

        public override Task<ResponseResult> Stop(RequestArgs request, ServerCallContext context)
        {
            return Task.FromResult(new ResponseResult { Data = _driver.Stop(request.Request) });
        }

        public override Task<ResponseResult> Get(RequestArgs request, ServerCallContext context)
        {
            return Task.FromResult(new ResponseResult { Data = _driver.Get(request.Request) });
        }

        public override Task<ResponseResult> Set(RequestArgs request, ServerCallContext context)
        {
            return Task.FromResult(new ResponseResult { Data = _driver.Set(request.Request) });
        }
    }
}

