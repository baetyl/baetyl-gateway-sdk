package com.baidu.bce.sdk.context;

import com.baidu.bce.sdk.context.models.access.AccessConfig;
import com.baidu.bce.sdk.context.models.access.AccessTemplate;
import com.baidu.bce.sdk.context.models.access.ModelMapping;
import com.baidu.bce.sdk.context.models.device.DeviceInfo;
import com.baidu.bce.sdk.context.models.device.DeviceProperty;
import com.baidu.bce.sdk.context.models.device.PropertyVisitor;
import com.baidu.bce.sdk.context.models.yaml.SubDeviceYaml;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.io.File;
import java.net.URL;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class DmContextTest {
    @Test
    public void testLoadYamlConfig() {
        URL ul = DmContextTest.class.getClassLoader().getResource("models.yml");
        Assertions.assertNotNull(ul);

        String modelPath = ul.getPath();
        File f = new File(modelPath);

        String path = f.getParent();
        String driverName = "custom-java";

        Context dm = new Context();
        dm.loadYamlConfig(path, driverName);

        // access_template.yaml
        Map<String, AccessTemplate> accYamlRes = dm.getAllAccessTemplates(driverName);

        List<DeviceProperty> dpList = new ArrayList<>();
        DeviceProperty dp = new DeviceProperty();
        dp.setName("temperature");
        dp.setId("1");
        dp.setType("float32");
        dp.setCurrent(null);
        dp.setExpect(null);
        dp.setVisitor(new PropertyVisitor("{\"name\":\"temperature\",\"type\":\"float32\",\"index\":0}"));
        dpList.add(dp);

        List<ModelMapping> mmList = new ArrayList<>();
        ModelMapping mm = new ModelMapping();
        mm.setAttribute("temperature");
        mm.setType("value");
        mm.setExpression("x1");
        mm.setPrecision(0);
        mm.setDeviation(0);
        mm.setSilentWin(0);
        mmList.add(mm);

        Map<String, AccessTemplate> accYamlExp = new HashMap<>();
        accYamlExp.put("custom-access-template", new AccessTemplate("custom-access-template", null, dpList, mmList));

        Assertions.assertEquals(accYamlExp, accYamlRes);

        // models.yaml
        Map<String, List<DeviceProperty>> modelYamlRes = dm.getAllDeviceModels(driverName);

        List<DeviceProperty> modelDpList = new ArrayList<>();
        DeviceProperty modelDp0 = new DeviceProperty();
        modelDp0.setName("temperature");
        modelDp0.setMode("ro");
        modelDp0.setType("float32");
        modelDpList.add(modelDp0);
        DeviceProperty modelDp1 = new DeviceProperty();
        modelDp1.setName("humidity");
        modelDp1.setMode("ro");
        modelDp1.setType("float32");
        modelDpList.add(modelDp1);
        DeviceProperty modelDp2 = new DeviceProperty();
        modelDp2.setName("pressure");
        modelDp2.setMode("rw");
        modelDp2.setType("float32");
        modelDpList.add(modelDp2);

        Map<String, List<DeviceProperty>> modelYamlExp = new HashMap<>();
        modelYamlExp.put("custom-simulator", modelDpList);

        Assertions.assertEquals(modelYamlExp, modelYamlRes);

        // sub_devices.yaml
        SubDeviceYaml subDeviceYamlRes = dm.getSubDeviceYamls().get(driverName);

        List<DeviceInfo> devices = new ArrayList<>();
        devices.add(new DeviceInfo("custom-zx",
                "1680267765qc3zxy",
                "custom-simulator",
                "custom-access-template",
                new AccessConfig("{\"device\": \"custom-dev-0\",\"interval\": 3000000000}")));
        SubDeviceYaml subDeviceYamlExp = new SubDeviceYaml(devices, "");

        Assertions.assertEquals(subDeviceYamlExp, subDeviceYamlRes);

        // getDriverNameByDevice
        String dnRes = dm.getDriverNameByDevice("custom-zx");
        Assertions.assertEquals(driverName, dnRes);

        // getAllDevices
        List<DeviceInfo> devList = dm.getAllDevices(driverName);

        DeviceInfo info = new DeviceInfo("custom-zx",
                "1680267765qc3zxy",
                "custom-simulator",
                "custom-access-template",
                new AccessConfig("{\"device\": \"custom-dev-0\",\"interval\": 3000000000}"));
        List<DeviceInfo> devExp = new ArrayList<>();
        devExp.add(info);

        Assertions.assertEquals(devExp, devList);
    }

}
