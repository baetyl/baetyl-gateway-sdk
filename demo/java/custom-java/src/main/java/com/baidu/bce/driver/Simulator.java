package com.baidu.bce.driver;

import com.baidu.bce.models.Point;
import com.baidu.bce.utils.L;

import java.util.Random;
import java.util.concurrent.TimeUnit;

public class Simulator {
    /**
     * SIMULATOR_INTERVAL: 10 seconds
     */
    private static final int SIMULATOR_INTERVAL = 10;
    private String name;
    private Point point;
    private Random random;

    public Simulator(String name) {
        this.random = new Random(System.currentTimeMillis());

        this.name = name;
        this.point = new Point();
        this.point.setTemperature(this.random.nextInt(100));
        this.point.setHumidity(this.random.nextInt(100));
        this.point.setPressure(this.random.nextInt(100));

        generateSimulateData();
    }

    public void set(int index, float val) {
        switch (index) {
            case Point.INDEX_TEMPERATURE -> this.point.setTemperature(val);
            case Point.INDEX_HUMIDITY -> this.point.setHumidity(val);
            case Point.INDEX_PRESSURE -> this.point.setPressure(val);
        }
    }

    public Object get(int index) throws Exception {
        switch (index) {
            case Point.INDEX_TEMPERATURE -> {
                return this.point.getTemperature();
            }
            case Point.INDEX_HUMIDITY -> {
                return this.point.getHumidity();
            }
            case Point.INDEX_PRESSURE -> {
                return this.point.getPressure();
            }
        }
        throw new Exception("unsupported point information");
    }

    private void generateSimulateData() {
        new Thread(() -> {
            while (true) {
                this.point.setTemperature(this.random.nextInt(100));
                this.point.setHumidity(this.random.nextInt(100));
                this.point.setPressure(this.random.nextInt(100));

                L.debug("Simulate generate random temperature = " + this.point.getTemperature());
                L.debug("Simulate generate random humidity = " + this.point.getHumidity());
                L.debug("Simulate generate random pressure = " + this.point.getPressure());
                try {
                    TimeUnit.SECONDS.sleep(SIMULATOR_INTERVAL);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }).start();
    }
}
