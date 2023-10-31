namespace DriverSDK.DmContext.Models
{
    public class SubDeviceYaml
    {
        public IList<DeviceInfo> Devices { get; set; } = new List<DeviceInfo>();
        public string Driver { get; set; } = "";

        public SubDeviceYaml() { }

        public SubDeviceYaml(IList<DeviceInfo> devices, string driver)
        {
            this.Devices = devices;
            this.Driver = driver;
        }
    }

    public class DeviceInfo
    {
        public string Name { get; set; } = "";
        public string Version { get; set; } = "";
        public string DeviceModel { get; set; } = "";
        public string AccessTemplate { get; set; } = "";
        public AccessConfig AccessConfig { get; set; } = new AccessConfig();

        public DeviceInfo() { }

        public DeviceInfo(string name, string version, string deviceModel, string accessTemplate, AccessConfig accessConfig)
        {
            this.Name = name;
            this.Version = version;
            this.DeviceModel = deviceModel;
            this.AccessTemplate = accessTemplate;
            this.AccessConfig = accessConfig;
        }
    }

    public class DeviceProperty
    {
        public string Name { get; set; } = "";
        public string Id { get; set; } = "";
        public string Type { get; set; } = "";
        public string Mode { get; set; } = "";
        public string Unit { get; set; } = "";
        public PropertyVisitor? Visitor { get; set; } = new PropertyVisitor();
        public string Format { get; set; } = "";
        public EnumType? EnumType { get; set; }
        public ArrayType? ArrayType { get; set; }
        public Dictionary<string, ObjectType>? ObjectType { get; set; }
        public List<string>? ObjectRequired { get; set; }
        public object? Current { get; set; }
        public object? Expect { get; set; }

        public DeviceProperty(){}

        public DeviceProperty(string name, string id, string type, string mode,
            string unit, PropertyVisitor visitor, string format, EnumType enumType,
            ArrayType arrayType, Dictionary<string, ObjectType> objectType,
            List<string> objectRequired, object current, object expect)
        {
            this.Name = name;
            this.Id = id;
            this.Type = type;
            this.Mode = mode;
            this.Unit = unit;
            this.Visitor = visitor;
            this.Format = format;
            this.EnumType = enumType;
            this.ArrayType = arrayType;
            this.ObjectType = objectType;
            this.ObjectRequired = objectRequired;
            this.Current = current;
            this.Expect = expect;
        }
    }

    public class PropertyVisitor
    {
        public string Custom { get; set; } = "";

        public PropertyVisitor() { }

        public PropertyVisitor(string custom)
        {
            this.Custom = custom;
        }
    }

    public class EnumValue
    {
        public string Name { get; set; } = "";
        public string Value { get; set; } = "";
        public string DisplayName { get; set; } = "";

        public EnumValue() { }

        public EnumValue(string name, string value, string displayName)
        {
            this.Name = name;
            this.Value = value;
            this.DisplayName = displayName;
        }
    }

    public class EnumType
    {
        public string Type { get; set; } = "";
        public EnumValue[]? Values { get; set; }

        public EnumType() { }

        public EnumType(string type, EnumValue[] values)
        {
            this.Type = type;
            this.Values = values;
        }
    }

    public class ArrayType
    {
        public string Type { get; set; } = "";
        public int Min { get; set; } = 0;
        public int Max { get; set; } = 0;
        public string Format { get; set; } = "";
        
        public ArrayType() { }

        public ArrayType(string type, int min, int max, string format)
        {
            this.Type = type;
            this.Min = min;
            this.Max = max;
            this.Format = format;
        }
    }

    public class ObjectType
    {
        public string DisplayName { get; set; } = "";
        public string Type { get; set; } = "";
        public string Format { get; set; } = "";

        public ObjectType() { }

        public ObjectType(string displayName, string type, string format)
        {
            this.DisplayName = displayName;
            this.Type = type;
            this.Format = format;
        }
    }

    public class Event
    {
        public string Type { get; set; } = "";
        public object Payload { get; set; } = new object();

        public Event()
        {
        }

        public Event(string type, object payload)
        {
            Type = type;
            Payload = payload;
        }
    }
}

