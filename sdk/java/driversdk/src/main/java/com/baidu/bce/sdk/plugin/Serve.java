package com.baidu.bce.sdk.plugin;

import com.baidu.bce.sdk.utils.ProcessInfo;
import io.grpc.Server;
import io.grpc.ServerBuilder;

import java.io.IOException;
import java.util.concurrent.TimeUnit;

public class Serve {
    // todo 支持 tls
    private final GrpcDriver grpcDriver;
    private final GrpcHealth grpcHealth;

    public Serve(IDriver driver) {
        this.grpcDriver = new GrpcDriver(driver);
        this.grpcHealth = new GrpcHealth();
    }

    // 每隔 5s 检查父进程是否在存在，不存在则关闭驱动子进程
    private void checker() {
        new Thread(() -> {
            while (true) {
                // 检查条件是否满足
                if (ProcessInfo.getPidOfParentProcess() == -1) {
                    System.out.println("ProcessInfo.getPidOfParentProcess() == -1");
                    // 条件不满足时退出程序
                    System.exit(0);
                }
                try {
                    // 每隔5秒钟进行一次检查
                    TimeUnit.SECONDS.sleep(5);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }).start();
    }

    public void start() {
        checker();

        try {
            Server server = ServerBuilder.forPort(0) // 使用0表示随机选择可用端口
                    .addService(this.grpcHealth) // go-plugin 要求有健康检查服务
                    .addService(this.grpcDriver) // 添加服务实现类
                    .build();
            // 启动服务器
            server.start();

            // go-plugin 要求必须调用系统打印函数输出下面格式的数据
            int port = server.getPort();
            System.out.println("1|1|tcp|127.0.0.1:" + port + "|grpc");

            // 等待服务器终止
            server.awaitTermination();
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}
