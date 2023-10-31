using DriverSDK.DmContext;
using DriverSDK.DmContext.Models;
using Newtonsoft.Json;

namespace DriverSDK.Tests.DmContext
{
    [TestFixture]
    public class ContextTests
    {
        private string _basePath;

        [SetUp]
        public void SetUp()
        {
            _basePath = Path.Combine(TestContext.CurrentContext.TestDirectory, "../../../TestData");
        }

        [Test]
        public void LoadYamlConfig_Success()
        {
            var dm = new Context();
            string driverName = "custom-csharp";
            dm.LoadYamlConfig(_basePath, driverName);

            // access_template.yaml
            IList<DeviceProperty> dpList = new List<DeviceProperty>();
            DeviceProperty dp = new()
            {
                Name = "temperature",
                Id = "1",
                Type = "float32",
                Visitor = new PropertyVisitor("{\"name\":\"temperature\",\"type\":\"float32\",\"index\":0}")
            };
            dpList.Add(dp);

            IList<ModelMapping> mmList = new List<ModelMapping>();
            ModelMapping mm = new()
            {
                Attribute = "temperature",
                Type = "value",
                Expression = "x1",
                Precision = 0,
                Deviation = 0,
                SilentWin = 0
            };
            mmList.Add(mm);

            var accYamlExp = new Dictionary<string, AccessTemplate>();
            AccessTemplate at = new()
            {
                Name = "custom-access-template",
                Properties = dpList,
                Mappings = mmList
            };
            accYamlExp.Add("custom-access-template", at);

            var accYamlRes = dm.GetAllAccessTemplates(driverName);

            string aye = JsonConvert.SerializeObject(accYamlExp);
            string ayr = JsonConvert.SerializeObject(accYamlRes);

            Assert.That(ayr, Is.EqualTo(aye));

            // models.yaml
            var modelYamlRes = dm.GetAllDeviceModels(driverName);

            IList<DeviceProperty> modelDpList = new List<DeviceProperty>();
            DeviceProperty modelDp0 = new()
            {
                Name = "temperature",
                Mode = "ro",
                Type = "float32"
            };
            modelDpList.Add(modelDp0);
            DeviceProperty modelDp1 = new()
            {
                Name = "humidity",
                Mode = "ro",
                Type = "float32"
            };
            modelDpList.Add(modelDp1);
            DeviceProperty modelDp2 = new()
            {
                Name = "pressure",
                Mode = "rw",
                Type = "float32"
            };
            modelDpList.Add(modelDp2);

            IDictionary<string, IList<DeviceProperty>> modelYamlExp = new Dictionary<string, IList<DeviceProperty>>
            {
                { "custom-simulator", modelDpList }
            };

            string mye = JsonConvert.SerializeObject(modelYamlExp);
            string myr = JsonConvert.SerializeObject(modelYamlRes);

            Assert.That(myr, Is.EqualTo(mye));

            // sub_devices.yaml
            SubDeviceYaml subDeviceYamlRes = dm.GetSubDeviceYamls()[driverName];

            IList<DeviceInfo> devices = new List<DeviceInfo>();
            devices.Add(new DeviceInfo("custom-zx",
                    "1680267765qc3zxy",
                    "custom-simulator",
                    "custom-access-template",
                    new AccessConfig("{\"device\": \"custom-dev-0\",\"interval\": 3000000000}")));
            SubDeviceYaml subDeviceYamlExp = new(devices, "");

            string sdye = JsonConvert.SerializeObject(subDeviceYamlExp);
            string sdyr = JsonConvert.SerializeObject(subDeviceYamlRes);

            Assert.That(sdyr, Is.EqualTo(sdye));

            string dnRes = dm.GetDriverNameByDevice("custom-zx");
            Assert.That(dnRes, Is.EqualTo(driverName));
        }
    }
}

