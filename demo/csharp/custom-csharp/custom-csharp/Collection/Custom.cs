using DriverSDK.DmContext;
using DriverSDK.DmContext.Models;
using Buildin.Csharp.Models;
using DriverSDK.Plugin;
using DriverSDK.Utils;

namespace Buildin.Csharp.Collecton
{
    public class Custom
    {
        public Context Ctx { get; set; }
        private DriverConfig _cfg;
        private IDictionary<string, Worker> _ws;
        private IDictionary<string, Device> _devs;
        private IReport _report;
        private IDictionary<string, Thread> _threads;

        public Custom(Context ctx, DriverConfig cfg, IReport report)
        {
            Ctx = ctx;
            _cfg = cfg;
            _report = report;

            IDictionary<string, DeviceInfo> infos = new Dictionary<string, DeviceInfo>();
            foreach (DeviceInfo info in Ctx.GetAllDevices(_cfg.DriverName))
            {
                infos.Add(info.Name, info);
            }

            _devs = new Dictionary<string, Device>();
            foreach (DeviceConfig item in _cfg.Devices)
            {
                if (infos.ContainsKey(item.Device))
                {
                    Device dev = new(item, infos[item.Device]);
                    _devs.Add(item.Device, dev);
                }
            }

            _ws = new Dictionary<string, Worker>();
            foreach (Job job in _cfg.Jobs)
            {
                if (_devs.ContainsKey(job.Device))
                {
                    _ws.Add(_devs[job.Device].GetInfo().Name,
                        new Worker(job, _devs[job.Device], _cfg.DriverName, _report));
                }
            }

            _threads = new Dictionary<string, Thread>();
        }

        public void Start()
        {
            foreach (string key in _ws.Keys)
            {
                Thread td = new(() =>
                {
                    _ws[key].Working();
                    lock (_threads)
                    {
                        _threads.Remove(key);
                    }
                });
                _threads.Add(key, td);
                td.Start();
            }
        }

        public void Restart()
        {
            lock (_threads)
            {
                foreach (Thread td in _threads.Values)
                {
                    td.Interrupt();
                }
                _threads.Clear();

                foreach (string key in _ws.Keys)
                {
                    Thread td = new(() =>
                    {
                        _ws[key].Working();
                        _threads.Remove(key);
                    });
                    _threads.Add(key, td);
                    td.Start();
                }
            }
        }

        public void Stop()
        {
            lock (_threads)
            {
                foreach (Thread td in _threads.Values)
                {
                    td.Interrupt();
                }
                _threads.Clear();
            }
        }

        public void Set(DeviceInfo? info, IDictionary<string, object> props)
        {
            if (info == null || !_ws.ContainsKey(info.Name))
            {
                return;
            }
            L.Debug("Custom set (name = " + info.Name + ")");

            Worker w = _ws[info.Name];
            foreach (KeyValuePair<string, object> e in props)
            {
                foreach (Property p in w.Job.Properties)
                {
                    if (e.Key.Equals(p.Name))
                    {
                        object val = e.Value;
                        if (val is double || val is float || val is int || val is long)
                        {
                            w.Device.Set(p.Index, Convert.ToDouble(val));
                        }
                        else
                        {
                            L.Debug("custom set " + e.Key + " error");
                        }
                    }
                }
            }
        }

        public void Event(DeviceInfo info, Event? e)
        {
            if (e != null && e.Type.Equals(Context.TYPE_REPORT_EVENT))
            {
                PropertyGet(info);
            }
            else
            {
                L.Debug("custom event " + info.Name);
            }
        }

        public void PropertyGet(DeviceInfo? info)
        {
            L.Debug("Custom get");
            if (info != null && _ws.ContainsKey(info.Name))
            {
                L.Debug("Custom get (name = " + info.Name + ")");
                Worker w = _ws[info.Name];
                try
                {
                    w.ReportProperty();
                }
                catch (Exception e)
                {
                    L.Error(e.ToString());
                }
            }
        }
    }
}

