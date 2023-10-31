using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;
using DriverSDK.DmContext.Models;

namespace DriverSDK.DmContext
{
    public class Context
    {
        public const string DEFAULT_SUB_DEVICE_CONF = "sub_devices.yml";
        public const string DEFAULT_DEVICE_MODEL_CONF = "models.yml";
        public const string DEFAULT_ACCESS_TEMPLATE_CONF = "access_template.yml";

        public const string TYPE_REPORT_EVENT = "report";
        public const string KEY_DRIVER_NAME = "driverName";
        public const string KEY_DEVICE_NAME = "deviceName";

        public const string MESSAGE_DEVICE_EVENT = "deviceEvent";
        public const string MESSAGE_DEVICE_REPORT = "deviceReport";
        public const string MESSAGE_DEVICE_DESIRE = "deviceDesire";
        public const string MESSAGE_DEVICE_DELTA = "deviceDelta";
        public const string MESSAGE_DEVICE_PROPERTY_GET = "thing.property.get";
        public const string MESSAGE_DEVICE_EVENT_REPORT = "thing.event.post";
        public const string MESSAGE_DEVICE_LIFECYCLE_REPORT = "thing.lifecycle.post";

        private IDictionary<string, IDictionary<string, IList<DeviceProperty>>> _modelYamls;
        private IDictionary<string, IDictionary<string, AccessTemplate>> _accessTemplateYamls;
        private IDictionary<string, SubDeviceYaml> _subDeviceYamls;
        private IDictionary<string, IDictionary<string, DeviceInfo>> _deviceInfos;
        private IDictionary<string, string> _deviceDriverMap;

        private string _driverConfigPathBase;

        public Context()
        {
            _modelYamls = new Dictionary<string, IDictionary<string, IList<DeviceProperty>>>();
            _accessTemplateYamls = new Dictionary<string, IDictionary<string, AccessTemplate>>();
            _subDeviceYamls = new Dictionary<string, SubDeviceYaml>();
            _deviceDriverMap = new Dictionary<string, string>();
            _deviceInfos = new Dictionary<string, IDictionary<string, DeviceInfo>>();
            _driverConfigPathBase = "";
        }

        public void LoadYamlConfig(string path, string driverName)
        {
            _driverConfigPathBase = path;

            var deserializer = new DeserializerBuilder()
                .IgnoreUnmatchedProperties()
                .WithNamingConvention(CamelCaseNamingConvention.Instance)
                .Build();

            var yaml = File.ReadAllText(Path.Combine(path, DEFAULT_ACCESS_TEMPLATE_CONF));
            var aty = deserializer.Deserialize<IDictionary<string, AccessTemplate>>(yaml);
            foreach (var item in aty)
            {
                item.Value.Name = item.Key;
            }
            _accessTemplateYamls[driverName] = aty;

            yaml = File.ReadAllText(Path.Combine(path, DEFAULT_DEVICE_MODEL_CONF));
            var my = deserializer.Deserialize<IDictionary<string, IList<DeviceProperty>>>(yaml);
            _modelYamls[driverName] = my;

            yaml = File.ReadAllText(Path.Combine(path, DEFAULT_SUB_DEVICE_CONF));
            var sdy = deserializer.Deserialize<SubDeviceYaml>(yaml);
            _subDeviceYamls[driverName] = sdy;

            var info = new Dictionary<string, DeviceInfo>();
            if (sdy.Devices != null)
            {
                foreach (var item in sdy.Devices)
                {
                    if (item.Name != null)
                    {
                        info[item.Name] = item;
                        _deviceDriverMap[item.Name] = driverName;
                    }
                }
            }
            _deviceInfos[driverName] = info;
        }

        public IList<DeviceInfo> GetAllDevices(string driverName)
        {
            var res = new List<DeviceInfo>();
            foreach (var name in _deviceInfos[driverName].Keys)
            {
                res.Add(_deviceInfos[driverName][name]);
            }
            return res;
        }

        public DeviceInfo? GetDevice(string driverName, string deviceName)
        {
            if (_deviceInfos.ContainsKey(driverName) && _deviceInfos[driverName].ContainsKey(deviceName))
            {
                return _deviceInfos[driverName][deviceName];
            }
            return null;
        }

        public string GetDriverNameByDevice(string deviceName)
        {
            return _deviceDriverMap[deviceName];
        }

        public IList<DeviceProperty>? GetDeviceModel(string driverName, string deviceModelName)
        {
            if (_modelYamls.ContainsKey(driverName))
            {
                return _modelYamls[driverName][deviceModelName];
            }
            return null;
        }

        public IDictionary<string, IList<DeviceProperty>> GetAllDeviceModels(string driverName)
        {
            return _modelYamls[driverName];
        }

        public AccessTemplate? GetAccessTemplate(string driverName, string accessTemplateName)
        {
            if (_accessTemplateYamls.ContainsKey(driverName))
            {
                return _accessTemplateYamls[driverName][accessTemplateName];
            }
            return null;
        }

        public IDictionary<string, AccessTemplate> GetAllAccessTemplates(string driverName)
        {
            return _accessTemplateYamls[driverName];
        }

        public IDictionary<string, SubDeviceYaml> GetSubDeviceYamls()
        {
            return _subDeviceYamls;
        }
    }
}

