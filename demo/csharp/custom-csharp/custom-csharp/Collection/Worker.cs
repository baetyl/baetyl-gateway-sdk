using DriverSDK.DmContext;
using DriverSDK.DmContext.Models;
using Buildin.Csharp.Models;
using DriverSDK.Plugin;
using DriverSDK.Utils;
using Newtonsoft.Json;

namespace Buildin.Csharp.Collecton
{
    public class Worker
    {
        public Job Job { set; get; }
        public Device Device { set; get; }
        public string DriverName { set; get; }
        public IReport Report { set; get; }

        public Worker(Job job, Device device, string driverName, IReport report)
        {
            Job = job;
            Device = device;
            DriverName = driverName;
            Report = report;
        }

        public void Working()
        {
            L.Debug("Worker interval " + Job.Interval + " ns");
            new Thread(() =>
            {
                while (true)
                {
                    try
                    {
                        ReportProperty();
                    }
                    catch (Exception e)
                    {
                        Console.WriteLine(e);
                    }
                    finally
                    {
                        try
                        {
                            Thread.Sleep(TimeSpan.FromMicroseconds(Job.Interval/1000));
                        }
                        catch (ThreadInterruptedException e)
                        {
                            Console.WriteLine(e);
                        }
                    }
                }
            }).Start();
        }

        public void ReportProperty()
        {
            L.Debug("Worker reportProperty");

            Dictionary<string, object> props = new();
            foreach (Property item in Job.Properties)
            {
                object val = Device.Get(item.Index);
                props.Add(item.Name, val);
            }

            Dictionary<string, string> meta = new()
            {
                { Context.KEY_DRIVER_NAME, DriverName },
                { Context.KEY_DEVICE_NAME, Device.GetInfo().Name }
            };
            Message msg = new Message(Context.MESSAGE_DEVICE_REPORT, meta, props);

            string jStr = JsonConvert.SerializeObject(msg);

            L.Debug("Worker reportProperty msg " + jStr);
            Report.Post(jStr);
        }
    }
}

