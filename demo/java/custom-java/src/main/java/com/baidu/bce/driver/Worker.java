package com.baidu.bce.driver;

import com.alibaba.fastjson2.JSON;
import com.baidu.bce.models.Job;
import com.baidu.bce.models.Property;
import com.baidu.bce.sdk.context.Context;
import com.baidu.bce.sdk.context.models.Message;
import com.baidu.bce.sdk.plugin.IReport;
import com.baidu.bce.utils.L;
import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;

@Data
@AllArgsConstructor
public class Worker {
    private Job job;
    private Device device;
    private String driverName;
    private IReport report;

    public void working() {
        L.debug("Worker interval " + this.job.getInterval() + " ns");
        new Thread(() -> {
            while (true) {
                try {
                    reportProperty();
                } catch (Exception e) {
                    e.printStackTrace();
                } finally {
                    try {
                        TimeUnit.NANOSECONDS.sleep(this.job.getInterval());
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
        }).start();
    }

    public void reportProperty() throws Exception {
        L.debug("Worker reportProperty");

        Map<String, Object> props = new HashMap<>();
        for (Property item : this.job.getProperties()) {
            Object val = this.device.get(item.getIndex());
            props.put(item.getName(), val);
        }

        Map<String, String> meta = new HashMap<>();
        meta.put(Context.KEY_DRIVER_NAME, this.driverName);
        meta.put(Context.KEY_DEVICE_NAME, this.device.getInfo().getName());
        Message msg = new Message(Context.MESSAGE_DEVICE_REPORT, meta, props);

        String jStr = JSON.toJSONString(msg);

        L.debug("Worker reportProperty msg " + jStr);
        this.report.post(jStr);
    }
}
