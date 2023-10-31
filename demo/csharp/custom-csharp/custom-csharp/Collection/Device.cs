using DriverSDK.DmContext.Models;
using Buildin.Csharp.Models;

namespace Buildin.Csharp.Collecton
{
    public class Device
    {
        private DeviceConfig cfg;
        private DeviceInfo info;
        private Simulator cli;

        public Device(DeviceConfig cfg, DeviceInfo info)
        {
            this.cfg = cfg;
            this.info = info;
            this.cli = new Simulator(cfg.Device);
        }

        public object Get(int index)
        {
            try
            {
                return this.cli.Get(index);
            }
            catch (Exception)
            {
                throw;
            }
        }

        public void Set(int index, double val)
        {
            this.cli.Set(index, val);
        }

        public DeviceInfo GetInfo()
        {
            return info;
        }

        public void SetInfo(DeviceInfo info)
        {
            this.info = info;
        }
    }
}

