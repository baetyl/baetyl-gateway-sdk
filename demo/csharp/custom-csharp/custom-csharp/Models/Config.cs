namespace Buildin.Csharp.Models
{
    public class DriverConfig
    {
        public string DriverName { get; set; } = "";
        public List<DeviceConfig> Devices { get; set; } = new List<DeviceConfig>();
        public List<Job> Jobs { get; set; } = new List<Job>();

        public DriverConfig() { }

        public DriverConfig(string driverName, List<DeviceConfig> devices, List<Job> jobs)
        {
            DriverName = driverName;
            Devices = devices;
            Jobs = jobs;
        }
    }

    public class DeviceConfig
    {
        public string Device { get; set; } = "";

        public DeviceConfig() { }

        public DeviceConfig(string device)
        {
            Device = device;
        }
    }

    public class Job
    {
        public string Device { get; set; }
        public long Interval { get; set; } // ns
        public List<Property> Properties { get; set; }

        public Job(string device, long interval, List<Property> properties)
        {
            Device = device;
            Interval = interval;
            Properties = properties;
        }
    }

    public class Property
    {
        public string Name { get; set; }
        public string Type { get; set; }
        public int Index { get; set; }

        public Property(string name, string type, int index)
        {
            Name = name;
            Type = type;
            Index = index;
        }
    }
}

