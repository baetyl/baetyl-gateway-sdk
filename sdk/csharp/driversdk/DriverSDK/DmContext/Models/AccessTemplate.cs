using System.Xml.Linq;

namespace DriverSDK.DmContext.Models
{
    public class AccessTemplate
    {
        public string Name { get; set; } = "";
        public string Version { get; set; } = "";
        public IList<DeviceProperty>? Properties { get; set; }
        public IList<ModelMapping>? Mappings { get; set; }

        public AccessTemplate() { }

        public AccessTemplate(string name, string version, IList<DeviceProperty> properties,  IList<ModelMapping> mappings)
        {
            this.Name = name;
            this.Version = version;
            this.Properties = properties;
            this.Mappings = mappings;
        }
    }

    public class AccessConfig
    {
        public string Custom { get; set; } = "";

        public AccessConfig() { }

        public AccessConfig(string custom)
        {
            this.Custom = custom;
        }
    }

    public class ModelMapping
    {
        public string Attribute { get; set; } = "";
        public string Type { get; set; } = "";
        public string Expression { get; set; } = "";
        public int Precision { get; set; } = 0;
        public double Deviation { get; set; } = 0;
        public int SilentWin { get; set; } = 0;

        public ModelMapping() { }

        public ModelMapping(string attribute, string type, string expression, int precision, double deviation, int silentWin) 
        {
            this.Attribute = attribute;
            this.Type = type;
            this.Expression = expression;
            this.Precision = precision;
            this.Deviation = deviation;
            this.SilentWin = silentWin;
        }
    }
}

