package com.baidu.bce.driver;

import com.baidu.bce.models.DeviceConfig;
import com.baidu.bce.models.DriverConfig;
import com.baidu.bce.models.Job;
import com.baidu.bce.models.Property;
import com.baidu.bce.sdk.context.Context;
import com.baidu.bce.sdk.context.models.device.DeviceInfo;
import com.baidu.bce.sdk.context.models.device.Event;
import com.baidu.bce.sdk.plugin.IReport;
import com.baidu.bce.utils.L;
import lombok.Data;

import java.util.HashMap;
import java.util.Map;

@Data
public class Custom {
    private final Context ctx;
    private DriverConfig cfg;
    private Map<String, Worker> ws;
    private Map<String, Device> devs;
    private IReport report;
    private final Map<String, Thread> threads;

    public Custom(Context ctx, DriverConfig cfg, IReport report) {
        this.ctx = ctx;
        this.cfg = cfg;
        this.report = report;

        Map<String, DeviceInfo> infos = new HashMap<>();
        for (DeviceInfo info : this.ctx.getAllDevices(this.cfg.getDriverName())) {
            infos.put(info.getName(), info);
        }

        this.devs = new HashMap<>();
        for (DeviceConfig item : this.cfg.getDevices()) {
            if (infos.containsKey(item.getDevice())) {
                Device dev = new Device(item, infos.get(item.getDevice()));
                this.devs.put(item.getDevice(), dev);
            }
        }

        this.ws = new HashMap<>();
        for (Job job : this.cfg.getJobs()) {
            Device dev = this.devs.get(job.getDevice());
            if (dev != null) {
                this.ws.put(dev.getInfo().getName(), new Worker(job, dev, this.cfg.getDriverName(), this.report));
            }
        }

        this.threads = new HashMap<>();
    }

    public void start() {
        for (String key : this.ws.keySet()) {
            Thread td = new Thread(() -> {
                ws.get(key).working();
                synchronized (this.threads) {
                    this.threads.remove(key);
                }
            });
            this.threads.put(key, td);
            td.start();
        }
    }

    public void restart() {
        synchronized (this.threads) {
            for (Thread td : this.threads.values()) {
                td.interrupt();
            }
            this.threads.clear();

            for (String key : this.ws.keySet()) {
                Thread td = new Thread(() -> {
                    ws.get(key).working();
                    this.threads.remove(key);
                });
                this.threads.put(key, td);
                td.start();
            }
        }
    }

    public void stop() {
        synchronized (this.threads) {
            for (Thread td : this.threads.values()) {
                td.interrupt();
            }
            this.threads.clear();
        }
    }

    public void set(DeviceInfo info, Map<String, Object> props) {
        L.debug("Custom set (name = " + info.getName()+")");
        if (!this.ws.containsKey(info.getName())) {
            return;
        }
        Worker w = this.ws.get(info.getName());
        for (Map.Entry<String, Object> e : props.entrySet()) {
            for (Property p : w.getJob().getProperties()) {
                if (e.getKey().equals(p.getName())) {
                    Object val = e.getValue();
                    if (val instanceof Double ||
                            val instanceof Float ||
                            val instanceof Integer ||
                            val instanceof Long) {
                        w.getDevice().set(p.getIndex(), Float.parseFloat(val.toString()));
                    } else {
                        L.debug("custom set " + e.getKey() + " error");
                    }
                }
            }
        }
    }

    public void event(DeviceInfo info, Event event) {
        if (event.getType().equals(Context.TYPE_REPORT_EVENT)) {
            propertyGet(info);
        } else {
            L.debug("custom event " + info.getName());
        }
    }

    public void propertyGet(DeviceInfo info) {
        L.debug("Custom get (name = " + info.getName()+")");
        if (this.ws.containsKey(info.getName())) {
            Worker w = this.ws.get(info.getName());
            try {
                w.reportProperty();
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
}
