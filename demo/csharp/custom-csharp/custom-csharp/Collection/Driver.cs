using DriverSDK.DmContext;
using DriverSDK.DmContext.Models;
using Buildin.Csharp.Models;
using DriverSDK.Plugin;
using DriverSDK.Utils;
using Newtonsoft.Json.Linq;

namespace Buildin.Csharp.Collecton
{
    public class Driver : IDriver
    {
        private string _driverName = "";
        private string _configPath = "";
        private DriverConfig _config = new();
        private IReport? _report;
        private Custom? _custom;

        public string GetDriverInfo(string data)
        {
            JObject info = new()
            {
                ["name"] = _driverName,
                ["configPath"] = _configPath,
                ["config"] = JToken.FromObject(_config)
            };

            return info.ToString();
        }

        public string SetConfig(string data)
        {
            L.Debug("Driver setConfig " + data);
            _configPath = data;
            return string.Format("plugin {0} setConfig success", _driverName);
        }

        public string Setup(string driver, IReport report)
        {
            L.Debug("Driver setup " + driver + " IReport " + report);
            _driverName = driver;
            _report = report;
            return string.Format("plugin {0} setup success", driver);
        }

        public string Start(string data)
        {
            L.Debug("Driver start " + data);
            try
            {
                Context ctx = new();
                ctx.LoadYamlConfig(_configPath, _driverName);
                _config = LoadConfig(ctx, _driverName);
                if (_report != null)
                {
                    _custom = new Custom(ctx, _config, _report);
                    _custom.Start();
                    return string.Format("plugin {0} start success", _driverName);
                }
                return string.Format("plugin {0} start failed", _driverName);
            }
            catch (Exception e)
            {
                L.Error("Driver start failed, " + e.Message);
                return string.Format("plugin {0} start failed, {1}", _driverName, e.Message);
            }
        }

        public string Restart(string data)
        {
            L.Debug("Driver restart " + data);
            if (_custom != null)
            {
                _custom.Restart();
                return string.Format("plugin {0} restart success", _driverName);
            }
            return string.Format("plugin {0} restart failed", _driverName);
        }

        public string Stop(string data)
        {
            L.Debug("Driver stop " + data);
            if (_custom != null)
            {
                _custom.Stop();
                return string.Format("plugin {0} stop success", _driverName);
            }
            return string.Format("plugin {0} stop failed", _driverName);
        }

        public string Get(string data)
        {
            L.Debug("Driver get " + data);

            if (_custom == null)
            {
                return string.Format("plugin {0} get failed, custom error", _driverName);
            }

            Message? msg = JObject.Parse(data).ToObject<Message>();
            if (msg == null)
            {
                return string.Format("plugin {0} get failed, msg parse error {1}", _driverName, data);
            }
            string driverName = msg.Meta[Context.KEY_DRIVER_NAME];
            string deviceName = msg.Meta[Context.KEY_DEVICE_NAME];

            if (string.IsNullOrEmpty(driverName) || string.IsNullOrEmpty(deviceName))
            {
                return string.Format("plugin {0} get failed", _driverName);
            }

            DeviceInfo? info = _custom.Ctx.GetDevice(driverName, deviceName);
            if (info == null)
            {
                return string.Format("plugin {0} get failed, device info error", _driverName);
            }

            switch (msg.Kind)
            {
                case Context.MESSAGE_DEVICE_EVENT:
                    var dt = Convert.ToString(msg.Content);
                    if (dt == null)
                    {
                        return string.Format("plugin {0} get failed, msg content error {1}", _driverName, msg.Content);
                    }
                    Event? ev = JObject.Parse(dt).ToObject<Event>();
                    _custom.Event(info, ev);
                    break;
                case Context.MESSAGE_DEVICE_PROPERTY_GET:
                    _custom.PropertyGet(info);
                    break;
                default:
                    L.Debug("driver get unsupported message type");
                    break;
            }
            return string.Format("plugin {0} get success", _driverName);
        }

        public string Set(string data)
        {
            L.Debug("Driver set " + data);
            if (_custom == null)
            {
                return string.Format("plugin {0} set failed, custom error", _driverName);
            }

            Message? msg = JObject.Parse(data).ToObject<Message>();
            if (msg == null)
            {
                return string.Format("plugin {0} set failed, msg parse error {1}", _driverName, data);
            }

            string driverName = msg.Meta[Context.KEY_DRIVER_NAME];
            string deviceName = msg.Meta[Context.KEY_DEVICE_NAME];

            if (string.IsNullOrEmpty(driverName) || string.IsNullOrEmpty(deviceName))
            {
                return string.Format("plugin {0} get failed", _driverName);
            }

            DeviceInfo? info = _custom.Ctx.GetDevice(driverName, deviceName);

            var dt = Convert.ToString(msg.Content);
            if (dt == null)
            {
                return string.Format("plugin {0} set failed, msg content error {1}", _driverName, msg.Content);
            }

            Dictionary<string, object>? origData = JObject.Parse(dt).ToObject<Dictionary<string, object>>();
            if (origData == null)
            {
                return string.Format("plugin {0} set failed, invalid msg content", _driverName);
            }
            Dictionary<string, object> props = new();
            foreach (KeyValuePair<string, object> entry in origData)
            {
                props[entry.Key] = entry.Value;
            }
            _custom.Set(info, props);
            return string.Format("plugin {0} set success", _driverName);
        }

        private DriverConfig LoadConfig(Context dm, string driverName)
        {
            List<DeviceConfig> devices = new();
            List<Job> jobs = new();

            foreach (DeviceInfo info in dm.GetAllDevices(driverName))
            {
                AccessConfig accessConfig = info.AccessConfig;
                if (accessConfig == null)
                {
                    continue;
                }
                DeviceConfig device = new()
                {
                    Device = info.Name
                };
                devices.Add(device);

                List<Property> jobProps = new();

                AccessTemplate? tpl = dm.GetAccessTemplate(driverName, info.AccessTemplate);
                if (tpl != null && tpl.Properties != null)
                {
                    foreach (DeviceProperty prop in tpl.Properties)
                    {
                        if (prop.Visitor != null)
                        {
                            string visitor = prop.Visitor.Custom;
                            if (!string.IsNullOrEmpty(visitor))
                            {
                                Property? jobProp = JObject.Parse(visitor).ToObject<Property>();
                                if (jobProp != null)
                                {
                                    jobProps.Add(jobProp);
                                }
                            }
                        }
                    }
                }

                Job? job = JObject.Parse(accessConfig.Custom).ToObject<Job>();
                if (job != null)
                {
                    job.Device = info.Name;
                    job.Properties = jobProps;

                    jobs.Add(job);
                }
            }
            return new DriverConfig(driverName, devices, jobs);
        }
    }
}

